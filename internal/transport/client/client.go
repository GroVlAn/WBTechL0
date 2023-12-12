package client

import (
	md "github.com/GroVlAn/WBTechL0/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

type Client struct {
	log *logrus.Logger
}

func NewClient(
	log *logrus.Logger,
) *Client {
	return &Client{
		log: log,
	}
}

// InitBaseMiddlewares Initialization chi middlewares
func (c *Client) initBaseMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
}

func (c *Client) Handler() *chi.Mux {
	router := chi.NewRouter()

	c.baseMiddleware(router)
	c.initBaseMiddlewares(router)

	router.Mount("/order", c.Order())

	return router
}

func (c *Client) baseMiddleware(r *chi.Mux) {
	r.Use(md.SkipFavicon)
}
