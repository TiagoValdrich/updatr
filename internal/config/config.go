package config

import (
	"github.com/tiagovaldrich/updatr/internal/logger"
	"go.uber.org/zap"
)

type Config struct {
	Logger logger.Logger
}

func NewConfig() *Config {
	cfg := &Config{}

	cfg.SetupLogger()

	return cfg
}

func (c *Config) SetupLogger() {
	c.setupDefaultLogger()
}

func (c *Config) setupDefaultLogger() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.CallerKey = ""
	config.EncoderConfig.TimeKey = ""

	zapLogger, err := config.Build()
	if err != nil {
		panic(err)
	}

	c.Logger = logger.NewZapLogger(logger.ZapLoggerParams{
		ZapLogger: zapLogger.Sugar(),
	})
}
