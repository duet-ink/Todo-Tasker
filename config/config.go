package config

import (
	"fmt"
	"log/slog"
	"os"
)

func New() string {
	_port := os.Getenv("PORT")
	if _port == "" {
		_port = "80"
	}
	_port = fmt.Sprintf(
		":%s", _port,
	)
	slog.SetDefault(
		slog.New(
			slog.NewJSONHandler(
				os.Stdout, nil,
			),
		),
	)
	return _port
}
