package cli

import (
	"flag"
	"os"
	"path/filepath"

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

func (h *Handler) ReadArguments() Arguments {
	executableDirPath := h.getDefaultPath()
	defaultConfigFilePath := filepath.Join(executableDirPath, "config.toml")

	arguments := Arguments{}
	arguments.Path = flag.String("path", executableDirPath, "Path to the directory to be updated")
	arguments.ConfigFilePath = flag.String("config", defaultConfigFilePath, "Path to the configuration file, by the default it will look for a config.toml file in the executable directory")
	flag.Parse()

	return arguments
}

func (h *Handler) getDefaultPath() string {
	executablePath, err := os.Executable()
	if err != nil {
		h.logger.Errorw("failed to get executable path", "error", err)

		panic(err)
	}

	return filepath.Dir(executablePath)
}
