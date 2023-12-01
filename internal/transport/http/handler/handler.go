package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

type HttpHandler struct {
	router chi.Router
	log    *logrus.Logger
}

func NewHttpHandler(router chi.Router, log *logrus.Logger) *HttpHandler {
	return &HttpHandler{
		router: router,
		log:    log,
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
