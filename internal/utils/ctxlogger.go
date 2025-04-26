package utils

import (
	"context"

	"go.uber.org/zap"
)

type ctxLoggerKey struct{}

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxLoggerKey{}, logger)
}

func LoggerFromContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(ctxLoggerKey{}).(*zap.Logger); ok {
		return logger
	}
	return zap.NewNop()
}
