package wtf

import (
	"github.com/rivo/tview"
)

// FocusTracker is used by the app to track which onscreen widget currently has focus,
// and to move focus between widgets.
type FocusTracker struct {
	App     *tview.Application
	Idx     int
	Widgets []TextViewer
}

/* -------------------- Exported Functions -------------------- */

// Next sets the focus on the next widget in the widget list. If the current widget is
// the last widget, sets focus on the first widget.
func (tracker *FocusTracker) Next() {
	tracker.blur(tracker.Idx)
	tracker.increment()
	tracker.focus(tracker.Idx)
}

// None removes focus from the currently-focused widget.
func (tracker *FocusTracker) None() {
	tracker.blur(tracker.Idx)
}

// Prev sets the focus on the previous widget in the widget list. If the current widget is
// the last widget, sets focus on the last widget.
func (tracker *FocusTracker) Prev() {
	tracker.blur(tracker.Idx)
	tracker.decrement()
	tracker.focus(tracker.Idx)
}

/* -------------------- Exported Functions -------------------- */

func (tracker *FocusTracker) blur(idx int) {
	view := tracker.Widgets[idx].TextView()
	view.Blur()
	view.SetBorderColor(ColorFor(Config.UString("wtf.colors.border.normal", "gray")))
}

func (tracker *FocusTracker) decrement() {
	tracker.Idx = tracker.Idx - 1

	if tracker.Idx < 0 {
		tracker.Idx = len(tracker.Widgets) - 1
	}
}

func (tracker *FocusTracker) focus(idx int) {
	view := tracker.Widgets[idx].TextView()
	tracker.App.SetFocus(view)
	view.SetBorderColor(ColorFor(Config.UString("wtf.colors.border.focus", "gray")))
}

func (tracker *FocusTracker) increment() {
	tracker.Idx = tracker.Idx + 1

	if tracker.Idx == len(tracker.Widgets) {
		tracker.Idx = 0
	}
}
