package jira

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

const HelpText = `
 Keyboard commands for Jira:

   /: Show/hide this help window
   j: Select the next item in the list
   k: Select the previous item in the list

   arrow down: Select the next item in the list
   arrow up:   Select the previous item in the list

   return: Open the selected issue in a browser
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.KeyboardWidget
	wtf.TextWidget

	app *tview.Application

	result   *SearchResult
	selected int
	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		HelpfulWidget:  wtf.NewHelpfulWidget(app, pages, HelpText),
		KeyboardWidget: wtf.NewKeyboardWidget(),
		TextWidget:     wtf.NewTextWidget(app, settings.common, true),

		app:      app,
		settings: settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.unselect()

	widget.View.SetScrollable(true)
	widget.View.SetRegions(true)

	widget.HelpfulWidget.SetView(widget.View)

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
	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	if widget.result == nil {
		return
	}

	str := fmt.Sprintf("%s- [green]%s[white]", widget.CommonSettings.Title, widget.settings.projects)

	widget.Redraw(str, widget.contentFrom(widget.result), false)
	widget.app.QueueUpdateDraw(func() {
		widget.View.Highlight(strconv.Itoa(widget.selected)).ScrollToHighlight()
	})
}

func (widget *Widget) next() {
	widget.selected++
	if widget.result != nil && widget.selected >= len(widget.result.Issues) {
		widget.selected = 0
	}
}

func (widget *Widget) prev() {
	widget.selected--
	if widget.selected < 0 && widget.result != nil {
		widget.selected = len(widget.result.Issues) - 1
	}
}

func (widget *Widget) openItem() {
	sel := widget.selected
	if sel >= 0 && widget.result != nil && sel < len(widget.result.Issues) {
		issue := &widget.result.Issues[widget.selected]
		wtf.OpenFile(widget.settings.domain + "/browse/" + issue.Key)
	}
}

func (widget *Widget) unselect() {
	widget.selected = -1
}

func (widget *Widget) contentFrom(searchResult *SearchResult) string {
	str := " [red]Assigned Issues[white]\n"

	for idx, issue := range searchResult.Issues {
		fmtStr := fmt.Sprintf(
			`["%d"][""][%s] [%s]%-6s[white] [green]%-10s[white] [yellow][%s][white] [%s]%s`,
			idx,
			widget.rowColor(idx),
			widget.issueTypeColor(&issue),
			issue.IssueFields.IssueType.Name,
			issue.Key,
			issue.IssueFields.IssueStatus.IName,
			widget.rowColor(idx),
			issue.IssueFields.Summary,
		)

		_, _, w, _ := widget.View.GetInnerRect()
		fmtStr = fmtStr + wtf.PadRow(len(issue.IssueFields.Summary), w+1)

		str = str + fmtStr + "\n"
	}

	return str
}

func (widget *Widget) rowColor(idx int) string {
	if widget.View.HasFocus() && (idx == widget.selected) {
		widget.settings.common.DefaultFocussedRowColor()
	}

	return widget.settings.common.RowColor(idx)
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
