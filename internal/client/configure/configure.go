// Package configure предназначен для настройки программы.
package configure

import (
	"flag"
	"net/url"

	"gophkeeper/internal/logger"

	"github.com/caarlos0/env"
	"go.uber.org/zap"
)

// Config хранит текущую конфигурацию сервиса.
type Config struct {
	Address    string `env:"ADDRESS" json:"address,omitempty"`         // адрес сервера grpc
	StaticPath string `env:"STATIC_PATH" json:"static_path,omitempty"` // путь до рабочей дирректории
}

func (cfg *Config) readFlags() {
	address := flag.String("g", "127.0.0.1:8080", "Сетевой адрес grpc host:port")
	staticPath := flag.String("s", "./", "Путь до файлов статики ")
	flag.Parse()

	if cfg.Address == "" {
		cfg.Address = *address
	}
	if cfg.StaticPath == "" {
		cfg.StaticPath = *staticPath
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

// ReadConfig читает конфигурацию клиента
func (cfg *Config) ReadConfig() bool {
	err := cfg.readEnv()
	if err != nil {
		return false
	}

	cfg.readFlags()

	return cfg.checkConfig()
}
