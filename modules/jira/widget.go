package jira

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.KeyboardWidget
	wtf.ScrollableWidget

	result   *SearchResult
	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:   wtf.NewKeyboardWidget(app, pages, settings.common),
		ScrollableWidget: wtf.NewScrollableWidget(app, settings.common, true),

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
		widget.Redraw(widget.CommonSettings.Title, err.Error(), true)
		return
	}
	widget.result = searchResult
	widget.SetItemCount(len(searchResult.Issues))
	widget.Render()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) Render() {
	if widget.result == nil {
		return
	}

	str := fmt.Sprintf("%s- [green]%s[white]", widget.CommonSettings.Title, widget.settings.projects)

	widget.Redraw(str, widget.contentFrom(widget.result), false)
}

func (widget *Widget) openItem() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.result != nil && sel < len(widget.result.Issues) {
		issue := &widget.result.Issues[sel]
		wtf.OpenFile(widget.settings.domain + "/browse/" + issue.Key)
	}
}

func (widget *Widget) contentFrom(searchResult *SearchResult) string {
	str := " [red]Assigned Issues[white]\n"

	for idx, issue := range searchResult.Issues {
		fmtStr := fmt.Sprintf(
			`["%d"][%s] [%s]%-6s[white] [green]%-10s[white] [yellow][%s][white] [%s]%s[""]`,
			idx,
			widget.RowColor(idx),
			widget.issueTypeColor(&issue),
			issue.IssueFields.IssueType.Name,
			issue.Key,
			issue.IssueFields.IssueStatus.IName,
			widget.RowColor(idx),
			issue.IssueFields.Summary,
		)

		_, _, w, _ := widget.View.GetInnerRect()
		fmtStr += wtf.PadRow(len(issue.IssueFields.Summary), w+1)

		str = str + fmtStr + "\n"
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
