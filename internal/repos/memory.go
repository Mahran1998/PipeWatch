package repos

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

type MemoryStore struct {
	mu    sync.RWMutex
	next  uint64
	repos []Repo
}

func NewMemoryStore() *MemoryStore { return &MemoryStore{} }

func (s *MemoryStore) Add(ctx context.Context, provider, fullName, baseURL string) (Repo, error) {
	_ = ctx
	id := atomic.AddUint64(&s.next, 1)

	r := Repo{
		ID:       id,
		Provider: provider,
		FullName: fullName,
		BaseURL:  baseURL,
		AddedAt:  time.Now().UTC(),
	}

	s.mu.Lock()
	s.repos = append(s.repos, r)
	s.mu.Unlock()

	return r, nil
}

func (s *MemoryStore) List(ctx context.Context) ([]Repo, error) {
	_ = ctx
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]Repo, len(s.repos))
	copy(out, s.repos)
	return out, nil
}

func (s *MemoryStore) Close() error { return nil }
