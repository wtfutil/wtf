package azurelogs

import (
	"fmt"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"

	"github.com/rivo/tview"

	"github.com/wtfutil/wtf/view"
)

const (
	defaultTableWidth  = 120
	minColumnWidth     = 8
	maxColumnWidth     = 30
	maxDisplayRows     = 50
	truncateMarker     = "..."
	sampleRowsForWidth = 15
)

type Widget struct {
	view.TextWidget
	settings   *Settings
	loading    bool
	lastError  error
	dataLoaded bool
	tableData  *TableResp
}

// NewWidget creates a new instance of a widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, _ *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.Common),
		settings:   settings,
	}

	widget.settings.RefreshInterval = 60 * time.Second

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	// Reset state to allow fresh data fetch
	widget.loading = false
	widget.lastError = nil
	widget.dataLoaded = false
	widget.tableData = nil

	widget.Redraw(widget.content)
}

/* -------------------- Helper Functions -------------------- */

func (widget *Widget) fetchDataAsync() {
	sess, err := Init(to.Ptr(widget.settings.Queryfile))
	if err != nil {
		widget.setError(fmt.Errorf("failed to initialize Azure session: %w", err))
		return
	}

	// Execute Azure query directly
	tableResp, err := RunQuery(sess)
	if err != nil {
		widget.setError(fmt.Errorf("failed to execute Azure query: %w", err))
		return
	}

	// Check if we have valid data structure
	if tableResp == nil || len(tableResp.Header) == 0 {
		widget.setError(fmt.Errorf("no table structure returned from query"))
		return
	}

	// Store the data and mark as loaded
	widget.tableData = tableResp
	widget.dataLoaded = true
	widget.loading = false
	widget.Redraw(widget.content)
}

// setError is a helper function to set error state and trigger redraw
func (widget *Widget) setError(err error) {
	widget.lastError = err
	widget.loading = false
	widget.Redraw(widget.content)
}

func (widget *Widget) renderTable(title string) (string, string, bool) {
	if widget.tableData == nil {
		return title, "[red]Error: No table data available[white]", true
	}

	// Calculate column widths and format table - headers are always shown when available
	colWidths := calculateAdaptiveColumnWidths(widget.tableData, defaultTableWidth)

	var sb strings.Builder
	// Always show headers when we have table structure
	widget.formatTableHeaders(&sb, widget.tableData.Header, colWidths)
	widget.formatTableSeparator(&sb, widget.tableData.Header, colWidths)

	// Show data rows if available, otherwise show informative message
	if len(widget.tableData.Rows) == 0 {
		sb.WriteString("[dim](No data rows returned)[white]\n")
	} else {
		widget.formatTableRows(&sb, widget.tableData.Rows, widget.tableData.Header, colWidths)
	}

	return title, sb.String(), false
}

// formatTableHeaders writes the table header row to the string builder
func (widget *Widget) formatTableHeaders(sb *strings.Builder, headers []string, colWidths []int) {
	for i, header := range headers {
		if i > 0 {
			sb.WriteString(" ¦")
		}
		headerText := header
		if i < len(colWidths) && len(headerText) > colWidths[i] {
			headerText = headerText[:colWidths[i]-len(truncateMarker)] + truncateMarker
		}
		_, _ = fmt.Fprintf(sb, "[lightblue]%-*s[white]", colWidths[i], headerText)
	}
	sb.WriteString("\n")
}

// formatTableSeparator writes the table separator row to the string builder
func (widget *Widget) formatTableSeparator(sb *strings.Builder, headers []string, colWidths []int) {
	for i := range headers {
		if i > 0 {
			sb.WriteString("---")
		}
		sb.WriteString(strings.Repeat("-", colWidths[i]))
	}
	sb.WriteString("\n")
}

