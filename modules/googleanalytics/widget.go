package googleanalytics

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	settings *Settings
}

func NewWidget(tviewApp *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),

		settings: settings,
	}

	return &widget
}

func (widget *Widget) Refresh() {
	websiteReports := widget.Fetch()
	contentTable := widget.createTable(websiteReports)

	widget.Redraw(func() (string, string, bool) { return widget.CommonSettings().Title, contentTable, false })
}
