-- name: CreateParameter :exec
INSERT INTO parameters ( ID, delete_at_days, percentage_pricing, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5 );

-- name: GetParameter :one
SELECT *
FROM parameters
WHERE parameters.id = $1;

-- name: GetParameters :many
SELECT *
FROM parameters;

-- name: UpdateParameter :exec
UPDATE parameters SET delete_at_days = $2, percentage_pricing = $3, updated_at = $4 WHERE parameters.id = $1;


