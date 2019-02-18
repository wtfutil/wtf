package circleci

import (
	"fmt"
	"os"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget
	*Client
}

const apiEnvKey = "WTF_CIRCLE_API_KEY"

func NewWidget(app *tview.Application) *Widget {
	apiKey := wtf.Config.UString(
		"wtf.mods.circleci.apiKey",
		os.Getenv(apiEnvKey),
	)

	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, "CircleCI", "circleci", false),
		Client:     NewClient(apiKey),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	builds, err := widget.Client.BuildsFor()

	widget.View.SetTitle(fmt.Sprintf("%s - Builds", widget.Name))

	var content string
	if err != nil {
		widget.View.SetWrap(true)
		content = err.Error()
	} else {
		widget.View.SetWrap(false)
		content = widget.contentFrom(builds)
	}

	widget.View.SetText(content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(builds []*Build) string {
	var str string
	for idx, build := range builds {
		if idx > 10 {
			return str
		}

		str = str + fmt.Sprintf(
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
