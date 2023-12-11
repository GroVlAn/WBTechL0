package datagenapp

import (
	"github.com/GroVlAn/WBTechL0/internal/config"
	"github.com/GroVlAn/WBTechL0/internal/server/servhttp"
	"github.com/GroVlAn/WBTechL0/internal/tools/logwrap"
	"github.com/GroVlAn/WBTechL0/internal/transport/natspub"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	pathConfig = "configs"
	nameConfig = "dataGeneratorConfig"
)

const (
	logFile    = "logs.txt"
	permission = 0644
)

type DataGeneratorApp struct {
	Runner
}

func (dga *DataGeneratorApp) initLogger() (*logwrap.Logger, *logrus.Logger) {
	logger := logwrap.NewLogger(logFile, permission)

	if err := logger.InitLogger(); err != nil {
		logrus.Fatalf(err.Error())
	}

	return logger, logger.Log
}

func (dga *DataGeneratorApp) initConfig(mode string) config.Config {
	if err := config.InitEnv(); err != nil {
		log.Fatalf("error initializing env: %s", err.Error())
	}

	if err := config.InitConfig(pathConfig, nameConfig); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	return config.NewConfig(mode)
}

func (dga *DataGeneratorApp) Run(mode string) {
	logger, logApp := dga.initLogger()

	defer func() {
		if err := logger.File.Close(); err != nil {
			logrus.Fatalf("error while closing file: %s", err.Error())
		}
	}()

	conf := dga.initConfig(mode)

	pub := natspub.NewPublish(conf, logApp)

	pub.Run()

	serv := servhttp.NewHttpServer(&conf.ServerConfig, http.NewServeMux())

	go func() {
		if err := serv.Start(); err != nil {
			logApp.Fatalf("error occurred while starting server: %s", err.Error())
		}
	}()

	logApp.Infoln("Service data generator is started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}

type Runner interface {
	Run(mode string)
}
