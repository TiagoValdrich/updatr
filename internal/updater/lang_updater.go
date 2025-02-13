package updater

import (
	"os"

	"github.com/tiagovaldrich/updatr/internal/plangs"
	"go.uber.org/zap"
)

type LangUpdater struct {
	logger         *zap.SugaredLogger
	directory      string
	executor       *Executor
	configLoader   *ConfigLoader
	configFilePath *string
}

// Probably good to rethink how this is instantiated, too many parameters already
func NewLangUpdater(logger *zap.SugaredLogger, directory string, configFilePath *string) *LangUpdater {
	langUpdater := &LangUpdater{
		logger:         logger,
		directory:      directory,
		configFilePath: configFilePath,
	}

	langUpdater.executor = NewExecutor(ExecutorParams{
		Logger:  logger,
		DirPath: directory,
	})

	langUpdater.configLoader = NewConfigLoader(logger)

	return langUpdater
}

func (lu *LangUpdater) Update() error {
	commandsToRun, err := lu.loadCommands()
	if err != nil {
		lu.logger.Errorw("failed to load commands", "error", err)

		return err
	}

	for _, command := range commandsToRun {
		lu.logger.Infow("running command", "command", command, "directory", lu.directory)

		if err := lu.executor.Run(command); err != nil {
			lu.logger.Errorw(
				"failed to run command",
				"error", err,
				"command", command,
				"directory", lu.directory,
			)

			return err
		}
	}

	return nil
}

func (lu *LangUpdater) loadCommands() ([]string, error) {
	programmingLanguage, err := lu.identifyProgrammingLanguage()
	if err != nil {
		lu.logger.Errorw("failed to identify programming language", "error", err)

		return nil, err
	}

	lu.logger.Infow("programming language identified", "programming_language", programmingLanguage)

	commands, err := lu.loadComandsFromProgramingLanguage(programmingLanguage)
	if err != nil {
		lu.logger.Errorw(
			"failed to load commands from programming language",
			"error", err,
			"programming_language", programmingLanguage,
		)

		return nil, err
	}

	return commands, nil
}

func (lu *LangUpdater) identifyProgrammingLanguage() (plangs.ProgrammingLanguage, error) {
	dirEntries, err := os.ReadDir(lu.directory)
	if err != nil {
		return "", err
	}

	for _, langIdentifier := range plangs.AvailableIdentifiers {
		programmingLanguage, err := langIdentifier.Identify(dirEntries)
		if err != nil {
			return "", err
		}

		if programmingLanguage != "" {
			return programmingLanguage, nil
		}

	}

	return "", ErrLangNotSupported
}

func (lu *LangUpdater) loadComandsFromProgramingLanguage(
	programmingLanguage plangs.ProgrammingLanguage,
) ([]string, error) {
	err := lu.configLoader.LoadConfig(lu.configFilePath)
	if err != nil {
		return DefaultOperations, err
	}

	lu.logger.Infow("config file loaded", "languageConfig", lu.configLoader.GetLanguageConfig())

	if lu.configLoader.IsLanguageAvailable(programmingLanguage.String()) {
		return lu.configLoader.GetCommandsForLanguage(programmingLanguage.String()), nil
	}

	return DefaultOperations, nil
}
