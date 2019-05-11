package wtf

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
)

type helpItem struct {
	Key  string
	Text string
}

// KeyboardWidget manages keyboard control for a widget
type KeyboardWidget struct {
	app      *tview.Application
	pages    *tview.Pages
	view     *tview.TextView
	settings *cfg.Common

	charMap map[string]func()
	keyMap  map[tcell.Key]func()
	helpMap []helpItem
}

// NewKeyboardWidget creates and returns a new instance of KeyboardWidget
func NewKeyboardWidget(app *tview.Application, pages *tview.Pages, settings *cfg.Common) KeyboardWidget {
	return KeyboardWidget{
		app:      app,
		pages:    pages,
		settings: settings,
		charMap:  make(map[string]func()),
		keyMap:   make(map[tcell.Key]func()),
		helpMap:  []helpItem{},
	}
}

// SetKeyboardChar sets a character/function combination that responds to key presses
// Example:
//
//    widget.SetKeyboardChar("d", widget.deleteSelectedItem)
//
func (widget *KeyboardWidget) SetKeyboardChar(char string, fn func(), helpText string) {
	widget.charMap[char] = fn
	widget.helpMap = append(widget.helpMap, helpItem{char, helpText})
}

// SetKeyboardKey sets a tcell.Key/function combination that responds to key presses
// Example:
//
//    widget.SetKeyboardKey(tcell.KeyCtrlD, widget.deleteSelectedItem)
//
func (widget *KeyboardWidget) SetKeyboardKey(key tcell.Key, fn func(), helpText string) {
	widget.keyMap[key] = fn
	widget.helpMap = append(widget.helpMap, helpItem{tcell.KeyNames[key], helpText})
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

func (widget *KeyboardWidget) HelpText() string {

	str := "Keyboard commands for " + widget.settings.Module.Type + "\n\n"

	for _, item := range widget.helpMap {
		str = str + fmt.Sprintf("%s: %s\n", item.Key, item.Text)
	}

	return str
}

func (widget *KeyboardWidget) SetView(view *tview.TextView) {
	widget.view = view
}

func (widget *KeyboardWidget) ShowHelp() {
	closeFunc := func() {
		widget.pages.RemovePage("help")
		widget.app.SetFocus(widget.view)
	}

	modal := NewBillboardModal(widget.HelpText(), closeFunc)

	widget.pages.AddPage("help", modal, false, true)
	widget.app.SetFocus(modal)

	widget.app.QueueUpdate(func() {
		widget.app.Draw()
	})
}
