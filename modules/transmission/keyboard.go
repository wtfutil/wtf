package transmission

// import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp, "Show/hide this help prompt")
	widget.SetKeyboardChar("m", nil, "Add new magnet torrent")
	widget.SetKeyboardChar("p", nil, "Pause torrent")
	widget.SetKeyboardChar("r", nil, "Remove torrent from list")
}
