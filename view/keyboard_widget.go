package view

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const refreshKeyChar = "r"

type helpItem struct {
	Key  string
	Text string
}

// KeyboardWidget manages keyboard control for a widget
type KeyboardWidget struct {
	pages    *tview.Pages
	settings *cfg.Common
	tviewApp *tview.Application
	view     *tview.TextView

	charMap  map[string]func()
	keyMap   map[tcell.Key]func()
	charHelp []helpItem
	keyHelp  []helpItem
	maxKey   int
}

// NewKeyboardWidget creates and returns a new instance of KeyboardWidget
func NewKeyboardWidget(tviewApp *tview.Application, pages *tview.Pages, settings *cfg.Common) *KeyboardWidget {
	keyWidget := &KeyboardWidget{
		tviewApp: tviewApp,
		pages:    pages,
		settings: settings,
		charMap:  make(map[string]func()),
		keyMap:   make(map[tcell.Key]func()),
		charHelp: []helpItem{},
		keyHelp:  []helpItem{},
	}

	keyWidget.initializeCommonKeyboardControls()

	return keyWidget
}

/* -------------------- Exported Functions --------------------- */

// AssignedChars returns a list of all the text characters assigned to an operation
func (widget *KeyboardWidget) AssignedChars() []string {
	chars := []string{}

	for char := range widget.charMap {
		chars = append(chars, char)
	}

	return chars
}

// HelpText returns the help text and keyboard command info for this widget
func (widget *KeyboardWidget) HelpText() string {
	str := " [green::b]Keyboard commands for " + strings.Title(widget.settings.Module.Type) + "[white]\n\n"

	for _, item := range widget.charHelp {
		str += fmt.Sprintf("  %s\t%s\n", item.Key, item.Text)
	}

	str += "\n\n"

	for _, item := range widget.keyHelp {
		str += fmt.Sprintf("  %-*s\t%s\n", widget.maxKey, item.Key, item.Text)
	}

	return str
}

// InitializeRefreshKeyboardControl assigns the module's explicit refresh function to
// the commom refresh key value
func (widget *KeyboardWidget) InitializeRefreshKeyboardControl(refreshFunc func()) {
	if refreshFunc != nil {
		widget.SetKeyboardChar(refreshKeyChar, refreshFunc, "Refresh widget")
	}
}

// InputCapture is the function passed to tview's SetInputCapture() function
// This is done during the main widget's creation process using the following code:
//
//    widget.View.SetInputCapture(widget.InputCapture)
//
func (widget *KeyboardWidget) InputCapture(event *tcell.EventKey) *tcell.EventKey {
	if event == nil {
		return nil
	}

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

// LaunchDocumentation opens the module docs in a browser
func (widget *KeyboardWidget) LaunchDocumentation() {
	path := widget.settings.DocPath
	if path == "" {
		path = widget.settings.Type
	}

	url := "https://wtfutil.com/modules/" + path
	utils.OpenFile(url)
}

// SetKeyboardChar sets a character/function combination that responds to key presses
// Example:
//
//    widget.SetKeyboardChar("d", widget.deleteSelectedItem)
//
func (widget *KeyboardWidget) SetKeyboardChar(char string, fn func(), helpText string) {
	if char == "" {
		return
	}

	// Check to ensure that the key trying to be used isn't already being used for something
	if _, ok := widget.charMap[char]; ok {
		panic(fmt.Sprintf("Key is already mapped to a keyboard command: %s\n", char))
	}

	widget.charMap[char] = fn
	widget.charHelp = append(widget.charHelp, helpItem{char, helpText})
}

// SetKeyboardKey sets a tcell.Key/function combination that responds to key presses
// Example:
//
//    widget.SetKeyboardKey(tcell.KeyCtrlD, widget.deleteSelectedItem)
//
func (widget *KeyboardWidget) SetKeyboardKey(key tcell.Key, fn func(), helpText string) {
	widget.keyMap[key] = fn
	widget.keyHelp = append(widget.keyHelp, helpItem{tcell.KeyNames[key], helpText})

	if len(tcell.KeyNames[key]) > widget.maxKey {
		widget.maxKey = len(tcell.KeyNames[key])
	}
}

// SetView assigns the passed-in tview.TextView view to this widget
func (widget *KeyboardWidget) SetView(view *tview.TextView) {
	widget.view = view
}

// ShowHelp displays the modal help dialog for a module
func (widget *KeyboardWidget) ShowHelp() {
	if widget.pages == nil {
		return
	}

	closeFunc := func() {
		widget.pages.RemovePage("help")
		widget.tviewApp.SetFocus(widget.view)
	}

	modal := NewBillboardModal(widget.HelpText(), closeFunc)

	widget.pages.AddPage("help", modal, false, true)
	widget.tviewApp.SetFocus(modal)

	widget.tviewApp.QueueUpdate(func() {
		widget.tviewApp.Draw()
	})
}

/* -------------------- Unexported Functions -------------------- */

// initializeCommonKeyboardControls sets up the keyboard controls that are common to
// all widgets that accept keyboard input
func (widget *KeyboardWidget) initializeCommonKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp, "Show/hide this help prompt")
	widget.SetKeyboardChar("\\", widget.LaunchDocumentation, "Open the documentation for this module in a browser")
}
