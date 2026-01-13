package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Mahran1998/pipewatch/internal/httpapi"
	"github.com/Mahran1998/pipewatch/internal/repos"
)

func main() {
	addr := env("PIPEWATCH_ADDR", ":8080")
	dbURL := os.Getenv("DATABASE_URL")

	var store repos.Store
	if dbURL != "" {
		pg, err := repos.NewPostgresStore(context.Background(), dbURL)
		if err != nil {
			log.Fatalf("failed to connect to postgres: %v", err)
		}
		store = pg
		log.Printf("repo store: postgres")
	} else {
		store = repos.NewMemoryStore()
		log.Printf("repo store: memory")
	}
	defer store.Close()

	srv := &http.Server{
		Addr:              addr,
		Handler:           httpapi.Router(store),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("pipewatch listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func env(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
