package utils

import (
	"go.uber.org/zap"
)

func NewLoggerWithoutStacktrace() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.DisableStacktrace = true         // 不记录 stacktrace
	cfg.EncoderConfig.StacktraceKey = "" // 移除 stacktrace 字段
	return cfg.Build()
}
