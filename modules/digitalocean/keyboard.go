package digitalocean

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(widget.Refresh)

	widget.SetKeyboardChar("?", widget.showInfo, "Show info about the selected droplet")

	widget.SetKeyboardChar("b", widget.dropletRestart, "Reboot the selected droplet")
	widget.SetKeyboardChar("j", widget.Prev, "Select previous item")
	widget.SetKeyboardChar("k", widget.Next, "Select next item")
	widget.SetKeyboardChar("p", widget.dropletEnabledPrivateNetworking, "Enable private networking for the selected drople")
	widget.SetKeyboardChar("s", widget.dropletShutDown, "Shut down the selected droplet")
	widget.SetKeyboardChar("u", widget.Unselect, "Clear selection")

	widget.SetKeyboardKey(tcell.KeyCtrlD, widget.dropletDestroy, "Destroy the selected droplet")
	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next item")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.showInfo, "Show info about the selected droplet")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous item")
}
