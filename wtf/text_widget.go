package wtf

import (
	"fmt"
	"time"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
)

var Config *config.Config

type Position struct {
	Top    int
	Left   int
	Width  int
	Height int
}

type TextWidget struct {
	Enabled     bool
	Name        string
	Position    Position
	RefreshedAt time.Time
	RefreshInt  int
	View        *tview.TextView
}

func NewTextWidget(name string, configKey string) TextWidget {
	widget := TextWidget{
		Enabled:    Config.UBool(fmt.Sprintf("wtf.%s.enabled", configKey), false),
		Name:       name,
		RefreshInt: Config.UInt(fmt.Sprintf("wtf.%s.refreshInterval", configKey)),
	}

	return widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *TextWidget) RefreshInterval() int {
	return widget.RefreshInt
}

func (widget *TextWidget) TextView() *tview.TextView {
	return widget.View
}
