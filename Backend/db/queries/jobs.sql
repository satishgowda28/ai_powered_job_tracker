-- name: CreateJob :one
INSERT INTO jobs (user_id, company, title, job_description, job_location, salary, job_url, status, notes)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetJobs :many
SELECT * FROM jobs
WHERE user_id = $1
ORDER BY applied_at DESC
LIMIT $2 OFFSET $3;

-- name: GetJobByID :one
SELECT * FROM jobs
WHERE id = $1 AND user_id = $2;

-- name: UpdateJobStatus :one
UPDATE jobs SET Status = $2, update_at = NOW()
WHERE id = $1
RETURNING *;