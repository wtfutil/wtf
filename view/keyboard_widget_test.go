package view

import (
	"testing"

	"github.com/gdamore/tcell"
	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/cfg"
)

func test() {}

func testKeyboardWidget() *KeyboardWidget {
	keyWid := NewKeyboardWidget(
		&cfg.Common{
			Module: cfg.Module{
				Name: "testWidget",
				Type: "testType",
			},
		},
	)
	return keyWid
}

func Test_SetKeyboardChar(t *testing.T) {
	tests := []struct {
		name     string
		char     string
		fn       func()
		helpText string
		mapChar  string
		expected bool
	}{
		{
			name:     "with blank char",
			char:     "",
			fn:       test,
			helpText: "help",
			mapChar:  "",
			expected: false,
		},
		{
			name:     "with undefined char",
			char:     "d",
			fn:       test,
			helpText: "help",
			mapChar:  "m",
			expected: false,
		},
		{
			name:     "with defined char",
			char:     "d",
			fn:       test,
			helpText: "help",
			mapChar:  "d",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keyWid := testKeyboardWidget()
			keyWid.SetKeyboardChar(tt.char, tt.fn, tt.helpText)

			actual := keyWid.charMap[tt.mapChar]

			if tt.expected != (actual != nil) {
				t.Errorf("\nexpected: %s\n     got: %T", "actual != nil", actual)
			}
		})
	}
}

func Test_SetKeyboardKey(t *testing.T) {
	tests := []struct {
		name     string
		key      tcell.Key
		fn       func()
		helpText string
		mapKey   tcell.Key
		expected bool
	}{
		{
			name:     "with undefined key",
			key:      tcell.KeyCtrlA,
			fn:       test,
			helpText: "help",
			mapKey:   tcell.KeyCtrlZ,
			expected: false,
		},
		{
			name:     "with defined key",
			key:      tcell.KeyCtrlA,
			fn:       test,
			helpText: "help",
			mapKey:   tcell.KeyCtrlA,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keyWid := testKeyboardWidget()
			keyWid.SetKeyboardKey(tt.key, tt.fn, tt.helpText)

			actual := keyWid.keyMap[tt.mapKey]

			if tt.expected != (actual != nil) {
				t.Errorf("\nexpected: %s\n     got: %T", "actual != nil", actual)
			}
		})
	}
}

func Test_InputCapture(t *testing.T) {
	tests := []struct {
		name     string
		before   func(keyWid *KeyboardWidget) *KeyboardWidget
		event    *tcell.EventKey
		expected *tcell.EventKey
	}{
		{
			name:     "with nil event",
			before:   func(keyWid *KeyboardWidget) *KeyboardWidget { return keyWid },
			event:    nil,
			expected: nil,
		},
		{
			name:     "with undefined event",
			before:   func(keyWid *KeyboardWidget) *KeyboardWidget { return keyWid },
			event:    tcell.NewEventKey(tcell.KeyRune, 'a', tcell.ModNone),
			expected: tcell.NewEventKey(tcell.KeyRune, 'a', tcell.ModNone),
		},
		{
			name: "with defined event and char handler",
			before: func(keyWid *KeyboardWidget) *KeyboardWidget {
				keyWid.SetKeyboardChar("a", test, "help")
				return keyWid
			},
			event:    tcell.NewEventKey(tcell.KeyRune, 'a', tcell.ModNone),
			expected: nil,
		},
		{
			name: "with defined event and key handler",
			before: func(keyWid *KeyboardWidget) *KeyboardWidget {
				keyWid.SetKeyboardKey(tcell.KeyRune, test, "help")
				return keyWid
			},
			event:    tcell.NewEventKey(tcell.KeyRune, 'a', tcell.ModNone),
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keyWid := testKeyboardWidget()
			keyWid = tt.before(keyWid)
			actual := keyWid.InputCapture(tt.event)

			if tt.expected == nil {
				if actual != nil {
					t.Errorf("\nexpected: %v\n     got: %v", tt.expected, actual.Rune())
				}
				return
			}

			if tt.expected.Rune() != actual.Rune() {
				t.Errorf("\nexpected: %v\n     got: %v", tt.expected.Rune(), actual.Rune())
			}
		})
	}
}

func Test_initializeCommonKeyboardControls(t *testing.T) {
	t.Run("nil refreshFunc", func(t *testing.T) {
		keyWid := testKeyboardWidget()

		assert.NotNil(t, keyWid.charMap["\\"])
	})
}

func Test_InitializeRefreshKeyboardControl(t *testing.T) {
	t.Run("nil refreshFunc", func(t *testing.T) {
		keyWid := testKeyboardWidget()
		keyWid.InitializeRefreshKeyboardControl(nil)

		assert.Nil(t, keyWid.charMap["r"])
	})

	t.Run("non-nil refreshFunc", func(t *testing.T) {
		keyWid := testKeyboardWidget()
		keyWid.InitializeRefreshKeyboardControl(func() {})

		assert.NotNil(t, keyWid.charMap["r"])
	})
}

func Test_HelpText(t *testing.T) {
	keyWid := testKeyboardWidget()
	keyWid.SetKeyboardChar("a", test, "a help")
	keyWid.SetKeyboardKey(tcell.KeyCtrlO, test, "keyCtrlO help")

	assert.NotNil(t, keyWid.HelpText())
}
