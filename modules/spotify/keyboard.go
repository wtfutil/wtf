package spotify

import (
	"time"

	"github.com/gdamore/tcell"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("l", widget.next, "Select next item")
	widget.SetKeyboardChar("h", widget.previous, "Select previous item")
	widget.SetKeyboardChar(" ", widget.playPause, "Play/pause song")

	widget.SetKeyboardKey(tcell.KeyDown, widget.next, "Select next item")
	widget.SetKeyboardKey(tcell.KeyUp, widget.previous, "Select previous item")
}

func (widget *Widget) previous() {
	widget.client.Previous()
	time.Sleep(time.Second * 1)
	widget.Refresh()
}

func (widget *Widget) next() {
	widget.client.Next()
	time.Sleep(time.Second * 1)
	widget.Refresh()
}

func (widget *Widget) playPause() {
	widget.client.PlayPause()
	time.Sleep(time.Second * 1)
	widget.Refresh()
}
