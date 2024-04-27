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
