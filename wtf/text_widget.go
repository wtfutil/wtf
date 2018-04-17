package wtf

import (
	"fmt"
	"time"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/color"
)

var Config *config.Config

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
		enabled:    Config.UBool(fmt.Sprintf("wtf.mods.%s.enabled", configKey), false),
		Name:       name,
		RefreshInt: Config.UInt(fmt.Sprintf("wtf.mods.%s.refreshInterval", configKey)),
		Position: Position{
			top:    Config.UInt(fmt.Sprintf("wtf.mods.%s.position.top", configKey)),
			left:   Config.UInt(fmt.Sprintf("wtf.mods.%s.position.left", configKey)),
			height: Config.UInt(fmt.Sprintf("wtf.mods.%s.position.height", configKey)),
			width:  Config.UInt(fmt.Sprintf("wtf.mods.%s.position.width", configKey)),
		},
	}

	widget.addView()

	return widget
}

/* -------------------- Exported Functions -------------------- */

//func (widget *TextWidget) SetBorder() {
  //var colorName string

	//if widget.View.HasFocus() {
		//colorName = Config.UString("wtf.border.normal")
 //} else {
		//colorName = Config.UString("wtf.border.focus")
 //}

 //widget.View.SetBorderColor(color.ColorFor(colorName))
//}

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

/* -------------------- Unexported Functions -------------------- */

func (widget *TextWidget) addView() {
	view := tview.NewTextView()

	view.SetBorder(true)
	view.SetBorderColor(color.ColorFor(Config.UString("wtf.border.normal")))
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name)
	view.SetWrap(false)

	widget.View = view
}

