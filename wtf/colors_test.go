package wtf

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func Test_ASCIItoTviewColors(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected string
	}{
		{
			name:     "with blank text",
			text:     "",
			expected: "",
		},
		{
			name:     "with no color",
			text:     "cat",
			expected: "cat",
		},
		{
			name:     "with defined color",
			text:     "[38;5;226mcat/\x1b[0m",
			expected: "[38;5;226mcat/[-]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ASCIItoTviewColors(tt.text)

			if tt.expected != actual {
				t.Errorf("\nexpected: %q\n     got: %q", tt.expected, actual)
			}
		})
	}
}

func Test_ColorFor(t *testing.T) {
	tests := []struct {
		name     string
		label    string
		expected tcell.Color
	}{
		{
			name:     "with no label",
			label:    "",
			expected: tcell.ColorDefault,
		},
		{
			name:     "with missing label",
			label:    "cat",
			expected: tcell.ColorDefault,
		},
		{
			name:     "with defined label",
			label:    "tomato",
			expected: tcell.ColorTomato,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ColorFor(tt.label)

			if tt.expected != actual {
				t.Errorf("\nexpected: %q\n     got: %q", tt.expected, actual)
			}
		})
	}
}
