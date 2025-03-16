package updater

import (
	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
)

type LanguageConfig struct {
	Commands       []string `toml:"commands"`
	IgnoreProjects []string `toml:"ignore_projects"`
}

type ConfigLoader struct {
	logger         *zap.SugaredLogger
	languageConfig map[string]LanguageConfig
}

func NewConfigLoader(logger *zap.SugaredLogger) *ConfigLoader {
	return &ConfigLoader{
		logger: logger,
	}
}

// Returns a map, where the keys are the programming language names
// and the values are the LanguageConfig structs
func (cl *ConfigLoader) GetLanguageConfig() map[string]LanguageConfig {
	return cl.languageConfig
}

func (cl *ConfigLoader) IsLanguageAvailable(language string) bool {
	_, ok := cl.languageConfig[language]

	return ok
}

func (cl *ConfigLoader) GetCommandsForLanguage(language string) []string {
	return cl.languageConfig[language].Commands
}

// Load into the ConfigLoader the configuration from the provided file path
func (cl *ConfigLoader) LoadConfig(filePath *string) error {
	if filePath == nil {
		cl.logger.Error("config file path is nil")

		return ErrInvalidConfigFile
	}

	if _, err := toml.DecodeFile(*filePath, &cl.languageConfig); err != nil {
		cl.logger.Errorw("failed to decode toml configuration file", "error", err)

		return err
	}

	return nil
}

func (cl *ConfigLoader) CanIgnoreProject(language string, projectName string) bool {
	ignoreProjects := cl.languageConfig[language].IgnoreProjects

	for _, project := range ignoreProjects {
		if project == projectName {
			return true
		}
	}

	return false
}
