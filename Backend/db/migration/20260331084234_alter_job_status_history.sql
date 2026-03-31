-- +goose Up
-- +goose StatementBegin
ALTER TABLE job_status_history
  ADD CONSTRAINT fk_old_status FOREIGN KEY (old_status) REFERENCES job_application_statuse(slug),
  ADD CONSTRAINT fk_new_status FOREIGN KEY (new_status) REFERENCES job_application_statuse(slug);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
