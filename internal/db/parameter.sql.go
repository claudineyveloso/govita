// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: parameter.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createParameter = `-- name: CreateParameter :exec
INSERT INTO parameters ( ID, delete_at_days, percentage_pricing, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5 )
`

type CreateParameterParams struct {
	ID                uuid.UUID `json:"id"`
	DeleteAtDays      int32     `json:"delete_at_days"`
	PercentagePricing int32     `json:"percentage_pricing"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (q *Queries) CreateParameter(ctx context.Context, arg CreateParameterParams) error {
	_, err := q.db.ExecContext(ctx, createParameter,
		arg.ID,
		arg.DeleteAtDays,
		arg.PercentagePricing,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const getParameter = `-- name: GetParameter :one
SELECT id, delete_at_days, percentage_pricing, created_at, updated_at
FROM parameters
WHERE parameters.id = $1
`

func (q *Queries) GetParameter(ctx context.Context, id uuid.UUID) (Parameter, error) {
	row := q.db.QueryRowContext(ctx, getParameter, id)
	var i Parameter
	err := row.Scan(
		&i.ID,
		&i.DeleteAtDays,
		&i.PercentagePricing,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getParameters = `-- name: GetParameters :many
SELECT id, delete_at_days, percentage_pricing, created_at, updated_at
FROM parameters
`

func (q *Queries) GetParameters(ctx context.Context) ([]Parameter, error) {
	rows, err := q.db.QueryContext(ctx, getParameters)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Parameter
	for rows.Next() {
		var i Parameter
		if err := rows.Scan(
			&i.ID,
			&i.DeleteAtDays,
			&i.PercentagePricing,
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

const updateParameter = `-- name: UpdateParameter :exec
UPDATE parameters SET delete_at_days = $2, percentage_pricing = $3, updated_at = $4 WHERE parameters.id = $1
`

type UpdateParameterParams struct {
	ID                uuid.UUID `json:"id"`
	DeleteAtDays      int32     `json:"delete_at_days"`
	PercentagePricing int32     `json:"percentage_pricing"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (q *Queries) UpdateParameter(ctx context.Context, arg UpdateParameterParams) error {
	_, err := q.db.ExecContext(ctx, updateParameter,
		arg.ID,
		arg.DeleteAtDays,
		arg.PercentagePricing,
		arg.UpdatedAt,
	)
	return err
}
