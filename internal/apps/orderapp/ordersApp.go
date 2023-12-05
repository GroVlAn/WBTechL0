package orderapp

import (
	"github.com/GroVlAn/WBTechL0/internal/config"
	"github.com/GroVlAn/WBTechL0/internal/database/postgres"
	"github.com/GroVlAn/WBTechL0/internal/repository/postgresrepos"
	"github.com/GroVlAn/WBTechL0/internal/server/servhttp"
	"github.com/GroVlAn/WBTechL0/internal/service"
	"github.com/GroVlAn/WBTechL0/internal/tools/logwrap"
	"github.com/GroVlAn/WBTechL0/internal/transport/rest"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

type OrdersApp struct {
	Runner
}

const (
	pathConfig = "configs"
	nameConfig = "ordersConfig"
)

const (
	logFile    = "logs.txt"
	permission = 0644
)

func (p *OrdersApp) Run(mode string) {
	logger := logwrap.NewLogger(logFile, permission)
	defer func() {
		if err := logger.File.Close(); err != nil {
			logrus.Fatalf("error while closing file: %s", err.Error())
		}
	}()

	if err := logger.InitLogger(); err != nil {
		logrus.Fatalf(err.Error())
	}

	log := logger.Log

	if err := config.InitEnv(); err != nil {
		log.Fatalf("error initializing env: %s", err.Error())
	}

	if err := config.InitConfig(pathConfig, nameConfig); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}
	conf := config.NewConfig(mode)

	db, err := postgres.NewPostgresqlDB(conf.PostgresConfig)

	if err != nil {
		log.Fatalf("DB error: %s", err.Error())
	}

	repo := postgresrepos.NewPostgresRepos(db)
	ser := service.NewService(
		repo.ProductRepository,
		repo.PaymentRepository,
		repo.DeliveryRepository,
		repo.OrderRepository,
	)

	httpHand := rest.NewHttpHandler(log, ser)
	serv := servhttp.NewHttpServer(&conf.ServerConfig, httpHand.Handler())

	go func() {
		if err := serv.Start(); err != nil {
			log.Fatalf("error occurred while starting server: %s", err.Error())
		}
	}()

	log.Infoln("Service orders is started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}

type Runner interface {
	Run(mode string)
}
