package jira

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.KeyboardWidget
	view.ScrollableWidget

	result   *SearchResult
	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:   view.NewKeyboardWidget(app, pages, settings.common),
		ScrollableWidget: view.NewScrollableWidget(app, settings.common, true),

		settings: settings,
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	searchResult, err := widget.IssuesFor(
		widget.settings.username,
		widget.settings.projects,
		widget.settings.jql,
	)

	if err != nil {
		widget.result = nil
		widget.Redraw(widget.CommonSettings().Title, err.Error(), true)
		return
	}
	widget.result = searchResult
	widget.SetItemCount(len(searchResult.Issues))
	widget.Render()
}

func (widget *Widget) Render() {
	if widget.result == nil {
		return
	}

	str := fmt.Sprintf("%s- [green]%s[white]", widget.CommonSettings().Title, widget.settings.projects)

	widget.Redraw(str, widget.contentFrom(widget.result), false)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) openItem() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.result != nil && sel < len(widget.result.Issues) {
		issue := &widget.result.Issues[sel]
		utils.OpenFile(widget.settings.domain + "/browse/" + issue.Key)
	}
}

func (widget *Widget) contentFrom(searchResult *SearchResult) string {
	str := " [red]Assigned Issues[white]\n"

	for idx, issue := range searchResult.Issues {
		row := fmt.Sprintf(
			`[%s] [%s]%-6s[white] [green]%-10s[white] [yellow][%s][white] [%s]%s`,
			widget.RowColor(idx),
			widget.issueTypeColor(&issue),
			issue.IssueFields.IssueType.Name,
			issue.Key,
			issue.IssueFields.IssueStatus.IName,
			widget.RowColor(idx),
			issue.IssueFields.Summary,
		)

		str += utils.HighlightableHelper(widget.View, row, idx, len(issue.IssueFields.Summary))
	}

	return str
}

func (widget *Widget) issueTypeColor(issue *Issue) string {
	switch issue.IssueFields.IssueType.Name {
	case "Bug":
		return "red"
	case "Story":
		return "blue"
	case "Task":
		return "orange"
	default:
		return "white"
	}
}
