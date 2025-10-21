package config

import (
	"fmt"
	"log/slog"
	"os"
)

func New() (string, string) {
	_port := os.Getenv("PORT")
	_adminPort := os.Getenv("ADMIN_PORT")
	if _port == "" {
		_port = "80"
	}
	_port = fmt.Sprintf(
		":%s", _port,
	)
	if _adminPort == "" {
		_adminPort = "4657"
	}
	_adminPort = fmt.Sprintf(
		":%s", _adminPort,
	)
	slog.SetDefault(
		slog.New(
			slog.NewJSONHandler(
				os.Stdout, nil,
			),
		),
	)
	return _port, _adminPort
}
