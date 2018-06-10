package jira

import (
	"fmt"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

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
	searchResult, err := IssuesFor(Config.UString("wtf.mods.jira.username"), getProjects(), Config.UString("wtf.mods.jira.jql", ""))

	widget.UpdateRefreshedAt()

	if err != nil {
		widget.View.SetWrap(true)
		widget.View.SetTitle(fmt.Sprintf("%s", widget.Name))
		fmt.Fprintf(widget.View, "%v", err)
	} else {
		widget.View.SetWrap(false)
		widget.View.SetTitle(
			fmt.Sprintf(
				"%s- [green]%s[white]",
				widget.Name,
				Config.UString("wtf.mods.jira.project"),
			),
		)
		widget.View.SetText(fmt.Sprintf("%s", widget.contentFrom(searchResult)))
	}
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

func getProjects() []string {
	// see if project is set to a single string
	configPath := "wtf.mods.jira.project"
	singleProject, err := Config.String(configPath)
	if err == nil {
		return []string{singleProject}
	}
	// else, assume list
	projList := Config.UList(configPath)
	var ret []string
	for _, proj := range projList {
		if str, ok := proj.(string); ok {
			ret = append(ret, str)
		}
	}
	return ret
}
