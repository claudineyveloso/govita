-- name: CreateSearch :exec
INSERT INTO searches ( ID, description, created_at, updated_at)
VALUES ($1, $2, $3, $4);

-- name: GetSearch :one
SELECT *
FROM searches
WHERE searches.id = $1;

-- name: GetSearches :many
SELECT *
FROM searches;

-- name: DeleteSearch :exec
DELETE FROM searches WHERE searches.id = $1;
