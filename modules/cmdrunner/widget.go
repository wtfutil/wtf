package cmdrunner

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	args     []string
	cmd      string
	settings *Settings
}

// NewWidget creates a new instance of the widget
func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, settings.common),

		args:     settings.args,
		cmd:      settings.cmd,
		settings: settings,
	}

	widget.View.SetWrap(true)

	return &widget
}

func (widget *Widget) content() (string, string, bool) {
	result := widget.execute()

	ansiTitle := tview.TranslateANSI(widget.CommonSettings().Title)
	if ansiTitle == defaultTitle {
		ansiTitle = tview.TranslateANSI(widget.String())
	}
	ansiResult := tview.TranslateANSI(result)

	return ansiTitle, ansiResult, false
}

// Refresh executes the command and updates the view with the results
func (widget *Widget) Refresh() {
	widget.Redraw(widget.content)
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
	return utils.ExecuteCommand(cmd)
}
