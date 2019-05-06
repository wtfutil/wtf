package spotify

import (
	"time"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp)
	widget.SetKeyboardChar("h", widget.previous)
	widget.SetKeyboardChar("l", widget.next)
	widget.SetKeyboardChar(" ", widget.playPause)
}

func (widget *Widget) previous() {
	widget.SpotifyClient.Previous()
	time.Sleep(time.Second * 1)
	widget.Refresh()
}

func (widget *Widget) next() {
	widget.SpotifyClient.Next()
	time.Sleep(time.Second * 1)
	widget.Refresh()
}

func (widget *Widget) playPause() {
	widget.SpotifyClient.PlayPause()
	time.Sleep(time.Second * 1)
	widget.Refresh()
}
