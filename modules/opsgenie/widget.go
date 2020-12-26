package opsgenie

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	settings *Settings
}

func NewWidget(tviewApp *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	onCallResponses, err := widget.Fetch(
		widget.settings.scheduleIdentifierType,
		widget.settings.schedule,
	)
	title := widget.CommonSettings().Title

	var content string
	wrap := false
	if err != nil {
		wrap = true
		content = err.Error()
	} else {

		for _, data := range onCallResponses {
			if (len(data.OnCallData.Recipients) == 0) && !widget.settings.displayEmpty {
				continue
			}

			var msg string
			if len(data.OnCallData.Recipients) == 0 {
				msg = " [gray]no one[white]\n\n"
			} else {
				msg = fmt.Sprintf(" %s\n\n", strings.Join(utils.NamesFromEmails(data.OnCallData.Recipients), ", "))
			}

			content += widget.cleanScheduleName(data.OnCallData.Parent.Name)
			content += msg
		}
	}

	return title, content, wrap
}

func (widget *Widget) cleanScheduleName(schedule string) string {
	cleanedName := strings.Replace(schedule, "_", " ", -1)
	return fmt.Sprintf(" [green]%s[white]\n", cleanedName)
}
