package app

import (
	"fmt"
	"sort"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
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
	Idx       int
	IsFocused bool
	Widgets   []wtf.Wtfable

	config   *config.Config
	tviewApp *tview.Application
}

func NewFocusTracker(tviewApp *tview.Application, widgets []wtf.Wtfable, config *config.Config) FocusTracker {
	focusTracker := FocusTracker{
		tviewApp:  tviewApp,
		Idx:       -1,
		IsFocused: false,
		Widgets:   widgets,

		config: config,
	}

	focusTracker.assignHotKeys()

	return focusTracker
}

/* -------------------- Exported Functions -------------------- */

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
			tracker.IsFocused = true
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

	tracker.IsFocused = true
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

	tracker.IsFocused = true
}

func (tracker *FocusTracker) Refocus() {
	tracker.focus(tracker.Idx)
}

/* -------------------- Unexported Functions -------------------- */

// AssignHotKeys assigns an alphabetic keyboard character to each focusable
// widget so that the widget can be brought into focus by pressing that keyboard key
// Valid numbers are between 1 and 9, inclusive
func (tracker *FocusTracker) assignHotKeys() {
	if !tracker.useNavShortcuts() {
		return
	}

	usedKeys := make(map[string]bool)
	focusables := tracker.focusables()

	// First, block out the explicitly-defined characters so they can't be automatically
	// assigned to other modules
	for _, focusable := range focusables {
		if focusable.FocusChar() != "" {
			usedKeys[focusable.FocusChar()] = true
		}
	}

	focusNum := 1

	// Range over all the modules and assign focus characters to any that are focusable
	// and don't have explicitly-defined focus characters
	for _, focusable := range focusables {
		if focusable.FocusChar() != "" {
			continue
		}

		if _, foundKey := usedKeys[fmt.Sprint(focusNum)]; foundKey {
			for ; foundKey; _, foundKey = usedKeys[fmt.Sprint(focusNum)] {
				focusNum++
			}
		}

		// Don't allow focus characters > "9"
		if focusNum >= 10 {
			break
		}

		focusable.SetFocusChar(fmt.Sprint(focusNum))
		focusNum++
	}
}

func (tracker *FocusTracker) blur(idx int) {
	widget := tracker.focusableAt(idx)
	if widget == nil {
		return
	}

	view := widget.TextView()
	view.Blur()

	view.SetBorderColor(
		wtf.ColorFor(
			widget.BorderColor(),
		),
	)

	tracker.IsFocused = false
}

func (tracker *FocusTracker) decrement() {
	tracker.Idx--

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
	view.SetBorderColor(
		wtf.ColorFor(
			widget.CommonSettings().Colors.BorderTheme.Focused,
		),
	)
	tracker.tviewApp.SetFocus(view)
}

func (tracker *FocusTracker) focusables() []wtf.Wtfable {
	focusable := []wtf.Wtfable{}

	for _, widget := range tracker.Widgets {
		if widget.Focusable() {
			focusable = append(focusable, widget)
		}
	}

	// Sort for deterministic ordering
	sort.SliceStable(focusable, func(i, j int) bool {
		iTop := focusable[i].CommonSettings().Top
		jTop := focusable[j].CommonSettings().Top

		if iTop < jTop {
			return true
		}
		if iTop == jTop {
			return focusable[i].CommonSettings().Left < focusable[j].CommonSettings().Left
		}
		return false
	})

	return focusable
}

func (tracker *FocusTracker) focusableAt(idx int) wtf.Wtfable {
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
		if widget.TextView() == tracker.tviewApp.GetFocus() {
			return widgetFocused
		}
	}

	return appBoardFocused
}

func (tracker *FocusTracker) increment() {
	tracker.Idx++

	if tracker.Idx == len(tracker.focusables()) {
		tracker.Idx = 0
	}
}

func (tracker *FocusTracker) useNavShortcuts() bool {
	return tracker.config.UBool("wtf.navigation.shortcuts", true)
}
