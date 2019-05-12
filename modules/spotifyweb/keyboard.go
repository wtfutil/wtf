package spotifyweb

import (
	"time"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp, "Show/hide this help prompt")
	widget.SetKeyboardChar("h", widget.selectPrevious, "Select previous item")
	widget.SetKeyboardChar("l", widget.selectNext, "Select next item")
	widget.SetKeyboardChar(" ", widget.playPause, "Play/pause")
	widget.SetKeyboardChar("s", widget.toggleShuffle, "Toggle shuffle")
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
