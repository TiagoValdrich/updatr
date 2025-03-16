package config

import "go.uber.org/zap"

type Config struct {
	Logger *zap.SugaredLogger
}

func NewConfig() *Config {
	cfg := &Config{}

	cfg.SetupLogger()

	return cfg
}

func (c *Config) SetupLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	sugarLogger := logger.Sugar()

	c.Logger = sugarLogger
}
