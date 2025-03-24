package cli

import (
	"flag"

	"github.com/tiagovaldrich/updatr/internal/logger"
)

type Handler struct {
	logger logger.Logger
}

func NewHandler(logger logger.Logger) *Handler {
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
