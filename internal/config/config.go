package config

import "go.uber.org/zap"

type Config struct {
	Logger *zap.SugaredLogger
}

func NewConfig() *Config {
	cfg := &Config{}

	cfg.SetupLogger()
	// here we are going to load environment variables
	// that will load custom configurations to the application

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
