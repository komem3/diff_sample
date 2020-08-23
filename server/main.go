package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"cloud.google.com/go/datastore"
	"github.com/go-chi/chi"
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

	r.Route("/api/message", func(r chi.Router) {
		r.Post("/", CreateHandler)
		r.Patch("/{id}", UpdateHandler)
	})

	// Create a route along /files that will serve contents from
	// the ./data/ folder.
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "dist"))
	FileServer(r, "/", filesDir)

	port := ":8080"
	log.Printf("[DEBUG] listen port %s\n", port)
	log.Panic(http.ListenAndServe(port, r))
}
