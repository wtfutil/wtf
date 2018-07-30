// +build windows

package security

import (
	"fmt"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget("Security", "security", false),
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

	widget.UpdateRefreshedAt()
	widget.View.SetText(widget.contentFrom(data))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(data *SecurityData) string {
	str := " [red]WiFi[white]\n"
	str = str + fmt.Sprintf(" %8s: %s\n", "Network", data.WifiName)
	str = str + fmt.Sprintf(" %8s: %s\n", "Crypto", data.WifiEncryption)
	str = str + "\n"
	str = str + " [red]Firewall[white]          [red]DNS[white]\n"
	str = str + fmt.Sprintf(" %8s: %4s %12s\n", "Enabled", data.FirewallEnabled, data.DnsAt(0))
	str = str + fmt.Sprintf(" %8s: %4s %12s\n", "Stealth", data.FirewallStealth, data.DnsAt(1))
	str = str + "\n"
	str = str + " [red]Users[white]\n"
	str = str + fmt.Sprintf(" %s", strings.Join(data.LoggedInUsers, ","))

	return str
}
