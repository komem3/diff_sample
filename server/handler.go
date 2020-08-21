package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"cloud.google.com/go/datastore"
	"github.com/go-chi/chi"
	"github.com/rs/xid"
)

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

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}