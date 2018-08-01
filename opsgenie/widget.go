package opsgenie

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
		TextWidget: wtf.NewTextWidget("OpsGenie", "opsgenie", false),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	data, err := Fetch()

	widget.UpdateRefreshedAt()
	widget.View.SetTitle(widget.ContextualTitle(widget.Name))

	var content string
	if err != nil {
		widget.View.SetWrap(true)
		content = err.Error()
	} else {
		widget.View.SetWrap(false)
		content = widget.contentFrom(data)
	}

	widget.View.SetText(content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(onCallResponse *OnCallResponse) string {
	str := ""

	displayEmpty := wtf.Config.UBool("wtf.mods.opsgenie.displayEmpty", true)

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
