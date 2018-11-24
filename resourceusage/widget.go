package resourceusage

import (
	"code.cloudfoundry.org/bytefmt"
	"fmt"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

var started = false
var ok = true
var prevStats []cpu.TimesStat

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

	cpuStats, err := cpu.Times(true)
	if err != nil {
		return
	}

	var stats = make([]wtf.Bar, len(cpuStats)+2)

	for i, stat := range cpuStats {
		prevStat := stat
		if len(prevStats) > i {
			prevStat = prevStats[i]
		} else {
			prevStats = append(prevStats, stat)
		}

		// based on htop algorithm described here: https://stackoverflow.com/a/23376195/1516085
		prevIdle := prevStat.Idle + prevStat.Iowait
		idle := stat.Idle + stat.Iowait

		prevNonIdle := prevStat.User + prevStat.Nice + prevStat.System + prevStat.Irq + prevStat.Softirq + prevStat.Steal
		nonIdle := stat.User + stat.Nice + stat.System + stat.Irq + stat.Softirq + stat.Steal

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

	//memInfo, err := linux.ReadMemInfo("/proc/meminfo")
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
