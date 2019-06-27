package hibp

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

// Widget is the container for hibp data
type Widget struct {
	wtf.TextWidget

	accounts []string
	settings *Settings
}

// NewWidget creates a new instance of a widget
func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := &Widget{
		TextWidget: wtf.NewTextWidget(app, settings.common, false),

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
	data, err := widget.Fetch(widget.settings.accounts)

	title := widget.CommonSettings.Title
	title = title + widget.sinceDateForTitle()

	var content string
	if err != nil {
		content = err.Error()
	} else {
		content = widget.contentFrom(data)
	}

	widget.Redraw(title, content, false)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(data []*Status) string {
	str := ""

	for _, stat := range data {
		color := widget.settings.colors.ok
		if stat.HasBeenCompromised() {
			color = widget.settings.colors.pwned
		}

		str = str + fmt.Sprintf(" [%s]%s[white]\n", color, stat.Account)
	}

	return str
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
