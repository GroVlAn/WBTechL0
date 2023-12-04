package rest

import (
	"github.com/GroVlAn/WBTechL0/internal/tools/writeresp"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (hh *HttpHandler) PaymentHandler() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/payment", func(r chi.Router) {
		r.Get("/{paymentID}", hh.Payment)
		r.Post("/", hh.CreatePayment)
		r.Delete("/{paymentID}", hh.CreatePayment)
	})

	return router
}

func (hh *HttpHandler) CreatePayment(w http.ResponseWriter, req *http.Request) {
	writeresp.Write(
		hh.log,
		[]byte("Create payment"),
		"error can not write payment response",
		w.Write,
	)
}

func (hh *HttpHandler) Payment(w http.ResponseWriter, req *http.Request) {
	writeresp.Write(
		hh.log,
		[]byte("Payment"),
		"error can not write payment response",
		w.Write,
	)
}

func (hh *HttpHandler) DeletePayment(w http.ResponseWriter, req *http.Request) {
	writeresp.Write(
		hh.log,
		[]byte("Payment"),
		"error can not write payment response",
		w.Write,
	)
}
