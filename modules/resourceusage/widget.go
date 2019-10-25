package resourceusage

import (
	"fmt"
	"math"
	"time"

	"code.cloudfoundry.org/bytefmt"
	"github.com/rivo/tview"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/wtfutil/wtf/view"
)

var (
	ok      = true
	started = false
)

// Widget define wtf widget to register widget later
type Widget struct {
	view.BarGraph

	app      *tview.Application
	settings *Settings
}

// NewWidget Make new instance of widget
func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		BarGraph: view.NewBarGraph(app, settings.common.Name, settings.common),

		app:      app,
		settings: settings,
	}

	widget.View.SetWrap(false)
	widget.View.SetWordWrap(false)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// MakeGraph - Load the dead drop stats
func MakeGraph(widget *Widget) {

	cpuStats, err := cpu.Percent(time.Duration(0), !widget.settings.cpuCombined)
	if err != nil {
		return
	}

	var stats = make([]view.Bar, len(cpuStats)+2)

	for i, stat := range cpuStats {
		// Stats sometimes jump outside the 0-100 range, possibly due to timing
		stat = math.Min(100, stat)
		stat = math.Max(0, stat)

		var label string
		if (widget.settings.cpuCombined) {
			label = "CPU"
		} else {
			label = fmt.Sprint(i)
		}
		
		bar := view.Bar{
			Label:      label,
			Percent:    int(stat),
			ValueLabel: fmt.Sprintf("%d%%", int(stat)),
			LabelColor: "red",
		}

		stats[i] = bar
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return
	}

	memIndex := len(cpuStats)

	usedMemLabel := bytefmt.ByteSize(memInfo.Used)
	totalMemLabel := bytefmt.ByteSize(memInfo.Total)

	if usedMemLabel[len(usedMemLabel)-1] == totalMemLabel[len(totalMemLabel)-1] {
		usedMemLabel = usedMemLabel[:len(usedMemLabel)-1]
	}

	stats[memIndex] = view.Bar{
		Label:      "Mem",
		Percent:    int(memInfo.UsedPercent),
		ValueLabel: fmt.Sprintf("%s/%s", usedMemLabel, totalMemLabel),
		LabelColor: "green",
	}

	swapIndex := len(cpuStats) + 1
	swapUsed := memInfo.SwapTotal - memInfo.SwapFree
	var swapPercent float64
	if memInfo.SwapTotal > 0 {
		swapPercent = float64(swapUsed) / float64(memInfo.SwapTotal)
	}

	usedSwapLabel := bytefmt.ByteSize(swapUsed)
	totalSwapLabel := bytefmt.ByteSize(memInfo.SwapTotal)

	if usedSwapLabel[len(usedSwapLabel)-1] == totalMemLabel[len(totalSwapLabel)-1] {
		usedSwapLabel = usedSwapLabel[:len(usedSwapLabel)-1]
	}

	stats[swapIndex] = view.Bar{
		Label:      "Swp",
		Percent:    int(swapPercent * 100),
		ValueLabel: fmt.Sprintf("%s/%s", usedSwapLabel, totalSwapLabel),
		LabelColor: "yellow",
	}

	widget.BarGraph.BuildBars(stats[:])

}

// Refresh & update after interval time
func (widget *Widget) Refresh() {

	if widget.Disabled() {
		return
	}

	widget.app.QueueUpdateDraw(func() {
		widget.View.Clear()	
		display(widget)
	})
}

/* -------------------- Unexported Functions -------------------- */

func display(widget *Widget) {
	MakeGraph(widget)
}
