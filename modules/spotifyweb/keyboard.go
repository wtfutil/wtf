package spotifyweb

import (
	"time"

	"github.com/gdamore/tcell"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("h", widget.selectPrevious, "Select previous item")
	widget.SetKeyboardChar("l", widget.selectNext, "Select next item")
	widget.SetKeyboardChar(" ", widget.playPause, "Play/pause")
	widget.SetKeyboardChar("s", widget.toggleShuffle, "Toggle shuffle")

	widget.SetKeyboardKey(tcell.KeyDown, widget.selectNext, "Select next item")
	widget.SetKeyboardKey(tcell.KeyUp, widget.selectPrevious, "Select previous item")
}

func (widget *Widget) selectPrevious() {
	err := widget.client.Previous()
	if err != nil {
		return
	}

	time.Sleep(time.Millisecond * 500)
	widget.Refresh()
}

func (widget *Widget) selectNext() {
	err := widget.client.Next()
	if err != nil {
		return
	}

	time.Sleep(time.Millisecond * 500)
	widget.Refresh()
}

func (widget *Widget) playPause() {
	var err error
	if widget.playerState.CurrentlyPlaying.Playing {
		err = widget.client.Pause()
	} else {
		err = widget.client.Play()
	}
	if err != nil {
		return
	}

	time.Sleep(time.Millisecond * 500)
	widget.Refresh()
}

func (widget *Widget) toggleShuffle() {
	widget.playerState.ShuffleState = !widget.playerState.ShuffleState
	err := widget.client.Shuffle(widget.playerState.ShuffleState)
	if err != nil {
		return
	}

	time.Sleep(time.Millisecond * 500)
	widget.Refresh()
}
