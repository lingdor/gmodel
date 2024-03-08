package mylog

import (
	"context"
	"log/slog"

	"github.com/go-logr/logr"
)

func GetLogger(ctx context.Context) *slog.Logger {

	logger := logr.FromContextAsSlogLogger(ctx)
	if logger == nil {
		logger = slog.Default()
	}
	return logger
}
