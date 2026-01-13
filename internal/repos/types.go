package repos

import (
	"context"
	"time"
)

type Repo struct {
	ID       uint64    `json:"id"`
	Provider string    `json:"provider"`
	FullName string    `json:"full_name"`
	BaseURL  string    `json:"base_url"`
	AddedAt  time.Time `json:"added_at"`
}

type Store interface {
	Add(ctx context.Context, provider, fullName, baseURL string) (Repo, error)
	List(ctx context.Context) ([]Repo, error)
	Close() error
}
