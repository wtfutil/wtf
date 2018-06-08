package gspreadsheets

import (
	"fmt"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
	sheets "google.golang.org/api/sheets/v4"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Google Spreadsheets ", "gspreadsheets", false),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	cells, _ := Fetch()

	widget.UpdateRefreshedAt()

	widget.View.SetText(fmt.Sprintf("%s", widget.contentFrom(cells)))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(valueRanges []*sheets.ValueRange) string {
	if valueRanges == nil {
		return "error 1"
	}

	valuesColor := Config.UString("wtf.mods.gspreadsheets.colors.values", "green")
	res := ""

	cells := wtf.ToStrs(Config.UList("wtf.mods.gspreadsheets.cells.names"))
	for i := 0; i < len(valueRanges); i++ {
		res = res + fmt.Sprintf("%s\t[%s]%s\n", cells[i], valuesColor, valueRanges[i].Values[0][0])
	}

	return res
}
