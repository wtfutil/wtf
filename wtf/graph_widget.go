package wtf

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/rivo/tview"
)

type GraphWidget struct {
	enabled   bool
	focusable bool

	Name        string
	RefreshedAt time.Time
	RefreshInt  int
	View        *tview.TextView

	Position

	Data [][2]int64
}

// NewGraphWidget initialize your fancy new graph
func NewGraphWidget(name string, configKey string, focusable bool) GraphWidget {
	widget := GraphWidget{
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

func (widget *GraphWidget) BorderColor() string {
	if widget.Focusable() {
		return Config.UString("wtf.colors.border.focusable", "red")
	}

	return Config.UString("wtf.colors.border.normal", "gray")
}

func (widget *GraphWidget) Disabled() bool {
	return !widget.Enabled()
}

func (widget *GraphWidget) Enabled() bool {
	return widget.enabled
}

func (widget *GraphWidget) Focusable() bool {
	return widget.enabled && widget.focusable
}

func (widget *GraphWidget) RefreshInterval() int {
	return widget.RefreshInt
}

func (widget *GraphWidget) TextView() *tview.TextView {
	return widget.View
}

/* -------------------- Unexported Functions -------------------- */

func (widget *GraphWidget) UpdateRefreshedAt() {
	widget.RefreshedAt = time.Now()
}

func (widget *GraphWidget) addView() {
	view := tview.NewTextView()

	view.SetBackgroundColor(ColorFor(Config.UString("wtf.colors.background", "black")))
	view.SetBorder(true)
	view.SetBorderColor(ColorFor(widget.BorderColor()))
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name)
	view.SetWrap(false)

	widget.View = view
}

// BuildBars will build a string of * to represent your data of [time][value]
// time should be passed as a int64
func (widget *GraphWidget) BuildBars(maxStars int, starChar string, data [][2]int64) {

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

	widget.View.SetText(buffer.String())

}

/* -------------------- Exported Functions -------------------- */
