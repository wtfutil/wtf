package opsgenie

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
	data, err := widget.Fetch(
		widget.settings.scheduleIdentifierType,
		widget.settings.schedule,
	)

	var content string
	wrap := false
	if err != nil {
		wrap = true
		content = err.Error()
	} else {
		content = widget.contentFrom(data)
	}

	widget.Redraw(widget.CommonSettings.Title, content, wrap)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(onCallResponses []*OnCallResponse) string {
	str := ""

	for _, data := range onCallResponses {
		if (len(data.OnCallData.Recipients) == 0) && (widget.settings.displayEmpty == false) {
			continue
		}

		var msg string
		if len(data.OnCallData.Recipients) == 0 {
			msg = " [gray]no one[white]\n\n"
		} else {
			msg = fmt.Sprintf(" %s\n\n", strings.Join(wtf.NamesFromEmails(data.OnCallData.Recipients), ", "))
		}

		str += widget.cleanScheduleName(data.OnCallData.Parent.Name)
		str += msg
	}

	return str
}

func (widget *Widget) cleanScheduleName(schedule string) string {
	cleanedName := strings.Replace(schedule, "_", " ", -1)
	return fmt.Sprintf(" [green]%s[white]\n", cleanedName)
}
