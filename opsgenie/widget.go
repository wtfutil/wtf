package opsgenie

import (
	"fmt"
	"strings"
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
		TextWidget: wtf.NewTextWidget(" ⏰ OpsGenie ", "opsgenie", false),
	}

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

func (widget *Widget) contentFrom(onCallResponse *OnCallResponse) string {
	str := ""

	hideEmpty := Config.UBool("wtf.mods.opsgenie.hideEmpty", false)

	for _, data := range onCallResponse.OnCallData {
		if (len(data.Recipients) == 0) && (hideEmpty == true) {
			continue
		}

		var msg string
		if len(data.Recipients) == 0 {
			msg = " [gray]no one[white]\n\n"
		} else {
			msg = fmt.Sprintf(" %s\n\n", strings.Join(wtf.NamesFromEmails(data.Recipients), ", "))
		}

		str = str + widget.cleanScheduleName(data.Parent.Name)
		str = str + msg
	}

	return str
}

func (widget *Widget) cleanScheduleName(schedule string) string {
	cleanedName := strings.Replace(schedule, "_", " ", -1)
	return fmt.Sprintf(" [green]%s[white]\n", cleanedName)
}
