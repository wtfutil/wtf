package lunarphase

import "github.com/gdamore/tcell/v2"

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("n", widget.NextDay, "Show next day lunar phase")
	widget.SetKeyboardChar("p", widget.PrevDay, "Show previous day lunar phase")
	widget.SetKeyboardChar("t", widget.Today, "Show today lunar phase")
	widget.SetKeyboardChar("N", widget.NextWeek, "Show next week lunar phase")
	widget.SetKeyboardChar("P", widget.PrevWeek, "Show previous week lunar phase")
	widget.SetKeyboardChar("o", widget.OpenMoonPhase, "Open 'Moon Phase for Today' in browser")

	widget.SetKeyboardKey(tcell.KeyLeft, widget.PrevDay, "Show previous day lunar phase")
	widget.SetKeyboardKey(tcell.KeyRight, widget.NextDay, "Show next day lunar phase")
	widget.SetKeyboardKey(tcell.KeyUp, widget.NextWeek, "Show next week lunar phase")
	widget.SetKeyboardKey(tcell.KeyDown, widget.PrevWeek, "Show previous week lunar phase")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.OpenMoonPhase, "Open 'Moon Phase for Today' in browser")
	widget.SetKeyboardKey(tcell.KeyCtrlD, widget.DisableWidget, "Disable/Enable this widget instance")
}
