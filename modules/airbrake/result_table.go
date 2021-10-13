package airbrake

import (
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type resultTable struct {
	propertyMap map[string]string

	colWidth0   int
	colWidth1   int
	tableHeight int
}

func newResultTable(result, message string) *resultTable {
	propTable := &resultTable{
		colWidth0:   20,
		colWidth1:   51,
		tableHeight: 15,
	}
	propTable.propertyMap = map[string]string{result: message}

	return propTable
}

func (propTable *resultTable) render() string {
	tbl := view.NewInfoTable(
		[]string{"Result", "Message"},
		propTable.propertyMap,
		propTable.colWidth0,
		propTable.colWidth1,
		propTable.tableHeight,
	)

	return tbl.Render() + utils.CenterText("Esc to close", 80)
}
