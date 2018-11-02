package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/CzarSimon/appv/pkg/schema"
)

const PROJECT_FILE = "appv.json"

var ErrNoProjectFound = errors.New("No project file found.")

type ProjectRepo interface {
	Get() (schema.Project, error)
	Save(project schema.Project) error
}

type fileRepo struct {
	cwd string
}

func NewProjectRepo() *fileRepo {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &fileRepo{
		cwd: cwd,
	}
}

func (r *fileRepo) Save(project schema.Project) error {
	rawJson, err := json.Marshal(project)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(PROJECT_FILE, rawJson, 0644)
}

func (r *fileRepo) Get() (schema.Project, error) {
	rawJson, err := ioutil.ReadFile(PROJECT_FILE)
	if err == os.ErrNotExist {
		return schema.EmptyProject, ErrNoProjectFound
	}
	if err != nil {
		return schema.EmptyProject, err
	}

	var project schema.Project

	err = json.Unmarshal(rawJson, &rawJson)
	return project, err
}
