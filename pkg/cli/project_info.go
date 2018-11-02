package cli

import (
	"fmt"

	"github.com/urfave/cli"
)

func (appv *appv) image(c *cli.Context) error {
	if !appv.initalized {
		return errAppIsNotInitalized
	}

	printToStdout(appv.project.Image(), c)
	return nil
}

func (appv *appv) name(c *cli.Context) error {
	if !appv.initalized {
		return errAppIsNotInitalized
	}

	printToStdout(appv.project.Name, c)
	return nil
}

func (appv *appv) version(c *cli.Context) error {
	if !appv.initalized {
		return errAppIsNotInitalized
	}

	printToStdout(appv.project.Version, c)
	return nil
}

func (appv *appv) registry(c *cli.Context) error {
	if !appv.initalized {
		return errAppIsNotInitalized
	}

	printToStdout(appv.project.Registry, c)
	return nil
}

func printToStdout(message string, c *cli.Context) {
	fmt.Print(message)
	if c.Bool("new-line") {
		fmt.Print("\n")
	}
}
