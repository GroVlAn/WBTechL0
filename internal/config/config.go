package config

import (
	"fmt"
	"github.com/GroVlAn/WBTechL0/internal/database/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

type HttpConfig struct {
	Port string
}

type ServerConfig struct {
	MaxHeaderBytes    int
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
}

type Config struct {
	HttpConfig     HttpConfig
	ServerConfig   ServerConfig
	PostgresConfig postgres.Config
}

const (
	maxHeaderBytes    = 1 << 20
	readHeaderTimeout = 10 * time.Second
	writeTimeout      = 10 * time.Second
)

func NewConfig(mode string) Config {
	password, ok := os.LookupEnv("DB_PASSWORD")

	if !ok {
		logrus.Fatal("Can not find db password in .env file")
	}

	return Config{
		HttpConfig: HttpConfig{
			Port: viper.GetString(fmt.Sprintf("%s.http.port", mode)),
		},
		ServerConfig: ServerConfig{
			MaxHeaderBytes:    maxHeaderBytes,
			ReadHeaderTimeout: readHeaderTimeout,
			WriteTimeout:      writeTimeout,
		},
		PostgresConfig: postgres.Config{
			Host:     viper.GetString(fmt.Sprintf("%s.db.postgres.host", mode)),
			Port:     viper.GetString(fmt.Sprintf("%s.db.postgres.port", mode)),
			Username: viper.GetString(fmt.Sprintf("%s.db.postgres.username", mode)),
			Password: password,
			DBName:   viper.GetString(fmt.Sprintf("%s.db.postgres.db_name", mode)),
			SSLMode:  viper.GetString(fmt.Sprintf("%s.db.postgres.sslmode", mode)),
		},
	}
}

func InitConfig(path string, nameConfig string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(nameConfig)

	return viper.ReadInConfig()
}

func InitEnv() error {
	return godotenv.Load()
}
