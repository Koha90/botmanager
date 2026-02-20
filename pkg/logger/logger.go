// Package logger provides application logger setup.
package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Logger struct {
	*slog.Logger
	closer io.Closer
}

// Close gracefully closes underlying resources (if any).
func (l *Logger) Close() error {
	if l.closer != nil {
		return l.closer.Close()
	}
	return nil
}

// Setup initializes logger, sets it as default and returns wrapper.
func Setup(env string) (*Logger, error) {
	var (
		handler slog.Handler
		closer  io.Closer
	)

	switch env {
	case envLocal:
		handler = slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelDebug,
			},
		)

	case envDev:
		handler = slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelDebug,
			},
		)

	case envProd:
		if err := os.MkdirAll("logs", 0o755); err != nil {
			return nil, fmt.Errorf("cannot create logs directory: %w", err)
		}

		file, err := os.OpenFile("logs/logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
		if err != nil {
			return nil, fmt.Errorf("cannot open log file: %w", err)
		}

		fileHandler := slog.NewJSONHandler(
			file,
			&slog.HandlerOptions{
				Level: slog.LevelInfo,
			},
		)

		stdoutHandler := slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelInfo,
			},
		)

		handler = slog.NewMultiHandler(fileHandler, stdoutHandler)
		closer = file

	default:
		handler = slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo},
		)
	}

	base := slog.New(handler)

	slog.SetDefault(base)

	return &Logger{
		Logger: base,
		closer: closer,
	}, nil
}
