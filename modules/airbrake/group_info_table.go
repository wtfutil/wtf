package airbrake

import (
	"fmt"
	"strconv"

	"github.com/wtfutil/wtf/view"
)

type groupInfoTable struct {
	group       *Group
	propertyMap map[string]string

	colWidth0   int
	colWidth1   int
	tableHeight int
}

func newGroupInfoTable(g *Group) *groupInfoTable {
	propTable := &groupInfoTable{
		group: g,

		colWidth0:   20,
		colWidth1:   51,
		tableHeight: 15,
	}

	propTable.propertyMap = propTable.buildPropertyMap()

	return propTable
}

func (propTable *groupInfoTable) buildPropertyMap() map[string]string {
	propMap := map[string]string{}

	g := propTable.group
	if g == nil {
		return propMap
	}
	propMap["1. First Seen"] = g.CreatedAt
	propMap["2. Last Seen"] = g.LastNoticeAt
	propMap["3. Occurrences"] = strconv.Itoa(int(g.NoticeCount))
	propMap["4. Environment"] = g.Context.Environment
	propMap["5. Severity"] = g.Context.Severity
	propMap["6. Muted"] = fmt.Sprintf("%v", g.Muted)
	propMap["7. File"] = g.File()

	return propMap
}

func (propTable *groupInfoTable) render() string {
	tbl := view.NewInfoTable(
		[]string{"Property", "Value"},
		propTable.propertyMap,
		propTable.colWidth0,
		propTable.colWidth1,
		propTable.tableHeight,
	)

	return tbl.Render()
}
