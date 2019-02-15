package wtf

import (
	"fmt"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
)

var Config *config.Config

type TextWidget struct {
	enabled   bool
	focusable bool
	focusChar string

	Name          string
	Configuration *config.Config
	RefreshInt    int
	View          *tview.TextView

	Position
}

func NewTextWidget(app *tview.Application, name string, configKey string, focusable bool) TextWidget {
	// Ignoring error, because.... why wouldn't we have a config?
	config, _ := Config.Get(fmt.Sprintf("wtf.mods.%s", configKey))

	focusCharValue := config.UInt("focusChar", -1)
	focusChar := string('0' + focusCharValue)
	if focusCharValue == -1 {
		focusChar = ""
	}

	widget := TextWidget{
		enabled:       config.UBool("enabled", false),
		focusable:     focusable,
		focusChar:     focusChar,
		Name:          config.UString("title"),
		Configuration: config,
		RefreshInt:    config.UInt("refreshInterval"),
	}

	widget.Position = NewPosition(
		config.UInt("position.top"),
		config.UInt("position.left"),
		config.UInt("position.width"),
		config.UInt("position.height"),
	)

	widget.addView(app, configKey)

	return widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *TextWidget) BorderColor() string {
	if widget.Focusable() {
		return Config.UString("wtf.colors.border.focusable", "red")
	}

	return Config.UString("wtf.colors.border.normal", "gray")
}

func (widget *TextWidget) ContextualTitle(defaultStr string) string {
	if widget.FocusChar() == "" {
		return fmt.Sprintf(" %s ", defaultStr)
	}

	return fmt.Sprintf(" %s [darkgray::u]%s[::-][green] ", defaultStr, widget.FocusChar())
}

func (widget *TextWidget) Disable() {
	widget.enabled = false
}

func (widget *TextWidget) Disabled() bool {
	return !widget.Enabled()
}

func (widget *TextWidget) Enabled() bool {
	return widget.enabled
}

func (widget *TextWidget) Focusable() bool {
	return widget.enabled && widget.focusable
}

func (widget *TextWidget) FocusChar() string {
	return widget.focusChar
}

func (widget *TextWidget) RefreshInterval() int {
	return widget.RefreshInt
}

func (widget *TextWidget) SetFocusChar(char string) {
	widget.focusChar = char
}

func (widget *TextWidget) TextView() *tview.TextView {
	return widget.View
}

/* -------------------- Unexported Functions -------------------- */

func (widget *TextWidget) addView(app *tview.Application, configKey string) {
	view := tview.NewTextView()

	view.SetBackgroundColor(ColorFor(
		widget.Configuration.UString("colors.background",
			Config.UString("wtf.colors.background", "black"),
		),
	))

	view.SetTextColor(ColorFor(
		widget.Configuration.UString(
			"colors.text",
			Config.UString("wtf.colors.text", "white"),
		),
	))

	view.SetTitleColor(ColorFor(
		widget.Configuration.UString(
			"colors.title",
			Config.UString("wtf.colors.title", "white"),
		),
	))

	view.SetBorder(true)
	view.SetBorderColor(ColorFor(widget.BorderColor()))
	view.SetChangedFunc(func() {
		app.Draw()
	})
	view.SetDynamicColors(true)
	view.SetTitle(widget.ContextualTitle(widget.Name))
	view.SetWrap(false)

	widget.View = view
}
