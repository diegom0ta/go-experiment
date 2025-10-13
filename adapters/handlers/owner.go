package handlers

import (
	"encoding/json"
	"errors"
	"experiment/adapters/controllers"
	"experiment/adapters/presenters"
	"experiment/adapters/presenters/input"
	"experiment/adapters/presenters/output"
	"net/http"
)

var ErrorOwnerAlreadyExists = errors.New("owner already exists")

type OwnerHandler struct {
	createOwnerController     controllers.CreateOwnerController
	createOwnerPresenter      presenters.CreateOwnerPresenter
	getOwnerByEmailController controllers.GetOwnerByEmailController
	getOwnerPresenter         presenters.GetOwnerPresenter
}

func NewOwnerHandler(createOwnerController controllers.CreateOwnerController, createOwnerPresenter presenters.CreateOwnerPresenter, getOwnerByEmailController controllers.GetOwnerByEmailController, getOwnerPresenter presenters.GetOwnerPresenter) *OwnerHandler {
	return &OwnerHandler{
		createOwnerController:     createOwnerController,
		createOwnerPresenter:      createOwnerPresenter,
		getOwnerByEmailController: getOwnerByEmailController,
		getOwnerPresenter:         getOwnerPresenter,
	}
}

func (h *OwnerHandler) CreateOwner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var ownerInput input.OwnerInput
	if err := json.NewDecoder(r.Body).Decode(&ownerInput); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.createOwnerController.HandleCreateOwner(ctx, &ownerInput); err.Error() == ErrorOwnerAlreadyExists.Error() {
		w.WriteHeader(http.StatusConflict)
		response := h.createOwnerPresenter.Present("Owner already exists")
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusCreated)
		response := h.createOwnerPresenter.Present("Owner created successfully")
		json.NewEncoder(w).Encode(response)
	}
}

func (h *OwnerHandler) GetOwnerByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	owner, err := h.getOwnerByEmailController.HandleGetOwnerByEmail(ctx, email)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var wallets []output.WalletOutput
	for _, w := range owner.Wallets {
		wallets = append(wallets, output.WalletOutput{
			ID:      w.ID,
			Name:    w.WalletName,
			Balance: w.Balance,
		})
	}

	o := output.GetOwnerOutput{
		ID:       owner.ID,
		Name:     owner.OwnerName,
		Email:    owner.Email,
		Document: owner.Document,
		Wallets:  wallets,
	}

	w.WriteHeader(http.StatusOK)

	response := h.getOwnerPresenter.Present(&o)
	json.NewEncoder(w).Encode(response)
}
