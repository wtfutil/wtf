package jira

import (
	"fmt"

	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Jira ", "jira", false),
	}

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

	var content string
	if err != nil {
		widget.View.SetWrap(true)
		widget.View.SetTitle(widget.Name)
		content = err.Error()
	} else {
		widget.View.SetWrap(false)
		widget.View.SetTitle(
			fmt.Sprintf(
				"%s- [green]%s[white]",
				widget.Name,
				wtf.Config.UString("wtf.mods.jira.project"),
			),
		)
		content = widget.contentFrom(searchResult)
	}

	widget.View.SetText(content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(searchResult *SearchResult) string {
	str := " [red]Assigned Issues[white]\n"

	for idx, issue := range searchResult.Issues {
		str = str + fmt.Sprintf(
			" [%s]%-6s[white] [green]%-10s[%s] %s\n",
			widget.issueTypeColor(&issue),
			issue.IssueFields.IssueType.Name,
			issue.Key,
			wtf.RowColor("jira", idx),
			issue.IssueFields.Summary,
		)
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
