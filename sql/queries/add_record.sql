-- name: CreateRecord :one
INSERT INTO records_2026 (
        record_date,
        record_title,
        record_type,
        updated_at
    )
VALUES ($1, $2, $3, NOW())
RETURNING *;
-- name: ResetRecords :exec
DELETE FROM records_2026;
-- name: GetAllRecords :many
SELECT *
FROM records_2026
ORDER BY record_date ASC;
-- name: DeleteRecord :exec
DELETE FROM records_2026
WHERE id = $1;