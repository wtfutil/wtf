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

func NewWidget(tviewApp *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),
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

	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	builds, err := widget.Client.BuildsFor()

	title := fmt.Sprintf("%s - Builds", widget.CommonSettings().Title)
	var str string
	wrap := false
	if err != nil {
		wrap = true
		str = err.Error()
	} else {
		for idx, build := range builds {
			if idx > widget.settings.numberOfBuilds {
				break
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
	}

	return title, str, wrap
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
