package wtf

import (
	"fmt"
	"time"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
)

var Config *config.Config

type Position struct {
	top    int
	left   int
	width  int
	height int
}

func (pos *Position) Top() int {
	return pos.top
}

func (pos *Position) Left() int {
	return pos.left
}

func (pos *Position) Width() int {
	return pos.width
}

func (pos *Position) Height() int {
	return pos.height
}

/* -------------------- TextWidget -------------------- */

type TextWidget struct {
	enabled bool

	Name        string
	RefreshedAt time.Time
	RefreshInt  int
	View        *tview.TextView

	Position
}

func NewTextWidget(name string, configKey string) TextWidget {
	widget := TextWidget{
		enabled:    Config.UBool(fmt.Sprintf("wtf.%s.enabled", configKey), false),
		Name:       name,
		RefreshInt: Config.UInt(fmt.Sprintf("wtf.%s.refreshInterval", configKey)),
		Position: Position{
			top:    Config.UInt(fmt.Sprintf("wtf.%s.position.top", configKey)),
			left:   Config.UInt(fmt.Sprintf("wtf.%s.position.left", configKey)),
			height: Config.UInt(fmt.Sprintf("wtf.%s.position.height", configKey)),
			width:  Config.UInt(fmt.Sprintf("wtf.%s.position.width", configKey)),
		},
	}

	return widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *TextWidget) Disabled() bool {
	return !widget.Enabled()
}

func (widget *TextWidget) Enabled() bool {
	return widget.enabled
}

func (widget *TextWidget) RefreshInterval() int {
	return widget.RefreshInt
}

func (widget *TextWidget) TextView() *tview.TextView {
	return widget.View
}
