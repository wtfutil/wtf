package twitch

import "github.com/gdamore/tcell/v2"

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("j", widget.Next, "Select next item")
	widget.SetKeyboardChar("k", widget.Prev, "Select previous item")
	widget.SetKeyboardChar("o", widget.openTwitch, "Open target URL in browser")
	widget.SetKeyboardChar("s", widget.openStreamlink, "Open target stream via streamlink (github.com/streamlink/streamlink)")

	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next item")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous item")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openTwitch, "Open stream in browser")
	widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect, "Clear selection")
}
