package progress

import (
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/cfg"
)

// createTestWidget builds a Widget with the given settings fields, wired up
// enough to exercise the pure calculation helpers.
func createTestWidget(minimum, maximum, current float64, showPercentage string, padding int) *Widget {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)

	settings := &Settings{
		common: &cfg.Common{
			Title: "Test Progress",
		},
		showPercentage: showPercentage,
		padding:        padding,
		minimum:        minimum,
		maximum:        maximum,
		current:        current,
	}

	return NewWidget(app, redrawChan, settings)
}

func TestNewWidget(t *testing.T) {
	widget := createTestWidget(0, 100, 50, "right", 1)

	assert.NotNil(t, widget)
	assert.Equal(t, 0.0, widget.minimum)
	assert.Equal(t, 100.0, widget.maximum)
	assert.Equal(t, 50.0, widget.current)
	assert.Equal(t, " ", widget.padding)
}

func TestCalcPercent(t *testing.T) {
	tests := []struct {
		name     string
		minimum  float64
		maximum  float64
		current  float64
		expected float64
	}{
		{
			name:     "maximum zero, current treated as percentage",
			minimum:  0,
			maximum:  0,
			current:  42,
			expected: 0.42,
		},
		{
			name:     "maximum zero, current at zero",
			minimum:  0,
			maximum:  0,
			current:  0,
			expected: 0,
		},
		{
			name:     "current within range",
			minimum:  0,
			maximum:  200,
			current:  50,
			expected: 0.25,
		},
		{
			name:     "current above maximum caps at 1",
			minimum:  0,
			maximum:  100,
			current:  150,
			expected: 1,
		},
		{
			name:     "current below minimum caps at 0",
			minimum:  10,
			maximum:  100,
			current:  5,
			expected: 0,
		},
		{
			name:     "current equal to maximum",
			minimum:  0,
			maximum:  100,
			current:  100,
			expected: 1,
		},
		{
			name:     "current equal to minimum",
			minimum:  10,
			maximum:  100,
			current:  10,
			expected: 0,
		},
		{
			name:     "non-zero minimum offsets range",
			minimum:  50,
			maximum:  150,
			current:  100,
			expected: 0.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := createTestWidget(tt.minimum, tt.maximum, tt.current, "right", 1)
			widget.calcPercent()
			assert.InDelta(t, tt.expected, widget.percent, 0.0001)
		})
	}
}

func TestFormatPercent(t *testing.T) {
	tests := []struct {
		name           string
		showPercentage string
		percent        float64
		expected       string
	}{
		{
			name:           "left adds trailing space",
			showPercentage: "left",
			percent:        0.5,
			expected:       "50% ",
		},
		{
			name:           "right adds leading space",
			showPercentage: "right",
			percent:        0.5,
			expected:       " 50%",
		},
		{
			name:           "none returns empty string",
			showPercentage: "none",
			percent:        0.5,
			expected:       "",
		},
		{
			name:           "default has no extra spaces",
			showPercentage: "above",
			percent:        0.5,
			expected:       "50%",
		},
		{
			name:           "zero percent",
			showPercentage: "right",
			percent:        0,
			expected:       " 0%",
		},
		{
			name:           "full percent rounds correctly",
			showPercentage: "left",
			percent:        1,
			expected:       "100% ",
		},
		{
			name:           "rounds to nearest whole number",
			showPercentage: "default",
			percent:        0.6666,
			expected:       "67%",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := createTestWidget(0, 100, 0, tt.showPercentage, 1)
			result := widget.formatPercent(tt.percent)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCalcBarWidth(t *testing.T) {
	tests := []struct {
		name           string
		showPercentage string
		padding        int
		percent        string
		innerWidth     int
	}{
		{
			name:           "no percentage display",
			showPercentage: "above",
			padding:        1,
			percent:        "50%",
			innerWidth:     40,
		},
		{
			name:           "left percentage reduces width",
			showPercentage: "left",
			padding:        1,
			percent:        "50% ",
			innerWidth:     40,
		},
		{
			name:           "right percentage reduces width",
			showPercentage: "right",
			padding:        2,
			percent:        " 100%",
			innerWidth:     60,
		},
		{
			name:           "zero padding",
			showPercentage: "none",
			padding:        0,
			percent:        "",
			innerWidth:     20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := createTestWidget(0, 100, 0, tt.showPercentage, tt.padding)
			widget.View.SetRect(0, 0, tt.innerWidth, 3)

			width := widget.calcBarWidth(tt.percent)

			expected := tt.innerWidth - tt.padding*2
			if tt.showPercentage == "left" || tt.showPercentage == "right" {
				expected -= len(tt.percent)
			}
			assert.Equal(t, expected, width)
		})
	}
}
