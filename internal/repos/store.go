package repos

import (
	"sync"
	"sync/atomic"
	"time"
)

type Repo struct {
	ID       uint64    `json:"id"`
	Provider string    `json:"provider"`  // github | gitlab (for now just a string)
	FullName string    `json:"full_name"` // e.g. "Mahran1998/pipewatch"
	BaseURL  string    `json:"base_url"`  // e.g. https://github.com
	AddedAt  time.Time `json:"added_at"`
}

type Store struct {
	mu    sync.RWMutex
	next  uint64
	repos []Repo
}

func NewStore() *Store { return &Store{} }

func (s *Store) Add(provider, fullName, baseURL string) Repo {
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

	return r
}

func (s *Store) List() []Repo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]Repo, len(s.repos))
	copy(out, s.repos)
	return out
}
