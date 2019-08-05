package bamboohr

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
	"github.com/wtfutil/wtf/wtf"
)

const APIURI = "https://api.bamboohr.com/api/gateway.php"

type Widget struct {
	view.TextWidget

	settings *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, settings.common, false),

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	client := NewClient(
		APIURI,
		widget.settings.apiKey,
		widget.settings.subdomain,
	)

	todayItems := client.Away(
		"timeOff",
		time.Now().Local().Format(wtf.DateFormat),
		time.Now().Local().Format(wtf.DateFormat),
	)

	widget.Redraw(widget.CommonSettings().Title, widget.contentFrom(todayItems), false)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(items []Item) string {
	if len(items) == 0 {
		return fmt.Sprintf("\n\n\n\n\n\n\n\n%s", utils.CenterText("[grey]no one[white]", 50))
	}

	str := ""
	for _, item := range items {
		str += widget.format(item)
	}

	return str
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
