package covid

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

// Widget is the struct that defines this module widget
type Widget struct {
	view.TextWidget

	settings *Settings
	err      error
}

// NewWidget creates a new widget for this module
func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := &Widget{
		TextWidget: view.NewTextWidget(app, nil, settings.Common),

		settings: settings,
	}

	return widget
}

// Refresh checks if this module widget is disabled
func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.Redraw(widget.content)
}

// Render renders this module widget
func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

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
