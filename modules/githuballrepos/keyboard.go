package githuballrepos

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/wtfutil/wtf/logger"
	"github.com/wtfutil/wtf/utils"
)

func (widget *Widget) initializeKeyboardControls() {

	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("j", widget.Next, "Select next PR")
	widget.SetKeyboardChar("k", widget.Prev, "Select previous PR")
	widget.SetKeyboardChar("l", widget.nextTab, "Next tab")
	widget.SetKeyboardChar("h", widget.prevTab, "Previous tab")
	widget.SetKeyboardChar("o", widget.openSelected, "Open selected PR in browser")

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next PR")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous PR")
	widget.SetKeyboardKey(tcell.KeyRight, widget.nextTab, "Next tab")
	widget.SetKeyboardKey(tcell.KeyLeft, widget.prevTab, "Previous tab")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openSelected, "Open selected PR in browser")
}

func (widget *Widget) Next() {
	logger.Log("Next called")
	if widget.data == nil || len(widget.data.MyPRs) == 0 {
		logger.Log("No data or empty MyPRs")
		return
	}
	widget.selectedPRIndex = (widget.selectedPRIndex + 1) % len(widget.data.MyPRs)
	widget.isItemSelected = true
	logger.Log(fmt.Sprintf("Next completed. New index: %d\n", widget.selectedPRIndex))
}

func (widget *Widget) Prev() {
	logger.Log("Prev called")
	if widget.data == nil || len(widget.data.MyPRs) == 0 {
		logger.Log("No data or empty MyPRs")
		return
	}
	if widget.selectedPRIndex <= 0 {
		widget.selectedPRIndex = len(widget.data.MyPRs) - 1
	} else {
		widget.selectedPRIndex--
	}
	widget.isItemSelected = true
	logger.Log(fmt.Sprintf("Prev completed. New index: %d\n", widget.selectedPRIndex))
}

func (widget *Widget) nextTab() {
	widget.currentTab = (widget.currentTab + 1) % 3
	widget.selectedPRIndex = -1
	widget.isItemSelected = false
	if !widget.testMode {
		widget.display()
	}
}

func (widget *Widget) prevTab() {
	widget.currentTab--
	if widget.currentTab < 0 {
		widget.currentTab = 2
	}
	widget.selectedPRIndex = -1
	widget.isItemSelected = false

	if !widget.testMode {
		widget.display()
	}
}

func (widget *Widget) openSelected() {
	var prs []PR
	switch widget.currentTab {
	case 0:
		prs = widget.data.MyPRs
	case 1:
		prs = widget.data.PRReviewRequests
	case 2:
		prs = widget.data.WatchedPRs
	}

	if widget.selectedPRIndex >= 0 && widget.selectedPRIndex < len(prs) {
		utils.OpenFile(prs[widget.selectedPRIndex].URL)
	}
}
