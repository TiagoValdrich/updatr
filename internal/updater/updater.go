package updater

import (
	"fmt"
	"os"

	"github.com/tiagovaldrich/updatr/internal/cli"

	"go.uber.org/zap"
)

type Updater struct {
	logger *zap.SugaredLogger
}

func NewUpdater(logger *zap.SugaredLogger) *Updater {
	return &Updater{
		logger: logger,
	}
}

func (u *Updater) Update(arguments cli.Arguments) error {
	if err := u.validatePath(arguments.Path); err != nil {
		return err
	}

	_, err := u.readDirectoriesOnPath(arguments.Path)
	if err != nil {
		return err
	}

	return nil
}

func (u *Updater) validatePath(path *string) error {
	if path == nil {
		return fmt.Errorf("the provided path is nil")
	}

	pathInfo, err := os.Stat(*path)

	if os.IsNotExist(err) {
		return fmt.Errorf("the provided path does not exist")
	}

	if !pathInfo.IsDir() {
		return fmt.Errorf("the provided path is not a directory")
	}

	return nil
}

func (u *Updater) readDirectoriesOnPath(path *string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(*path)
	if err != nil {
		u.logger.Errorw("failed to read directory", "error", err, "path", *path)

		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			u.logger.Infow("found directory", "name", file.Name())
		} else {
			u.logger.Infow("found file", "name", file.Name())
		}
	}

	return files, nil
}
