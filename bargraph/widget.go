package bargraph

/**************
This is a demo bargraph that just populates some random date/val data
*/

import (
	"github.com/rivo/tview"
	"math/rand"
	"time"

	"github.com/senorprogrammer/wtf/wtf"
)

var started = false
var ok = true

// Widget define wtf widget to register widget later
type Widget struct {
	wtf.BarGraph
}

// NewWidget Make new instance of widget
func NewWidget(app *tview.Application) *Widget {
	widget := Widget{
		BarGraph: wtf.NewBarGraph(app, "Sample Bar Graph", "bargraph", false),
	}

	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// MakeGraph - Load the dead drop stats
func MakeGraph(widget *Widget) {

	//this could come from config
	const lineCount = 8
	var stats [lineCount]wtf.Bar

	barTime := time.Now()
	for i := 0; i < lineCount; i++ {
		barTime = barTime.Add(time.Duration(time.Minute))

		bar := wtf.Bar{
			Label:   barTime.Format("15:04"),
			Percent: rand.Intn(100-5) + 5,
		}

		stats[i] = bar
	}

	widget.BarGraph.BuildBars(stats[:])

}

// Refresh & update after interval time
func (widget *Widget) Refresh() {

	if widget.Disabled() {
		return
	}

	widget.View.Clear()

	display(widget)

}

/* -------------------- Unexported Functions -------------------- */

func display(widget *Widget) {
	MakeGraph(widget)
}
