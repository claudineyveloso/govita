package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/claudineyveloso/govita.git/internal/services/healthy"
	"github.com/claudineyveloso/govita.git/internal/services/parameter"
	"github.com/claudineyveloso/govita.git/internal/services/search"
	"github.com/claudineyveloso/govita.git/internal/services/user"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	r := mux.NewRouter()
	healthy.RegisterRoutes(r)

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(r)

	parameterStore := parameter.NewStore(s.db)
	parameterHandler := parameter.NewHandler(parameterStore)
	parameterHandler.RegisterRoutes(r)

	searchStore := search.NewStore(s.db)
	searchHandler := search.NewHandler(searchStore)
	searchHandler.RegisterRoutes(r)

	fmt.Println("Server started on http://localhost:8080")

	return http.ListenAndServe("localhost:8080",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		)(r))
}
