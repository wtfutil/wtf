package cmdrunner

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	args   []string
	cmd    string
	result string
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" CmdRunner ", "cmdrunner", false),

		args: wtf.ToStrs(wtf.Config.UList("wtf.mods.cmdrunner.args")),
		cmd:  wtf.Config.UString("wtf.mods.cmdrunner.cmd"),
	}

	widget.View.SetWrap(true)

	return &widget
}

func (widget *Widget) Refresh() {
	widget.UpdateRefreshedAt()
	widget.execute()

	title := wtf.Config.UString("wtf.mods.cmdrunner.title", widget.String())
	widget.View.SetTitle(title)

	widget.View.SetText(widget.result)
}

func (widget *Widget) String() string {
	args := strings.Join(widget.args, " ")
	return fmt.Sprintf(" %s %s ", widget.cmd, args)
}

func (widget *Widget) execute() {
	cmd := exec.Command(widget.cmd, widget.args...)
	widget.result = wtf.ExecuteCommand(cmd)
}
