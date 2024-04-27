package parameter

import (
	"fmt"
	"net/http"

	"github.com/claudineyveloso/govita.git/internal/types"
	"github.com/claudineyveloso/govita.git/internal/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler struct {
	ParameterStore types.ParameterStore
}

func NewHandler(parameterStore types.ParameterStore) *Handler {
	return &Handler{ParameterStore: parameterStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/create_parameter", h.handleCreateParameter).Methods(http.MethodPost)
}

func (h *Handler) handleCreateParameter(w http.ResponseWriter, r *http.Request) {
	var parameter types.Parameter
	if err := utils.ParseJSON(r, &parameter); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err := h.ParameterStore.CreateParameter(parameter)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, parameter)
}

func (h *Handler) handleGetParameters(w http.ResponseWriter, r *http.Request) {
	users, err := h.ParameterStore.GetParameters()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter parâmetros: %v", err), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) handleGetParameter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["parameterID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID do Parâmetro ausente"))
		return
	}
	parsedParameterID, err := uuid.Parse(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID do Parâmetro inválido"))
		return
	}

	user, err := h.ParameterStore.GetParameterByID(parsedParameterID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, user)
}
