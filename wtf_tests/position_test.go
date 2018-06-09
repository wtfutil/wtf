package wtf_tests

import (
	"testing"

	. "github.com/senorprogrammer/wtf/wtf"
)

func TestPosition(t *testing.T) {
	pos := NewPosition(0, 1, 2, 3)

	if pos.Top() != 0 {
		t.Fatalf("Expected 0 but got %d", pos.Top())
	}

	if pos.Left() != 1 {
		t.Fatalf("Expected 1 but got %d", pos.Left())
	}

	if pos.Width() != 2 {
		t.Fatalf("Expected 2 but got %d", pos.Width())
	}

	if pos.Height() != 3 {
		t.Fatalf("Expected 3 but got %d", pos.Height())
	}
}
