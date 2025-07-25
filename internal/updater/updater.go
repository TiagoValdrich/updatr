package updater

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/tiagovaldrich/updatr/internal/cli"
	"github.com/tiagovaldrich/updatr/internal/logger"
)

type Updater struct {
	logger    logger.Logger
	arguments cli.Arguments
}

func NewUpdater(logger logger.Logger, arguments cli.Arguments) *Updater {
	return &Updater{
		logger:    logger,
		arguments: arguments,
	}
}

func (u *Updater) Update() error {
	if err := u.validatePath(u.arguments.Path); err != nil {
		return err
	}

	dirEntries, err := u.readDirectoriesOnPath(u.arguments.Path)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}

	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			wg.Add(1)

			u.logger.Info("directory found, CD into it", "directory", dirEntry.Name())

			go func() {
				defer wg.Done()

				err := u.runCommandsInDirectory(u.arguments.Path, dirEntry.Name())
				if err != nil {
					u.logger.Error(
						"failed to run commands in directory",
						"error", err,
						"directory", dirEntry.Name(),
						"path", *u.arguments.Path,
					)
				}
			}()
		}
	}

	wg.Wait()

	return nil
}

func (u *Updater) validatePath(path *string) error {
	if path == nil {
		return fmt.Errorf("the provided path is nil")
	}

	if u.hasUserHomeDir(*path) {
		newPath, err := u.replaceHomeDirAlias(*path)
		if err != nil {
			u.logger.Error("failed to replace home dir alias", "error", err)

			return err
		}

		*path = newPath
	}

	pathInfo, err := os.Stat(*path)
	if err != nil {
		u.logger.Error(
			"failed to get path info",
			"error", err,
			"path", *path,
		)

		if os.IsNotExist(err) {
			return fmt.Errorf("the provided path does not exist")
		}

		return err
	}

	if !pathInfo.IsDir() {
		return fmt.Errorf("the provided path is not a directory")
	}

	return nil
}

func (u *Updater) readDirectoriesOnPath(path *string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(*path)
	if err != nil {
		u.logger.Error("failed to read directory", "error", err, "path", *path)

		return nil, err
	}

	return files, nil
}

func (u *Updater) runCommandsInDirectory(path *string, directory string) error {
	dirFullPath := filepath.Join(*path, directory)

	return NewLangUpdater(u.logger, dirFullPath, u.arguments.ConfigFilePath).Update()
}

func (u *Updater) hasUserHomeDir(path string) bool {
	return len(path) > 0 && path[0] == UserHomeDirAlias
}

func (u *Updater) replaceHomeDirAlias(path string) (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	newPath := strings.Replace(path, string(UserHomeDirAlias), userHomeDir, 1)

	return newPath, nil
}
