package digitalclock

import (
	"strings"
	"testing"
)

func TestMergeLines(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{"single line", []string{"hello"}, "hello"},
		{"multiple lines", []string{"a", "b", "c"}, "a\nb\nc"},
		{"empty slice", []string{}, ""},
		{"empty strings", []string{"", "", ""}, "\n\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeLines(tt.input)
			if result != tt.expected {
				t.Errorf("mergeLines(%v) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRenderClock(t *testing.T) {
	tests := []struct {
		name       string
		font       string
		hourFormat string
		color      string
	}{
		{"digital 24hr", "digitalfont", "24", "white"},
		{"big 24hr", "bigfont", "24", "green"},
		{"bold 12hr", "boldfont", "12", "red"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := Settings{
				font:       tt.font,
				hourFormat: tt.hourFormat,
				color:      tt.color,
			}
			result, needBorder := renderClock(settings)
			if result == "" {
				t.Error("renderClock returned empty string")
			}
			if !strings.Contains(result, tt.color) {
				t.Errorf("renderClock missing color %q in output", tt.color)
			}
			// Digital font has 3 rows <= minRowsForBorder(3), so needs border
			if tt.font == "digitalfont" && !needBorder {
				t.Error("digital font should need border")
			}
			// Big font has 6 rows > 3, no border needed
			if tt.font == "bigfont" && needBorder {
				t.Error("big font should not need border")
			}
		})
	}
}

func TestGetHourMinute(t *testing.T) {
	// 24-hour format
	result24 := getHourMinute("24")
	if result24 == "" {
		t.Error("getHourMinute(24) returned empty")
	}
	// In 24-hour format, AMPM is " " so no AM/PM letter should appear at the end
	if strings.HasSuffix(result24, "A") || strings.HasSuffix(result24, "P") {
		t.Errorf("getHourMinute(24) = %q, should not end with AM/PM indicator", result24)
	}

	// 12-hour format
	result12 := getHourMinute("12")
	if result12 == "" {
		t.Error("getHourMinute(12) returned empty")
	}
	// Should contain AM or PM indicator
	if !strings.Contains(result12, "A") && !strings.Contains(result12, "P") {
		t.Errorf("getHourMinute(12) = %q, expected AM or PM indicator", result12)
	}
}

func TestGetDate(t *testing.T) {
	withPrefix := getDate("2006-01-02", true)
	if !strings.HasPrefix(withPrefix, "Date: ") {
		t.Errorf("getDate with prefix = %q, want 'Date: ' prefix", withPrefix)
	}

	withoutPrefix := getDate("2006-01-02", false)
	if strings.HasPrefix(withoutPrefix, "Date: ") {
		t.Errorf("getDate without prefix = %q, should not have 'Date: ' prefix", withoutPrefix)
	}
}

func TestGetUTC(t *testing.T) {
	result := getUTC()
	if !strings.HasPrefix(result, "UTC: ") {
		t.Errorf("getUTC() = %q, want 'UTC: ' prefix", result)
	}
}

func TestGetEpoch(t *testing.T) {
	result := getEpoch()
	if !strings.HasPrefix(result, "Epoch: ") {
		t.Errorf("getEpoch() = %q, want 'Epoch: ' prefix", result)
	}
}

func TestRenderWidget(t *testing.T) {
	tests := []struct {
		name           string
		withDate       bool
		withUTC        bool
		withEpoch      bool
		withDatePrefix bool
		expectDate     bool
		expectUTC      bool
		expectEpoch    bool
	}{
		{"all options", true, true, true, true, true, true, true},
		{"clock only", false, false, false, false, false, false, false},
		{"date without prefix", true, false, false, false, true, false, false},
		{"utc only", false, true, false, false, false, true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := Settings{
				font:           "digitalfont",
				hourFormat:     "24",
				color:          "white",
				dateFormat:     "2006-01-02",
				withDate:       tt.withDate,
				withUTC:        tt.withUTC,
				withEpoch:      tt.withEpoch,
				withDatePrefix: tt.withDatePrefix,
			}

			result := renderWidget(settings)

			if result == "" {
				t.Fatal("renderWidget returned empty string")
			}
			if tt.expectDate && !strings.Contains(result, "Date:") && !strings.Contains(result, "202") {
				t.Error("expected date in output")
			}
			if tt.expectUTC && !strings.Contains(result, "UTC:") {
				t.Error("expected UTC in output")
			}
			if tt.expectEpoch && !strings.Contains(result, "Epoch:") {
				t.Error("expected Epoch in output")
			}
			if !tt.expectUTC && strings.Contains(result, "UTC:") {
				t.Error("did not expect UTC in output")
			}
			if !tt.expectEpoch && strings.Contains(result, "Epoch:") {
				t.Error("did not expect Epoch in output")
			}
		})
	}
}
