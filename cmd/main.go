package main

import (
	"log/slog"

	"github.com/vtigo/memorias/internals/adapters/http"
	"github.com/vtigo/memorias/internals/infra"
)

func main() {
	infra.SetupLogger(nil)
	slog.Info("starting memories...")
	
	addr := "0.0.0.0:3333"

	slog.Info("listening...")
	err := http.StartServer(addr)
	if err != nil {
		panic(err)
	}
}
