// Package logger provide setting of logger.
package logger

import (
	"fmt"
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// SetupLogger provide setup of log and his levels.
func SetupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug},
			),
		)

	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug},
			),
		)
	case envProd:
		if err := os.MkdirAll("logs", 0o755); err != nil {
			panic(fmt.Errorf("не удалось создать директорию логов: %w", err))
		}

		file, err := os.OpenFile("logs/logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
		if err != nil {
			panic(fmt.Errorf("не удалось открыть файл логов: %w", err))
		}

		logger = slog.New(
			slogmulti.Fanout(
				slog.NewJSONHandler(file, &slog.HandlerOptions{Level: slog.LevelInfo}),
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
			),
		)

	default:
		logger = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo},
			))
	}

	return logger
}
