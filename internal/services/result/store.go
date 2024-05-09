package result

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

func (s *Store) CreateResult(result types.Result) error {
	queries := db.New(s.db)
	ctx := context.Background()

	result.ID = uuid.New()
	now := time.Now()
	result.CreatedAt = now
	result.UpdatedAt = now

	createResultParams := db.CreateResultParams{
		ID:          result.ID,
		ImageUrl:    result.ImageUrl,
		Description: result.Description,
		Font:        result.Font,
		Price:       result.Price,
		Promotion:   result.Promotion,
		SearchID:    result.SearchID,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}
	if err := queries.CreateResult(ctx, createResultParams); err != nil {
		fmt.Println("Erro ao criar um resultado da busca:", err)
		return err
	}
	return nil
}

func (s *Store) GetResults() ([]*types.Result, error) {
	queries := db.New(s.db)
	ctx := context.Background()

	dbResults, err := queries.GetResults(ctx)
	if err != nil {
		fmt.Println("Erro ao buscar parâmetros:", err)
		return nil, err
	}

	var results []*types.Result
	for _, dbResult := range dbResults {
		result := &types.Result{
			ID:          dbResult.ID,
			ImageUrl:    dbResult.ImageUrl,
			Description: dbResult.Description,
			Font:        dbResult.Font,
			Price:       dbResult.Price,
			Promotion:   dbResult.Promotion,
			SearchID:    dbResult.SearchID,
			CreatedAt:   dbResult.CreatedAt,
			UpdatedAt:   dbResult.UpdatedAt,
		}
		results = append(results, result)
	}

	return results, nil
}

func (s *Store) GetResultByID(id uuid.UUID) (*types.Result, error) {
	queries := db.New(s.db)
	ctx := context.Background()

	dbResult, err := queries.GetResult(ctx, id)
	if err != nil {
		fmt.Println("Erro ao buscar um resultado:", err)
		return nil, err
	}

	result := &types.Result{
		ID:          dbResult.ID,
		ImageUrl:    dbResult.ImageUrl,
		Description: dbResult.Description,
		Font:        dbResult.Font,
		Price:       dbResult.Price,
		Promotion:   dbResult.Promotion,
		SearchID:    dbResult.SearchID,
		CreatedAt:   dbResult.CreatedAt,
		UpdatedAt:   dbResult.UpdatedAt,
	}

	return result, nil
}


func (s *Store) DeleteResult(id uuid.UUID) error {
	queries := db.New(s.db)
	ctx := context.Background()

	if err := queries.DeleteResult(ctx, id); err != nil {
		fmt.Println("Erro ao deletar o resultado:", err)
		return err
	}

	return nil
}

