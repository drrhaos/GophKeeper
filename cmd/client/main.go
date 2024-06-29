package main

import (
	"gophkeeper/internal/client/configure"
	"gophkeeper/internal/client/tuiclient"
	"gophkeeper/internal/logger"
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

	formUI := tuiclient.Form{}

	formUI.NewForm(cfg)
}
