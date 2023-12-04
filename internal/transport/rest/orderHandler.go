package rest

import (
	"github.com/GroVlAn/WBTechL0/internal/tools/writeresp"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (hh *HttpHandler) OrderHandler() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/order", func(r chi.Router) {
		r.Get("/{orderID}", hh.Order)
		r.Post("/", hh.CreateOrder)
		r.Delete("/{orderID}", hh.DeleteOrder)
	})

	return router
}

func (hh *HttpHandler) CreateOrder(w http.ResponseWriter, req *http.Request) {
	writeresp.Write(
		hh.log,
		[]byte("Create order"),
		"error can not write order response",
		w.Write,
	)
}

func (hh *HttpHandler) Order(w http.ResponseWriter, req *http.Request) {
	writeresp.Write(
		hh.log,
		[]byte("Order"),
		"error can not write order response",
		w.Write,
	)
}

func (hh *HttpHandler) DeleteOrder(w http.ResponseWriter, req *http.Request) {
	writeresp.Write(
		hh.log,
		[]byte("Order"),
		"error can not write order response",
		w.Write,
	)
}
