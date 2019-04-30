package opsgenie

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	app      *tview.Application
	settings *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, settings.common, false),

		app:      app,
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
	if err != nil {
		widget.View.SetWrap(true)
		content = err.Error()
	} else {
		widget.View.SetWrap(false)
		content = widget.contentFrom(data)
	}

	widget.app.QueueUpdateDraw(func() {
		widget.View.SetTitle(widget.ContextualTitle(widget.CommonSettings.Title))
		widget.View.SetText(content)
	})
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

		str = str + widget.cleanScheduleName(data.OnCallData.Parent.Name)
		str = str + msg
	}

	return str
}

func (widget *Widget) cleanScheduleName(schedule string) string {
	cleanedName := strings.Replace(schedule, "_", " ", -1)
	return fmt.Sprintf(" [green]%s[white]\n", cleanedName)
}
