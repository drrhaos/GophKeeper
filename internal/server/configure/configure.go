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
	Address     string `env:"ADDRESS" json:"address,omitempty"`           // адрес сервера grpc
	AddressRest string `env:"ADDRESS_REST" json:"address_rest,omitempty"` // адрес сервера rest
	StaticPath  string `env:"STATIC_PATH" json:"static_path,omitempty"`   // путь до статических файлов
	DatabaseDsn string `env:"DATABASE_DSN" json:"database_dsn,omitempty"` // DSN базы данных
	SecretKey   string `env:"SECRET_KEY" json:"secret_key,omitempty"`     // ключ шифрования
}

func (cfg *Config) readFlags() {
	address := flag.String("g", "127.0.0.1:8080", "Сетевой адрес grpc host:port")
	addressRest := flag.String("r", "127.0.0.1:8081", "Сетевой адрес rest host:port")
	staticPath := flag.String("s", "../../swagger-ui/", "Путь до файлов статики ")
	databaseDsn := flag.String("d", "",
		"Сетевой адрес базя данных postgres://postgres:postgres@postgres:5432/praktikum?sslmode=disable")
	secretKey := flag.String("k", "test", "Сетевой адрес host:port")
	flag.Parse()

	if cfg.Address == "" {
		cfg.Address = *address
	}

	if cfg.AddressRest == "" {
		cfg.AddressRest = *addressRest
	}

	if cfg.StaticPath == "" {
		cfg.StaticPath = *staticPath
	}

	if cfg.DatabaseDsn == "" {
		cfg.DatabaseDsn = *databaseDsn
	}

	if cfg.SecretKey == "" {
		cfg.SecretKey = *secretKey
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

	if cfg.AddressRest == "" {
		cfg.AddressRest = "127.0.0.1:8081"
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
