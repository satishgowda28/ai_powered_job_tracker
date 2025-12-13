-- name: CreateRefreshToken :one
INSERT INTO user_refresh_token (token, user_id, created_at, updated_at, expires_at)
VALUES ($1, $2, NOW(), NOW(), $3)
RETURNING *;

-- name: RevokeRefreshToken :one
UPDATE user_refresh_token SET revoked_at = NOW(), updated_at = NOW()
WHERE token = $1
RETURNING *;

-- name: GetUserFromRefreshToken :one
SELECT users.* from users
JOIN user_refresh_token ON users.id = user_refresh_token.user_id
WHERE user_refresh_token.token = $1
AND revoked_at IS NULL
AND expires_at > NOW();