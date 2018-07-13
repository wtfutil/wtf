package todoist

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

const checkWidth = 4

func (w *Widget) display() {
	proj := w.CurrentProject()

	if proj == nil {
		return
	}

	w.View.SetTitle(fmt.Sprintf("%s- [green]%s[white] ", w.Name, proj.Project.Name))
	str := wtf.SigilStr(len(w.projects), w.idx, w.View) + "\n"

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

		str = str + row + wtf.PadRow((checkWidth+len(item.Content)), (checkWidth+maxLen+1)) + "\n"
	}

	w.View.Clear()
	w.View.SetText(str)
}
