package gspreadsheets

import (
	"fmt"

	"github.com/wtfutil/wtf/wtf"
	sheets "google.golang.org/api/sheets/v4"
)

type Widget struct {
	wtf.TextWidget

	settings *Settings
}

func NewWidget(refreshChan chan<- string, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(refreshChan, settings.common, false),

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	cells, _ := widget.Fetch()

	widget.View.SetText(widget.contentFrom(cells))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(valueRanges []*sheets.ValueRange) string {
	if valueRanges == nil {
		return "error 1"
	}

	res := ""

	cells := wtf.ToStrs(widget.settings.cellNames)
	for i := 0; i < len(valueRanges); i++ {
		res = res + fmt.Sprintf("%s\t[%s]%s\n", cells[i], widget.settings.colors.values, valueRanges[i].Values[0][0])
	}

	return res
}
