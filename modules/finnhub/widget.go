package finnhub

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

// Widget ..
type Widget struct {
	view.TextWidget
	*Client

	settings *Settings
}

// NewWidget ..
func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, settings.common),
		Client:     NewClient(settings.symbols, settings.apiKey),

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	quotes, err := widget.Client.Getquote()

	title := fmt.Sprintf("%s - from finnhub api", widget.CommonSettings().Title)
	var str string
	wrap := false
	if err != nil {
		wrap = true
		str = err.Error()
	} else {
		for idx, q := range quotes {
			if idx > 10 {
				break
			}

			str += fmt.Sprintf(
				"[%s]: %s \n",
				q.Stock,
				q.C,
			)
		}
	}

	return title, str, wrap
}
