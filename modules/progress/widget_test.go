package progress

import (
	"os/exec"
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
		colors: colors{
			gradientA: "#56ab2f",
			gradientB: "#a8e063",
		},
	}

	widget := NewWidget(app, redrawChan, settings)
	widget.View.SetRect(0, 0, 40, 3)

	return widget
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
			// maximum == 0 branch: the current > 100 guard sets percent to 1,
			// but it is then unconditionally overwritten by current/100.
			name:     "maximum zero, current above 100",
			minimum:  0,
			maximum:  0,
			current:  150,
			expected: 1.5,
		},
		{
			// maximum == 0 branch: the current < 0 guard sets percent to 0,
			// but it is then unconditionally overwritten by current/100.
			name:     "maximum zero, current below zero",
			minimum:  0,
			maximum:  0,
			current:  -10,
			expected: -0.1,
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

func TestBuildProgressBar(t *testing.T) {
	tests := []struct {
		name  string
		solid string
	}{
		{name: "gradient fill", solid: ""},
		{name: "solid fill", solid: "#ff0000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := createTestWidget(0, 100, 50, "right", 1)
			widget.settings.solid = tt.solid

			bar := widget.buildProgressBar(" 50%")

			assert.NotNil(t, bar)
			assert.Equal(t, widget.calcBarWidth(" 50%"), bar.Width)
		})
	}
}

func TestContent(t *testing.T) {
	tests := []struct {
		name             string
		showPercentage   string
		err              error
		expectedContains string
		expectedAbsent   string
	}{
		{
			name:             "error takes precedence",
			showPercentage:   "right",
			err:              assert.AnError,
			expectedContains: "[red]Error: " + assert.AnError.Error(),
		},
		{
			name:             "left percentage",
			showPercentage:   "left",
			expectedContains: "50% ",
		},
		{
			name:             "right percentage",
			showPercentage:   "right",
			expectedContains: " 50%",
		},
		{
			name:             "above percentage",
			showPercentage:   "above",
			expectedContains: "50%",
		},
		{
			name:             "below percentage",
			showPercentage:   "below",
			expectedContains: "50%",
		},
		{
			name:           "none hides percentage text",
			showPercentage: "none",
			expectedAbsent: "50%",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := createTestWidget(0, 100, 50, tt.showPercentage, 1)
			widget.percent = 0.5
			widget.err = tt.err

			content := widget.content()

			if tt.expectedContains != "" {
				assert.Contains(t, content, tt.expectedContains)
			}
			if tt.expectedAbsent != "" {
				assert.NotContains(t, content, tt.expectedAbsent)
			}
		})
	}
}

func TestDisplay(t *testing.T) {
	tests := []struct {
		name           string
		showPercentage string
		expectedTitle  string
	}{
		{
			name:           "titleLeft prepends percentage",
			showPercentage: "titleLeft",
			expectedTitle:  "50% Test Progress",
		},
		{
			name:           "titleRight appends percentage",
			showPercentage: "titleRight",
			expectedTitle:  "Test Progress 50%",
		},
		{
			name:           "other modes leave title untouched",
			showPercentage: "right",
			expectedTitle:  "Test Progress",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := createTestWidget(0, 100, 50, tt.showPercentage, 1)
			widget.percent = 0.5

			widget.display()

			select {
			case <-widget.RedrawChan:
			default:
				t.Fatal("expected display() to signal RedrawChan")
			}

			assert.Contains(t, widget.View.GetTitle(), tt.expectedTitle)
		})
	}
}

// findTestShell locates a shell binary capable of running `<shell> -c <cmd>`
// so execValueCmd/Refresh can be exercised on both Unix (sh/bash) and
// Windows (PowerShell accepts -c as a -Command alias) CI runners.
func findTestShell(t *testing.T) string {
	t.Helper()

	for _, candidate := range []string{"sh", "bash", "pwsh", "powershell"} {
		if path, err := exec.LookPath(candidate); err == nil {
			return path
		}
	}

	t.Skip("no compatible shell found to exercise execValueCmd")
	return ""
}

