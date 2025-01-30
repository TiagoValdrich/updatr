package updater

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

	dirEntries, err := u.readDirectoriesOnPath(arguments.Path)
	if err != nil {
		return err
	}

	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			u.logger.Infow("directory found, CD into it", "directory", dirEntry.Name())

			err := u.runCommandsInDirectory(arguments.Path, dirEntry.Name())
			if err != nil {
				u.logger.Errorw(
					"failed to run commands in directory",
					"error", err,
					"directory", dirEntry.Name(),
					"path", *arguments.Path,
				)
			}
		}
	}

	return nil
}

func (u *Updater) validatePath(path *string) error {
	if path == nil {
		return fmt.Errorf("the provided path is nil")
	}

	if u.hasUserHomeDir(*path) {
		newPath, err := u.replaceHomeDirAlias(*path)
		if err != nil {
			u.logger.Errorw("failed to replace home dir alias", "error", err)

			return err
		}

		*path = newPath
	}

	pathInfo, err := os.Stat(*path)
	if err != nil {
		u.logger.Errorw(
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
		u.logger.Errorw("failed to read directory", "error", err, "path", *path)

		return nil, err
	}

	return files, nil
}

func (u *Updater) runCommandsInDirectory(path *string, directory string) error {
	dirFullPath := filepath.Join(*path, directory)

	cmd := exec.Command("ls", "-la")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = dirFullPath

	if err := cmd.Run(); err != nil {
		u.logger.Errorw(
			"failed to run command in directory",
			"error", err,
			"command", "ls -la",
			"dirPath", dirFullPath,
		)

		return err
	}

	return nil
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
