package system

import (
	"fmt"
	"time"

	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	systemInfo *SystemInfo
	Date       string
	Version    string
}

func NewWidget(date, version string) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" System ", "system", false),

		Date:    date,
		Version: version,
	}

	widget.systemInfo = NewSystemInfo()

	return &widget
}

func (widget *Widget) Refresh() {
	widget.UpdateRefreshedAt()

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
	} else {
		return str.Format("Jan _2, 15:04")
	}
}
