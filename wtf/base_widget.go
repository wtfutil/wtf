package wtf

import (
	//"fmt"
	"time"
)

type BaseWidget struct {
	Name        string
	RefreshedAt time.Time
	RefreshInt  int
}

func (widget *BaseWidget) RefreshInterval() int {
	return widget.RefreshInt
}
