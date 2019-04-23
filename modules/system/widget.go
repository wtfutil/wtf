package system

import (
	"fmt"
	"time"

	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	Date       string
	Version    string
	settings   *Settings
	systemInfo *SystemInfo
}

func NewWidget(refreshChan chan<- string, date, version string, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(refreshChan, settings.common, false),

		Date:     date,
		settings: settings,
		Version:  version,
	}

	widget.systemInfo = NewSystemInfo()

	return &widget
}

func (widget *Widget) Refresh() {
	widget.View.SetText(
		fmt.Sprintf(
			"%8s: %s\n%8s: %s\n\n%8s: %s\n%8s: %s",
			"Built",
			widget.prettyDate(),
			"Vers",
			widget.Version,
			"OS",
			widget.systemInfo.ProductVersion,
			"Build",
			widget.systemInfo.BuildVersion,
		),
	)
}

func (widget *Widget) prettyDate() string {
	str, err := time.Parse(wtf.TimestampFormat, widget.Date)

	if err != nil {
		return err.Error()
	}

	return str.Format("Jan _2, 15:04")
}
