package httpapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Mahran1998/pipewatch/internal/repos"
)

type addRepoRequest struct {
	Provider string `json:"provider"`
	FullName string `json:"full_name"`
	BaseURL  string `json:"base_url"`
}

func reposHandler(store repos.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			list, err := store.List(r.Context())
			if err != nil {
				http.Error(w, "db error", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(list)
			return

		case http.MethodPost:
			var req addRepoRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "invalid JSON body", http.StatusBadRequest)
				return
			}

			req.Provider = strings.TrimSpace(req.Provider)
			req.FullName = strings.TrimSpace(req.FullName)
			req.BaseURL = strings.TrimSpace(req.BaseURL)

			if req.Provider == "" || req.FullName == "" || req.BaseURL == "" {
				http.Error(w, "provider, full_name, base_url are required", http.StatusBadRequest)
				return
			}

			created, err := store.Add(r.Context(), req.Provider, req.FullName, req.BaseURL)
			if err != nil {
				http.Error(w, "db error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_ = json.NewEncoder(w).Encode(created)
			return

		default:
			w.Header().Set("Allow", "GET, POST")
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}
