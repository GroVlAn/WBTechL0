package orderapp

import (
	"github.com/GroVlAn/WBTechL0/internal/config"
	"github.com/GroVlAn/WBTechL0/internal/database/postgres"
	"github.com/GroVlAn/WBTechL0/internal/repository/postgresrepos"
	"github.com/GroVlAn/WBTechL0/internal/server/servhttp"
	"github.com/GroVlAn/WBTechL0/internal/service"
	"github.com/GroVlAn/WBTechL0/internal/tools/logwrap"
	"github.com/GroVlAn/WBTechL0/internal/transport/rest"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
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

func (p *OrdersApp) initLogger() (*logwrap.Logger, *logrus.Logger) {
	logger := logwrap.NewLogger(logFile, permission)

	if err := logger.InitLogger(); err != nil {
		logrus.Fatalf(err.Error())
	}

	return logger, logger.Log
}

func (p *OrdersApp) initConfig(mode string) config.Config {
	if err := config.InitEnv(); err != nil {
		log.Fatalf("error initializing env: %s", err.Error())
	}

	if err := config.InitConfig(pathConfig, nameConfig); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	return config.NewConfig(mode)
}

func (p *OrdersApp) initDB(conf *config.Config) *sqlx.DB {
	db, err := postgres.NewPostgresqlDB(conf.PostgresConfig)

	if err != nil {
		log.Fatalf("DB error: %s", err.Error())
	}

	return db
}

func (p *OrdersApp) initRSH(log *logrus.Logger, db *sqlx.DB) *rest.HttpHandler {
	repo := postgresrepos.NewPostgresRepos(log, db)
	ser := service.NewService(
		log,
		repo.ProductRepository,
		repo.PaymentRepository,
		repo.DeliveryRepository,
		repo.OrderRepository,
	)

	return rest.NewHttpHandler(
		log,
		ser.ProductService,
		ser.DeliveryService,
		ser.PaymentService,
		ser.OrderService,
	)
}

func (p *OrdersApp) Run(mode string) {
	logger, log := p.initLogger()
	defer func() {
		if err := logger.File.Close(); err != nil {
			logrus.Fatalf("error while closing file: %s", err.Error())
		}
	}()

	conf := p.initConfig(mode)
	db := p.initDB(&conf)
	httpHand := p.initRSH(log, db)

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
