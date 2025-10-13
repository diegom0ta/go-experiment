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

	ownerHandler := handlers.NewOwnerHandler(createOwnerController, createOwnerPresenter, getOwnerController, getOwnerPresenter)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	mux.HandleFunc("/owner/create", ownerHandler.CreateOwner)

	mux.HandleFunc("/owner", ownerHandler.GetOwnerByEmail)
}

func (r *Router) Start() {
	mux := http.NewServeMux()
	r.srv.Handler = mux
	r.SetupRoutes(mux)
}
