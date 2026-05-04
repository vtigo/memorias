package http

import (
	"fmt"
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
	
	slog.Info(fmt.Sprintf("listening at %s", addr))
	return http.ListenAndServe(addr, nil)
}

func BuildAddress(baseURL string, port int) string {
	return fmt.Sprintf("%s:%v", baseURL, port)
}
