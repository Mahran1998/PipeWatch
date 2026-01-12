package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Mahran1998/pipewatch/internal/httpapi"
)

func main() {
	addr := env("PIPEWATCH_ADDR", ":8080")

	srv := &http.Server{
		Addr:              addr,
		Handler:           httpapi.Router(),
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
