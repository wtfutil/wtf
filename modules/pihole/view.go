package pihole

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func getSummaryView(c http.Client, settings *Settings) string {
	var err error

	var s Status

	s, err = getStatus(c, settings)
	if err != nil {
		return err.Error()
	}

	var sb strings.Builder

	buf := new(bytes.Buffer)

	switch strings.ToLower(s.Status) {
	case "disabled":
		sb.WriteString(" [white]Status [red]DISABLED\n")
	case "enabled":
		sb.WriteString(" [white]Status [green]ENABLED\n")
	default:
		sb.WriteString(" [white]Status [yellow]UNKNOWN\n")
	}

	summaryTable := createTable([]string{}, buf)
	summaryTable.Append([]string{"Domain blocklist", s.DomainsBeingBlocked, "Queries today", s.DNSQueriesToday})
	summaryTable.Append([]string{"Ads blocked today", fmt.Sprintf("%s (%s%%)", s.AdsBlockedToday, s.AdsPercentageToday), "Cached queries", s.QueriesCached})
	summaryTable.Append([]string{"Blocklist Age", fmt.Sprintf("%dd %dh %dm", s.GravityLastUpdated.Relative.Days,
		s.GravityLastUpdated.Relative.Hours, s.GravityLastUpdated.Relative.Minutes), "Forwarded queries", s.QueriesForwarded})
	summaryTable.Render()

	sb.WriteString(buf.String())

	return sb.String()
}

func getTopItemsView(c http.Client, settings *Settings) string {
	var err error

	var ti TopItems

	ti, err = getTopItems(c, settings)
	if err != nil {
		return err.Error()
	}

	buf := new(bytes.Buffer)

	var sb strings.Builder

	tiTable := createTable([]string{"Top Queries", "", "Top Ads", ""}, buf)

	largest := len(ti.TopAds)
	if len(ti.TopQueries) > largest {
		largest = len(ti.TopQueries)
	}

	sortedTiQueries := sortMapByIntVal(ti.TopQueries)

	sortedTiAds := sortMapByIntVal(ti.TopAds)

	for x := 0; x < largest; x++ {
		tiQVal := []string{"", ""}
		if len(sortedTiQueries) > x {
			tiQVal = []string{shorten(sortedTiQueries[x][0], settings.maxDomainWidth), sortedTiQueries[x][1]}
		}

		tiAVal := []string{"", ""}

		if len(sortedTiAds) > x {
			tiAVal = []string{shorten(sortedTiAds[x][0], settings.maxDomainWidth), sortedTiAds[x][1]}
		}

		tiTable.Append([]string{tiQVal[0], tiQVal[1], tiAVal[0], tiAVal[1]})
	}

	tiTable.Render()
	sb.WriteString(buf.String())

	return sb.String()
}

func getTopClientsView(c http.Client, settings *Settings) string {
	tc, err := getTopClients(c, settings)
	if err != nil {
		return err.Error()
	}

	var tq QueryTypes

	tq, err = getQueryTypes(c, settings)
	if err != nil {
		return err.Error()
	}

	buf := new(bytes.Buffer)

	tcTable := createTable([]string{"Top Clients", "", "Top Query Types", ""}, buf)

	sortedTcQueries := sortMapByIntVal(tc.TopSources)

	sortedTopQT := sortMapByFloatVal(tq.QueryTypes)

	largest := len(tc.TopSources)

	if len(tq.QueryTypes) > largest {
		largest = len(tq.QueryTypes)
	}

	if settings.showTopClients < largest {
		largest = settings.showTopClients
	}

	for x := 0; x < largest; x++ {
		tcVal := []string{"", ""}

		if len(sortedTcQueries) > x {
			tcVal = []string{sortedTcQueries[x][0], sortedTcQueries[x][1]}
		}

		tqtVal := []string{"", ""}

		if len(sortedTopQT) > x && sortedTopQT[x][0] != "" {
			tqtVal = []string{sortedTopQT[x][0], sortedTopQT[x][1] + "%"}
		}

		tcTable.Append([]string{tcVal[0], tcVal[1], tqtVal[0], tqtVal[1]})
	}

	tcTable.Render()

	var sb strings.Builder

	sb.WriteString(buf.String())

	return sb.String()
}

func shorten(s string, limit int) string {
	if len(s) > limit {
		return s[:limit] + "..."
	}

	return s
}

func createTable(header []string, buf io.Writer) *tablewriter.Table {
	table := tablewriter.NewWriter(buf)

	if len(header) != 0 {
		table.SetHeader(header)
		table.SetHeaderLine(false)
		table.SetHeaderAlignment(0)
	}

	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(true)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetTablePadding(" ")
	table.SetNoWhiteSpace(false)

	return table
}

func sortMapByIntVal(m map[string]int) (sorted [][]string) {
	type kv struct {
		Key   string
		Value int
	}

	ss := make([]kv, len(m))
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss {
		sorted = append(sorted, []string{kv.Key, strconv.Itoa(kv.Value)})
	}

	return
}

func sortMapByFloatVal(m map[string]float32) (sorted [][]string) {
	type kv struct {
		Key   string
		Value float32
	}

	ss := make([]kv, len(m))

	for k, v := range m {
		if k == "" || v == 0.00 {
			continue
		}

		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss {
		sorted = append(sorted, []string{kv.Key, fmt.Sprintf("%.2f", kv.Value)})
	}

	return
}
