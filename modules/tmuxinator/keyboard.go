package tmuxinator

import (
	"github.com/gdamore/tcell/v2"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("j", widget.Next, "Select next project")
	widget.SetKeyboardChar("k", widget.Prev, "Select previous project")

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next project")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous project")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.startProject, "Start or go to project")
	widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect, "Clear selection")
}

// Next cycles the currently highlighted text down
func (widget *Widget) Next() {
	widget.Selected++
	if widget.Selected >= widget.maxItems {
		widget.Selected = 0
	}
	// widget.View.Highlight(strconv.Itoa(widget.Selected))
	// widget.View.ScrollToHighlight()
	widget.display()
}

// Prev cycles the currently highlighted text up
func (widget *Widget) Prev() {
	widget.Selected--
	if widget.Selected < 0 {
		widget.Selected = widget.maxItems - 1
	}
	// widget.View.Highlight(strconv.Itoa(widget.Selected))
	// widget.View.ScrollToHighlight()
	widget.display()
}

// Unselect stops highlighting the text and jumps the scroll position to the top
func (widget *Widget) Unselect() {
	widget.Selected = -1
	widget.View.Highlight()
	widget.View.ScrollToBeginning()
}

func (widget *Widget) startProject() {
	if widget.Selected >= 0 && len(widget.Items) > 0 {
		projectName := widget.Items[widget.Selected]

		startTmuxProject(projectName)
	}
}

