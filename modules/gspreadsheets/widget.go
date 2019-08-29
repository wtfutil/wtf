package gspreadsheets

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
	sheets "google.golang.org/api/sheets/v4"
)

type Widget struct {
	view.TextWidget

	settings *Settings
	cells    []*sheets.ValueRange
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, settings.common, false),

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	cells, _ := widget.Fetch()
	widget.cells = cells

	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	if widget.cells == nil {
		return widget.CommonSettings().Title, "error 1", false
	}

	res := ""

	cells := utils.ToStrs(widget.settings.cellNames)
	for i := 0; i < len(widget.cells); i++ {
		res += fmt.Sprintf("%s\t[%s]%s\n", cells[i], widget.settings.colors.values, widget.cells[i].Values[0][0])
	}

	return widget.CommonSettings().Title, res, false
}
