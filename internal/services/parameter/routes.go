package parameter

import (
	"encoding/json"
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
	router.HandleFunc("/update_parameter", h.handleUpdateParameter).Methods(http.MethodPut)
	router.HandleFunc("/get_parameters", h.handleGetParameters).Methods(http.MethodGet)
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

func (h *Handler) handleUpdateParameter(w http.ResponseWriter, r *http.Request) {
	var parameter types.Parameter
	if err := utils.ParseJSON(r, &parameter); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err := h.ParameterStore.UpdateParameter(parameter)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	response := map[string]interface{}{
		"data":    parameter,
		"message": "Registro alterado com sucesso",
		"status":  http.StatusOK,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonResponse)
}

func (h *Handler) handleGetParameters(w http.ResponseWriter, r *http.Request) {
	parameters, err := h.ParameterStore.GetParameters()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter parâmetros: %v", err), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, parameters)
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

	parameter, err := h.ParameterStore.GetParameterByID(parsedParameterID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, parameter)
}
