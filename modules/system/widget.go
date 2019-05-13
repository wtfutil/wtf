package system

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	Date    string
	Version string

	settings   *Settings
	systemInfo *SystemInfo
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings, date, version string) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, pages, settings.common, false),

		Date: date,

		settings: settings,
		Version:  version,
	}

	widget.SetRefreshFunction(widget.Refresh)

	widget.systemInfo = NewSystemInfo()

	return &widget
}

func (widget *Widget) Refresh() {
	content := fmt.Sprintf(
		"%8s: %s\n%8s: %s\n\n%8s: %s\n%8s: %s",
		"Built",
		widget.prettyDate(),
		"Vers",
		widget.Version,
		"OS",
		widget.systemInfo.ProductVersion,
		"Build",
		widget.systemInfo.BuildVersion,
	)
	widget.Redraw(widget.CommonSettings.Title, content, false)
}

func (widget *Widget) prettyDate() string {
	str, err := time.Parse(wtf.TimestampFormat, widget.Date)

	if err != nil {
		return err.Error()
	}

	return str.Format("Jan _2, 15:04")
}
