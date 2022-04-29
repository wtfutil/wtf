package view

import (
	"testing"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
)

func testTextWidget() TextWidget {
	txtWid := NewTextWidget(
		tview.NewApplication(),
		make(chan bool),
		tview.NewPages(),
		&cfg.Common{
			Module: cfg.Module{
				Name: "test widget",
			},
		},
	)
	return txtWid
}

func Test_Bordered(t *testing.T) {
	tests := []struct {
		name     string
		before   func(txtWid TextWidget) TextWidget
		expected bool
	}{
		{
			name: "without border",
			before: func(txtWid TextWidget) TextWidget {
				txtWid.bordered = false
				return txtWid
			},
			expected: false,
		},
		{
			name: "with border",
			before: func(txtWid TextWidget) TextWidget {
				txtWid.bordered = true
				return txtWid
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txtWid := testTextWidget()
			txtWid = tt.before(txtWid)
			actual := txtWid.Bordered()

			if tt.expected != actual {
				t.Errorf("\nexpected: %t\n     got: %t", tt.expected, actual)
			}
		})
	}
}

func Test_Disabled(t *testing.T) {
	tests := []struct {
		name     string
		before   func(txtWid TextWidget) TextWidget
		expected bool
	}{
		{
			name: "when not enabled",
			before: func(txtWid TextWidget) TextWidget {
				txtWid.enabled = false
				return txtWid
			},
			expected: true,
		},
		{
			name: "when enabled",
			before: func(txtWid TextWidget) TextWidget {
				txtWid.enabled = true
				return txtWid
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txtWid := testTextWidget()
			txtWid = tt.before(txtWid)
			actual := txtWid.Disabled()

			if tt.expected != actual {
				t.Errorf("\nexpected: %t\n     got: %t", tt.expected, actual)
			}
		})
	}
}

func Test_Enabled(t *testing.T) {
	tests := []struct {
		name     string
		before   func(txtWid TextWidget) TextWidget
		expected bool
	}{
		{
			name: "when not enabled",
			before: func(txtWid TextWidget) TextWidget {
				txtWid.enabled = false
				return txtWid
			},
			expected: false,
		},
		{
			name: "when enabled",
			before: func(txtWid TextWidget) TextWidget {
				txtWid.enabled = true
				return txtWid
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txtWid := testTextWidget()
			txtWid = tt.before(txtWid)
			actual := txtWid.Enabled()

			if tt.expected != actual {
				t.Errorf("\nexpected: %t\n     got: %t", tt.expected, actual)
			}
		})
	}
}

func Test_Focusable(t *testing.T) {
	tests := []struct {
		name     string
		before   func(txtWid TextWidget) TextWidget
		expected bool
	}{
		{
			name: "when not focusable",
			before: func(txtWid TextWidget) TextWidget {
				txtWid.enabled = false
				txtWid.focusable = false
				return txtWid
			},
			expected: false,
		},
		{
			name: "when not focusable",
			before: func(txtWid TextWidget) TextWidget {
				txtWid.enabled = false
				txtWid.focusable = true
				return txtWid
			},
			expected: false,
		},
		{
			name: "when not focusable",
			before: func(txtWid TextWidget) TextWidget {
				txtWid.enabled = true
				txtWid.focusable = false
				return txtWid
			},
			expected: false,
		},
		{
			name: "when focusable",
			before: func(txtWid TextWidget) TextWidget {
				txtWid.enabled = true
				txtWid.focusable = true
				return txtWid
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txtWid := testTextWidget()
			txtWid = tt.before(txtWid)
			actual := txtWid.Focusable()

			if tt.expected != actual {
				t.Errorf("\nexpected: %t\n     got: %t", tt.expected, actual)
			}
		})
	}
}

func Test_Name(t *testing.T) {
	txtWid := testTextWidget()
	actual := txtWid.Name()
	expected := "test widget"

	if expected != actual {
		t.Errorf("\nexpected: %s\n     got: %s", expected, actual)
	}
}

func Test_String(t *testing.T) {
	txtWid := testTextWidget()
	actual := txtWid.String()
	expected := "test widget"

	if expected != actual {
		t.Errorf("\nexpected: %s\n     got: %s", expected, actual)
	}
}
