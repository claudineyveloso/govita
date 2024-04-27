package types

import (
	"time"

	"github.com/google/uuid"
)

type Parameter struct {
	ID                uuid.UUID `json:"id"`
	DeleteAtDays      int32     `json:"delete_at_days"`
	PercentagePricing int32     `json:"percentage_pricing"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type ParameterStore interface {
	CreateParameter(Parameter) error
	GetParameters() ([]*Parameter, error)
	GetParameterByID(id uuid.UUID) (*Parameter, error)
}
