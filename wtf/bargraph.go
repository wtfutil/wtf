package wtf

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/rivo/tview"
)

//BarGraph lets make graphs
type BarGraph struct {
	enabled     bool
	focusable   bool
	starChar    string
	maxStars    int
	Name        string
	RefreshedAt time.Time
	RefreshInt  int
	View        *tview.TextView

	Position

	Data [][2]int64
}

// NewBarGraph initialize your fancy new graph
func NewBarGraph(name string, configKey string, focusable bool) BarGraph {
	widget := BarGraph{
		enabled:    Config.UBool(fmt.Sprintf("wtf.mods.%s.enabled", configKey), false),
		focusable:  focusable,
		starChar:   Config.UString(fmt.Sprintf("wtf.mods.%s.graphIcon", configKey), name),
		maxStars:   Config.UInt(fmt.Sprintf("wtf.mods.%s.graphStars", configKey), 20),
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

func (widget *BarGraph) UpdateRefreshedAt() {
	widget.RefreshedAt = time.Now()
}

func (widget *BarGraph) addView() {
	view := tview.NewTextView()

	view.SetBackgroundColor(colorFor(Config.UString("wtf.colors.background", "black")))
	view.SetBorder(true)
	view.SetBorderColor(colorFor(widget.BorderColor()))
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name)
	view.SetWrap(false)

	widget.View = view
}

// BuildBars will build a string of * to represent your data of [time][value]
// time should be passed as a int64
func (widget *BarGraph) BuildBars(data [][2]int64) {

	widget.View.SetText(BuildStars(data, widget.maxStars, widget.starChar))

}

//BuildStars build the string to display
func BuildStars(data [][2]int64, maxStars int, starChar string) string {
	var buffer bytes.Buffer

	//counter to inintialize min value
	var count int

	//store the max value from the array
	var maxValue int

	//store the min value from the array
	var minValue int

	//just getting min and max values
	for i := range data {

		var val = int(data[i][0])

		//initialize the min value
		if count == 0 {
			minValue = val
		}
		count++

		//update max value
		if val > maxValue {
			maxValue = val
		}

		//update minValue
		if val < minValue {
			minValue = val
		}

	}

	// each number = how many stars?
	var starRatio = float64(maxStars) / float64((maxValue - minValue))

	//build the stars
	for i := range data {
		var val = int(data[i][0])

		//how many stars for this one?
		var starCount = int(float64((val - minValue)) * starRatio)

		if starCount == 0 {
			starCount = 1
		}
		//build the actual string
		var stars = strings.Repeat(starChar, starCount)

		//parse the time
		var t = time.Unix(int64(data[i][1]/1000), 0)

		//write the line
		buffer.WriteString(fmt.Sprintf("%s -\t [red]%s[white] - (%d)\n", t.Format("Jan 02, 2006"), stars, val))
	}

	return buffer.String()
}

/* -------------------- Exported Functions -------------------- */
