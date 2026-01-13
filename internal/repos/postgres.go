package repos

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStore struct {
	pool *pgxpool.Pool
}

func NewPostgresStore(ctx context.Context, dbURL string) (*PostgresStore, error) {
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, err
	}
	return &PostgresStore{pool: pool}, nil
}

func (s *PostgresStore) Close() error {
	s.pool.Close()
	return nil
}

func (s *PostgresStore) Add(ctx context.Context, provider, fullName, baseURL string) (Repo, error) {
	var r Repo
	err := s.pool.QueryRow(ctx, `
		INSERT INTO repos (provider, full_name, base_url)
		VALUES ($1, $2, $3)
		ON CONFLICT (provider, full_name) DO UPDATE
		SET base_url = EXCLUDED.base_url
		RETURNING id, provider, full_name, base_url, added_at
	`, provider, fullName, baseURL).Scan(&r.ID, &r.Provider, &r.FullName, &r.BaseURL, &r.AddedAt)
	return r, err
}

func (s *PostgresStore) List(ctx context.Context) ([]Repo, error) {
	rows, err := s.pool.Query(ctx, `
		SELECT id, provider, full_name, base_url, added_at
		FROM repos
		ORDER BY id ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Repo
	for rows.Next() {
		var r Repo
		if err := rows.Scan(&r.ID, &r.Provider, &r.FullName, &r.BaseURL, &r.AddedAt); err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, rows.Err()
}
