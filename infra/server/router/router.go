package router

import (
	"experiment/adapters/controllers"
	"experiment/adapters/handlers"
	"experiment/adapters/presenters"
	"experiment/infra/server"
	"experiment/repository"
	"experiment/services"
	"experiment/usecases"
	"fmt"
	"net/http"
)

type Router struct {
	srv *server.Server
}

func NewRouter(srv *server.Server) *Router {
	return &Router{srv: srv}
}

func (r *Router) SetupRoutes(mux *http.ServeMux) {
	ownerRepo := repository.NewOwnerRepository()
	ownerCache := services.NewOwnerCache()

	createOwnerUC := usecases.NewCreateOwnerUseCase(ownerRepo, ownerCache)
	createOwnerController := controllers.NewCreateOwnerController(createOwnerUC)
	createOwnerPresenter := presenters.NewCreateOwnerPresenter()

	getOwnerByEmailUC := usecases.NewGetOwnerByEmailUseCase(ownerRepo, ownerCache)
	getOwnerController := controllers.NewGetOwnerByEmailController(getOwnerByEmailUC)
	getOwnerPresenter := presenters.NewGetOwnerPresenter()

	walletRepo := repository.NewWalletRepository()
	walletCache := services.NewWalletCache()

	createWalletUC := usecases.NewCreateWalletUseCase(walletRepo, walletCache, getOwnerByEmailUC, ownerCache)
	createWalletController := controllers.NewCreateWalletController(createWalletUC)
	createWalletPresenter := presenters.NewCreateWalletPresenter()

	walletHandler := handlers.NewWalletHandler(createWalletController, createWalletPresenter)

	ownerHandler := handlers.NewOwnerHandler(createOwnerController, createOwnerPresenter, getOwnerController, getOwnerPresenter)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "OK")
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/owner/create", ownerHandler.CreateOwner)
	mux.HandleFunc("/owner", ownerHandler.GetOwnerByEmail)
	mux.HandleFunc("/wallet/create", walletHandler.CreateWallet)
}

func (r *Router) Start() {
	mux := http.NewServeMux()
	r.srv.Handler = mux
	r.SetupRoutes(mux)
}
