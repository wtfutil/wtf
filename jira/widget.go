package jira

import (
	"fmt"
	"time"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

var Config *config.Config

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget("JIRA", "jira"),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	searchResult, err := IssuesFor(Config.UString("wtf.mods.jira.username"))

	widget.View.Clear()

	if err != nil {
		widget.View.SetWrap(true)
		widget.View.SetTitle(fmt.Sprintf(" %s ", widget.Name))
		fmt.Fprintf(widget.View, "%v", err)
	} else {
		widget.View.SetWrap(false)
		widget.View.SetTitle(
			fmt.Sprintf(
				" %s: [green]%s[white] ",
				widget.Name,
				Config.UString("wtf.mods.jira.project"),
			),
		)
		fmt.Fprintf(widget.View, "%s", widget.contentFrom(searchResult))
	}

	widget.RefreshedAt = time.Now()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(searchResult *SearchResult) string {
	str := " [red]Assigned Issues[white]\n"

	for _, issue := range searchResult.Issues {
		str = str + fmt.Sprintf(
			" [%s]%-6s[white] [green]%-10s[white] %s\n",
			widget.issueTypeColor(&issue),
			issue.IssueFields.IssueType.Name,
			issue.Key,
			issue.IssueFields.Summary,
		)
	}

	return str
}

func (widget *Widget) issueTypeColor(issue *Issue) string {
	var color string

	switch issue.IssueFields.IssueType.Name {
	case "Bug":
		color = "red"
	case "Story":
		color = "blue"
	case "Task":
		color = "orange"
	default:
		color = "white"
	}

	return color
}
