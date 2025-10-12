package handlers

import (
	"encoding/json"
	"experiment/adapters/controllers"
	"experiment/adapters/presenters"
	"experiment/adapters/presenters/input"
	"net/http"
)

type OwnerHandler struct {
	createOwnerController controllers.CreateOwnerController
	createOwnerPresenter  presenters.CreateOwnerPresenter
}

func NewOwnerHandler(createOwnerController controllers.CreateOwnerController, createOwnerPresenter presenters.CreateOwnerPresenter) *OwnerHandler {
	return &OwnerHandler{
		createOwnerController: createOwnerController,
		createOwnerPresenter:  createOwnerPresenter,
	}
}

func (h *OwnerHandler) CreateOwner(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var ownerInput input.OwnerInput
	if err := json.NewDecoder(r.Body).Decode(&ownerInput); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.createOwnerController.HandleCreateOwner(&ownerInput); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := h.createOwnerPresenter.Present("Owner created successfully")
	json.NewEncoder(w).Encode(response)
}
