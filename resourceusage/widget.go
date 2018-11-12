package resourceusage

import (
	"code.cloudfoundry.org/bytefmt"
	"fmt"
	"github.com/c9s/goprocinfo/linux"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

var started = false
var ok = true
var prevStats []linux.CPUStat

// Widget define wtf widget to register widget later
type Widget struct {
	wtf.BarGraph
}

// NewWidget Make new instance of widget
func NewWidget(app *tview.Application) *Widget {
	widget := Widget{
		BarGraph: wtf.NewBarGraph(app,  "Resource Usage", "resourceusage", false),
	}

	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// MakeGraph - Load the dead drop stats
func MakeGraph(widget *Widget) {

	cpuStat, err := linux.ReadStat("/proc/stat")
	if err != nil {
		return
	}

	var stats = make([]wtf.Bar, len(cpuStat.CPUStats)+2)

	for i, stat := range cpuStat.CPUStats {
		prevStat := stat
		if len(prevStats) > i {
			prevStat = prevStats[i]
		} else {
			prevStats = append(prevStats, stat)
		}

		// based on htop algorithm described here: https://stackoverflow.com/a/23376195/1516085
		prevIdle := prevStat.Idle + prevStat.IOWait
		idle := stat.Idle + stat.IOWait

		prevNonIdle := prevStat.User + prevStat.Nice + prevStat.System + prevStat.IRQ + prevStat.SoftIRQ + prevStat.Steal
		nonIdle := stat.User + stat.Nice + stat.System + stat.IRQ + stat.SoftIRQ + stat.Steal

		prevTotal := prevIdle + prevNonIdle
		total := idle + nonIdle

		// differentiate: actual value minus the previous one
		difference := total - prevTotal
		idled := idle - prevIdle

		percentage := float64(0)
		if difference > 0 {
			percentage = float64(difference-idled) / float64(difference)
		}

		bar := wtf.Bar{
			Label:      fmt.Sprint(i),
			Percent:    int(percentage * 100),
			ValueLabel: fmt.Sprintf("%d%%", int(percentage*100)),
		}

		stats[i] = bar
		prevStats[i] = stat
	}

	memInfo, err := linux.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return
	}

	memIndex := len(cpuStat.CPUStats)
	memUsed := memInfo.MemTotal - memInfo.MemAvailable
	memPercent := float64(memUsed) / float64(memInfo.MemTotal)

	usedMemLabel := bytefmt.ByteSize(memUsed * bytefmt.KILOBYTE)
	totalMemLabel := bytefmt.ByteSize(memInfo.MemTotal * bytefmt.KILOBYTE)

	if usedMemLabel[len(usedMemLabel)-1] == totalMemLabel[len(totalMemLabel)-1] {
		usedMemLabel = usedMemLabel[:len(usedMemLabel)-1]
	}

	stats[memIndex] = wtf.Bar{
		Label:      "Mem",
		Percent:    int(memPercent * 100),
		ValueLabel: fmt.Sprintf("%s/%s", usedMemLabel, totalMemLabel),
	}

	swapIndex := len(cpuStat.CPUStats) + 1
	swapUsed := memInfo.SwapTotal - memInfo.SwapFree
	swapPercent := float64(swapUsed) / float64(memInfo.SwapTotal)

	usedSwapLabel := bytefmt.ByteSize(swapUsed * bytefmt.KILOBYTE)
	totalSwapLabel := bytefmt.ByteSize(memInfo.SwapTotal * bytefmt.KILOBYTE)

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
