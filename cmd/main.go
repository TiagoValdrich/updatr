package main

import (
	"fmt"
	"updatr/internal/cli"
)

func main() {
	cliHandler := cli.NewHandler()
	cliHandler.ReadArguments()

	fmt.Println("Path:", *cliHandler.Path)
}
