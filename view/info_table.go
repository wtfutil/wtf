package view

import (
	"bytes"
	"sort"

	"github.com/olekukonko/tablewriter"
)

/*
	An InfoTable is a two-column table of properties/values:

	-------------------------- -------------------------------------------------
			 PROPERTY                                VALUE
	-------------------------- -------------------------------------------------
	 CPUs                       1
	 Created                    2019-12-12T18:39:09Z
	 Disk                       25
	 Features                   ipv6
	 Image                      18.04.3 (LTS) x64 (Ubuntu)
	 Memory                     1024
	 Region                     Toronto 1 (tor1)
	-------------------------- -------------------------------------------------
*/

// InfoTable contains the internal guts of an InfoTable
type InfoTable struct {
	buf       *bytes.Buffer
	tblWriter *tablewriter.Table
}

// NewInfoTable creates and returns the stringified contents of a two-column table
func NewInfoTable(headers []string, dataMap map[string]string, colWidth0, colWidth1, tableHeight int) *InfoTable {
	tbl := &InfoTable{
		buf: new(bytes.Buffer),
	}

	tbl.tblWriter = tablewriter.NewWriter(tbl.buf)

	tbl.tblWriter.SetHeader(headers)
	tbl.tblWriter.SetBorder(true)
	tbl.tblWriter.SetCenterSeparator(" ")
	tbl.tblWriter.SetColumnSeparator(" ")
	tbl.tblWriter.SetRowSeparator("-")
	tbl.tblWriter.SetAlignment(tablewriter.ALIGN_LEFT)
	tbl.tblWriter.SetColMinWidth(0, colWidth0)
	tbl.tblWriter.SetColMinWidth(1, colWidth1)

	keys := []string{}
	for key := range dataMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	// Enumerate over the alphabetically-sorted keys to render the property values
	for _, key := range keys {
		tbl.tblWriter.Append([]string{key, dataMap[key]})
	}

	// Pad the table with extra rows to push it to the bottom
	paddingAmt := tableHeight - len(dataMap) - 1
	if paddingAmt > 0 {
		for i := 0; i < paddingAmt; i++ {
			tbl.tblWriter.Append([]string{"", ""})
		}
	}

	return tbl
}

// Render returns the stringified version of the table
func (tbl *InfoTable) Render() string {
	tbl.tblWriter.Render()
	return tbl.buf.String()
}
