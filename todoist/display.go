package todoist

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

const checkWidth = 4

func (widget *Widget) display() {
	proj := widget.CurrentProject()

	if proj == nil {
		return
	}

	title := fmt.Sprintf("[green]%s[white]", proj.Project.Name)
	widget.View.SetTitle(widget.ContextualTitle(title))

	str := wtf.SigilStr(len(widget.projects), widget.idx, widget.View) + "\n"

	maxLen := proj.LongestLine()

	for index, item := range proj.tasks {
		foreColor, backColor := "white", wtf.Config.UString("wtf.colors.background", "black")

		if index == proj.index {
			foreColor = wtf.Config.UString("wtf.colors.highlight.fore", "black")
			backColor = wtf.Config.UString("wtf.colors.highlight.back", "orange")
		}

		row := fmt.Sprintf(
			"[%s:%s]| | %s[white]",
			foreColor,
			backColor,
			tview.Escape(item.Content),
		)

		_, _, w, _ := widget.View.GetInnerRect()
		if w > maxLen {
			maxLen = w
		}

		str = str + row + wtf.PadRow((checkWidth+len(item.Content)), (checkWidth+maxLen+1)) + "\n"
	}

	//widget.View.Clear()
	widget.View.SetText(str)
}
