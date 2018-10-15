package wtf

import (
	"github.com/rivo/tview"
)

type FocusState int

const (
	widgetFocused FocusState = iota
	appBoardFocused
	neverFocused
)

// FocusTracker is used by the app to track which onscreen widget currently has focus,
// and to move focus between widgets.
type FocusTracker struct {
	App     *tview.Application
	Idx     int
	Widgets []Wtfable
}

/* -------------------- Exported Functions -------------------- */

// AssignHotKeys assigns an alphabetic keyboard character to each focusable
// widget so that the widget can be brought into focus by pressing that keyboard key
func (tracker *FocusTracker) AssignHotKeys() {
	if !tracker.useNavShortcuts() {
		return
	}

	usedKeys := make(map[string]bool)
	focusables := tracker.focusables()
	i := 1

	for _, focusable := range focusables {
		if focusable.FocusChar() != "" {
			usedKeys[focusable.FocusChar()] = true
		}
	}
	for _, focusable := range focusables {
		if focusable.FocusChar() != "" {
			continue
		}
		if _, foundKey := usedKeys[string('0'+i)]; foundKey {
			for ; foundKey; _, foundKey = usedKeys[string('0'+i)] {
				i++
			}
		}

		// Don't have nav characters > "9"
		if i >= 10 {
			break
		}

		focusable.SetFocusChar(string('0' + i))
		i++
	}
}

func (tracker *FocusTracker) FocusOn(char string) bool {
	if !tracker.useNavShortcuts() {
		return false
	}

	if tracker.focusState() == appBoardFocused {
		return false
	}

	hasFocusable := false

	for idx, focusable := range tracker.focusables() {
		if focusable.FocusChar() == char {
			tracker.blur(tracker.Idx)
			tracker.Idx = idx
			tracker.focus(tracker.Idx)

			hasFocusable = true
			break
		}
	}

	return hasFocusable
}

// Next sets the focus on the next widget in the widget list. If the current widget is
// the last widget, sets focus on the first widget.
func (tracker *FocusTracker) Next() {
	if tracker.focusState() == appBoardFocused {
		return
	}

	tracker.blur(tracker.Idx)
	tracker.increment()
	tracker.focus(tracker.Idx)
}

// None removes focus from the currently-focused widget.
func (tracker *FocusTracker) None() {
	if tracker.focusState() == appBoardFocused {
		return
	}

	tracker.blur(tracker.Idx)
}

// Prev sets the focus on the previous widget in the widget list. If the current widget is
// the last widget, sets focus on the last widget.
func (tracker *FocusTracker) Prev() {
	if tracker.focusState() == appBoardFocused {
		return
	}

	tracker.blur(tracker.Idx)
	tracker.decrement()
	tracker.focus(tracker.Idx)
}

func (tracker *FocusTracker) Refocus() {
	tracker.focus(tracker.Idx)
}

/* -------------------- Unexported Functions -------------------- */

func (tracker *FocusTracker) blur(idx int) {
	widget := tracker.focusableAt(idx)
	if widget == nil {
		return
	}

	view := widget.TextView()
	view.Blur()

	view.SetBorderColor(ColorFor(widget.BorderColor()))
}

func (tracker *FocusTracker) decrement() {
	tracker.Idx = tracker.Idx - 1

	if tracker.Idx < 0 {
		tracker.Idx = len(tracker.focusables()) - 1
	}
}

func (tracker *FocusTracker) focus(idx int) {
	widget := tracker.focusableAt(idx)
	if widget == nil {
		return
	}

	view := widget.TextView()
	view.SetBorderColor(ColorFor(Config.UString("wtf.colors.border.focused", "gray")))

	tracker.App.SetFocus(view)
	tracker.App.Draw()
}

func (tracker *FocusTracker) focusables() []Wtfable {
	focusable := []Wtfable{}

	for _, widget := range tracker.Widgets {
		if widget.Focusable() {
			focusable = append(focusable, widget)
		}
	}

	return focusable
}

func (tracker *FocusTracker) focusableAt(idx int) Wtfable {
	if idx < 0 || idx >= len(tracker.focusables()) {
		return nil
	}

	return tracker.focusables()[idx]
}

func (tracker *FocusTracker) focusState() FocusState {
	if tracker.Idx < 0 {
		return neverFocused
	}

	for _, widget := range tracker.Widgets {
		if widget.TextView() == tracker.App.GetFocus() {
			return widgetFocused
		}
	}

	return appBoardFocused
}

func (tracker *FocusTracker) increment() {
	tracker.Idx = tracker.Idx + 1

	if tracker.Idx == len(tracker.focusables()) {
		tracker.Idx = 0
	}
}

func (tracker *FocusTracker) useNavShortcuts() bool {
	return Config.UBool("wtf.navigation.shortcuts", true)
}
