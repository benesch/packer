package compress

import (
	"fmt"
	"os"
)

const BuilderId = "johnbellone.compress"

type Artifact struct {
	Path     string
	Provider string
}

func NewArtifact(provider, path string) *Artifact {
	return &Artifact{
		Path:     path,
		Provider: provider,
	}
}

func (*Artifact) BuilderId() string {
	return BuilderId
}

func (self *Artifact) Id() string {
	return ""
}

func (self *Artifact) Files() []string {
	return []string{self.Path}
}

func (self *Artifact) String() string {
	return fmt.Sprintf("'%s' compressing: %s", self.Provider, self.Path)
}

func (self *Artifact) Destroy() error {
	return os.Remove(self.Path)
}