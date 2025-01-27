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
	arguments := Arguments{}
	arguments.Path = flag.String("path", h.getDefaultPath(), "")
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
