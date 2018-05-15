package system

import (
	"fmt"
	"os/exec"
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

	widget.buildSystemInfo()

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

func (widget *Widget) buildSystemInfo() {
	arg := []string{}

	cmd := exec.Command("sw_vers", arg...)
	str := wtf.ExecuteCommand(cmd)

	widget.systemInfo = NewSystemInfo(str)
}
