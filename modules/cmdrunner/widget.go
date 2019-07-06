package cmdrunner

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	args     []string
	cmd      string
	settings *Settings
}

// NewWidget creates a new instance of the widget
func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, settings.common, false),

		args:     settings.args,
		cmd:      settings.cmd,
		settings: settings,
	}

	widget.View.SetWrap(true)

	return &widget
}

// Refresh executes the command and updates the view with the results
func (widget *Widget) Refresh() {
	result := widget.execute()

	ansiTitle := tview.TranslateANSI(widget.CommonSettings().Title)
	if ansiTitle == defaultTitle {
		ansiTitle = tview.TranslateANSI(widget.String())
	}
	ansiResult := tview.TranslateANSI(result)

	widget.Redraw(ansiTitle, ansiResult, false)
}

// String returns the string representation of the widget
func (widget *Widget) String() string {
	args := strings.Join(widget.args, " ")

	if args != "" {
		return fmt.Sprintf(" %s %s ", widget.cmd, args)
	}

	return fmt.Sprintf(" %s ", widget.cmd)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) execute() string {
	cmd := exec.Command(widget.cmd, widget.args...)
	return wtf.ExecuteCommand(cmd)
}
