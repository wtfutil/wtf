package jenkins

import (
	"fmt"
	"github.com/senorprogrammer/wtf/wtf"
	"os"
)

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget("Jenkins", "jenkins", false),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	view, err := Create(
		wtf.Config.UString("wtf.mods.jenkins.url"),
		wtf.Config.UString("wtf.mods.jenkins.user"),
		os.Getenv("WTF_JENKINS_API_KEY"),
	)

	widget.UpdateRefreshedAt()
	//widget.View.Clear()

	var content string
	if err != nil {
		widget.View.SetWrap(true)
		widget.View.SetTitle(fmt.Sprintf(" %s ", widget.Name))
		content = err.Error()
	} else {
		widget.View.SetWrap(false)
		widget.View.SetTitle(fmt.Sprintf(" %s: [green] ", widget.Name))
		content = widget.contentFrom(view)
	}

	widget.View.SetText(content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(view *View) string {
	str := fmt.Sprintf(" [red]%s[white]\n", view.Name)

	for _, job := range view.Jobs {
		str = str + fmt.Sprintf(
			" [%s]%-6s[white]\n",
			widget.jobColor(&job),
			job.Name,
		)
	}

	return str
}

func (widget *Widget) jobColor(job *Job) string {
	switch job.Color {
	case "blue":
		return "blue"
	case "red":
		return "red"
	default:
		return "white"
	}
}
