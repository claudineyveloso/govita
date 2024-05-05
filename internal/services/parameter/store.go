package parameter

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/claudineyveloso/govita.git/internal/db"
	"github.com/claudineyveloso/govita.git/internal/types"
	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateParameter(parameter types.Parameter) error {
	queries := db.New(s.db)
	ctx := context.Background()

	parameter.ID = uuid.New()
	now := time.Now()
	parameter.CreatedAt = now
	parameter.UpdatedAt = now

	createParameterParams := db.CreateParameterParams{
		ID:                parameter.ID,
		DeleteAtDays:      parameter.DeleteAtDays,
		PercentagePricing: parameter.PercentagePricing,
		CreatedAt:         parameter.CreatedAt,
		UpdatedAt:         parameter.UpdatedAt,
	}

	if err := queries.CreateParameter(ctx, createParameterParams); err != nil {
		fmt.Println("Erro ao criar um parâmetro:", err)
		return err
	}
	return nil
}

func (s *Store) UpdateParameter(parameter types.Parameter) error {
  queries := db.New(s.db)
  ctx := context.Background()

  now := time.Now()
  parameter.UpdatedAt = now

  updateParameterParams := db.UpdateParameterParams{
    ID:                parameter.ID,
    DeleteAtDays:      parameter.DeleteAtDays,
    PercentagePricing: parameter.PercentagePricing,
    UpdatedAt:         parameter.UpdatedAt,
  }

  if err := queries.UpdateParameter(ctx, updateParameterParams); err != nil {
    fmt.Println("Erro ao atualizar um parâmetros:", err)
    return err
  }
  return nil
}


func (s *Store) GetParameterByID(id uuid.UUID) (*types.Parameter, error) {
	queries := db.New(s.db)
	ctx := context.Background()

	dbParameter, err := queries.GetParameter(ctx, id)
	if err != nil {
		fmt.Println("Erro ao buscar um parâmetro:", err)
		return nil, err
	}

	parameter := &types.Parameter{
		ID:                dbParameter.ID,
		DeleteAtDays:      dbParameter.DeleteAtDays,
		PercentagePricing: dbParameter.PercentagePricing,
		CreatedAt:         dbParameter.CreatedAt,
		UpdatedAt:         dbParameter.UpdatedAt,
	}

	return parameter, nil
}

func (s *Store) GetParameters() ([]*types.Parameter, error) {
	queries := db.New(s.db)
	ctx := context.Background()

	dbParameters, err := queries.GetParameters(ctx)
	if err != nil {
		fmt.Println("Erro ao buscar parâmetros:", err)
		return nil, err
	}

	var parameters []*types.Parameter
	for _, dbParameter := range dbParameters {
		parameter := &types.Parameter{
			ID:                dbParameter.ID,
			DeleteAtDays:      dbParameter.DeleteAtDays,
			PercentagePricing: dbParameter.PercentagePricing,
			CreatedAt:         dbParameter.CreatedAt,
			UpdatedAt:         dbParameter.UpdatedAt,
		}
		parameters = append(parameters, parameter)
	}

	return parameters, nil
}
