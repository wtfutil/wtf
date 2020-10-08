package backend

import (
	"github.com/olebedev/config"
)

type Backend interface {
	Title() string
	Setup(*config.Config)
	BuildProjects() []*Project
	GetProject(string) *Project
	LoadTasks(string) ([]Task, error)
	CloseTask(*Task) error
	DeleteTask(*Task) error
	Sources() []string
}
