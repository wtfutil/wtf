package digitalocean

import "github.com/gdamore/tcell"

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeCommonControls(nil)

	widget.SetKeyboardChar("i", widget.showInfo, "Show info about the selected droplet")
	widget.SetKeyboardChar("j", widget.Prev, "Select previous item")
	widget.SetKeyboardChar("k", widget.Next, "Select next item")
	widget.SetKeyboardChar("r", widget.restart, "Reboot the selected droplet")
	widget.SetKeyboardChar("s", widget.shutDown, "Shut down the selected droplet")
	widget.SetKeyboardChar("u", widget.Unselect, "Clear selection")

	widget.SetKeyboardKey(tcell.KeyCtrlD, widget.destroySelectedDroplet, "Destroy the selected droplet")
	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next item")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous item")
}
