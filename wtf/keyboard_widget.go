package wtf

import (
	"github.com/gdamore/tcell"
)

// KeyboardWidget manages keyboard control for a widget
type KeyboardWidget struct {
	charMap map[string]func()
	keyMap  map[tcell.Key]func()
}

// NewKeyboardWidget creates and returns a new instance of KeyboardWidget
func NewKeyboardWidget() KeyboardWidget {
	return KeyboardWidget{
		charMap: make(map[string]func()),
		keyMap:  make(map[tcell.Key]func()),
	}
}

// SetKeyboardChar sets a character/function combination that responds to key presses
// Example:
//
//    widget.SetKeyboardChar("d", widget.deleteSelectedItem)
//
func (widget *KeyboardWidget) SetKeyboardChar(char string, fn func()) {
	widget.charMap[char] = fn
}

// SetKeyboardKey sets a tcell.Key/function combination that responds to key presses
// Example:
//
//    widget.SetKeyboardKey(tcell.KeyCtrlD, widget.deleteSelectedItem)
//
func (widget *KeyboardWidget) SetKeyboardKey(key tcell.Key, fn func()) {
	widget.keyMap[key] = fn
}

// InputCapture is the function passed to tview's SetInputCapture() function
// This is done during the main widget's creation process using the following code:
//
//    widget.View.SetInputCapture(widget.InputCapture)
//
func (widget *KeyboardWidget) InputCapture(event *tcell.EventKey) *tcell.EventKey {
	fn := widget.charMap[string(event.Rune())]
	if fn != nil {
		fn()
		return nil
	}

	fn = widget.keyMap[event.Key()]
	if fn != nil {
		fn()
		return nil
	}

	return event
}
