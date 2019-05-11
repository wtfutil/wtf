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
	str := widget.settings.common.SigilStr(len(widget.projects), widget.idx, width) + "\n"

	maxLen := proj.LongestLine()

	for index, item := range proj.tasks {
		foreColor, backColor := widget.settings.common.Colors.Text, widget.settings.common.Colors.Background

		if index == proj.index {
			foreColor = widget.settings.common.Colors.HighlightFore
			backColor = widget.settings.common.Colors.HighlightBack
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

	widget.TextWidget.Redraw(title, str, false)
}
