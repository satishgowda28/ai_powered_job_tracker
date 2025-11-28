-- name: InsertJobStatusHistory :exec
INSERT INTO job_status_history (job_id, old_status, new_status)
VALUES ($1, $2, $3);

-- name: GetHistoryForJob :many
SELECT * FROM job_status_history
WHERE job_id = $1
ORDER BY changed_at DESC;