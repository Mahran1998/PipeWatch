package httpapi

import (
	"net/http"

	"github.com/Mahran1998/pipewatch/internal/repos"
)

func Router(store repos.Store) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/repos", reposHandler(store))
	return mux
}
