package covid

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"

)

type Widget struct {
	view.TextWidget

	settings  *Settings
	err       error
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, nil, settings.Common),

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
	title := defaultTitle
	if widget.CommonSettings().Title != "" {
		title = widget.CommonSettings().Title
	}

	latest, err := LatestCases()
	var str string
	if err != nil {
		widget.err = err
	} else {
		str = fmt.Sprintf("[%s]Covid[white]\n", widget.settings.Colors.Subheading)
		str += fmt.Sprintf("%s: %d\n", "Confirmed", latest.Confirmed)
		str += fmt.Sprintf("%s: %d\n", "Deaths", latest.Deaths)
		str += fmt.Sprintf("%s: %d\n", "Recovered", latest.Recovered)
	}
	return title, str, true
}
