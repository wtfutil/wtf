package gspreadsheets

import (
	"fmt"

	"github.com/senorprogrammer/wtf/wtf"
	sheets "google.golang.org/api/sheets/v4"
)

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget("Google Spreadsheets", "gspreadsheets", false),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	cells, _ := Fetch()

	widget.UpdateRefreshedAt()

	widget.View.SetText(widget.contentFrom(cells))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(valueRanges []*sheets.ValueRange) string {
	if valueRanges == nil {
		return "error 1"
	}

	valuesColor := wtf.Config.UString("wtf.mods.gspreadsheets.colors.values", "green")
	res := ""

	cells := wtf.ToStrs(wtf.Config.UList("wtf.mods.gspreadsheets.cells.names"))
	for i := 0; i < len(valueRanges); i++ {
		res = res + fmt.Sprintf("%s\t[%s]%s\n", cells[i], valuesColor, valueRanges[i].Values[0][0])
	}

	return res
}
