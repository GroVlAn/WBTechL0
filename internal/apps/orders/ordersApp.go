package ordersApp

import (
	"github.com/GroVlAn/WBTechL0/internal/config"
	"github.com/GroVlAn/WBTechL0/internal/server/http"
	"github.com/GroVlAn/WBTechL0/internal/transport/http/handler"
	"github.com/go-chi/chi/v5"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type OrdersApp struct {
	Runner
}

func (p *OrdersApp) Run() {
	if err := config.InitConfig("configs", "ordersConfig"); err != nil {
		log.Fatalf("error initialisation config: %s", err.Error())
	}

	chiRouter := chi.NewRouter()
	conf := config.NewConfig("dev")
	httpHand := handler.NewHttpHandler(chiRouter)
	httpHand.InitBaseMiddlewares()
	serv := http.NewHttpServer(&conf, httpHand.Handler())

	go func() {
		if err := serv.Start(); err != nil {
			log.Fatalf("error occurred while starting server: %s", err.Error())
		}
	}()

	log.Println("Service orders is started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}

type Runner interface {
	Run()
}
