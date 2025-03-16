package updater

import "errors"

var (
	ErrLangNotSupported  = errors.New("language not supported")
	ErrInvalidConfigFile = errors.New("invalid configuration file")
)
