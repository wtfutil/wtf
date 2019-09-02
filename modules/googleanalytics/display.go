package googleanalytics

import (
	"fmt"
	"strings"
	"time"
)

func (widget *Widget) createTable(websiteReports []websiteReport) string {
	content := ""

	if len(websiteReports) == 0 {
		return content
	}

	if websiteReports[0].RealtimeReport != nil {
		content += "Realtime Visitor Counts\n"
		for _, websiteReport := range websiteReports {
			websiteRow := fmt.Sprintf(" %-20s", websiteReport.Name)

			if websiteReport.RealtimeReport == nil {
				websiteRow += fmt.Sprintf("No data found for given ViewId.")
			} else {
				if len(websiteReport.RealtimeReport.Rows) == 0 {
					websiteRow += "-"
				} else {
					websiteRow += fmt.Sprintf("%-10s", websiteReport.RealtimeReport.Rows[0][0])
				}
			}

			content += websiteRow + "\n"
		}

		content += "\n"
		content += "Historical Visitor Counts\n"
	}

	content += widget.createHeader()

	for _, websiteReport := range websiteReports {
		websiteRow := ""

		for _, report := range websiteReport.Report.Reports {
			websiteRow += fmt.Sprintf(" %-20s", websiteReport.Name)
			reportRows := report.Data.Rows
			noDataMonth := widget.settings.months - len(reportRows)

			// Fill in requested months with no data from query
			if noDataMonth > 0 {
				websiteRow += strings.Repeat("-         ", noDataMonth)
			}

			if reportRows == nil {
				websiteRow += fmt.Sprintf("No data found for given ViewId.")
			} else {
				for _, row := range reportRows {
					metrics := row.Metrics

					for _, metric := range metrics {
						websiteRow += fmt.Sprintf("%-10s", metric.Values[0])
					}
				}
			}

			content += websiteRow + "\n"
		}
	}

	return content
}

func (widget *Widget) createHeader() string {
	// Creates the table header of consisting of Months
	currentMonth := int(time.Now().Month())
	widgetStartMonth := currentMonth - widget.settings.months + 1
	header := "                     "

	for i := widgetStartMonth; i < currentMonth+1; i++ {
		header += fmt.Sprintf("%-10s", time.Month(i))
	}
	header += "\n"

	return header
}
