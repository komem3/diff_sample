package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/go-chi/chi"
	"github.com/rs/xid"
)

var dsClient *datastore.Client

func main() {
	ctx := context.Background()

	var err error
	dsClient, err = datastore.NewClient(ctx, "my-project")
	if err != nil {
		log.Fatal(err)
	}
	defer dsClient.Close()

	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Post("/create", CreateHandler)
	})

	port := ":8080"
	log.Printf("[DEBUG] listen port %s\n", port)
	log.Panic(http.ListenAndServe(port, r))
}

type Message struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type Error struct {
	Message string `json:"message"`
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var m Message
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		ErrorResp(w, http.StatusInternalServerError, err)
		return
	}
	m.ID = xid.New().String()

	if _, err := dsClient.Mutate(ctx,
		datastore.NewInsert(datastore.NameKey("Message", m.ID, nil), &m),
	); err != nil {
		ErrorResp(w, http.StatusInternalServerError, err)
		return
	}

	if err := json.NewEncoder(w).Encode(&m); err != nil {
		ErrorResp(w, http.StatusInternalServerError, err)
		return
	}
}

func ErrorResp(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(&Error{Message: err.Error()}); err != nil {
		log.Panic(err)
	}
	log.Printf("[ERROR] %v\n", err)
}
