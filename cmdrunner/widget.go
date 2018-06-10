package cmdrunner

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget

	args   []string
	cmd    string
	result string
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" CmdRunner ", "cmdrunner", false),

		args: wtf.ToStrs(Config.UList("wtf.mods.cmdrunner.args")),
		cmd:  Config.UString("wtf.mods.cmdrunner.cmd"),
	}

	widget.View.SetWrap(true)

	return &widget
}

func (widget *Widget) Refresh() {
	widget.UpdateRefreshedAt()
	widget.execute()
	widget.View.SetTitle(fmt.Sprintf(" %s ", widget))

	widget.View.SetText(fmt.Sprintf("%s", widget.result))
}

func (widget *Widget) String() string {
	args := strings.Join(widget.args, " ")
	return fmt.Sprintf("%s %s", widget.cmd, args)
}

func (widget *Widget) execute() {
	cmd := exec.Command(widget.cmd, widget.args...)
	widget.result = wtf.ExecuteCommand(cmd)
}
