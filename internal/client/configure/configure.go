// Package configure предназначен для настройки программы.
package configure

import (
	"flag"
	"net/url"
	"os"
	"path/filepath"

	"gophkeeper/internal/logger"

	"github.com/caarlos0/env"
	"go.uber.org/zap"
)

// Config хранит текущую конфигурацию сервиса.
type Config struct {
	Address    string `env:"ADDRESS" json:"address,omitempty"`         // адрес сервера grpc
	Secret     string `env:"SECRET" json:"secret,omitempty"`           // ключ шифрования
	StaticPath string `env:"STATIC_PATH" json:"static_path,omitempty"` // путь до рабочей дирректории
}

func (cfg *Config) readFlags() error {
	dirHomeName, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	workDir := filepath.Join(dirHomeName, ".gophkeeper")
	err = os.MkdirAll(workDir, os.ModePerm)
	if err != nil {
		return err
	}

	address := flag.String("g", "127.0.0.1:8080", "Сетевой адрес grpc host:port")
	secret := flag.String("s", "test", "Ключ шифрования")
	staticPath := flag.String("w", workDir, "Путь до рабочей дирректории ")
	flag.Parse()

	if cfg.Address == "" {
		cfg.Address = *address
	}

	if cfg.Secret == "" {
		cfg.Secret = *secret
	}

	if cfg.StaticPath == "" {
		cfg.StaticPath = *staticPath
	}
	return nil
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

	err = cfg.readFlags()
	if err != nil {
		return false
	}

	return cfg.checkConfig()
}
