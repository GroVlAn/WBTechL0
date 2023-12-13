package clientapp

import (
	"github.com/GroVlAn/WBTechL0/internal/config"
	"github.com/GroVlAn/WBTechL0/internal/server/servhttp"
	"github.com/GroVlAn/WBTechL0/internal/tools/logwrap"
	"github.com/GroVlAn/WBTechL0/internal/transport/client"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	pathConfig = "configs"
	nameConfig = "clientConfig"
)

const (
	logFile    = "logs.txt"
	permission = 0644
)

type ClientApp struct {
}

// initLogger init custom logger
func (ca *ClientApp) initLogger() (*logwrap.Logger, *logrus.Logger) {
	logger := logwrap.NewLogger(logFile, permission)

	if err := logger.InitLogger(); err != nil {
		logrus.Fatalf(err.Error())
	}

	return logger, logger.Log
}

// initConfig method for initializing app's config
func (ca *ClientApp) initConfig(mode string) config.Config {
	if err := config.InitEnv(); err != nil {
		log.Fatalf("error initializing env: %s", err.Error())
	}

	if err := config.InitConfig(pathConfig, nameConfig); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	return config.NewConfig(mode)
}

func (ca *ClientApp) Run(mode string) {
	logger, logApp := ca.initLogger()

	defer func() {
		if err := logger.File.Close(); err != nil {
			logrus.Fatalf("error while closing file: %s", err.Error())
		}
	}()

	conf := ca.initConfig(mode)

	cl := client.NewClient(logApp)

	serv := servhttp.NewHttpServer(&conf.ServerConfig, cl.Handler())

	go func() {
		if err := serv.Start(); err != nil {
			logApp.Fatalf("error occurred while starting server: %s", err.Error())
		}
	}()

	logApp.Infoln("Service client is started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}
