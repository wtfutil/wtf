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

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, settings.common),

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
	str := " [red]WiFi[white]\n"
	str += fmt.Sprintf(" %8s: %s\n", "Network", data.WifiName)
	str += fmt.Sprintf(" %8s: %s\n", "Crypto", data.WifiEncryption)
	str += "\n"

	str += " [red]Firewall[white]\n"
	str += fmt.Sprintf(" %8s: %4s\n", "Status", data.FirewallEnabled)
	str += fmt.Sprintf(" %8s: %4s\n", "Stealth", data.FirewallStealth)
	str += "\n"

	str += " [red]Users[white]\n"
	str += fmt.Sprintf("  %s", strings.Join(data.LoggedInUsers, "\n  "))
	str += "\n\n"

	str += " [red]DNS[white]\n"
	str += fmt.Sprintf("  %12s\n", data.DnsAt(0))
	str += fmt.Sprintf("  %12s\n", data.DnsAt(1))
	str += "\n"

	return widget.CommonSettings().Title, str, false
}

func (widget *Widget) labelColor(label string) string {
	switch label {
	case "on":
		return "green"
	case "off":
		return "red"
	default:
		return "white"
	}
}
