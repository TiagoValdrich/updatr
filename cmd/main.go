package main

import (
	"github.com/tiagovaldrich/updatr/internal/cli"
	"github.com/tiagovaldrich/updatr/internal/config"
	"github.com/tiagovaldrich/updatr/internal/updater"
)

func main() {
	cfg := config.NewConfig()

	cliHandler := cli.NewHandler(cfg.Logger)

	arguments := cliHandler.ReadArguments()

	updater := updater.NewUpdater(cfg.Logger, arguments)
	if err := updater.Update(); err != nil {
		panic(err)
	}
}
