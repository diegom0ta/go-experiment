package handlers

import (
	"encoding/json"
	"errors"
	"experiment/adapters/controllers"
	"experiment/adapters/presenters"
	"experiment/adapters/presenters/input"
	"net/http"
)

var ErrWalletAlreadyExists = errors.New("wallet already exists")

type WalletHandler struct {
	createWalletController controllers.CreateWalletController
	createWalletPresenter  presenters.CreateWalletPresenter
}

func NewWalletHandler(createWalletController controllers.CreateWalletController, createWalletPresenter presenters.CreateWalletPresenter) *WalletHandler {
	return &WalletHandler{
		createWalletController: createWalletController,
		createWalletPresenter:  createWalletPresenter,
	}
}

func (h *WalletHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var walletInput input.WalletInput
	if err := json.NewDecoder(r.Body).Decode(&walletInput); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	err := h.createWalletController.HandleCreateWallet(ctx, email, &walletInput)
	switch {
	case err != nil && err.Error() == ErrWalletAlreadyExists.Error():
		w.WriteHeader(http.StatusConflict)
		response := h.createWalletPresenter.Present("Wallet already exists")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	case err != nil:
		http.Error(w, "internal server error", http.StatusInternalServerError)
	default:
		w.WriteHeader(http.StatusCreated)
		response := h.createWalletPresenter.Present("Wallet created successfully")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}
}
