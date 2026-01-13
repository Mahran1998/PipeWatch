package httpapi

import (
	"net/http"

	"github.com/Mahran1998/pipewatch/internal/repos"
)

func Router() http.Handler {
	store := repos.NewStore()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/repos", reposHandler(store))
	return mux
}
