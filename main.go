package main

import (
	"github.com/CzarSimon/appv/pkg/cli"
)

func main() {
	app := cli.New()

	err := app.Run()
	cli.CheckErrAndExit(err)
}
