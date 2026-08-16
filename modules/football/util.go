package football

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

func createTable(header []string, buf *bytes.Buffer) *tablewriter.Table {

	table := tablewriter.NewWriter(buf)
	if len(header) != 0 {
		table.SetHeader(header)
	}
	table.SetBorder(false)
	table.SetCenterSeparator(" ")
	table.SetColumnSeparator(" ")
	table.SetRowSeparator(" ")
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	return table
}

func parseDateString(d string) string {

	return fmt.Sprintf("🕙 %s", strings.Replace(d, "T", " ", 1))
}

func getDateString(offset int) string {

	today := time.Now()
	return today.AddDate(0, 0, offset).Format("2006-01-02")

}

// scoreString renders one side's goal count. v4 leaves scores null until the
// match produces them, so a missing value is shown as a dash rather than as
// a misleading 0.
func scoreString(goals *int) string {

	if goals == nil {
		return "-"
	}

	return strconv.Itoa(*goals)
}

// totalTable returns the overall table from a standings response. v4 returns
// TOTAL, HOME and AWAY tables for a league competition; taking the first
// entry unconditionally only picks the overall table by accident of ordering.
func totalTable(standings []Standing) []Table {

	for _, standing := range standings {
		if standing.Type == "" || strings.EqualFold(standing.Type, "TOTAL") {
			return standing.Table
		}
	}

	return nil
}
