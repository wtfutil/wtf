package wtf

import (
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/color"
)

type FocusTracker struct {
	App     *tview.Application
	Idx     int
	Widgets []TextViewer
}

/* -------------------- Exported Functions -------------------- */

func (tracker *FocusTracker) Next() {
	tracker.blur(tracker.Idx)
	tracker.increment()
	tracker.focus(tracker.Idx)
}

func (tracker *FocusTracker) None() {
	tracker.blur(tracker.Idx)
}

func (tracker *FocusTracker) Prev() {
	tracker.blur(tracker.Idx)
	tracker.decrement()
	tracker.focus(tracker.Idx)
}

/* -------------------- Exported Functions -------------------- */

func (tracker *FocusTracker) blur(idx int) {
	view := tracker.Widgets[idx].TextView()
	view.Blur()
	view.SetBorderColor(color.ColorFor(Config.UString("wtf.border.normal")))
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
	view.SetBorderColor(color.ColorFor(Config.UString("wtf.border.focus")))
}

func (tracker *FocusTracker) increment() {
	tracker.Idx = tracker.Idx + 1

	if tracker.Idx == len(tracker.Widgets) {
		tracker.Idx = 0
	}
}
