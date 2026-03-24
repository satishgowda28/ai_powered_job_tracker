-- +goose Up
-- +goose StatementBegin
ALTER TABLE jobs
  ALTER COLUMN status TYPE TEXT,
  ALTER COLUMN status SET DEFAULT 'applied',
  ALTER COLUMN status SET NOT NULL,
  ADD CONSTRAINT fk_jobs_status
  FOREIGN KEY (status)
  REFERENCES job_application_statuse(slug);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
