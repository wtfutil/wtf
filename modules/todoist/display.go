package todoist

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
)

func (widget *Widget) content() (string, string, bool) {
	proj := widget.CurrentProject()

	if proj == nil {
		return widget.CommonSettings().Title, "", false
	}

	if proj.err != nil {
		return widget.CommonSettings().Title, proj.err.Error(), true
	}

	title := fmt.Sprintf("[green]%s[white]", proj.Project.Name)

	str := ""

	for idx, item := range proj.tasks {
		row := fmt.Sprintf(
			`[%s]| | %s[%s]`,
			widget.RowColor(idx),
			tview.Escape(item.Content),
			widget.RowColor(idx),
		)

		str += utils.HighlightableHelper(widget.View, row, idx, len(item.Content))
	}
	return title, str, false
}

func (widget *Widget) display() {
	widget.ScrollableWidget.Redraw(widget.content)
}
