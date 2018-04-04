package opsgenie

import (
	"fmt"
	"strings"
	"time"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

var Config *config.Config

type Widget struct {
	wtf.BaseWidget
	View *tview.TextView
}

func NewWidget() *Widget {
	widget := Widget{
		BaseWidget: wtf.BaseWidget{
			Name:        "OpsGenie",
			RefreshedAt: time.Now(),
			RefreshInt:  Config.UInt("wtf.opsgenie.refreshInterval", 21600),
		},
	}

	widget.addView()
	go wtf.Refresh(&widget)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	data := Fetch()

	widget.View.SetTitle(" ⏰ On Call ")
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

func (widget *Widget) contentFrom(onCallResponse *OnCallResponse) string {
	str := "\n"

	for _, data := range onCallResponse.OnCallData {
		str = str + fmt.Sprintf(" [green]%s[white]\n", widget.cleanScheduleName(data.Parent.Name))

		if len(data.Recipients) == 0 {
			str = str + " [gray]no one[white]\n"
		} else {
			str = str + fmt.Sprintf(" %s\n", strings.Join(widget.namesFromEmails(data.Recipients), ", "))
		}

		str = str + "\n"
	}

	return str
}

func (widget *Widget) cleanScheduleName(schedule string) string {
	return strings.Replace(schedule, "_", " ", -1)
}

func (widget *Widget) namesFromEmails(emails []string) []string {
	names := []string{}

	for _, email := range emails {
		parts := strings.Split(email, "@")
		names = append(names, strings.Title(strings.Replace(parts[0], ".", " ", -1)))
	}

	return names
}
