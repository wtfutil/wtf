package flighty

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/rivo/tview"
)

func (widget *Widget) display() {
	widget.Redraw(widget.content)
}

// This uses Bubbletea's Table to render the data (https://github.com/charmbracelet/bubbletea)
// This is not an efficient or effective use of this; do not take this as cannon
func (widget *Widget) content() (string, string, bool) {
	title := widget.CommonSettings().Title
	setWrap := false

	cols := []table.Column{
		{Title: "Owner", Width: 16},
		{Title: "ICAO24", Width: 6},
		{Title: "Call Sign", Width: 10},
		{Title: "Departed", Width: 10},
		{Title: "Arriving", Width: 10},
	}

	rows := []table.Row{}
	for _, flight := range widget.Data {
		row := table.Row{
			flight.Owner,
			flight.ICAO24,
			flight.CallSign,
			flight.EstDepartureAirport,
			flight.EstArrivalAirport,
		}
		rows = append(rows, row)
	}

	tbl := table.New(
		table.WithColumns(cols),
		table.WithRows(rows),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected.Reverse(true)

	tbl.SetStyles(s)

	return title, tview.TranslateANSI(tbl.View()), setWrap
}
