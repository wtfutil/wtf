package system

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	systemInfo *SystemInfo
	Date       string
	Version    string
}

func NewWidget(app *tview.Application, date, version string) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, "System", "system", false),

		Date:    date,
		Version: version,
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
