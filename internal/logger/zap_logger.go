package logger

import "go.uber.org/zap"

type ZapLogger struct {
	logger *zap.SugaredLogger
}

// Error implements Logger.
func (z *ZapLogger) Error(msg string, fields ...interface{}) {
	z.logger.Errorf(msg, fields...)
}

// Info implements Logger.
func (z *ZapLogger) Info(msg string, fields ...interface{}) {
	z.logger.Infof(msg, fields...)
}

// Warn implements Logger.
func (z *ZapLogger) Warn(msg string, fields ...interface{}) {
	z.logger.Warnf(msg, fields...)
}
