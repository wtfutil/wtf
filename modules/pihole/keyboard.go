package pihole

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("d", widget.disable, "disable Pi-hole")
	widget.SetKeyboardChar("e", widget.enable, "enable Pi-hole")
}
