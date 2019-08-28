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
	err      error
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
		widget.err = err
		widget.result = nil
		return
	}
	widget.err = nil
	widget.result = searchResult
	widget.SetItemCount(len(searchResult.Issues))
	widget.Render()
}

func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) openItem() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.result != nil && sel < len(widget.result.Issues) {
		issue := &widget.result.Issues[sel]
		utils.OpenFile(widget.settings.domain + "/browse/" + issue.Key)
	}
}

func (widget *Widget) content() (string, string, bool) {
	if widget.err != nil {
		return widget.CommonSettings().Title, widget.err.Error(), true
	}

	title := fmt.Sprintf("%s- [green]%s[white]", widget.CommonSettings().Title, widget.settings.projects)
	str := " [red]Assigned Issues[white]\n"

	if widget.result == nil || len(widget.result.Issues) == 0 {
		return title, "No results to display", false
	}

	for idx, issue := range widget.result.Issues {
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

	return title, str, false
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
