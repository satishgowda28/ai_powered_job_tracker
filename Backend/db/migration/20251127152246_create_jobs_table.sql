-- +goose Up
-- +goose StatementBegin
CREATE TABLE jobs(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID REFERENCES users(id) ON DELETE CASCADE,
  company TEXT NOT NULL,
  title TEXT NOT NULL,
  job_description TEXT,
  job_location TEXT,
  salary INT,
  job_url TEXT,
  status TEXT NOT NULL DEFAULT 'applied',
  notes TEXT,
  applied_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE jobs;
-- +goose StatementEnd
