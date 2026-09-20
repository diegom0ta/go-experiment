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

var ErrWalletAlreadyExists = errors.New("wallet already exists")

type WalletHandler struct {
	createWalletController      controllers.CreateWalletController
	createWalletPresenter       presenters.CreateWalletPresenter
	getWalletsByOwnerController controllers.GetWalletsByOwnerController
	getWalletsByOwnerPresenter  presenters.GetWalletsByOwnerPresenter
}

func NewWalletHandler(createWalletController controllers.CreateWalletController,
	createWalletPresenter presenters.CreateWalletPresenter,
	getWalletsByOwnerController controllers.GetWalletsByOwnerController,
	getWalletsByOwnerPresenter presenters.GetWalletsByOwnerPresenter) *WalletHandler {
	return &WalletHandler{
		createWalletController:      createWalletController,
		createWalletPresenter:       createWalletPresenter,
		getWalletsByOwnerController: getWalletsByOwnerController,
		getWalletsByOwnerPresenter:  getWalletsByOwnerPresenter,
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

func (h *WalletHandler) GetOwnerWallets(w http.ResponseWriter, r *http.Request) {
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

	wallets, err := h.getWalletsByOwnerController.HandleGetWalletsByOwner(ctx, email)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var walletsOutput []output.WalletOutput
	for _, wallet := range wallets {
		walletsOutput = append(walletsOutput, output.WalletOutput{
			ID:      wallet.ID,
			Name:    wallet.WalletName,
			Balance: wallet.Balance,
		})
	}

	response := h.getWalletsByOwnerPresenter.Present(walletsOutput)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
