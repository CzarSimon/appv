package cli

import (
	"fmt"
	"strings"

	"github.com/CzarSimon/appv/pkg/schema"
	"github.com/urfave/cli"
)

const initalSugestedVersion = "0.1.0"

func (appv *appv) init(c *cli.Context) error {
	if appv.initalized {
		return errAppAlreadyInitalized
	}

	fmt.Println("Please describe your new project")
	project := getProjectFromUser()

	return appv.porjectRepo.Save(project)
}

func getProjectFromUser() schema.Project {
	var project schema.Project
	project.Name = stringFromStdin("Name")
	project.Version = stringFromStdinWithDefault("Version", initalSugestedVersion)
	project.Registry = stringFromStdin("Container registry")

	if !confirmProject(project) {
		return getProjectFromUser()
	}

	return project
}

func confirmProject(project schema.Project) bool {
	fmt.Printf("Does this look ok?\n%s\n", project)
	value := stringFromStdin("No to redo, any other key to confirm")
	confirmation := strings.ToLower(value)

	return confirmation != "no"
}

func stringFromStdin(description string) string {
	fmt.Printf("%s: ", description)
	var value string
	fmt.Scanln(&value)
	return value
}

func stringFromStdinWithDefault(description, defaultValue string) string {
	value := stringFromStdin(fmt.Sprintf("%s (%s)", description, defaultValue))
	if value == "" {
		return defaultValue
	}
	return value
}
