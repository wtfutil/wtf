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
		widget.SetItemCount(0)
	} else {
		widget.err = nil
		widget.result = searchResult
		widget.SetItemCount(len(searchResult.Issues))
	}
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

const MaxIssueTypeLength = 7
const MaxStatusNameLength = 14

func (widget *Widget) content() (string, string, bool) {
	if widget.err != nil {
		return widget.CommonSettings().Title, widget.err.Error(), true
	}

	title := fmt.Sprintf("%s- [green]%s[white]", widget.CommonSettings().Title, widget.settings.projects)

	str := " [red]Assigned Issues[white]\n"

	if widget.result == nil || len(widget.result.Issues) == 0 {
		return title, "No results to display", false
	}

	longestIssueTypeLength, longestKeyLength, longestStatusNameLength := getLongestColumnLengths(widget.result.Issues)

	for idx, issue := range widget.result.Issues {
		row := fmt.Sprintf(
			`[%s] [%s]%-*s[white] [green]%-*s[white] [yellow]%-*s[white] [%s]%s`,
			widget.RowColor(idx),
			widget.issueTypeColor(&issue),
			longestIssueTypeLength+1,
			trimToMaxLength(issue.IssueFields.IssueType.Name, MaxIssueTypeLength),
			longestKeyLength+1,
			issue.Key,
			longestStatusNameLength+1,
			trimToMaxLength(issue.IssueFields.IssueStatus.IName, MaxStatusNameLength),
			widget.RowColor(idx),
			issue.IssueFields.Summary,
		)

		str += utils.HighlightableHelper(widget.View, row, idx, len(issue.IssueFields.Summary))
	}

	return title, str, false
}

func getLongestColumnLengths(issues []Issue) (int, int, int) {
	longestIssueTypeLength := 0
	longestKeyLength := 0
	longestStatusNameLength := 0
	for _, issue := range issues {
		issueTypeLength := len(issue.IssueFields.IssueType.Name)
		if issueTypeLength > longestIssueTypeLength {
			longestIssueTypeLength = issueTypeLength
		}

		issueKeyLength := len(issue.Key)
		if issueKeyLength > longestKeyLength {
			longestKeyLength = len("WTF-XXX") // issueKeyLength
		}

		statusNameLength := len(issue.IssueFields.IssueStatus.IName)
		if statusNameLength > longestStatusNameLength {
			longestStatusNameLength = statusNameLength
		}
	}

	if longestIssueTypeLength > MaxIssueTypeLength {
		longestIssueTypeLength = MaxIssueTypeLength
	}

	if longestStatusNameLength > MaxStatusNameLength {
		longestStatusNameLength = MaxStatusNameLength
	}

	return longestIssueTypeLength, longestKeyLength, longestStatusNameLength
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

func trimToMaxLength(text string, maxLength int) string {
	if len(text) <= maxLength {
		return text
	} else {
		return text[:maxLength]
	}
}
