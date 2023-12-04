package rest

import (
	"github.com/GroVlAn/WBTechL0/internal/tools/writeresp"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (hh *HttpHandler) DeliveryHandler() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/delivery", func(r chi.Router) {
		r.Get("/{deliveryID}", hh.Delivery)
		r.Post("/", hh.CreateDelivery)
		r.Delete("/{deliveryID}", hh.DeleteDelivery)
	})

	return router
}

func (hh *HttpHandler) CreateDelivery(w http.ResponseWriter, req *http.Request) {
	writeresp.Write(
		hh.log,
		[]byte("Create delivery"),
		"error can not write delivery response",
		w.Write,
	)
}

func (hh *HttpHandler) Delivery(w http.ResponseWriter, req *http.Request) {
	writeresp.Write(
		hh.log,
		[]byte("Delivery"),
		"error can not write delivery response",
		w.Write,
	)
}

func (hh *HttpHandler) DeleteDelivery(w http.ResponseWriter, req *http.Request) {
	writeresp.Write(
		hh.log,
		[]byte("Delivery"),
		"error can not write delivery response",
		w.Write,
	)
}
