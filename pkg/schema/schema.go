package schema

import (
	"encoding/json"
	"fmt"
)

var EmptyProject = Project{}

// Project holds the metadata about a project tracked by appv.
type Project struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Registry string `json:"registry"`
}

func (p Project) String() string {
	rawJson, _ := json.MarshalIndent(p, "", "    ")
	return fmt.Sprint(string(rawJson))
}

// Image returns the container image name that the project ought to create.
func (p Project) Image() string {
	return fmt.Sprintf("%s/%s:%s", p.Registry, p.Name, p.Version)
}
