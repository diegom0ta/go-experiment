package router

import (
	"experiment/adapters/controllers"
	"experiment/adapters/handlers"
	"experiment/adapters/presenters"
	"experiment/infra/server"
	"experiment/repository"
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
	createOwnerUC := usecases.NewCreateOwnerUseCase(ownerRepo)
	createOwnerController := controllers.NewCreateOwnerController(createOwnerUC)
	createOwnerPresenter := presenters.NewCreateOwnerPresenter()
	ownerHandler := handlers.NewOwnerHandler(createOwnerController, createOwnerPresenter)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	mux.HandleFunc("/owner/create", ownerHandler.CreateOwner)
}

func (r *Router) Start() {
	mux := http.NewServeMux()
	r.srv.Handler = mux
	r.SetupRoutes(mux)
}
