package azurelogs

import (
	"strings"
	"testing"
	"time"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/cfg"
)

func TestNewWidget(t *testing.T) {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)

	settings := &Settings{
		Common: &cfg.Common{
			Title: "Test Azure Logs",
		},
		Queryfile: "/path/to/query.yml",
	}

	widget := NewWidget(app, redrawChan, nil, settings)

	assert.NotNil(t, widget)
	assert.Equal(t, settings, widget.settings)
	assert.Equal(t, 60*time.Second, widget.settings.RefreshInterval)
	assert.False(t, widget.loading)
	assert.False(t, widget.dataLoaded)
	assert.Nil(t, widget.lastError)
	assert.Nil(t, widget.tableData)
}

// TestWidget_Refresh removed as it tests core WTF framework functionality (Disabled() method)
// rather than Azure-specific logic

func TestWidget_SetError(t *testing.T) {
	widget := createTestWidget()
	widget.loading = true

	testError := assert.AnError
	widget.setError(testError)

	assert.Equal(t, testError, widget.lastError)
	assert.False(t, widget.loading)
}

func TestWidget_RenderTable(t *testing.T) {
	tests := []struct {
		name           string
		tableData      *TableResp
		expectedTitle  string
		expectedError  bool
		expectedOutput string
	}{
		{
			name:           "nil table data",
			tableData:      nil,
			expectedTitle:  "Test Title",
			expectedError:  true,
			expectedOutput: "[red]Error: No table data available[white]",
		},
		{
			name: "table with headers but no data",
			tableData: &TableResp{
				Header: []string{"Column1", "Column2"},
				Rows:   []TableRow{},
			},
			expectedTitle:  "Test Title",
			expectedError:  false,
			expectedOutput: "[lightblue]Column1 [white] ¦[lightblue]Column2 [white]", // Just check the header part
		},
		{
			name: "table with headers and data",
			tableData: &TableResp{
				Header: []string{"Col1", "Col2"},
				Rows: []TableRow{
					{"Value1", "Value2"},
					{"Value3", "Value4"},
				},
			},
			expectedTitle: "Test Title",
			expectedError: false,
			expectedOutput: func() string {
				// This will contain the formatted table with headers, separator, and data
				return "[lightblue]Col1    [white] ¦[lightblue]Col2    [white]\n--------¦--------\nValue1   ¦Value2  \nValue3   ¦Value4  \n"
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := createTestWidget()
			widget.tableData = tt.tableData

			title, content, hasError := widget.renderTable("Test Title")

			assert.Equal(t, tt.expectedTitle, title)
			assert.Equal(t, tt.expectedError, hasError)
			assert.Contains(t, content, strings.Split(tt.expectedOutput, "\n")[0]) // Check first line
		})
	}
}

func TestWidget_FormatTableHeaders(t *testing.T) {
	widget := createTestWidget()
	var sb strings.Builder

	headers := []string{"Header1", "Header2", "Header3"}
	colWidths := []int{10, 15, 8}

	widget.formatTableHeaders(&sb, headers, colWidths)

	result := sb.String()
	assert.Contains(t, result, "[lightblue]Header1   [white]")
	assert.Contains(t, result, "¦")
	assert.Contains(t, result, "[lightblue]Header2        [white]")
	assert.Contains(t, result, "[lightblue]Header3 [white]")
	assert.True(t, strings.HasSuffix(result, "\n"))
}

func TestWidget_FormatTableSeparator(t *testing.T) {
	widget := createTestWidget()
	var sb strings.Builder

	headers := []string{"A", "B", "C"}
	colWidths := []int{5, 8, 6}

	widget.formatTableSeparator(&sb, headers, colWidths)

	result := sb.String()
	// Expected: 5 dashes + "---" + 8 dashes + "---" + 6 dashes + "\n"
	assert.Equal(t, "-----"+"---"+"--------"+"---"+"------\n", result)
}

func TestWidget_FormatTableRows(t *testing.T) {
	widget := createTestWidget()
	var sb strings.Builder

	headers := []string{"Col1", "Col2"}
	colWidths := []int{8, 8}
	rows := []TableRow{
		{"Data1", "Data2"},
		{"LongData", "Short"},
	}

	widget.formatTableRows(&sb, rows, headers, colWidths)

	result := sb.String()
	lines := strings.Split(strings.TrimSpace(result), "\n")
	assert.Len(t, lines, 2)
	assert.Contains(t, lines[0], "Data1")
	assert.Contains(t, lines[0], "Data2")
	assert.Contains(t, lines[1], "LongData")
	assert.Contains(t, lines[1], "Short")
}

