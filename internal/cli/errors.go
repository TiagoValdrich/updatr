package cli

import "errors"

var (
	ErrPathNotProvided       = errors.New("path to the directory not informed. Please use the --path flag")
	ErrConfigFileNotProvided = errors.New("config file path not informed. Please use the --config flag")
)
