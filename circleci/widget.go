package circleci

import (
	"fmt"
	"github.com/senorprogrammer/wtf/wtf"
)

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

func buildColor(build *Build) string {
	switch build.Status {
	case "failed":
		return "red"
	case "running":
		return "yellow"
	case "success":
		return "green"
	default:
		return "white"
	}
}
