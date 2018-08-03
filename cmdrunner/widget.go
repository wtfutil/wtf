package cmdrunner

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/rivo/tview"
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

	title := tview.TranslateANSI(wtf.Config.UString("wtf.mods.cmdrunner.title", widget.String()))
	widget.View.SetTitle(title)

	widget.View.SetText(widget.result)
}

func (widget *Widget) String() string {
	args := strings.Join(widget.args, " ")

	if args != "" {
		return fmt.Sprintf(" %s %s ", widget.cmd, args)
	}

	return fmt.Sprintf(" %s ", widget.cmd)
}

func (widget *Widget) execute() {
	cmd := exec.Command(widget.cmd, widget.args...)
	widget.result = tview.TranslateANSI(wtf.ExecuteCommand(cmd))
}
