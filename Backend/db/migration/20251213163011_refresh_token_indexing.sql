-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_user_refresh_token_user_id ON user_refresh_token(user_id);
CREATE INDEX idx_user_refresh_token_expiry ON user_refresh_token(expires_at, revoked_at) WHERE revoked_at IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
