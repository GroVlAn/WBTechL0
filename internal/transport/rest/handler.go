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
	dServ    service.DeliveryService
	pmtServ  service.PaymentService
	orServ   service.OrderService
}

func NewHttpHandler(
	log *logrus.Logger,
	prodServ service.ProductService,
	dServ service.DeliveryService,
	pmtServ service.PaymentService,
	orServ service.OrderService,
) *HttpHandler {
	return &HttpHandler{
		log:      log,
		prodServ: prodServ,
		dServ:    dServ,
		pmtServ:  pmtServ,
		orServ:   orServ,
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

	r.Mount(basePath+"/product", hh.ProductHandler())
	r.Mount(basePath+"/delivery", hh.DeliveryHandler())
	r.Mount(basePath+"/payment", hh.PaymentHandler())
	r.Mount(basePath+"/order", hh.OrderHandler())

	return r
}

func (hh *HttpHandler) baseMiddleware(r *chi.Mux) {
	r.Use(md.SkipFavicon)
	r.Use(md.Cors)
}
