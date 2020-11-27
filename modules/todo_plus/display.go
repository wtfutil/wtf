package todo_plus

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

	if proj.Err != nil {
		return widget.CommonSettings().Title, proj.Err.Error(), true
	}

	title := fmt.Sprintf(
		"[%s]%s[white]",
		widget.settings.Colors.TextTheme.Title,
		proj.Name)

	str := ""

	for idx, item := range proj.Tasks {
		row := fmt.Sprintf(
			`[%s]| | %s[%s]`,
			widget.RowColor(idx),
			tview.Escape(item.Name),
			widget.RowColor(idx),
		)

		str += utils.HighlightableHelper(widget.View, row, idx, len(item.Name))
	}
	return title, str, false
}

func (widget *Widget) display() {
	widget.ScrollableWidget.Redraw(widget.content)
}
