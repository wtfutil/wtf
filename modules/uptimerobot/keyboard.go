package uptimerobot

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp, "Show/hide this help widget")
	widget.SetKeyboardChar("r", widget.Refresh, "Refresh widget")
}
