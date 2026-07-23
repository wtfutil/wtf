package ping

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// expectedStatusLine mirrors the padding logic in Widget.content so tests
// stay correct if the minimum column width or padding character changes.
func expectedStatusLine(nameWidth int, label, status string) string {
	return fmt.Sprintf("[white]%-*s: %s", nameWidth, label, status)
}

func Test_Widget_content(t *testing.T) {
	tests := []struct {
		name     string
		hosts    []Host
		expected string
	}{
		{
			name:     "no hosts produces empty content",
			hosts:    []Host{},
			expected: "",
		},
		{
			name: "single host up",
			hosts: []Host{
				{Label: "example.com", Hostname: "example.com", Up: true},
			},
			expected: expectedStatusLine(12, "example.com", "[green]Up"),
		},
		{
			name: "single host down",
			hosts: []Host{
				{Label: "example.com", Hostname: "example.com", Up: false},
			},
			expected: expectedStatusLine(12, "example.com", "[red]DOWN"),
		},
		{
			name: "multiple hosts mixed status",
			hosts: []Host{
				{Label: "up-host", Hostname: "up.example.com", Up: true},
				{Label: "down-host", Hostname: "down.example.com", Up: false},
			},
			// nameWidth stays at the 12-char minimum since both labels are shorter.
			expected: strings.Join([]string{
				expectedStatusLine(12, "up-host", "[green]Up"),
				expectedStatusLine(12, "down-host", "[red]DOWN"),
			}, "\n"),
		},
		{
			name: "long label widens name column",
			hosts: []Host{
				{Label: "a-very-long-hostname-label", Hostname: "long.example.com", Up: true},
			},
			// nameWidth grows to len(label)+2 since it exceeds the 12-char minimum.
			expected: expectedStatusLine(28, "a-very-long-hostname-label", "[green]Up"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := &Widget{hosts: tt.hosts}

			assert.Equal(t, tt.expected, widget.content())
		})
	}
}

func Test_Widget_content_shortLabelsPadToMinimumWidth(t *testing.T) {
	widget := &Widget{
		hosts: []Host{
			{Label: "a", Hostname: "a.example.com", Up: true},
		},
	}

	content := widget.content()

	// nameWidth defaults to 12 when all labels are shorter than that.
	assert.Equal(t, expectedStatusLine(12, "a", "[green]Up"), content)
}
