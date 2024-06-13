// Package configure предназначен для настройки программы.
package configure

import (
	"flag"
	"gophkeeper/internal/logger"
	"net/url"

	"github.com/caarlos0/env"
	"go.uber.org/zap"
)

// Config хранит текущую конфигурацию сервиса.
type Config struct {
	Address         string `env:"ADDRESS" json:"address,omitempty"`                               // адрес сервера
	DatabaseDsn     string `env:"DATABASE_DSN" json:"database_dsn,omitempty"`                     // DSN базы данных
}

func (cfg *Config) readFlags() {
	address := flag.String("a", "", "Сетевой адрес host:port")
	databaseDsn := flag.String("d", "",
		"Сетевой адрес базя данных postgres://postgres:postgres@postgres:5432/praktikum?sslmode=disable")
	flag.Parse()
	
	if *address != "" {
		cfg.Address = *address
	}

	if *databaseDsn != "" {
		cfg.DatabaseDsn = *databaseDsn
	}
}

func (cfg *Config) readEnv() error {
	var tmpCfg Config
	err := env.Parse(&tmpCfg)
	if err != nil {
		logger.Log.Warn("Не удалось найти переменные окружения", zap.Error(err))
		return err
	}
	return nil
}

func (cfg *Config) checkConfig() bool {
	if cfg.Address == "" {
		cfg.Address = "127.0.0.1:8080"
	}

	_, errURL := url.ParseRequestURI("http://" + cfg.Address)
	if errURL != nil {
		logger.Log.Error("неверный формат адреса")
		return false
	}

	return true
}

// ReadConfig читает конфигурацию сервера
func (cfg *Config) ReadConfig() bool {
	err := cfg.readEnv()
	if err != nil {
		return false
	}

	cfg.readFlags()

	return cfg.checkConfig()
}
