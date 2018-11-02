package cli

import (
	"github.com/CzarSimon/appv/pkg/schema"
	"github.com/urfave/cli"
)

func (appv *appv) update(c *cli.Context) error {
	project, err := appv.porjectRepo.Get()
	if err != nil {
		return err
	}

	newProject := getProjectUpdateFromUser(project)
	return appv.porjectRepo.Save(newProject)
}

func getProjectUpdateFromUser(oldProject schema.Project) schema.Project {
	var newProject schema.Project
	newProject.Name = stringFromStdinWithDefault(
		"Name", oldProject.Name)
	newProject.Version = stringFromStdinWithDefault(
		"Version", oldProject.Version)
	newProject.Registry = stringFromStdinWithDefault(
		"Container registry", oldProject.Registry)

	if !confirmProject(newProject) {
		return getProjectUpdateFromUser(oldProject)
	}

	return newProject
}