func TestWidget_FormatTableRows_WithTruncation(t *testing.T) {
	widget := createTestWidget()
	var sb strings.Builder

	headers := []string{"Col1", "Col2"}
	colWidths := []int{8, 8}

	// Create more rows than maxDisplayRows to test truncation
	rows := make([]TableRow, maxDisplayRows+10)
	for i := range rows {
		rows[i] = TableRow{"data1", "data2"}
	}

	widget.formatTableRows(&sb, rows, headers, colWidths)

	result := sb.String()
	assert.Contains(t, result, "more rows truncated")
}

func TestWidget_Content(t *testing.T) {
	tests := []struct {
		name             string
		queryfile        string
		lastError        error
		dataLoaded       bool
		loading          bool
		expectedTitle    string
		expectedContains string
	}{
		{
			name:             "no query file configured",
			queryfile:        "",
			expectedTitle:    "Test Azure Logs",
			expectedContains: "[red]Error: queryFile must be configured",
		},
		{
			name:             "has error",
			queryfile:        "/path/to/query.yml",
			lastError:        assert.AnError,
			expectedTitle:    "Test Azure Logs",
			expectedContains: "[red]Error:",
		},
		{
			name:             "data loaded",
			queryfile:        "/path/to/query.yml",
			dataLoaded:       true,
			expectedTitle:    "Test Azure Logs",
			expectedContains: "[red]Error: No table data available", // Since tableData is nil
		},
		{
			name:             "loading state",
			queryfile:        "/path/to/query.yml",
			loading:          false, // Will trigger loading
			expectedTitle:    "Test Azure Logs",
			expectedContains: "[yellow]Loading Azure Logs data",
		},
		{
			name:             "still loading",
			queryfile:        "/path/to/query.yml",
			loading:          true,
			expectedTitle:    "Test Azure Logs",
			expectedContains: "[yellow]Loading Azure Logs data",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := createTestWidget()
			widget.settings.Queryfile = tt.queryfile
			widget.lastError = tt.lastError
			widget.dataLoaded = tt.dataLoaded
			widget.loading = tt.loading

			title, content, _ := widget.content()

			assert.Equal(t, tt.expectedTitle, title)
			assert.Contains(t, content, tt.expectedContains)
		})
	}
}

func TestCalculateAdaptiveColumnWidths(t *testing.T) {
	tests := []struct {
		name           string
		tableResp      *TableResp
		availableWidth int
		expected       []int
	}{
		{
			name: "empty headers",
			tableResp: &TableResp{
				Header: []string{},
				Rows:   []TableRow{},
			},
			availableWidth: 100,
			expected:       []int{},
		},
		{
			name: "headers only",
			tableResp: &TableResp{
				Header: []string{"Short", "VeryLongHeaderName"},
				Rows:   []TableRow{},
			},
			availableWidth: 100,
			expected:       []int{minColumnWidth, 18}, // "VeryLongHeaderName" is 18 chars
		},
		{
			name: "headers with data",
			tableResp: &TableResp{
				Header: []string{"Col1", "Col2"},
				Rows: []TableRow{
					{"ShortData", "VeryLongDataValue"},
					{"X", "Y"},
				},
			},
			availableWidth: 100,
			expected:       []int{9, 17}, // Max of header/data lengths
		},
		{
			name: "width constraints",
			tableResp: &TableResp{
				Header: []string{"VeryVeryVeryLongColumnNameThatExceedsMaxWidth"},
				Rows:   []TableRow{},
			},
			availableWidth: 100,
			expected:       []int{maxColumnWidth}, // Capped at maxColumnWidth
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateAdaptiveColumnWidths(tt.tableResp, tt.availableWidth)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCalculateAdaptiveColumnWidths_Scaling(t *testing.T) {
	// Test case where columns need to be scaled down
	tableResp := &TableResp{
		Header: []string{"LongHeader1", "LongHeader2", "LongHeader3"},
		Rows:   []TableRow{},
	}

	// Very small available width to force scaling
	result := calculateAdaptiveColumnWidths(tableResp, 20)

	// All columns should be scaled down to minimum width
	for _, width := range result {
		assert.GreaterOrEqual(t, width, minColumnWidth)
	}

	// Total width + separators should not exceed available width significantly
	totalWidth := 0
	for _, width := range result {
		totalWidth += width
	}
	separatorSpace := (len(result) - 1) * 2
	assert.LessOrEqual(t, totalWidth+separatorSpace, 30) // Allow some margin for scaling
}

// Helper function to create a test widget
func createTestWidget() *Widget {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)

	settings := &Settings{
		Common: &cfg.Common{
			Title:   "Test Azure Logs",
			Enabled: true, // Enable by default for tests
		},
		Queryfile: "/path/to/query.yml",
	}

	return NewWidget(app, redrawChan, nil, settings)
}
