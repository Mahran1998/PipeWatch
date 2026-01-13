package httpapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mahran1998/pipewatch/internal/repos"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	Router(repos.NewMemoryStore()).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}
