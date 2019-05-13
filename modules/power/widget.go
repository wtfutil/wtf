package power

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	Battery *Battery

	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, pages, settings.common, false),

		Battery: NewBattery(),

		settings: settings,
	}

	widget.SetRefreshFunction(widget.Refresh)

	widget.View.SetWrap(true)

	return &widget
}

func (widget *Widget) Refresh() {
	widget.Battery.Refresh()

	content := ""
	content = content + fmt.Sprintf(" %10s: %s\n", "Source", powerSource())
	content = content + "\n"
	content = content + widget.Battery.String()

	widget.Redraw(widget.CommonSettings.Title, content, true)
}
