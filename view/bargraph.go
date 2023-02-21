package view

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/wtf"
)

// BarGraph defines the data required to make a bar graph
type BarGraph struct {
	maxStars int
	starChar string

	*Base
	*KeyboardWidget

	View *tview.TextView
}

// Bar defines a single row in the bar graph
type Bar struct {
	Label      string
	Percent    int
	ValueLabel string
	LabelColor string
}

// NewBarGraph creates and returns an instance of BarGraph
func NewBarGraph(tviewApp *tview.Application, redrawChan chan bool, _ string, commonSettings *cfg.Common) BarGraph {
	widget := BarGraph{
		Base:           NewBase(tviewApp, redrawChan, nil, commonSettings),
		KeyboardWidget: NewKeyboardWidget(commonSettings),

		maxStars: commonSettings.Config.UInt("graphStars", 20),
		starChar: commonSettings.Config.UString("graphIcon", "|"),
	}

	widget.View = widget.createView(widget.bordered)

	return widget
}

/* -------------------- Exported Functions -------------------- */

// BuildBars will build a string of * to represent your data of [time][value]
// time should be passed as a int64
func (widget *BarGraph) BuildBars(data []Bar) {
	widget.View.SetText(BuildStars(data, widget.maxStars, widget.starChar))
	widget.Base.RedrawChan <- true
}

// BuildStars build the string to display
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
		if label == "" {
			label = fmt.Sprint(bar.Percent)
		}

		labelColor := bar.LabelColor
		if labelColor == "" {
			labelColor = "default"
		}

		//write the line
		_, err := buffer.WriteString(
			fmt.Sprintf(
				"%s%s[[%s]%s[default]%s] %s\n",
				bar.Label,
				strings.Repeat(" ", longestLabel-len(bar.Label)),
				labelColor,
				strings.Repeat(starChar, starCount),
				strings.Repeat(" ", maxStars-starCount),
				label,
			),
		)
		if err != nil {
			return ""
		}
	}

	return buffer.String()
}

func (widget *BarGraph) TextView() *tview.TextView {
	return widget.View
}

/* -------------------- Unexported Functions -------------------- */

func (widget *BarGraph) createView(bordered bool) *tview.TextView {
	view := tview.NewTextView()

	view.SetBackgroundColor(wtf.ColorFor(widget.commonSettings.Colors.WidgetTheme.Background))
	view.SetBorder(bordered)
	view.SetBorderColor(wtf.ColorFor(widget.BorderColor()))
	view.SetDynamicColors(true)
	view.SetTitle(widget.ContextualTitle(widget.CommonSettings().Title))
	view.SetTitleColor(wtf.ColorFor(widget.commonSettings.Colors.TextTheme.Title))
	view.SetWrap(false)

	return view
}
