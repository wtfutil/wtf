package wtf

import (
	"bytes"
	"fmt"
	"github.com/rivo/tview"
	"strings"
)

//BarGraph lets make graphs
type BarGraph struct {
	enabled    bool
	focusable  bool
	starChar   string
	maxStars   int
	Name       string
	RefreshInt int
	View       *tview.TextView

	Position
}

type Bar struct {
	Label      string
	Percent    int
	ValueLabel string
}

// NewBarGraph initialize your fancy new graph
func NewBarGraph(app *tview.Application, name string, configKey string, focusable bool) BarGraph {
	widget := BarGraph{
		enabled:    Config.UBool(fmt.Sprintf("wtf.mods.%s.enabled", configKey), false),
		focusable:  focusable,
		starChar:   Config.UString(fmt.Sprintf("wtf.mods.%s.graphIcon", configKey), "|"),
		maxStars:   Config.UInt(fmt.Sprintf("wtf.mods.%s.graphStars", configKey), 20),
		Name:       Config.UString(fmt.Sprintf("wtf.mods.%s.title", configKey), name),
		RefreshInt: Config.UInt(fmt.Sprintf("wtf.mods.%s.refreshInterval", configKey), 1),
	}

	widget.Position = NewPosition(
		Config.UInt(fmt.Sprintf("wtf.mods.%s.position.top", configKey)),
		Config.UInt(fmt.Sprintf("wtf.mods.%s.position.left", configKey)),
		Config.UInt(fmt.Sprintf("wtf.mods.%s.position.width", configKey)),
		Config.UInt(fmt.Sprintf("wtf.mods.%s.position.height", configKey)),
	)

	widget.addView(app, configKey)

	return widget
}

func (widget *BarGraph) BorderColor() string {
	if widget.Focusable() {
		return Config.UString("wtf.colors.border.focusable", "red")
	}

	return Config.UString("wtf.colors.border.normal", "gray")
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

func (widget *BarGraph) addView(app *tview.Application, configKey string) {
	view := tview.NewTextView()

	view.SetBackgroundColor(ColorFor(Config.UString("wtf.colors.background", "black")))
	view.SetBorder(true)
	view.SetBorderColor(ColorFor(widget.BorderColor()))
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name)
	view.SetTitleColor(ColorFor(
		Config.UString(
			fmt.Sprintf("wtf.mods.%s.colors.title", configKey),
			Config.UString("wtf.colors.title", "white"),
		),
	))
	view.SetWrap(false)

	widget.View = view
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
