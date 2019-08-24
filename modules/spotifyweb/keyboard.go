package spotifyweb

import (
	"time"

	"github.com/gdamore/tcell"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeCommonControls(widget.Refresh)

	widget.SetKeyboardChar("h", widget.selectPrevious, "Select previous item")
	widget.SetKeyboardChar("l", widget.selectNext, "Select next item")
	widget.SetKeyboardChar(" ", widget.playPause, "Play/pause")
	widget.SetKeyboardChar("s", widget.toggleShuffle, "Toggle shuffle")

	widget.SetKeyboardKey(tcell.KeyDown, widget.selectNext, "Select next item")
	widget.SetKeyboardKey(tcell.KeyUp, widget.selectPrevious, "Select previous item")
}

func (widget *Widget) selectPrevious() {
	widget.client.Previous()
	time.Sleep(time.Millisecond * 500)
	widget.Refresh()
}

func (widget *Widget) selectNext() {
	widget.client.Next()
	time.Sleep(time.Millisecond * 500)
	widget.Refresh()
}

func (widget *Widget) playPause() {
	if widget.playerState.CurrentlyPlaying.Playing {
		widget.client.Pause()
	} else {
		widget.client.Play()
	}
	time.Sleep(time.Millisecond * 500)
	widget.Refresh()
}

func (widget *Widget) toggleShuffle() {
	widget.playerState.ShuffleState = !widget.playerState.ShuffleState
	widget.client.Shuffle(widget.playerState.ShuffleState)
	time.Sleep(time.Millisecond * 500)
	widget.Refresh()
}
