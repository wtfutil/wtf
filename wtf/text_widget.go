package wtf

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

type TextWidget struct {
	bordered        bool
	commonSettings  *cfg.Common
	enabled         bool
	focusable       bool
	focusChar       string
	name            string
	quitChan        chan bool
	refreshing      bool
	refreshInterval int
	app             *tview.Application

	View *tview.TextView
}

func NewTextWidget(app *tview.Application, commonSettings *cfg.Common, focusable bool) TextWidget {
	widget := TextWidget{
		commonSettings: commonSettings,

		app:             app,
		bordered:        commonSettings.Bordered,
		enabled:         commonSettings.Enabled,
		focusable:       focusable,
		focusChar:       commonSettings.FocusChar(),
		name:            commonSettings.Name,
		quitChan:        make(chan bool),
		refreshing:      false,
		refreshInterval: commonSettings.RefreshInterval,
	}

	widget.View = widget.addView()
	widget.View.SetBorder(widget.bordered)

	return widget
}

/* -------------------- Exported Functions -------------------- */

// Bordered returns whether or not this widget should be drawn with a border
func (widget *TextWidget) Bordered() bool {
	return widget.bordered
}

func (widget *TextWidget) BorderColor() string {
	if widget.Focusable() {
		return widget.commonSettings.Colors.BorderFocusable
	}

	return widget.commonSettings.Colors.BorderNormal
}

func (widget *TextWidget) CommonSettings() *cfg.Common {
	return widget.commonSettings
}

func (widget *TextWidget) ConfigText() string {
	return utils.HelpFromInterface(cfg.Common{})
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

func (widget *TextWidget) HelpText() string {
	return fmt.Sprintf("\n  There is no help available for widget %s", widget.commonSettings.Module.Type)
}

func (widget *TextWidget) QuitChan() chan bool {
	return widget.quitChan
}

func (widget *TextWidget) Name() string {
	return widget.name
}

// Refreshing returns TRUE if the widget is currently refreshing its data, FALSE if it is not
func (widget *TextWidget) Refreshing() bool {
	return widget.refreshing
}

// RefreshInterval returns how often, in seconds, the widget will return its data
func (widget *TextWidget) RefreshInterval() int {
	return widget.refreshInterval
}

func (widget *TextWidget) SetFocusChar(char string) {
	widget.focusChar = char
}

func (widget *TextWidget) Stop() {
	widget.enabled = false
	widget.quitChan <- true
}

func (widget *TextWidget) String() string {
	return widget.name
}

func (widget *TextWidget) TextView() *tview.TextView {
	return widget.View
}

func (widget *TextWidget) Redraw(title, text string, wrap bool) {
	widget.app.QueueUpdateDraw(func() {
		widget.View.Clear()
		widget.View.SetWrap(wrap)
		widget.View.SetTitle(widget.ContextualTitle(title))
		widget.View.SetText(text)
	})
}

/* -------------------- Unexported Functions -------------------- */

func (widget *TextWidget) addView() *tview.TextView {
	view := tview.NewTextView()

	view.SetBackgroundColor(ColorFor(widget.commonSettings.Colors.Background))
	view.SetBorderColor(ColorFor(widget.BorderColor()))
	view.SetTextColor(ColorFor(widget.commonSettings.Colors.Text))
	view.SetTitleColor(ColorFor(widget.commonSettings.Colors.Title))

	view.SetBorder(true)
	view.SetDynamicColors(true)
	view.SetWrap(false)

	return view
}
