package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/CzarSimon/appv/pkg/repository"
	"github.com/CzarSimon/appv/pkg/schema"
	"github.com/urfave/cli"
)

var (
	errAppIsNotInitalized   = errors.New("Project is not tracked by appv. Have you run appv init?")
	errAppAlreadyInitalized = errors.New("Project is already initalized. You maybe want to run: appv update?")
)

const (
	appName    = "appv"
	appUsage   = "cli tool for keeping track of a package version."
	appVersion = "0.9"
)

type App interface {
	Run() error
}

func New() App {
	return newAppv()
}

type appv struct {
	porjectRepo repository.ProjectRepo
	project     schema.Project
	initalized  bool
}

func newAppv() *appv {
	porjectRepo := repository.NewProjectRepo()

	project, err := porjectRepo.Get()
	if err != repository.ErrNoProjectFound {
		CheckErrAndExit(err)
	}

	return &appv{
		porjectRepo: porjectRepo,
		project:     project,
		initalized:  (err != repository.ErrNoProjectFound),
	}
}

func (appv *appv) Run() error {
	cliApp := createCli(appv)
	return cliApp.Run(os.Args)
}

func createCli(appv *appv) *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appUsage
	app.Version = appVersion
	app.Commands = appv.getCommands()
	return app
}

func (appv *appv) getCommands() []cli.Command {
	return []cli.Command{
		cli.Command{
			Name:   "init",
			Usage:  "Creates an appv file for the project.",
			Action: appv.init,
		},
		cli.Command{
			Name:   "update",
			Usage:  "Prompts the user for updates.",
			Action: appv.update,
		},
		cli.Command{
			Name:   "image",
			Usage:  "Prints the appv image name",
			Action: appv.image,
			Flags:  getOutputFlags(),
		},
		cli.Command{
			Name:   "name",
			Usage:  "Prints the appv project name",
			Action: appv.name,
			Flags:  getOutputFlags(),
		},
		cli.Command{
			Name:   "version",
			Usage:  "Prints the appv project version",
			Action: appv.version,
			Flags:  getOutputFlags(),
		},
		cli.Command{
			Name:   "registry",
			Usage:  "Prints the appv project container registry",
			Action: appv.registry,
			Flags:  getOutputFlags(),
		},
	}
}

func getOutputFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "new-line, n",
			Usage: "If set the cli output adds a linebreak.",
		},
	}
}

func CheckErrAndExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
