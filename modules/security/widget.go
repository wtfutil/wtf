package security

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	settings *Settings
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.Common),

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {

	if widget.Disabled() {
		return
	}

	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	data := NewSecurityData()
	data.Fetch()
	var str string

	if data.WifiName != "" {
		str += fmt.Sprintf(" [%s]WiFi[white]\n", widget.settings.Colors.Subheading)
		str += fmt.Sprintf(" %8s: %s\n", "Network", data.WifiName)
		str += fmt.Sprintf(" %8s: %s\n", "Crypto", data.WifiEncryption)
		str += "\n"
	}

	str += fmt.Sprintf(" [%s]Firewall[white]\n", widget.settings.Colors.Subheading)
	str += fmt.Sprintf(" %8s: %4s\n", "Status", data.FirewallEnabled)
	str += fmt.Sprintf(" %8s: %4s\n", "Stealth", data.FirewallStealth)
	str += "\n"

	str += fmt.Sprintf(" [%s]Users[white]\n", widget.settings.Colors.Subheading)
	str += fmt.Sprintf("  %s", strings.Join(data.LoggedInUsers, "\n  "))
	str += "\n\n"

	str += fmt.Sprintf(" [%s]DNS[white]\n", widget.settings.Colors.Subheading)
	// If no DNS servers are found, display a single line of 'n/a'
	if len(data.Dns) == 0 {
		str += fmt.Sprintf(" %6s\n", "n/a")
	} else {
		for _, ip := range data.Dns {
			str += fmt.Sprintf(" %12s\n", ip)
		}
	}
	str += "\n"

	return widget.CommonSettings().Title, str, false
}
