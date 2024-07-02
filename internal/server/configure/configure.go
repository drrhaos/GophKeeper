// Package configure предназначен для настройки программы.
package configure

import (
	"flag"
	"os"

	"gophkeeper/internal/logger"

	"github.com/caarlos0/env"
	"go.uber.org/zap"
)

// Config хранит текущую конфигурацию сервиса.
type Config struct {
	Port        string `env:"PORT" json:"address,omitempty"`              // порт сервера grpc
	PortRest    string `env:"PORT_REST" json:"address_rest,omitempty"`    // порт сервера rest
	StaticPath  string `env:"STATIC_PATH" json:"static_path,omitempty"`   // путь до статических файлов
	WorkPath    string `env:"WORK_PATH" json:"work_path,omitempty"`       // путьдо рабочей дирректории
	DatabaseDsn string `env:"DATABASE_DSN" json:"database_dsn,omitempty"` // DSN базы данных
	SecretKey   string `env:"SECRET_KEY" json:"secret_key,omitempty"`     // ключ шифрования
}

func (cfg *Config) readFlags() {
	port := flag.String("g", "8080", "Сетевой порт grpc")
	portRest := flag.String("r", "8081", "Сетевой порт rest")
	staticPath := flag.String("s", "../../swagger-ui/", "Путь до файлов статики ")
	workPath := flag.String("w", "./data", "Путь до рабочей дирректории")
	databaseDsn := flag.String("d", "",
		"Сетевой адрес базя данных postgres://postgres:postgres@postgres:5432/praktikum?sslmode=disable")
	secretKey := flag.String("k", "test", "Сетевой адрес host:port")
	flag.Parse()

	if cfg.Port == "" {
		cfg.Port = *port
	}

	if cfg.PortRest == "" {
		cfg.PortRest = *portRest
	}

	if cfg.StaticPath == "" {
		cfg.StaticPath = *staticPath
	}

	if cfg.WorkPath == "" {
		cfg.WorkPath = *workPath
	}

	if cfg.DatabaseDsn == "" {
		cfg.DatabaseDsn = *databaseDsn
	}

	if cfg.SecretKey == "" {
		cfg.SecretKey = *secretKey
	}
}

func (cfg *Config) readEnv() error {
	err := env.Parse(cfg)
	if err != nil {
		logger.Log.Warn("Не удалось найти переменные окружения", zap.Error(err))
		return err
	}

	return nil
}

func (cfg *Config) checkConfig() bool {
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	if cfg.PortRest == "" {
		cfg.PortRest = "8081"
	}

	err := os.MkdirAll(cfg.WorkPath, os.ModePerm)
	if err != nil {
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
