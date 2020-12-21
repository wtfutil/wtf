package bamboohr

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
	"github.com/wtfutil/wtf/wtf"
)

const apiURI = "https://api.bamboohr.com/api/gateway.php"

type Widget struct {
	view.TextWidget

	settings *Settings
	items    []Item
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
	client := NewClient(
		apiURI,
		widget.settings.apiKey,
		widget.settings.subdomain,
	)

	widget.items = client.Away(
		"timeOff",
		time.Now().Local().Format(wtf.DateFormat),
		time.Now().Local().Format(wtf.DateFormat),
	)

	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	str := ""
	if len(widget.items) == 0 {
		str = fmt.Sprintf("\n\n\n\n\n\n\n\n%s", utils.CenterText("[grey]no one[white]", 50))
	} else {
		for _, item := range widget.items {
			str += widget.format(item)
		}
	}

	return widget.CommonSettings().Title, str, false
}

func (widget *Widget) format(item Item) string {
	var str string

	if item.IsOneDay() {
		str = fmt.Sprintf(" [green]%s[white]\n %s\n\n", item.Name(), item.PrettyEnd())
	} else {
		str = fmt.Sprintf(" [green]%s[white]\n %s - %s\n\n", item.Name(), item.PrettyStart(), item.PrettyEnd())
	}

	return str
}
