package rest

import (
	md "github.com/GroVlAn/WBTechL0/internal/middleware"
	"github.com/GroVlAn/WBTechL0/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

const (
	basePath = "/api"
)

type HttpHandler struct {
	log      *logrus.Logger
	prodServ service.ProductService
}

func NewHttpHandler(
	log *logrus.Logger,
	prodServ service.ProductService,
) *HttpHandler {
	return &HttpHandler{
		log:      log,
		prodServ: prodServ,
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

	hh.baseMiddleware(r)
	hh.initBaseMiddlewares(r)

	r.Mount(basePath, hh.ProductHandler())

	return r
}

func (hh *HttpHandler) baseMiddleware(r *chi.Mux) {
	r.Use(md.SkipFavicon)
}
