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

// Widget define wtf widget to register widget later
type Widget struct {
	settings *Settings
	tviewApp *tview.Application
	view.BarGraph
}

// NewWidget Make new instance of widget
func NewWidget(tviewApp *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		BarGraph: view.NewBarGraph(tviewApp, settings.Name, settings.Common),

		tviewApp: tviewApp,
		settings: settings,
	}

	widget.View.SetWrap(false)
	widget.View.SetWordWrap(false)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// MakeGraph - Load the dead drop stats
func MakeGraph(widget *Widget) {
	cpuStats, memInfo := getDataFromSystem(widget)

	var itemsCount = 0
	if widget.settings.showCPU {
		itemsCount += len(cpuStats)
	}

	if widget.settings.showMem {
		itemsCount++
	}

	if widget.settings.showSwp {
		itemsCount++
	}

	var stats = make([]view.Bar, itemsCount)
	var nextIndex = 0

	if widget.settings.showCPU && len(cpuStats) > 0 {
		for i, stat := range cpuStats {
			// Stats sometimes jump outside the 0-100 range, possibly due to timing
			stat = math.Min(100, stat)
			stat = math.Max(0, stat)

			var label string
			if widget.settings.cpuCombined {
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

			stats[nextIndex] = bar
			nextIndex++
		}
	}

	if widget.settings.showMem {
		usedMemLabel := bytefmt.ByteSize(memInfo.Used)
		totalMemLabel := bytefmt.ByteSize(memInfo.Total)

		if usedMemLabel[len(usedMemLabel)-1] == totalMemLabel[len(totalMemLabel)-1] {
			usedMemLabel = usedMemLabel[:len(usedMemLabel)-1]
		}

		stats[nextIndex] = view.Bar{
			Label:      "Mem",
			Percent:    int(memInfo.UsedPercent),
			ValueLabel: fmt.Sprintf("%s/%s", usedMemLabel, totalMemLabel),
			LabelColor: "green",
		}
		nextIndex++
	}

	if widget.settings.showSwp {
		swapUsed := memInfo.SwapTotal - memInfo.SwapFree
		var swapPercent float64
		if memInfo.SwapTotal > 0 {
			swapPercent = float64(swapUsed) / float64(memInfo.SwapTotal)
		}

		usedSwapLabel := bytefmt.ByteSize(swapUsed)
		totalSwapLabel := bytefmt.ByteSize(memInfo.SwapTotal)

		if usedSwapLabel[len(usedSwapLabel)-1] == totalSwapLabel[len(totalSwapLabel)-1] {
			usedSwapLabel = usedSwapLabel[:len(usedSwapLabel)-1]
		}

		stats[nextIndex] = view.Bar{
			Label:      "Swp",
			Percent:    int(swapPercent * 100),
			ValueLabel: fmt.Sprintf("%s/%s", usedSwapLabel, totalSwapLabel),
			LabelColor: "yellow",
		}
	}

	widget.BarGraph.BuildBars(stats)

}

// Refresh & update after interval time
func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.tviewApp.QueueUpdateDraw(func() {
		widget.View.Clear()
		display(widget)
	})
}

/* -------------------- Unexported Functions -------------------- */

func display(widget *Widget) {
	MakeGraph(widget)
}

func getDataFromSystem(widget *Widget) (cpuStats []float64, memInfo mem.VirtualMemoryStat) {
	if widget.settings.showCPU {
		rCPUStats, err := cpu.Percent(time.Duration(0), !widget.settings.cpuCombined)
		if err == nil {
			cpuStats = rCPUStats
		}
	}

	if widget.settings.showMem || widget.settings.showSwp {
		rMemInfo, err := mem.VirtualMemory()
		if err == nil {
			memInfo = *rMemInfo
		}
	}

	return cpuStats, memInfo
}
