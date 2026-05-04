package main

import (
	"log/slog"

	"github.com/vtigo/memorias/internals/adapters/http"
	"github.com/vtigo/memorias/internals/infra"
)

func main() {
	config, err := infra.LoadConfig()
	if err != nil {
		panic(err)
	}

	infra.SetupLogger(config.LogLevel, nil)
	slog.Info("starting memories...")
	
	addr := http.BuildAddress(config.BaseURL, config.Port)
	err = http.StartServer(addr)
	if err != nil {
		panic(err)
	}
}
