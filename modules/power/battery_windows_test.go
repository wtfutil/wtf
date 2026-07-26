//go:build windows

package power

import (
	"strings"
	"testing"
)

func Test_formatStatus(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		expected string
	}{
		{
			name:     "discharging",
			code:     "1",
			expected: "[yellow]discharging[white]",
		},
		{
			name:     "AC connected",
			code:     "2",
			expected: "[white]AC connected[white]",
		},
		{
			name:     "fully charged",
			code:     "3",
			expected: "[white]fully charged[white]",
		},
		{
			name:     "low battery",
			code:     "4",
			expected: "[yellow]low[white]",
		},
		{
			name:     "critical battery",
			code:     "5",
			expected: "[red]critical[white]",
		},
		{
			name:     "charging",
			code:     "6",
			expected: "[green]charging[white]",
		},
		{
			name:     "charging high",
			code:     "7",
			expected: "[green]charging (high)[white]",
		},
		{
			name:     "charging low",
			code:     "8",
			expected: "[green]charging (low)[white]",
		},
		{
			name:     "critical charging",
			code:     "9",
			expected: "[red]critical (charging)[white]",
		},
		{
			name:     "unknown code",
			code:     "99",
			expected: "[white]unknown[white]",
		},
		{
			name:     "empty code",
			code:     "",
			expected: "[white]unknown[white]",
		},
		{
			name:     "code with whitespace",
			code:     " 6 ",
			expected: "[green]charging[white]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			battery := NewBattery()
			actual := battery.formatStatus(tt.code)
			if actual != tt.expected {
				t.Errorf("formatStatus(%q) = %q, want %q", tt.code, actual, tt.expected)
			}
		})
	}
}

func Test_parseWMIOutput(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		wantNoBattery  bool
		wantContains   []string
		wantNotContain []string
	}{
		{
			name:          "empty output means no battery",
			input:         "",
			wantNoBattery: true,
		},
		{
			name:          "whitespace-only output means no battery",
			input:         "   \n  \t  ",
			wantNoBattery: true,
		},
		{
			name: "valid battery output with charging",
			input: `EstimatedChargeRemaining : 85
BatteryStatus           : 6
EstimatedRunTime        : 120`,
			wantContains: []string{"Charge", "Remaining", "State", "120", "charging"},
		},
		{
			name: "valid battery output discharging",
			input: `EstimatedChargeRemaining : 45
BatteryStatus           : 1
EstimatedRunTime        : 90`,
			wantContains: []string{"Charge", "90", "discharging"},
		},
		{
			name: "remaining time zero shows dash",
			input: `EstimatedChargeRemaining : 100
BatteryStatus           : 3
EstimatedRunTime        : 0`,
			wantContains: []string{"Remaining", "- min"},
		},
		{
			name: "remaining time empty shows dash",
			input: `EstimatedChargeRemaining : 100
BatteryStatus           : 2
EstimatedRunTime        : `,
			wantContains: []string{"- min"},
		},
		{
			name: "non-numeric charge uses raw value with percent",
			input: `EstimatedChargeRemaining : N/A
BatteryStatus           : 2
EstimatedRunTime        : 60`,
			wantContains: []string{"N/A%"},
		},
		{
			name: "low charge value formatted correctly",
			input: `EstimatedChargeRemaining : 15
BatteryStatus           : 4
EstimatedRunTime        : 20`,
			wantContains: []string{"Charge", "State", "[yellow]low[white]"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			battery := NewBattery()
			result := battery.parseWMIOutput(tt.input)

			if tt.wantNoBattery {
				if result != " no battery found" {
					t.Errorf("expected no battery message, got %q", result)
				}
				return
			}

			for _, want := range tt.wantContains {
				if !strings.Contains(result, want) {
					t.Errorf("result should contain %q, got:\n%s", want, result)
				}
			}

			for _, notWant := range tt.wantNotContain {
				if strings.Contains(result, notWant) {
					t.Errorf("result should NOT contain %q, got:\n%s", notWant, result)
				}
			}
		})
	}
}

func Test_NewBattery(t *testing.T) {
	battery := NewBattery()
	if battery == nil {
		t.Fatal("NewBattery() returned nil")
	}
	if battery.String() != "" {
		t.Errorf("new battery should have empty result, got %q", battery.String())
	}
}

func Test_BatteryString(t *testing.T) {
	battery := NewBattery()
	battery.result = "test content"
	if battery.String() != "test content" {
		t.Errorf("String() = %q, want %q", battery.String(), "test content")
	}
}
