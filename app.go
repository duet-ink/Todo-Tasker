package main

import (
	"Todo-Tasker/config"
	"Todo-Tasker/server"
	"log/slog"
	"net/http"
)

type serverType struct{
	Addr string
	Handler *http.ServeMux
}

func main() {
	_port, _adminPort := config.New()

	go func(_port string) {
		_mux := server.NewAdmin()
		_server := serverType{
			Addr:    _port,
			Handler: _mux,
		}.log()

		if err := _server.ListenAndServe(); err != nil {
			slog.Error(err.Error())
		}
	}(_adminPort)

	_mux := server.New()
	_server := serverType{
		Addr:    _port,
		Handler: _mux,
	}.log()

	if err := _server.ListenAndServe(); err != nil {
		slog.Error(err.Error())
	}
}

func (_server serverType) log() http.Server {
	slog.Info(
		"Starting server...",
		"url",
		"http://localhost"+_server.Addr,
	)
	return http.Server{
		Addr: _server.Addr,
		Handler: _server.Handler,
	}
}
