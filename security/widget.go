package security

import (
	"fmt"
	"time"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" 🤺 Security ", "security"),
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

	widget.View.Clear()
	fmt.Fprintf(widget.View, "%s", widget.contentFrom(data))

	widget.RefreshedAt = time.Now()
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

	return str
}
