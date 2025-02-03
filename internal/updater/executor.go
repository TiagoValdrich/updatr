package updater

import (
	"os"
	"os/exec"

	"go.uber.org/zap"
)

type ExecutorParams struct {
	Logger  *zap.SugaredLogger
	Input   *os.File
	Output  *os.File
	Error   *os.File
	DirPath string
}

type Executor struct {
	logger  *zap.SugaredLogger
	Input   *os.File
	Output  *os.File
	Error   *os.File
	DirPath string
}

func NewExecutor(params ExecutorParams) *Executor {
	executor := &Executor{
		logger:  params.Logger,
		Input:   params.Input,
		Output:  params.Output,
		Error:   params.Error,
		DirPath: params.DirPath,
	}

	executor.setDefaultParams()

	return executor
}

func (e *Executor) setDefaultParams() {
	if e.Input == nil {
		e.Input = os.Stdin
	}

	if e.Output == nil {
		e.Output = os.Stdout
	}

	if e.Error == nil {
		e.Error = os.Stderr
	}
}

func (e *Executor) Run(command string) error {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdin = e.Input
	cmd.Stdout = e.Output
	cmd.Stderr = e.Error
	cmd.Dir = e.DirPath

	if err := cmd.Run(); err != nil {
		e.logger.Errorw(
			"failed to run command",
			"error", err,
			"command", command,
		)

		return err
	}

	return nil
}
