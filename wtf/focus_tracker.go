package wtf

import (
	"github.com/rivo/tview"
)

type FocusState int

const (
	Widget FocusState = iota
	NonWidget
	NeverFocused
)

// FocusTracker is used by the app to track which onscreen widget currently has focus,
// and to move focus between widgets.
type FocusTracker struct {
	App     *tview.Application
	Idx     int
	Widgets []Wtfable
}

/* -------------------- Exported Functions -------------------- */

// Next sets the focus on the next widget in the widget list. If the current widget is
// the last widget, sets focus on the first widget.
func (tracker *FocusTracker) Next() {
	if tracker.focusState() == NonWidget {
		return
	}

	tracker.blur(tracker.Idx)
	tracker.increment()
	tracker.focus(tracker.Idx)
}

// None removes focus from the currently-focused widget.
func (tracker *FocusTracker) None() {
	if tracker.focusState() == NonWidget {
		return
	}

	tracker.blur(tracker.Idx)
}

// Prev sets the focus on the previous widget in the widget list. If the current widget is
// the last widget, sets focus on the last widget.
func (tracker *FocusTracker) Prev() {
	if tracker.focusState() == NonWidget {
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

	tracker.App.SetFocus(view)
	view.SetBorderColor(ColorFor(Config.UString("wtf.colors.border.focused", "gray")))
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

func (tracker *FocusTracker) increment() {
	tracker.Idx = tracker.Idx + 1

	if tracker.Idx == len(tracker.focusables()) {
		tracker.Idx = 0
	}
}

// widgetHasFocus returns true if one of the widgets currently has the app's focus,
// false if none of them do (ie: perhaps a modal dialog currently has it instead)
// If there's no index, it returns true because focus has never been assigned
func (tracker *FocusTracker) focusState() FocusState {
	if tracker.Idx < 0 {
		return NeverFocused
	}

	for _, widget := range tracker.Widgets {
		if widget.TextView() == tracker.App.GetFocus() {
			return Widget
		}
	}

	return NonWidget
}
