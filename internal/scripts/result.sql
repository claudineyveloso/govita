-- name: CreateResult :exec
INSERT INTO results ( ID, image_url, description, font, price, promotion, search_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: GetResult :one
SELECT *
FROM results
WHERE results.id = $1;

-- name: GetResults :many
SELECT *
FROM results;
