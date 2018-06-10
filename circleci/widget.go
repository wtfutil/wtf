package circleci

import (
	"fmt"
	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" CircleCI ", "circleci", false),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	builds, err := BuildsFor()

	widget.UpdateRefreshedAt()

	widget.View.SetTitle(fmt.Sprintf("%s - Builds", widget.Name))

	if err != nil {
		widget.View.SetWrap(true)
		fmt.Fprintf(widget.View, "%v", err)
	} else {
		widget.View.SetWrap(false)
		widget.View.SetText(fmt.Sprintf("%s", widget.contentFrom(builds)))
	}
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

func buildColor(b *Build) string {
	var color string

	switch b.Status {
	case "failed":
		color = "red"
	case "running":
		color = "yellow"
	case "success":
		color = "green"
	default:
		color = "white"
	}

	return color
}