func TestExecValueCmd(t *testing.T) {
	widget := createTestWidget(0, 100, 0, "right", 1)

	t.Run("shell undefined returns error", func(t *testing.T) {
		widget.shell = ""

		_, err := widget.execValueCmd("echo 42")
		assert.ErrorIs(t, err, errShellUndefined)
	})

	shell := findTestShell(t)
	widget.shell = shell

	t.Run("valid numeric output", func(t *testing.T) {
		val, err := widget.execValueCmd("echo 42")
		assert.NoError(t, err)
		assert.Equal(t, 42.0, val)
	})

	t.Run("non-numeric output returns error", func(t *testing.T) {
		_, err := widget.execValueCmd("echo not-a-number")
		assert.Error(t, err)
	})

	t.Run("command failure returns error", func(t *testing.T) {
		_, err := widget.execValueCmd("exit 1")
		assert.Error(t, err)
	})
}

func TestRefresh(t *testing.T) {
	t.Run("no commands configured just recalculates percent", func(t *testing.T) {
		widget := createTestWidget(0, 100, 75, "right", 1)

		widget.Refresh()

		select {
		case <-widget.RedrawChan:
		default:
			t.Fatal("expected Refresh() to trigger a redraw")
		}

		assert.Nil(t, widget.err)
		assert.InDelta(t, 0.75, widget.percent, 0.0001)
	})

	t.Run("commands populate values", func(t *testing.T) {
		shell := findTestShell(t)

		widget := createTestWidget(0, 0, 0, "right", 1)
		widget.shell = shell
		widget.settings.minimumCmd = "echo 0"
		widget.settings.maximumCmd = "echo 100"
		widget.settings.currentCmd = "echo 25"

		widget.Refresh()

		select {
		case <-widget.RedrawChan:
		default:
			t.Fatal("expected Refresh() to trigger a redraw")
		}

		assert.Nil(t, widget.err)
		assert.Equal(t, 0.0, widget.minimum)
		assert.Equal(t, 100.0, widget.maximum)
		assert.Equal(t, 25.0, widget.current)
		assert.InDelta(t, 0.25, widget.percent, 0.0001)
	})

	t.Run("minimumCmd failure sets error and stops early", func(t *testing.T) {
		shell := findTestShell(t)

		widget := createTestWidget(0, 100, 0, "right", 1)
		widget.shell = shell
		widget.settings.minimumCmd = "exit 1"
		widget.settings.maximumCmd = "echo 100"

		widget.Refresh()

		select {
		case <-widget.RedrawChan:
		default:
			t.Fatal("expected Refresh() to trigger a redraw even on error")
		}

		assert.Error(t, widget.err)
		assert.Contains(t, widget.err.Error(), "minimumCmd execution failed")
	})

	t.Run("maximumCmd failure sets error and stops early", func(t *testing.T) {
		shell := findTestShell(t)

		widget := createTestWidget(0, 100, 0, "right", 1)
		widget.shell = shell
		widget.settings.maximumCmd = "exit 1"
		widget.settings.currentCmd = "echo 50"

		widget.Refresh()

		select {
		case <-widget.RedrawChan:
		default:
			t.Fatal("expected Refresh() to trigger a redraw even on error")
		}

		assert.Error(t, widget.err)
		assert.Contains(t, widget.err.Error(), "maximumCmd execution failed")
	})

	t.Run("currentCmd failure sets error", func(t *testing.T) {
		shell := findTestShell(t)

		widget := createTestWidget(0, 100, 0, "right", 1)
		widget.shell = shell
		widget.settings.currentCmd = "echo not-a-number"

		widget.Refresh()

		select {
		case <-widget.RedrawChan:
		default:
			t.Fatal("expected Refresh() to trigger a redraw even on error")
		}

		assert.Error(t, widget.err)
		assert.Contains(t, widget.err.Error(), "currentCmd execution failed")
	})
}
