package security

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	settings *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, settings.common, false),

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {

	if widget.Disabled() {
		return
	}

	data := NewSecurityData()
	data.Fetch()

	widget.Redraw(widget.CommonSettings().Title, widget.contentFrom(data), false)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(data *SecurityData) string {
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

	return str
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
