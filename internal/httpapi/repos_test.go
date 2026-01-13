package httpapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRepos_POST_then_GET(t *testing.T) {
	h := Router()

	// POST /repos
	body := []byte(`{"provider":"github","full_name":"Mahran1998/pipewatch","base_url":"https://github.com"}`)
	req := httptest.NewRequest(http.MethodPost, "/repos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("POST expected 201, got %d, body=%s", rec.Code, rec.Body.String())
	}

	// GET /repos
	req2 := httptest.NewRequest(http.MethodGet, "/repos", nil)
	rec2 := httptest.NewRecorder()

	h.ServeHTTP(rec2, req2)
	if rec2.Code != http.StatusOK {
		t.Fatalf("GET expected 200, got %d", rec2.Code)
	}

	var got []map[string]any
	if err := json.Unmarshal(rec2.Body.Bytes(), &got); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("expected 1 repo, got %d", len(got))
	}
}
