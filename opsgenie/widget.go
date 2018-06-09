package opsgenie

import (
	"fmt"
	"strings"

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
		TextWidget: wtf.NewTextWidget(" OpsGenie ", "opsgenie", false),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	data, err := Fetch()

	widget.UpdateRefreshedAt()
	widget.View.SetTitle(widget.Name)

	if err != nil {
		widget.View.SetWrap(true)
		widget.View.SetText(fmt.Sprintf("%s", err))
	} else {
		widget.View.SetWrap(false)
		widget.View.SetText(fmt.Sprintf("%s", widget.contentFrom(data)))
	}
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(onCallResponse *OnCallResponse) string {
	str := ""

	displayEmpty := Config.UBool("wtf.mods.opsgenie.displayEmpty", true)

	for _, data := range onCallResponse.OnCallData {
		if (len(data.Recipients) == 0) && (displayEmpty == false) {
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
