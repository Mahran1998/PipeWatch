CREATE TABLE IF NOT EXISTS repos (
  id BIGSERIAL PRIMARY KEY,
  provider TEXT NOT NULL,
  full_name TEXT NOT NULL,
  base_url TEXT NOT NULL,
  added_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_repos_provider ON repos(provider);
CREATE UNIQUE INDEX IF NOT EXISTS uq_repos_provider_fullname ON repos(provider, full_name);
