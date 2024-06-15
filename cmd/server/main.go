package main

import (
	"gophkeeper/internal/logger"
	"gophkeeper/internal/server/configure"
	"gophkeeper/internal/server/grpcmode"
	"gophkeeper/internal/server/restmode"
)

const flagLogLevel = "info"

func main() {
	err := logger.Initialize(flagLogLevel)
	if err != nil {
		panic(err)
	}

	var cfg configure.Config
	ok := cfg.ReadConfig()
	if !ok {
		logger.Log.Panic("Error read config")
	}

	go grpcmode.Run(cfg)
	restmode.Run(cfg)
}
