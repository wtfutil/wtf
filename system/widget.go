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
	BuiltAt    string
	Version    string
}

func NewWidget(builtAt, version string) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Build ", "system", false),

		BuiltAt: builtAt,
		Version: version,
	}

	widget.systemInfo = NewSystemInfo()

	return &widget
}

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.View.Clear()

	fmt.Fprintf(
		widget.View,
		"%8s: %s\n%8s: %s\n\n%8s: %s\n%8s: %s",
		"Built",
		widget.prettyBuiltAt(),
		"Vers",
		widget.Version,
		"OS",
		widget.systemInfo.ProductVersion,
		"Build",
		widget.systemInfo.BuildVersion,
	)

	widget.RefreshedAt = time.Now()
}

func (widget *Widget) prettyBuiltAt() string {
	str, err := time.Parse(wtf.TimestampFormat, widget.BuiltAt)
	if err != nil {
		return err.Error()
	} else {
		return str.Format("Jan _2, 15:04")
	}
}
