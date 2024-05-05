// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: search.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSearch = `-- name: CreateSearch :exec
INSERT INTO searches ( ID, description, created_at, updated_at)
VALUES ($1, $2, $3, $4)
`

type CreateSearchParams struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (q *Queries) CreateSearch(ctx context.Context, arg CreateSearchParams) error {
	_, err := q.db.ExecContext(ctx, createSearch,
		arg.ID,
		arg.Description,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const getSearch = `-- name: GetSearch :one
SELECT id, description, created_at, updated_at
FROM searches
WHERE searches.id = $1
`

func (q *Queries) GetSearch(ctx context.Context, id uuid.UUID) (Search, error) {
	row := q.db.QueryRowContext(ctx, getSearch, id)
	var i Search
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSearches = `-- name: GetSearches :many
SELECT id, description, created_at, updated_at
FROM searches
`

func (q *Queries) GetSearches(ctx context.Context) ([]Search, error) {
	rows, err := q.db.QueryContext(ctx, getSearches)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Search
	for rows.Next() {
		var i Search
		if err := rows.Scan(
			&i.ID,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}