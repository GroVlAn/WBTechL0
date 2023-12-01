package http

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
func NewHttpServer(conf *config.Config, handler http.Handler) *ServerHttp {
	return &ServerHttp{
		httpServer: &http.Server{
			Addr:              ":" + conf.HttpConfig.Port,
			Handler:           handler,
			MaxHeaderBytes:    conf.ServerConfig.MaxHeaderBytes,
			ReadHeaderTimeout: conf.ServerConfig.ReadHeaderTimeout,
			WriteTimeout:      conf.ServerConfig.WriteTimeout,
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
