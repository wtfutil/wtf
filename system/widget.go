package system

import (
	"fmt"
	"time"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget

	systemInfo *SystemInfo
	Date       string
	Version    string
}

func NewWidget(date, version string) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" System and Build Info ", "system", false),

		Date:    date,
		Version: version,
	}

	widget.systemInfo = NewSystemInfo()

	return &widget
}

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.UpdateRefreshedAt()
	widget.View.Clear()

	fmt.Fprintf(
		widget.View,
		" %s: %s\n %s: %s\n %s: %s\n %s: %s\n %s: %s\n %s: %s\n %s: %s\n %s: %s\n %s: %d",
		"Built",
		widget.prettyDate(),
		"Version",
		widget.Version,
		"GoOS",
		widget.systemInfo.GoOs,
		"OS",
		widget.systemInfo.OS,
		"Platform",
		widget.systemInfo.Platform,
		"Kernel",
		widget.systemInfo.Kernel,
		"Kernel Version",
		widget.systemInfo.Version,
		"Hostname",
		widget.systemInfo.Hostname,
		"CPUs",
		widget.systemInfo.CPUs,
	)
}

func (widget *Widget) prettyDate() string {
	//if the date is not set in the build, print empty string instead of error
	if widget.Date == "" {
		return ""
	}
	str, err := time.Parse(wtf.TimestampFormat, widget.Date)
	if err != nil {
		return err.Error()
	} else {
		return str.Format("Jan _2, 15:04")
	}
}