// formatTableRows writes the table data rows to the string builder
func (widget *Widget) formatTableRows(sb *strings.Builder, rows []TableRow, headers []string, colWidths []int) {
	maxRows := maxDisplayRows
	rowCount := len(rows)
	if rowCount > maxRows {
		rowCount = maxRows
	}

	for rowIdx := 0; rowIdx < rowCount; rowIdx++ {
		row := rows[rowIdx]
		for colIdx, cell := range row {
			if colIdx >= len(headers) {
				break
			}

			if colIdx > 0 {
				sb.WriteString(" ¦")
			}

			cellText := strings.TrimSpace(cell)
			if colIdx < len(colWidths) && len(cellText) > colWidths[colIdx] {
				cellText = cellText[:colWidths[colIdx]-len(truncateMarker)] + truncateMarker
			}

			_, _ = fmt.Fprintf(sb, "%-*s", colWidths[colIdx], cellText)
		}
		sb.WriteString("\n")
	}

	if len(rows) > maxRows {
		_, _ = fmt.Fprintf(sb, "\n[gray]... (%d more rows truncated for display)[white]\n", len(rows)-maxRows)
	}
}

// calculateAdaptiveColumnWidths computes optimal column widths based on content and available space
func calculateAdaptiveColumnWidths(tr *TableResp, availableWidth int) []int {
	if len(tr.Header) == 0 {
		return []int{}
	}

	// Calculate content-based widths
	widths := make([]int, len(tr.Header))

	// Start with header widths
	for i, header := range tr.Header {
		widths[i] = len(header)
	}

	// Check data rows to find maximum content width per column (if any rows exist)
	if len(tr.Rows) > 0 {
		maxRows := sampleRowsForWidth // Sample first N rows for width calculation
		rowCount := len(tr.Rows)
		if rowCount > maxRows {
			rowCount = maxRows
		}

		for rowIdx := 0; rowIdx < rowCount; rowIdx++ {
			row := tr.Rows[rowIdx]
			for colIdx, cell := range row {
				if colIdx >= len(widths) {
					break
				}
				cellLength := len(strings.TrimSpace(cell))
				if cellLength > widths[colIdx] {
					widths[colIdx] = cellLength
				}
			}
		}
	}

	// Apply minimum and maximum constraints
	totalWidth := 0
	for i := range widths {
		if widths[i] < minColumnWidth {
			widths[i] = minColumnWidth
		}
		if widths[i] > maxColumnWidth {
			widths[i] = maxColumnWidth
		}
		totalWidth += widths[i]
	}

	// Add space for separators: (n-1) * 2 chars for " ¦"
	separatorSpace := (len(widths) - 1) * 2
	totalUsed := totalWidth + separatorSpace

	// If we exceed available width, proportionally reduce columns
	if totalUsed > availableWidth {
		scaleFactor := float64(availableWidth-separatorSpace) / float64(totalWidth)
		for i := range widths {
			widths[i] = int(float64(widths[i]) * scaleFactor)
			if widths[i] < minColumnWidth {
				widths[i] = minColumnWidth
			}
		}
	}

	return widths
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	title := widget.CommonSettings().Title

	// Check if query file is configured
	if widget.settings.Queryfile == "" {
		return title, "[red]Error: queryFile must be configured in widget settings[white]\n\n", false
	}

	// If we have a previous error, show it immediately
	if widget.lastError != nil {
		return title, fmt.Sprintf("[red]Error: %v[white]\n\n[dim]Press 'r' to retry[white]", widget.lastError), true
	}

	// If data is already loaded, show it
	if widget.dataLoaded {
		return widget.renderTable(title)
	}

	// Show loading text while fetching data
	if !widget.loading {
		widget.loading = true

		// Start async data fetch
		go widget.fetchDataAsync()
		return title, "[yellow]Loading Azure Logs data...[white]\n\n[dim]• Initializing Azure session\n• Executing query on workspace\n• Processing results[white]", false
	}

	// Still loading, show loading text
	return title, "[yellow]Loading Azure Logs data...[white]\n\n[dim]• Initializing Azure session\n• Executing query on workspace\n• Processing results[white]", false
}
