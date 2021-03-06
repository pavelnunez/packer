package docker

import (
	"fmt"
	"strings"
)

// ImportArtifact is an Artifact implementation for when a container is
// exported from docker into a single flat file.
type ImportArtifact struct {
	BuilderIdValue string
	Driver         Driver
	IdValue        string

	// StateData should store data such as GeneratedData
	// to be shared with post-processors
	StateData map[string]interface{}
}

func (a *ImportArtifact) BuilderId() string {
	return a.BuilderIdValue
}

func (*ImportArtifact) Files() []string {
	return nil
}

func (a *ImportArtifact) Id() string {
	return a.IdValue
}

func (a *ImportArtifact) String() string {
	tags := a.StateData["docker_tags"]
	if tags == nil {
		return fmt.Sprintf("Imported Docker image: %s", a.Id())
	}
	cast := tags.([]interface{})
	names := []string{}
	for _, name := range cast {
		if n, ok := name.(string); ok {
			names = append(names, n)
		}
	}
	return fmt.Sprintf("Imported Docker image: %s with tags %s",
		a.Id(), strings.Join(names, " "))
}

func (a *ImportArtifact) State(name string) interface{} {
	return a.StateData[name]
}

func (a *ImportArtifact) Destroy() error {
	return a.Driver.DeleteImage(a.Id())
}
