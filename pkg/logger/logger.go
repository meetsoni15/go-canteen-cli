package logger

import (
	"log/slog"
	"os"
)

// New is initialled s logger
func New() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug, ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		// Remove all attributes except the message
		if a.Key == slog.MessageKey {
			return a
		}
		return slog.Attr{}
	}}))
	slog.SetDefault(logger)
	return logger
}
