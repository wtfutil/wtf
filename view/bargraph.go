package view

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/wtf"
)

//BarGraph lets make graphs
type BarGraph struct {
	commonSettings *cfg.Common
	enabled        bool
	focusable      bool
	key            string
	maxStars       int
	name           string
	quitChan       chan bool
	refreshing     bool
	starChar       string

	RefreshInt int
	View       *tview.TextView
}

type Bar struct {
	Label      string
	Percent    int
	ValueLabel string
}

// NewBarGraph initialize your fancy new graph
func NewBarGraph(app *tview.Application, name string, settings *cfg.Common, focusable bool) BarGraph {
	widget := BarGraph{
		enabled:        settings.Enabled,
		focusable:      focusable,
		maxStars:       settings.Config.UInt("graphStars", 20),
		name:           settings.Title,
		quitChan:       make(chan bool),
		starChar:       settings.Config.UString("graphIcon", "|"),
		commonSettings: settings,

		RefreshInt: settings.RefreshInterval,
	}

	widget.View = widget.addView()
	widget.View.SetChangedFunc(func() {
		app.Draw()
	})

	return widget
}

func (widget *BarGraph) BorderColor() string {
	if widget.Focusable() {
		return widget.commonSettings.Colors.BorderFocusable
	}

	return widget.commonSettings.Colors.BorderNormal
}

func (widget *BarGraph) CommonSettings() *cfg.Common {
	return widget.commonSettings
}

func (widget *BarGraph) Disable() {
	widget.enabled = false
}

func (widget *BarGraph) Disabled() bool {
	return !widget.Enabled()
}

func (widget *BarGraph) Enabled() bool {
	return widget.enabled
}

func (widget *BarGraph) Focusable() bool {
	return widget.enabled && widget.focusable
}

func (widget *BarGraph) FocusChar() string {
	return ""
}

func (widget *BarGraph) Key() string {
	return widget.key
}

func (widget *BarGraph) Name() string {
	return widget.name
}

func (widget *BarGraph) QuitChan() chan bool {
	return widget.quitChan
}

// Refreshing returns TRUE if the widget is currently refreshing its data, FALSE if it is not
func (widget *BarGraph) Refreshing() bool {
	return widget.refreshing
}

// RefreshInterval returns how often, in seconds, the widget will return its data
func (widget *BarGraph) RefreshInterval() int {
	return widget.RefreshInt
}

func (widget *BarGraph) SetFocusChar(char string) {
	return
}

func (widget *BarGraph) Stop() {
	widget.enabled = false
	widget.quitChan <- true
}

func (widget *BarGraph) TextView() *tview.TextView {
	return widget.View
}

func (widget *BarGraph) HelpText() string {
	return "No help available for this widget"
}

func (widget *BarGraph) ConfigText() string {
	return ""
}

// BuildBars will build a string of * to represent your data of [time][value]
// time should be passed as a int64
func (widget *BarGraph) BuildBars(data []Bar) {
	widget.View.SetText(BuildStars(data, widget.maxStars, widget.starChar))
}

//BuildStars build the string to display
func BuildStars(data []Bar, maxStars int, starChar string) string {
	var buffer bytes.Buffer

	// the number of characters in the longest label
	var longestLabel int

	//just getting min and max values
	for _, bar := range data {

		if len(bar.Label) > longestLabel {
			longestLabel = len(bar.Label)
		}

	}

	// each number = how many stars?
	var starRatio = float64(maxStars) / 100

	//build the stars
	for _, bar := range data {

		//how many stars for this one?
		var starCount = int(float64(bar.Percent) * starRatio)

		label := bar.ValueLabel
		if len(label) == 0 {
			label = fmt.Sprint(bar.Percent)
		}

		//write the line
		buffer.WriteString(
			fmt.Sprintf(
				"%s%s[[red]%s[white]%s] %s\n",
				bar.Label,
				strings.Repeat(" ", longestLabel-len(bar.Label)),
				strings.Repeat(starChar, starCount),
				strings.Repeat(" ", maxStars-starCount),
				label,
			),
		)
	}

	return buffer.String()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *BarGraph) addView() *tview.TextView {
	view := tview.NewTextView()

	view.SetBackgroundColor(wtf.ColorFor(widget.commonSettings.Colors.Background))
	view.SetBorder(true)
	view.SetBorderColor(wtf.ColorFor(widget.BorderColor()))
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name())
	view.SetTitleColor(wtf.ColorFor(widget.commonSettings.Colors.Title))
	view.SetWrap(false)

	return view
}
