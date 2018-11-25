package resourceusage

import (
	"code.cloudfoundry.org/bytefmt"
	"fmt"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"time"
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
		BarGraph: wtf.NewBarGraph(app, "Resource Usage", "resourceusage", false),
	}

	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// MakeGraph - Load the dead drop stats
func MakeGraph(widget *Widget) {

	cpuStats, err := cpu.Percent(time.Duration(0), true)
	if err != nil {
		return
	}

	var stats = make([]wtf.Bar, len(cpuStats)+2)

	for i, stat := range cpuStats {
		bar := wtf.Bar{
			Label:      fmt.Sprint(i),
			Percent:    int(stat),
			ValueLabel: fmt.Sprintf("%d%%", int(stat)),
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

	stats[memIndex] = wtf.Bar{
		Label:      "Mem",
		Percent:    int(memInfo.UsedPercent),
		ValueLabel: fmt.Sprintf("%s/%s", usedMemLabel, totalMemLabel),
	}

	swapIndex := len(cpuStats) + 1
	swapUsed := memInfo.SwapTotal - memInfo.SwapFree
	swapPercent := float64(swapUsed) / float64(memInfo.SwapTotal)

	usedSwapLabel := bytefmt.ByteSize(swapUsed)
	totalSwapLabel := bytefmt.ByteSize(memInfo.SwapTotal)

	if usedSwapLabel[len(usedSwapLabel)-1] == totalMemLabel[len(totalSwapLabel)-1] {
		usedSwapLabel = usedSwapLabel[:len(usedSwapLabel)-1]
	}

	stats[swapIndex] = wtf.Bar{
		Label:      "Swp",
		Percent:    int(swapPercent * 100),
		ValueLabel: fmt.Sprintf("%s/%s", usedSwapLabel, totalSwapLabel),
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
