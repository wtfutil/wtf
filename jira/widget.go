package jira

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
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
	wtf.TextWidget

	result   *SearchResult
	selected int
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget("Jira", "jira", true),
	}

	widget.HelpfulWidget.SetView(widget.View)
	widget.unselect()

	widget.View.SetInputCapture(widget.keyboardIntercept)
	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	searchResult, err := IssuesFor(
		wtf.Config.UString("wtf.mods.jira.username"),
		getProjects(),
		wtf.Config.UString("wtf.mods.jira.jql", ""),
	)

	widget.UpdateRefreshedAt()

	if err != nil {
		widget.result = nil
		widget.View.SetWrap(true)
		widget.View.SetTitle(widget.Name)
		widget.View.SetText(err.Error())
	} else {
		widget.result = searchResult
	}

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	if widget.result == nil {
		return
	}
	widget.View.SetWrap(false)

	str := fmt.Sprintf("%s- [green]%s[white]", widget.Name, wtf.Config.UString("wtf.mods.jira.project"))

	widget.View.SetTitle(widget.ContextualTitle(str))
	widget.View.SetText(fmt.Sprintf("%s", widget.contentFrom(widget.result)))
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
		wtf.OpenFile(wtf.Config.UString("wtf.mods.jira.domain") + "/browse/" + issue.Key)
	}
}

func (widget *Widget) unselect() {
	widget.selected = -1
}

func (widget *Widget) contentFrom(searchResult *SearchResult) string {
	str := " [red]Assigned Issues[white]\n"

	for idx, issue := range searchResult.Issues {
		fmtStr := fmt.Sprintf(
			"[%s] [%s]%-6s[white] [green]%-10s[white] [%s]%s",
			widget.rowColor(idx),
			widget.issueTypeColor(&issue),
			issue.IssueFields.IssueType.Name,
			issue.Key,
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
		return wtf.DefaultFocussedRowColor()
	}
	return wtf.RowColor("jira", idx)
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

func getProjects() []string {
	// see if project is set to a single string
	configPath := "wtf.mods.jira.project"
	singleProject, err := wtf.Config.String(configPath)
	if err == nil {
		return []string{singleProject}
	}
	// else, assume list
	projList := wtf.Config.UList(configPath)
	var ret []string
	for _, proj := range projList {
		if str, ok := proj.(string); ok {
			ret = append(ret, str)
		}
	}
	return ret
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "/":
		widget.ShowHelp()
	case "j":
		// Select the next item down
		widget.next()
		widget.display()
		return nil
	case "k":
		// Select the next item up
		widget.prev()
		widget.display()
		return nil
	}

	switch event.Key() {
	case tcell.KeyDown:
		// Select the next item down
		widget.next()
		widget.display()
		return nil
	case tcell.KeyEnter:
		widget.openItem()
		return nil
	case tcell.KeyEsc:
		// Unselect the current row
		widget.unselect()
		widget.display()
		return event
	case tcell.KeyUp:
		// Select the next item up
		widget.prev()
		widget.display()
		return nil
	default:
		// Pass it along
		return event
	}
}
