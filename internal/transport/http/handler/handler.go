package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HttpHandler struct {
	router chi.Router
}

func NewHttpHandler(router chi.Router) *HttpHandler {
	return &HttpHandler{
		router: router,
	}
}

// InitBaseMiddlewares Initialization chi middlewares
func (hh *HttpHandler) InitBaseMiddlewares() {
	hh.router.Use(middleware.RequestID)
	hh.router.Use(middleware.RealIP)
	hh.router.Use(middleware.Logger)
	hh.router.Use(middleware.Recoverer)
}

// Handler function for create routs and return chi router
func (hh *HttpHandler) Handler() chi.Router {

	return hh.router
}
