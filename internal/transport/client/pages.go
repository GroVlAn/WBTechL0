package client

import (
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
)

func (c *Client) Order() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Get("/", c.OrderPage)
		r.Get("/find", c.FindOrder)
	})

	return router
}

func (c *Client) OrderPage(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/order/index.html"))

	errEx := tmpl.Execute(w, "")

	if errEx != nil {
		c.log.Errorf("can not execute order Page: %s", errEx.Error())
	}
}

func (c *Client) FindOrder(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/order/find/index.html"))

	errEx := tmpl.Execute(w, "")

	if errEx != nil {
		c.log.Errorf("can not execute order Page: %s", errEx.Error())
	}
}
