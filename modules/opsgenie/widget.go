package opsgenie

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget
}

func NewWidget(app *tview.Application) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, "OpsGenie", "opsgenie", false),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	data, err := Fetch(
		wtf.Config.UString("wtf.mods.opsgenie.scheduleIdentifierType"),
		getSchedules(),
	)
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

func getSchedules() []string {
	// see if schedule is set to a single string
	configPath := "wtf.mods.opsgenie.schedule"
	singleSchedule, err := wtf.Config.String(configPath)
	if err == nil {
		return []string{singleSchedule}
	}
	// else, assume list
	scheduleList := wtf.Config.UList(configPath)
	var ret []string
	for _, schedule := range scheduleList {
		if str, ok := schedule.(string); ok {
			ret = append(ret, str)
		}
	}
	return ret
}

func (widget *Widget) contentFrom(onCallResponses []*OnCallResponse) string {
	str := ""

	displayEmpty := wtf.Config.UBool("wtf.mods.opsgenie.displayEmpty", true)

	for _, data := range onCallResponses {
		if (len(data.OnCallData.Recipients) == 0) && (displayEmpty == false) {
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
