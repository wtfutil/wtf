package hibp

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

// Widget is the container for hibp data
type Widget struct {
	view.TextWidget

	settings *Settings
	statuses []*Status
	err      error
}

// NewWidget creates a new instance of a widget
func NewWidget(tviewApp *tview.Application, settings *Settings) *Widget {
	widget := &Widget{
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),

		settings: settings,
	}

	return widget
}

/* -------------------- Exported Functions -------------------- */

// Fetch retrieves HIBP data from the HIBP API
func (widget *Widget) Fetch(accounts []string) ([]*Status, error) {
	data := []*Status{}

	for _, account := range accounts {
		stat, err := widget.fetchForAccount(account, widget.settings.since)
		if err != nil {
			return nil, err
		}

		data = append(data, stat)
	}

	return data, nil
}

// Refresh updates the data for this widget and displays it onscreen
func (widget *Widget) Refresh() {
	statuses, err := widget.Fetch(widget.settings.accounts)

	if err != nil {
		widget.err = err
		widget.statuses = nil
	} else {
		widget.err = nil
		widget.statuses = statuses
	}

	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	title := widget.CommonSettings().Title
	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	title += widget.sinceDateForTitle()
	str := ""

	for _, status := range widget.statuses {
		color := widget.settings.colors.ok

		if status.HasBeenCompromised() {
			color = widget.settings.colors.pwned
		}

		if status != nil {
			str += fmt.Sprintf(" [%s]%s[white]\n", color, status.Account)
		}
	}

	return title, str, false
}

func (widget *Widget) sinceDateForTitle() string {
	dateStr := ""

	if widget.settings.HasSince() {
		sinceStr := ""

		dt, err := widget.settings.SinceDate()
		if err != nil {
			sinceStr = widget.settings.since
		} else {
			sinceStr = dt.Format("Jan _2, 2006")
		}

		dateStr = dateStr + " since " + sinceStr
	}

	return dateStr
}
