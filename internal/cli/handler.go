package cli

import (
	"flag"

	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.SugaredLogger
}

func NewHandler(logger *zap.SugaredLogger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) ReadArguments() (Arguments, error) {
	arguments := Arguments{}
	arguments.Path = flag.String("path", "", "Path to the directory that will be scanned for the updates")
	arguments.ConfigFilePath = flag.String("config", "", "Path to the configuration(.toml) file")
	flag.Parse()

	if *arguments.Path == "" {
		return arguments, ErrPathNotProvided
	}

	if *arguments.ConfigFilePath == "" {
		return arguments, ErrConfigFileNotProvided
	}

	return arguments, nil
}
