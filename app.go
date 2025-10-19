package main

import (
	"Todo-Tasker/config"
	"Todo-Tasker/server"
	"log/slog"
	"net/http"
)

func main() {
	_mux := server.New()
	_port := config.New()
	_server := http.Server{
		Addr:    _port,
		Handler: _mux,
	}
	slog.Info(
		"Starting server...",
		"url",
		"http://localhost"+_server.Addr,
	)
	if err := _server.ListenAndServe(); err != nil {
		slog.Error(err.Error())
	}
}
