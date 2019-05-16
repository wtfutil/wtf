package todoist

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

const checkWidth = 4

func (widget *Widget) display() {
	proj := widget.CurrentProject()

	if proj == nil {
		return
	}

	title := fmt.Sprintf("[green]%s[white]", proj.Project.Name)

	_, _, width, _ := widget.View.GetRect()
	str := widget.settings.common.SigilStr(len(widget.projects), widget.Idx, width) + "\n"

	maxLen := proj.LongestLine()

	for index, item := range proj.tasks {
		row := fmt.Sprintf(
			`["%d"][""][%s]| | %s[%s]`,
			index,
			widget.RowColor(index),
			tview.Escape(item.Content),
			widget.RowColor(index),
		)

		_, _, w, _ := widget.View.GetInnerRect()
		if w > maxLen {
			maxLen = w
		}

		str = str + row + wtf.PadRow((checkWidth+len(item.Content)), (checkWidth+maxLen+1)) + `[""]` + "\n"
	}

	widget.ScrollableWidget.Redraw(title, str, false)
}
