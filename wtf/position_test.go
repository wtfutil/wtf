package wtf

import (
	"testing"
)

func Test_NewPosition(t *testing.T) {
	pos := NewPosition(0, 1, 2, 3)

	if pos.Height() != 3 {
		t.Fatalf("Expected 3 but got %d", pos.Height())
	}

	if pos.Left() != 1 {
		t.Fatalf("Expected 1 but got %d", pos.Left())
	}

	if pos.Top() != 0 {
		t.Fatalf("Expected 0 but got %d", pos.Top())
	}

	if pos.Width() != 2 {
		t.Fatalf("Expected 2 but got %d", pos.Width())
	}
}

func Test_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		height   int
		left     int
		top      int
		width    int
		expected bool
	}{
		{
			name:     "valid position",
			height:   2,
			left:     0,
			top:      1,
			width:    2,
			expected: true,
		},
		{
			name:     "invalid height",
			height:   0,
			left:     0,
			top:      1,
			width:    2,
			expected: false,
		},
		{
			name:     "invalid left",
			height:   2,
			left:     -1,
			top:      1,
			width:    2,
			expected: false,
		},
		{
			name:     "invalid top",
			height:   2,
			left:     0,
			top:      -1,
			width:    2,
			expected: false,
		},
		{
			name:     "invalid width",
			height:   2,
			left:     0,
			top:      1,
			width:    0,
			expected: false,
		},
	}

	for _, tt := range tests {
		pos := NewPosition(tt.top, tt.left, tt.width, tt.height)
		actual := pos.IsValid()

		if actual != tt.expected {
			t.Errorf("%s: expected: %v, got: %v", tt.name, tt.expected, actual)
		}
	}
}
