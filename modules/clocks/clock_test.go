package clocks

import (
	"testing"
	"time"
)

func TestSanitizeLocation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"no spaces", "America/New_York", "America/New_York"},
		{"single space", "America/New York", "America/New_York"},
		{"multiple spaces", "Some Place With Spaces", "Some_Place_With_Spaces"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sanitizeLocation(tt.input)
			if result != tt.expected {
				t.Errorf("sanitizeLocation(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestBuildClock(t *testing.T) {
	tests := []struct {
		name      string
		label     string
		location  string
		expectErr bool
	}{
		{"valid timezone", "NYC", "America/New_York", false},
		{"valid with space", "NYC", "America/New York", false},
		{"UTC", "UTC", "UTC", false},
		{"invalid timezone", "Bad", "Invalid/Timezone", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clock, err := BuildClock(tt.label, tt.location)
			if tt.expectErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if clock.Label != tt.label {
					t.Errorf("Label = %q, want %q", clock.Label, tt.label)
				}
			}
		})
	}
}

func TestNewClock(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	clock := NewClock("Test", loc)

	if clock.Label != "Test" {
		t.Errorf("Label = %q, want %q", clock.Label, "Test")
	}
	if clock.Location != loc {
		t.Errorf("Location = %v, want %v", clock.Location, loc)
	}
}

func TestClockToLocal(t *testing.T) {
	nyLoc, _ := time.LoadLocation("America/New_York")
	clock := NewClock("NYC", nyLoc)

	utcTime := time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)
	localTime := clock.ToLocal(utcTime)

	if localTime.Location() != nyLoc {
		t.Errorf("ToLocal location = %v, want %v", localTime.Location(), nyLoc)
	}

	// During EDT (June), NY is UTC-4
	expectedHour := 8
	if localTime.Hour() != expectedHour {
		t.Errorf("ToLocal hour = %d, want %d", localTime.Hour(), expectedHour)
	}
}

func TestClockDateAndTime(t *testing.T) {
	utcLoc, _ := time.LoadLocation("UTC")
	clock := NewClock("UTC", utcLoc)

	// Date and Time use LocalTime() which calls time.Now(), so we just
	// verify they return non-empty strings with the given format
	dateResult := clock.Date("Jan 2")
	if dateResult == "" {
		t.Error("Date returned empty string")
	}

	timeResult := clock.Time("15:04 MST")
	if timeResult == "" {
		t.Error("Time returned empty string")
	}
}
