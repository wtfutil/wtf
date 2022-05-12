package power

import (
	"fmt"
	"runtime"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

const (
	msgNoBattery       = " no battery found"
	productNameTrimLen = 14
)

type Widget struct {
	view.TextWidget

	Battery        *Battery
	ManagedDevices *ManagedDevices

	settings *Settings
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.Common),

		Battery:        NewBattery(),
		ManagedDevices: NewManagedDevices(),

		settings: settings,
	}

	widget.View.SetWrap(true)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.Battery.Refresh()

	// Handle the reading of connected battery-driven devices
	switch runtime.GOOS {
	case "darwin":
		widget.ManagedDevices.Refresh()
	case "linux":
	case "windows":
	default:
	}

	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	content := fmt.Sprintf(" %14s: %s\n", "Source", powerSource())

	if widget.Battery.String() != msgNoBattery {
		content += widget.Battery.String()
		content += "\n"
	}

	content += "\n"

	for _, manDev := range widget.ManagedDevices.Devices {
		if manDev.HasBattery() {
			percent := utils.ColorizePercent(float64(manDev.BatteryPercent()))

			content += fmt.Sprintf(" %s: %s", manDev.Product()[:productNameTrimLen], percent)
			content += "\n"
		}
	}

	return widget.CommonSettings().Title, content, true
}
