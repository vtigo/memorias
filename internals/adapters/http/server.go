package http

import (
	"io"
	"log/slog"
	"net/http"
)

func StartServer(addr string) error {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
		slog.Info("request", "method", req.Method, "path", req.URL.Path)
	}

	http.HandleFunc("/hello", helloHandler)

	return http.ListenAndServe(addr, nil)
}
