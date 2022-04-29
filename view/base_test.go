package view

import (
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/cfg"
)

func Benchmark_ContextualTitle(b *testing.B) {
	b.ReportAllocs()

	base := NewBase(
		tview.NewApplication(),
		make(chan bool),
		tview.NewPages(),
		&cfg.Common{},
	)
	base.SetFocusChar("a")

	defaultStr := "This is test"

	for i := 0; i < b.N; i++ {
		_ = base.ContextualTitle(defaultStr)
	}
}

func Test_ContextualTitle(t *testing.T) {
	tests := []struct {
		name       string
		defaultStr string
		focusChar  string
		expected   string
	}{
		{
			name:       "with empty defaultStr and empty focusChar",
			defaultStr: "",
			focusChar:  "",
			expected:   "",
		},
		{
			name:       "with valid defaultStr and empty focusChar",
			defaultStr: "cats",
			focusChar:  "",
			expected:   " cats ",
		},
		{
			name:       "with empty defaultStr and valid focusChar",
			defaultStr: "",
			focusChar:  "a",
			expected:   " [darkgray::u]a[::-][white] ",
		},
		{
			name:       "with valid defaultStr and valid focusChar",
			defaultStr: "cats",
			focusChar:  "a",
			expected:   " cats [darkgray::u]a[::-][white] ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base := NewBase(
				tview.NewApplication(),
				make(chan bool),
				tview.NewPages(),
				&cfg.Common{},
			)
			base.SetFocusChar(tt.focusChar)

			actual := base.ContextualTitle(tt.defaultStr)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
