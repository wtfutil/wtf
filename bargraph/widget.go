package bargraph

/**************
This is a demo bargraph that just populates some random date/val data
*/

import (
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
func NewWidget() *Widget {
	widget := Widget{
		BarGraph: wtf.NewBarGraph("Sample Bar Graph", "bargraph", false),
	}

	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// MakeGraph - Load the dead drop stats
func MakeGraph(widget *Widget) {

	//this could come from config
	const lineCount = 20
	var stats [lineCount][2]int64

	for i := lineCount - 1; i >= 0; i-- {

		stats[i][1] = time.Now().AddDate(0, 0, i*-1).Unix() * 1000
		stats[i][0] = int64(rand.Intn(120-5) + 5)

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
