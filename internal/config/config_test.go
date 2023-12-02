package config

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

const (
	pathConfig = "../../configs"
	nameConfig = "ordersConfig"
)

func TestNewConfig(t *testing.T) {
	if err := InitEnv("../../.env"); err != nil {
		log.Fatalf("error initializing env: %s", err.Error())
	}

	if err := InitConfig(pathConfig, nameConfig); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	devConf := NewConfig("dev")

	assert.IsType(t, &Config{}, &devConf)

	assert.Equal(t, "7010", devConf.ServerConfig.Port)
	assert.Equal(t, 1<<20, devConf.ServerConfig.MaxHeaderBytes)
	assert.Equal(t, 10*time.Second, devConf.ServerConfig.ReadHeaderTimeout)
	assert.Equal(t, 10*time.Second, devConf.ServerConfig.WriteTimeout)

	assert.Equal(t, "localhost", devConf.PostgresConfig.Host)
	assert.Equal(t, "5432", devConf.PostgresConfig.Port)
	assert.Equal(t, "root", devConf.PostgresConfig.Username)
	assert.Equal(t, "wb_orders", devConf.PostgresConfig.DBName)
	assert.Equal(t, "disable", devConf.PostgresConfig.SSLMode)
	assert.False(t, len(devConf.PostgresConfig.Password) == 0)

	prodConf := NewConfig("prod")

	assert.Equal(t, "8010", prodConf.ServerConfig.Port)
	assert.Equal(t, "wb_orders_db", prodConf.PostgresConfig.Host)
	assert.Equal(t, "5432", prodConf.PostgresConfig.Port)
	assert.Equal(t, "root", prodConf.PostgresConfig.Username)
	assert.Equal(t, "wb_orders", prodConf.PostgresConfig.DBName)
	assert.Equal(t, "disable", prodConf.PostgresConfig.SSLMode)
	assert.False(t, len(prodConf.PostgresConfig.Password) == 0)
}
