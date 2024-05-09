package search

import (
	"fmt"
	"net/http"

	"github.com/claudineyveloso/govita.git/internal/types"
	"github.com/claudineyveloso/govita.git/internal/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler struct {
	SearchStore types.SearchStore
}

func NewHandler(searchStore types.SearchStore) *Handler {
	return &Handler{SearchStore: searchStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/create_search", h.handleCreateSearch).Methods(http.MethodPost)
	router.HandleFunc("/get_searches", h.handleGetSearches).Methods(http.MethodGet)
	router.HandleFunc("/delete_search/{searchID}", h.handleDeleteSearch).Methods(http.MethodDelete)
}

func (h *Handler) handleCreateSearch(w http.ResponseWriter, r *http.Request) {
	var search types.SearchPayload
	if err := utils.ParseJSON(r, &search); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.SearchStore.CreateSearch(search); err != nil {
		http.Error(w, "Erro ao criar a busca", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, search)
}

func (h *Handler) handleGetSearches(w http.ResponseWriter, r *http.Request) {
	searches, err := h.SearchStore.GetSearches()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter parâmetros: %v", err), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, searches)
}

func (h *Handler) handleDeleteSearch(w http.ResponseWriter, r *http.Request) {
	search := mux.Vars(r)
	str, ok := search["searchID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID da Busca ausente"))
		return
	}
	parsedSearchID, err := uuid.Parse(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID da Busca inválida"))
		return
	}

	if err := h.SearchStore.DeleteSearch(parsedSearchID); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}
