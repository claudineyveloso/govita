package search

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

func (s *Store) CreateSearch(search types.SearchPayload) error {
	queries := db.New(s.db)
	ctx := context.Background()

	search.ID = uuid.New()
	now := time.Now()
	search.CreatedAt = now
	search.UpdatedAt = now

	createSearchParams := db.CreateSearchParams{
		ID:          search.ID,
		Description: search.Description,
		CreatedAt:   search.CreatedAt,
		UpdatedAt:   search.UpdatedAt,
	}

	if err := queries.CreateSearch(ctx, createSearchParams); err != nil {
		fmt.Println("Erro ao criar um busca:", err)
		return err
	}

	return nil
}

func (s *Store) GetSearches() ([]*types.Search, error) {
	queries := db.New(s.db)
	ctx := context.Background()

	searches, err := queries.GetSearches(ctx)
	if err != nil {
		fmt.Println("Erro ao buscar as buscas:", err)
		return nil, err
	}

	var searchs []*types.Search
	for _, search := range searches {
		searchs = append(searchs, &types.Search{
			ID:          search.ID,
			Description: search.Description,
			CreatedAt:   search.CreatedAt,
			UpdatedAt:   search.UpdatedAt,
		})
		searches = append(searches, search)
	}

	return searchs, nil
}

func (s *Store) GetSearchByID(id uuid.UUID) (*types.Search, error) {
	queries := db.New(s.db)
	ctx := context.Background()

	dbsearch, err := queries.GetSearch(ctx, id)
	if err != nil {
		fmt.Println("Erro ao buscar a busca:", err)
		return nil, err
	}

	search := &types.Search{
		ID:          dbsearch.ID,
		Description: dbsearch.Description,
		CreatedAt:   dbsearch.CreatedAt,
		UpdatedAt:   dbsearch.UpdatedAt,
	}
	return search, nil
}

func (s *Store) DeleteSearch(id uuid.UUID) error {
	queries := db.New(s.db)
	ctx := context.Background()

	if err := queries.DeleteSearch(ctx, id); err != nil {
		fmt.Println("Erro ao deletar a busca:", err)
		return err
	}

	return nil
}
