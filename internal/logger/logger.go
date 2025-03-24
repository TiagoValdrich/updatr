package logger

import (
	"log"

	"go.uber.org/zap"
)

type ZapLoggerParams struct {
	ZapLogger *zap.SugaredLogger
}

type StdLoggerParams struct {
	Logger *log.Logger
}

type Logger interface {
	Info(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
}

func NewZapLogger(params ZapLoggerParams) Logger {
	return &ZapLogger{
		logger: params.ZapLogger,
	}
}
