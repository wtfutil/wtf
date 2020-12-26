package system

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	Date    string
	Version string

	settings   *Settings
	systemInfo *SystemInfo
}

func NewWidget(tviewApp *tview.Application, date, version string, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),

		Date: date,

		settings: settings,
		Version:  version,
	}

	widget.systemInfo = NewSystemInfo()

	return &widget
}

func (widget *Widget) display() (string, string, bool) {
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

	return widget.CommonSettings().Title, content, false
}

func (widget *Widget) Refresh() {
	widget.Redraw(widget.display)
}

func (widget *Widget) prettyDate() string {
	str, err := time.Parse(utils.TimestampFormat, widget.Date)

	if err != nil {
		return err.Error()
	}

	return str.Format("Jan _2, 15:04")
}
