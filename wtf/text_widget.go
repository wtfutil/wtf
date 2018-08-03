package wtf

import (
	"fmt"
	"time"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
)

var Config *config.Config

type TextWidget struct {
	enabled   bool
	focusable bool
	focusChar string

	Name        string
	RefreshedAt time.Time
	RefreshInt  int
	View        *tview.TextView

	Position
}

func NewTextWidget(name string, configKey string, focusable bool) TextWidget {
	widget := TextWidget{
		enabled:   Config.UBool(fmt.Sprintf("wtf.mods.%s.enabled", configKey), false),
		focusable: focusable,

		Name:       Config.UString(fmt.Sprintf("wtf.mods.%s.title", configKey), name),
		RefreshInt: Config.UInt(fmt.Sprintf("wtf.mods.%s.refreshInterval", configKey)),
	}

	widget.Position = NewPosition(
		Config.UInt(fmt.Sprintf("wtf.mods.%s.position.top", configKey)),
		Config.UInt(fmt.Sprintf("wtf.mods.%s.position.left", configKey)),
		Config.UInt(fmt.Sprintf("wtf.mods.%s.position.width", configKey)),
		Config.UInt(fmt.Sprintf("wtf.mods.%s.position.height", configKey)),
	)

	widget.addView()

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

func (widget *TextWidget) addView() {
	view := tview.NewTextView()

	view.SetBackgroundColor(colorFor(Config.UString("wtf.colors.background", "black")))
	view.SetBorder(true)
	view.SetBorderColor(colorFor(widget.BorderColor()))
	view.SetDynamicColors(true)
	view.SetTitle(widget.ContextualTitle(widget.Name))
	view.SetWrap(false)

	widget.View = view
}

func (widget *TextWidget) UpdateRefreshedAt() {
	widget.RefreshedAt = time.Now()
}
