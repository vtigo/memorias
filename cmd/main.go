package main

import (
	"log/slog"

	"github.com/vtigo/memorias/internals/adapters/http"
	// "github.com/vtigo/memorias/internals/adapters/storaging"
	"github.com/vtigo/memorias/internals/infra"
	// "github.com/vtigo/memorias/internals/ports"
)

func main() {
	config, err := infra.LoadConfig()
	if err != nil {
		panic(err)
	}

	infra.SetupLogger(config.LogLevel, nil)
	slog.Info("starting memories...")

	// slog.Info("setting up storager")
	// var storager ports.FileStorager
	// if config.S3Bucket != nil {
	// 	// TODO: set up s3 storager
	// 	storager = storaging.NewLocalStorager("data")
	// } else {
	// 	storager = storaging.NewLocalStorager("data")
	// }
	
	addr := http.BuildAddress(config.BaseURL, config.Port)
	err = http.StartServer(addr)
	if err != nil {
		panic(err)
	}
}
