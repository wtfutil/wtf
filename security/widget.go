package security

import (
	"fmt"
	//"sort"
	"time"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

var Config *config.Config

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget("Security", "security"),
	}

	widget.addView()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	data := Fetch()

	widget.View.SetTitle(" 🤺 Security ")
	widget.RefreshedAt = time.Now()

	widget.View.Clear()
	fmt.Fprintf(widget.View, "%s", widget.contentFrom(data))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) addView() {
	view := tview.NewTextView()

	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorGray)
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name)
	view.SetWrap(false)

	widget.View = view
}

func (widget *Widget) contentFrom(data map[string]string) string {
	str := " [red]WiFi[white]\n"
	str = str + fmt.Sprintf(" %8s: %s\n", "Network", data["Network"])
	str = str + fmt.Sprintf(" %8s: %s\n", "Crypto", data["Encryption"])
	str = str + "\n"
	str = str + " [red]Firewall[white]\n"
	str = str + fmt.Sprintf(" %8s: %s\n", "Enabled", data["Enabled"])
	str = str + fmt.Sprintf(" %8s: %s\n", "Stealth", data["Stealth"])
	str = str + "\n"

	return str
}
