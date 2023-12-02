package servhttp

import (
	"context"
	"github.com/GroVlAn/WBTechL0/internal/config"
	"net/http"
)

type ServerHttp struct {
	httpServer *http.Server
}

/*
NewHttpServer
:param conf *config.Config - application config
:param handler http.Handler - some handler

function for create and initialization http server
*/
func NewHttpServer(conf *config.ServerConfig, handler http.Handler) *ServerHttp {
	return &ServerHttp{
		httpServer: &http.Server{
			Addr:              ":" + conf.Port,
			Handler:           handler,
			MaxHeaderBytes:    conf.MaxHeaderBytes,
			ReadHeaderTimeout: conf.ReadHeaderTimeout,
			WriteTimeout:      conf.WriteTimeout,
		},
	}
}

// Start method for start http server
func (s *ServerHttp) Start() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown method for shutdown http server
func (s *ServerHttp) Shutdown(cxt context.Context) error {
	return s.httpServer.Shutdown(cxt)
}
