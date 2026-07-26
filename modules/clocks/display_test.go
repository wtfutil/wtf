package clocks

import (
	"strings"
	"testing"
	"time"
)

func TestLabelWidth(t *testing.T) {
	tests := []struct {
		name     string
		labels   []string
		expected int
	}{
		{"all short labels", []string{"NYC", "LA"}, 12},
		{"one exceeds minimum", []string{"San Francisco Bay"}, 19},
		{"exceeds minimum by 1", []string{"1234567890123"}, 15},
		{"empty list", []string{}, 12},
	}

	loc, _ := time.LoadLocation("UTC")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clocks := make([]Clock, len(tt.labels))
			for i, label := range tt.labels {
				clocks[i] = NewClock(label, loc)
			}
			result := labelWidth(clocks)
			if result != tt.expected {
				t.Errorf("labelWidth() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestFormatClocks_Empty(t *testing.T) {
	result := formatClocks([]Clock{}, "Jan 2", "15:04", func(int) string { return "white" })
	if !strings.Contains(result, "no timezone data available") {
		t.Errorf("empty clocks should show 'no timezone data available', got: %q", result)
	}
}

func TestFormatClocks_WithClocks(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	clocks := []Clock{
		NewClock("UTC", loc),
		NewClock("Also UTC", loc),
	}

	colorCalled := []int{}
	rowColor := func(idx int) string {
		colorCalled = append(colorCalled, idx)
		return "green"
	}

	result := formatClocks(clocks, "Jan 2", "15:04 MST", rowColor)

	// Should contain both labels
	if !strings.Contains(result, "UTC") {
		t.Error("output should contain 'UTC' label")
	}
	if !strings.Contains(result, "Also UTC") {
		t.Error("output should contain 'Also UTC' label")
	}

	// Should use the color function
	if !strings.Contains(result, "[green]") {
		t.Error("output should contain color directive [green]")
	}

	// rowColor should be called for each clock
	if len(colorCalled) != 2 {
		t.Errorf("rowColor called %d times, want 2", len(colorCalled))
	}

	// Each line ends with [white] reset
	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	for _, line := range lines {
		if !strings.HasSuffix(line, "[white]") {
			t.Errorf("line missing [white] suffix: %q", line)
		}
	}
}
