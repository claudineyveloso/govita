package result

import (
	"fmt"
	"net/http"

	"github.com/claudineyveloso/govita.git/internal/types"
	"github.com/claudineyveloso/govita.git/internal/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler struct {
	ResultStore types.ResultStore
}

func NewHandler(resultStore types.ResultStore) *Handler {
	return &Handler{ResultStore: resultStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/create_result", h.handleCreateResult).Methods(http.MethodPost)
	router.HandleFunc("/get_results", h.handleGetResults).Methods(http.MethodGet)
	router.HandleFunc("/get_result", h.handleGetResult).Methods(http.MethodGet)
	router.HandleFunc("/delete_result/{resultID}", h.handleDeleteResult).Methods(http.MethodDelete)
}

func (h *Handler) handleCreateResult(w http.ResponseWriter, r *http.Request) {
	var result types.Result
	if err := utils.ParseJSON(r, &result); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err := h.ResultStore.CreateResult(result)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, result)
}

func (h *Handler) handleGetResults(w http.ResponseWriter, r *http.Request) {
	results, err := h.ResultStore.GetResults()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter resultado da busca: %v", err), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, results)
}

func (h *Handler) handleGetResult(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["resultID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID do Resultado ausente"))
		return
	}
	parsedResultID, err := uuid.Parse(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID do Resultado inválido"))
		return
	}
	result, err := h.ResultStore.GetResultByID(parsedResultID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, result)
}

func (h *Handler) handleDeleteResult(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["resultID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID do Resultado ausente"))
		return
	}
	parsedResultID, err := uuid.Parse(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID do Resultado inválido"))
		return
	}
	err = h.ResultStore.DeleteResult(parsedResultID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, nil)
}
