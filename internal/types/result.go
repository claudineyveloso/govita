package types

import (
	"time"

	"github.com/google/uuid"
)

type Result struct {
	ID          uuid.UUID `json:"id"`
	ImageUrl    string    `json:"image_url"`
	Description string    `json:"description"`
	Font        string    `json:"font"`
	Price       float64   `json:"price"`
	Promotion   bool      `json:"promotion"`
	SearchID    uuid.UUID `json:"search_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ResultStore interface {
	CreateResult(Result) error
}
