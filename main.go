package main

import (
	"github.com/tiagovaldrich/updatr/internal/cli"
	"github.com/tiagovaldrich/updatr/internal/config"
	"github.com/tiagovaldrich/updatr/internal/updater"
)

func main() {
	cfg := config.NewConfig()

	cliHandler := cli.NewHandler(cfg.Logger)

	arguments, err := cliHandler.ReadArguments()
	if err != nil {
		cfg.Logger.Error("failed to read cli arguments", err)

		return
	}

	updater := updater.NewUpdater(cfg.Logger, arguments)
	if err := updater.Update(); err != nil {
		panic(err)
	}
}
