package clocks

import (
	"testing"
	"time"
)

func buildTestClocks() ClockCollection {
	nyLoc, _ := time.LoadLocation("America/New_York")
	londonLoc, _ := time.LoadLocation("Europe/London")
	tokyoLoc, _ := time.LoadLocation("Asia/Tokyo")

	return ClockCollection{
		Clocks: []Clock{
			NewClock("New York", nyLoc),
			NewClock("London", londonLoc),
			NewClock("Tokyo", tokyoLoc),
		},
	}
}

func TestSortedAlphabetically(t *testing.T) {
	coll := buildTestClocks()
	coll.SortedAlphabetically()

	expected := []string{"London", "New York", "Tokyo"}
	for i, clock := range coll.Clocks {
		if clock.Label != expected[i] {
			t.Errorf("position %d: got %q, want %q", i, clock.Label, expected[i])
		}
	}
}

func TestSortedChronologically(t *testing.T) {
	coll := buildTestClocks()
	coll.SortedChronologically()

	// Chronological: NY (UTC-4/5) < London (UTC+0/1) < Tokyo (UTC+9)
	// Earlier time = first
	if len(coll.Clocks) != 3 {
		t.Fatalf("expected 3 clocks, got %d", len(coll.Clocks))
	}
	if coll.Clocks[0].Label != "New York" {
		t.Errorf("first clock = %q, want New York (earliest timezone)", coll.Clocks[0].Label)
	}
	if coll.Clocks[2].Label != "Tokyo" {
		t.Errorf("last clock = %q, want Tokyo (latest timezone)", coll.Clocks[2].Label)
	}
}

func TestSortedReverseChronologically(t *testing.T) {
	coll := buildTestClocks()
	coll.SortedReverseChronologically()

	if coll.Clocks[0].Label != "Tokyo" {
		t.Errorf("first clock = %q, want Tokyo (latest timezone)", coll.Clocks[0].Label)
	}
	if coll.Clocks[2].Label != "New York" {
		t.Errorf("last clock = %q, want New York (earliest timezone)", coll.Clocks[2].Label)
	}
}

func TestSorted(t *testing.T) {
	tests := []struct {
		name      string
		sortOrder string
		firstExp  string
	}{
		{"natural preserves order", "natural", "New York"},
		{"alphabetical", "alphabetical", "London"},
		{"chronological", "chronological", "New York"},
		{"reversechronological", "reversechronological", "Tokyo"},
		{"default is alphabetical", "unknown", "London"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coll := buildTestClocks()
			result := coll.Sorted(tt.sortOrder)
			if result[0].Label != tt.firstExp {
				t.Errorf("Sorted(%q) first = %q, want %q", tt.sortOrder, result[0].Label, tt.firstExp)
			}
		})
	}
}

func TestEmptyCollection(t *testing.T) {
	coll := ClockCollection{Clocks: []Clock{}}
	result := coll.Sorted("alphabetical")
	if len(result) != 0 {
		t.Errorf("expected empty result, got %d clocks", len(result))
	}
}
