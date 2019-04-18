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
		TextWidget: wtf.NewTextWidget(app, settings.common.Name, settings.common.ConfigKey, false),

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

	widget.View.SetText(widget.contentFrom(data))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(data *SecurityData) string {
	str := " [red]WiFi[white]\n"
	str = str + fmt.Sprintf(" %8s: %s\n", "Network", data.WifiName)
	str = str + fmt.Sprintf(" %8s: %s\n", "Crypto", data.WifiEncryption)
	str = str + "\n"

	str = str + " [red]Firewall[white]\n"
	str = str + fmt.Sprintf(" %8s: %4s\n", "Status", data.FirewallEnabled)
	str = str + fmt.Sprintf(" %8s: %4s\n", "Stealth", data.FirewallStealth)
	str = str + "\n"

	str = str + " [red]Users[white]\n"
	str = str + fmt.Sprintf("  %s", strings.Join(data.LoggedInUsers, "\n  "))
	str = str + "\n\n"

	str = str + " [red]DNS[white]\n"
	str = str + fmt.Sprintf("  %12s\n", data.DnsAt(0))
	str = str + fmt.Sprintf("  %12s\n", data.DnsAt(1))
	str = str + "\n"

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
