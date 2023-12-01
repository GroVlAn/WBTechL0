package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

type HttpHandler struct {
	log *logrus.Logger
}

func NewHttpHandler(log *logrus.Logger) *HttpHandler {
	return &HttpHandler{
		log: log,
	}
}

// InitBaseMiddlewares Initialization chi middlewares
func (hh *HttpHandler) initBaseMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
}

// Handler function for create routs and return chi router
func (hh *HttpHandler) Handler() *chi.Mux {
	r := chi.NewRouter()

	hh.initBaseMiddlewares(r)

	return r
}
