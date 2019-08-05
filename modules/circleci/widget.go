package circleci

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget
	*Client

	settings *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, settings.common, false),
		Client:     NewClient(settings.apiKey),

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	builds, err := widget.Client.BuildsFor()

	title := fmt.Sprintf("%s - Builds", widget.CommonSettings().Title)
	var content string
	wrap := false
	if err != nil {
		wrap = true
		content = err.Error()
	} else {
		content = widget.contentFrom(builds)
	}

	widget.Redraw(title, content, wrap)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(builds []*Build) string {
	var str string
	for idx, build := range builds {
		if idx > 10 {
			return str
		}

		str += fmt.Sprintf(
			"[%s] %s-%d (%s) [white]%s\n",
			buildColor(build),
			build.Reponame,
			build.BuildNum,
			build.Branch,
			build.AuthorName,
		)
	}

	return str
}

func buildColor(build *Build) string {
	switch build.Status {
	case "failed":
		return "red"
	case "running":
		return "yellow"
	case "success":
		return "green"
	case "fixed":
		return "green"
	default:
		return "white"
	}
}
