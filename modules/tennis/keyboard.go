package tennis

import "github.com/gdamore/tcell/v2"

// statusCycle is the order the 'l'/'h' keys move through the API's match
// statuses.
var statusCycle = []string{"live", "upcoming", "completed"}

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("l", widget.nextStatus, "Show the next match status")
	widget.SetKeyboardChar("h", widget.prevStatus, "Show the previous match status")

	widget.SetKeyboardKey(tcell.KeyRight, widget.nextStatus, "Show the next match status")
	widget.SetKeyboardKey(tcell.KeyLeft, widget.prevStatus, "Show the previous match status")
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) nextStatus() {
	widget.settings.status = shiftStatus(widget.settings.status, 1)
	widget.Refresh()
}

func (widget *Widget) prevStatus() {
	widget.settings.status = shiftStatus(widget.settings.status, -1)
	widget.Refresh()
}

// shiftStatus returns the status `offset` places away from `current` in
// statusCycle, wrapping at both ends.
func shiftStatus(current string, offset int) string {
	idx := 0
	for i, status := range statusCycle {
		if status == current {
			idx = i
			break
		}
	}

	next := (idx + offset) % len(statusCycle)
	if next < 0 {
		next += len(statusCycle)
	}

	return statusCycle[next]
}
