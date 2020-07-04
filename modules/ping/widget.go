package ping

import (
	"os/user"
	"sort"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.KeyboardWidget
	view.MultiSourceWidget
	view.TextWidget

	settings *Settings
}

// NewWidget creates a new instance of a widget
func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, settings.common),
		settings:   settings,
	}

	widget.settings.common.RefreshInterval = 30
	widget.View.SetInputCapture(widget.InputCapture)

	widget.SetDisplayFunction(widget.Refresh)
	widget.View.SetWordWrap(true)
	widget.View.SetWrap(settings.wrapText)

	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.Redraw(widget.content)
}

func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	title := widget.CommonSettings().Title

	cUser, _ := user.Current()
	if cUser != nil && cUser.Uid != "0" {
		return title, "module requires wtfutil is run as root", true
	}

	targets := parseTargets(widget.settings.targets)

	var outList results

	var ch = make(chan string, len(targets))

	for _, t := range targets {
		go func(t target) {
			ch <- getPingResult(widget, t, widget.settings.logging)
		}(t)
	}

	var res string

	for i := 1; i <= len(widget.settings.targets); i++ {
		res = <-ch
		outList = append(outList, res)
	}

	if widget.settings.useEmoji {
		sort.Strings(outList)
	} else {
		sort.Sort(outList)
	}

	output := strings.Join(outList, "\n")

	return title, output, widget.settings.wrapText
}
