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
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" ⏰ OpsGenie ", "opsgenie"),
	}

	widget.addView()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	data, err := Fetch()

	widget.View.SetTitle(" ⏰ On Call ")

	widget.View.Clear()

	if err != nil {
		widget.View.SetWrap(true)
		fmt.Fprintf(widget.View, "%s", err)
	} else {
		widget.View.SetWrap(false)
		fmt.Fprintf(widget.View, "%s", widget.contentFrom(data))
	}

	widget.RefreshedAt = time.Now()
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
	str := ""

	for _, data := range onCallResponse.OnCallData {
		str = str + fmt.Sprintf(" [green]%s[white]\n", widget.cleanScheduleName(data.Parent.Name))

		if len(data.Recipients) == 0 {
			str = str + " [gray]no one[white]\n"
		} else {
			str = str + fmt.Sprintf(" %s\n", strings.Join(wtf.NamesFromEmails(data.Recipients), ", "))
		}

		str = str + "\n"
	}

	return str
}

func (widget *Widget) cleanScheduleName(schedule string) string {
	return strings.Replace(schedule, "_", " ", -1)
}
