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
		widget.settings.Colors.Title,
		proj.Name)

	str := ""

	for idx, item := range proj.Tasks {
		displayName := tview.Escape(item.Name)

		if item.Prefix != "" {
			displayName = fmt.Sprintf("[yellow]%s[-] %s",
				tview.Escape(item.Prefix), displayName)
		}

		if item.DateSuffix != "" {
			if item.Overdue {
				displayName += fmt.Sprintf("[red]%s[-]", item.DateSuffix)
			} else {
				displayName += item.DateSuffix
			}
		}

		row := fmt.Sprintf(
			`[%s]| | %s[%s]`,
			widget.RowColor(idx),
			displayName,
			widget.RowColor(idx),
		)

		str += utils.HighlightableHelper(widget.View, row, idx, len(item.Name))
	}
	return title, str, false
}

func (widget *Widget) display() {
	widget.Redraw(widget.content)
}
