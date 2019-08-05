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
	maxStars int
	starChar string

	Base
	View *tview.TextView
}

type Bar struct {
	Label      string
	Percent    int
	ValueLabel string
}

// NewBarGraph initialize your fancy new graph
func NewBarGraph(app *tview.Application, name string, commonSettings *cfg.Common, focusable bool) BarGraph {
	widget := BarGraph{
		Base: NewBase(app, commonSettings, focusable),

		maxStars: commonSettings.Config.UInt("graphStars", 20),
		starChar: commonSettings.Config.UString("graphIcon", "|"),
	}

	widget.View = widget.addView()
	widget.View.SetBorder(widget.bordered)

	return widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *BarGraph) TextView() *tview.TextView {
	return widget.View
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
