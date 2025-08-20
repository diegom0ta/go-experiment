package router

import (
	"experiment/infra/server"
	"fmt"
	"net/http"
)

type Router struct {
	srv *server.Server
}

func NewRouter(srv *server.Server) *Router {
	return &Router{srv: srv}
}

func (r *Router) Start() {
	mux := http.NewServeMux()
	r.srv.Handler = mux

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})
}
