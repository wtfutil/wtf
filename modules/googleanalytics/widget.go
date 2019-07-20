package googleanalytics

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	settings *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, settings.common, false),

		settings: settings,
	}

	return &widget
}

func (widget *Widget) Refresh() {
	websiteReports := widget.Fetch()
	contentTable := widget.createTable(websiteReports)

	widget.Redraw(widget.CommonSettings().Title, contentTable, false)
}
