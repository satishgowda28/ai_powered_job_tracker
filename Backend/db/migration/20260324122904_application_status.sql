-- +goose Up
-- +goose StatementBegin
CREATE TABLE job_application_statuse(
  slug TEXT PRIMARY KEY,
  display_name TEXT NOT NULL,
  color_hex TEXT NOT NULL,
  description TEXT
);
-- INSERT INTO application_statuses (slug, display_name, color_hex) VALUES
-- ('applied', 'Applied', '#3498db'),
-- ('interview', 'Interviewing', '#f1c40f'),
-- ('offer', 'Offer Received', '#2ecc71'),
-- ('rejected', 'Rejected', '#e74c3c');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
