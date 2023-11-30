package config

import (
	"fmt"
	"github.com/spf13/viper"
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
	HttpConfig
	ServerConfig
}

const (
	maxHeaderBytes    = 1 << 20
	readHeaderTimeout = 10 * time.Second
	writeTimeout      = 10 * time.Second
)

func NewConfig(mode string) Config {
	return Config{
		HttpConfig{
			Port: viper.GetString(fmt.Sprintf("%s.http.port", mode)),
		},
		ServerConfig{
			MaxHeaderBytes:    maxHeaderBytes,
			ReadHeaderTimeout: readHeaderTimeout,
			WriteTimeout:      writeTimeout,
		},
	}
}

func InitConfig(path string, nameConfig string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(nameConfig)

	return viper.ReadInConfig()
}
