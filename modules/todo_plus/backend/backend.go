package backend

import (
	"github.com/olebedev/config"
)

type Backend interface {
	Title() string
	Setup(*config.Config)
	NewProject(int) *Project
	LoadTasks(int) ([]Task, error)
	CloseTask(*Task) error
	DeleteTask(*Task) error
}
