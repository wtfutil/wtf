package power

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	Battery *Battery

	app      *tview.Application
	settings *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, settings.common, false),

		Battery: NewBattery(),

		app:      app,
		settings: settings,
	}

	widget.View.SetWrap(true)

	return &widget
}

func (widget *Widget) Refresh() {
	widget.Battery.Refresh()

	content := ""
	content = content + fmt.Sprintf(" %10s: %s\n", "Source", powerSource())
	content = content + "\n"
	content = content + widget.Battery.String()

	widget.app.QueueUpdateDraw(func() {
		widget.View.SetText(content)
	})
}
