package arpansagovau

import (
	"fmt"
	"strings"
	"testing"
)

func TestFormatLocationData_UVLevels(t *testing.T) {
	tests := []struct {
		name          string
		loc           *location
		wantLevel     string
		wantColor     string
		wantContains  []string
		wantAbsent    []string
	}{
		{
			name:         "low UV",
			loc:          &location{name: "adl", index: 1.5, time: "10:00 AM", date: "25/07/2026", status: "ok"},
			wantLevel:    "(LOW)",
			wantColor:    "[green]",
			wantContains: []string{"Location: adl", "UV index:", "1.50", "Local time: 10:00 AM 25/07/2026", "Detector status: ok"},
		},
		{
			name:         "moderate UV lower bound",
			loc:          &location{name: "bri", index: 2.5, time: "11:00 AM", date: "25/07/2026", status: "ok"},
			wantLevel:    "(MODERATE)",
			wantColor:    "[yellow]",
			wantContains: []string{"2.50"},
		},
		{
			name:         "moderate UV upper",
			loc:          &location{name: "can", index: 5.4, time: "12:00 PM", date: "25/07/2026", status: "ok"},
			wantLevel:    "(MODERATE)",
			wantColor:    "[yellow]",
			wantContains: []string{"5.40"},
		},
		{
			name:         "high UV",
			loc:          &location{name: "syd", index: 6.0, time: "1:00 PM", date: "25/07/2026", status: "ok"},
			wantLevel:    "(HIGH)",
			wantColor:    "[orange]",
			wantContains: []string{"6.00"},
		},
		{
			name:         "very high UV",
			loc:          &location{name: "dar", index: 9.0, time: "2:00 PM", date: "25/07/2026", status: "ok"},
			wantLevel:    "(VERY HIGH)",
			wantColor:    "[red]",
			wantContains: []string{"9.00"},
		},
		{
			name:         "extreme UV",
			loc:          &location{name: "tow", index: 11.0, time: "12:00 PM", date: "25/07/2026", status: "ok"},
			wantLevel:    "(EXTREME)",
			wantColor:    "[fuchsia]",
			wantContains: []string{"11.00"},
		},
		{
			name:         "extreme UV at boundary",
			loc:          &location{name: "tow", index: 10.5, time: "12:00 PM", date: "25/07/2026", status: "ok"},
			wantLevel:    "(EXTREME)",
			wantColor:    "[fuchsia]",
			wantContains: []string{"10.50"},
		},
		{
			name:         "very high UV at boundary",
			loc:          &location{name: "per", index: 7.5, time: "11:30 AM", date: "25/07/2026", status: "ok"},
			wantLevel:    "(VERY HIGH)",
			wantColor:    "[red]",
			wantContains: []string{"7.50"},
		},
		{
			name:         "high UV at boundary",
			loc:          &location{name: "mel", index: 5.5, time: "11:30 AM", date: "25/07/2026", status: "ok"},
			wantLevel:    "(HIGH)",
			wantColor:    "[orange]",
			wantContains: []string{"5.50"},
		},
		{
			name:         "zero UV",
			loc:          &location{name: "hob", index: 0.0, time: "7:00 PM", date: "25/07/2026", status: "ok"},
			wantLevel:    "(LOW)",
			wantColor:    "[green]",
			wantContains: []string{"0.00"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := formatLocationData(tc.loc)

			if !strings.Contains(result, tc.wantLevel) {
				t.Errorf("expected level %q in output:\n%s", tc.wantLevel, result)
			}
			if !strings.Contains(result, tc.wantColor) {
				t.Errorf("expected color %q in output:\n%s", tc.wantColor, result)
			}
			for _, s := range tc.wantContains {
				if !strings.Contains(result, s) {
					t.Errorf("expected %q in output:\n%s", s, result)
				}
			}
		})
	}
}

func TestFormatLocationData_EmptyName(t *testing.T) {
	loc := &location{name: "", index: 5.0, status: "ok"}
	result := formatLocationData(loc)
	expected := "[red]No data?"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestFormatLocationData_StatusNotOk(t *testing.T) {
	tests := []struct {
		name   string
		loc    *location
		expect string
	}{
		{
			name:   "unavailable status",
			loc:    &location{name: "mel", index: 0.0, status: "unavailable"},
			expect: "[red]Data unavailable for mel",
		},
		{
			name:   "error status",
			loc:    &location{name: "adl", index: 0.0, status: "error"},
			expect: "[red]Data unavailable for adl",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := formatLocationData(tc.loc)
			if result != tc.expect {
				t.Errorf("expected %q, got %q", tc.expect, result)
			}
		})
	}
}

func TestFormatLocationData_IndexFormatting(t *testing.T) {
	loc := &location{name: "syd", index: 3.14159, time: "12:00 PM", date: "01/01/2026", status: "ok"}
	result := formatLocationData(loc)
	expected := fmt.Sprintf("%.2f", float32(3.14159))
	if !strings.Contains(result, expected) {
		t.Errorf("expected formatted index %q in output:\n%s", expected, result)
	}
}
