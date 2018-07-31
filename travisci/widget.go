package travisci

import (
	"fmt"
	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget("TravisCI", "travisci", false),
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

func (widget *Widget) contentFrom(builds *Builds) string {
	var str string
	for _, build := range builds.Builds {
		str = str + fmt.Sprintf(
			"[%s] %s-%s (%s) [white]%s\n",
			buildColor(&build),
			build.Repository.Name,
			build.Number,
			build.Branch.Name,
			build.CreatedBy.Login,
		)
	}

	return str
}

func buildColor(build *Build) string {
	switch build.State {
	case "broken":
		return "red"
	case "failed":
		return "red"
	case "failing":
		return "red"
	case "pending":
		return "yellow"
	case "started":
		return "yellow"
	case "fixed":
		return "green"
	case "passed":
		return "green"
	default:
		return "white"
	}
}
