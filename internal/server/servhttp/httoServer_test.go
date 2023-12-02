package servhttp

import (
	"github.com/GroVlAn/WBTechL0/internal/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

var expectedServerHttp *ServerHttp = &ServerHttp{
	httpServer: &http.Server{
		Addr:              ":8010",
		Handler:           &http.ServeMux{},
		MaxHeaderBytes:    1 << 20,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	},
}

func TestNewHttpServer(t *testing.T) {
	conf := config.ServerConfig{
		Port:              "8010",
		MaxHeaderBytes:    1 << 20,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	hand := http.ServeMux{}

	serv := NewHttpServer(&conf, &hand)

	assert.Equal(t, expectedServerHttp, serv)
}
