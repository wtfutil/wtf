package power

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

const (
	msgNoBattery = " no battery found"
)

type Widget struct {
	view.TextWidget

	Battery *Battery

	settings *Settings
}

func NewWidget(tviewApp *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),

		Battery: NewBattery(),

		settings: settings,
	}

	widget.View.SetWrap(true)

	return &widget
}

func (widget *Widget) content() (string, string, bool) {
	content := fmt.Sprintf(" %10s: %s\n", "Source", powerSource())
	content += "\n"

	if widget.Battery.String() == msgNoBattery {
		content += "[grey]"
		content += widget.Battery.String()
		content += "[white]"
	} else {
		content += widget.Battery.String()
	}

	return widget.CommonSettings().Title, content, true
}

func (widget *Widget) Refresh() {
	widget.Battery.Refresh()
	widget.Redraw(widget.content)
}
