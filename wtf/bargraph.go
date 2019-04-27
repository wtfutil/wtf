package wtf

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
)

//BarGraph lets make graphs
type BarGraph struct {
	enabled   bool
	focusable bool
	key       string
	maxStars  int
	name      string
	starChar  string

	RefreshInt int
	View       *tview.TextView
	settings   *cfg.Common

	Position
}

type Bar struct {
	Label      string
	Percent    int
	ValueLabel string
}

// NewBarGraph initialize your fancy new graph
func NewBarGraph(app *tview.Application, name string, settings *cfg.Common, focusable bool) BarGraph {
	widget := BarGraph{
		enabled:    settings.Enabled,
		focusable:  focusable,
		maxStars:   settings.Config.UInt("graphStars", 20),
		name:       settings.Title,
		starChar:   settings.Config.UString("graphIcon", "|"),
		RefreshInt: settings.RefreshInterval,
		settings:   settings,
	}

	widget.Position = NewPosition(
		settings.Position.Top,
		settings.Position.Left,
		settings.Position.Width,
		settings.Position.Height,
	)

	widget.View = widget.addView()
	widget.View.SetChangedFunc(func() {
		app.Draw()
	})

	return widget
}

func (widget *BarGraph) BorderColor() string {
	if widget.Focusable() {
		return widget.settings.Colors.BorderFocusable
	}

	return widget.settings.Colors.BorderNormal
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

// IsPositionable returns TRUE if the widget has valid position parameters, FALSE if it has
// invalid position parameters (ie: cannot be placed onscreen)
func (widget *BarGraph) IsPositionable() bool {
	return widget.Position.IsValid()
}

func (widget *BarGraph) Key() string {
	return widget.key
}

func (widget *BarGraph) Name() string {
	return widget.name
}

func (widget *BarGraph) RefreshInterval() int {
	return widget.RefreshInt
}

func (widget *BarGraph) SetFocusChar(char string) {
	return
}

func (widget *BarGraph) TextView() *tview.TextView {
	return widget.View
}

/* -------------------- Unexported Functions -------------------- */

func (widget *BarGraph) addView() *tview.TextView {
	view := tview.NewTextView()

	view.SetBackgroundColor(ColorFor(widget.settings.Colors.Background))
	view.SetBorder(true)
	view.SetBorderColor(ColorFor(widget.BorderColor()))
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name())
	view.SetTitleColor(ColorFor(widget.settings.Colors.Title))
	view.SetWrap(false)

	return view
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

/* -------------------- Exported Functions -------------------- */
