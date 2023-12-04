package rest

import (
	"github.com/GroVlAn/WBTechL0/internal/tools/writeresp"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (hh *HttpHandler) ProductHandler() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/product", func(r chi.Router) {
		r.Get("/{productID}", hh.Product)
		r.Post("/", hh.CreateProduct)
		r.Delete("/{productID}", hh.DeleteProduct)
	})

	return router
}

func (hh *HttpHandler) CreateProduct(w http.ResponseWriter, req *http.Request) {
	writeresp.Write(
		hh.log,
		[]byte("Create product"),
		"error can not write product response",
		w.Write,
	)
}

func (hh *HttpHandler) Product(w http.ResponseWriter, req *http.Request) {
	writeresp.Write(
		hh.log,
		[]byte("Product"),
		"error can not write product response",
		w.Write,
	)
}

func (hh *HttpHandler) DeleteProduct(w http.ResponseWriter, req *http.Request) {
	writeresp.Write(
		hh.log,
		[]byte("Product"),
		"error can not write product response",
		w.Write,
	)
}
