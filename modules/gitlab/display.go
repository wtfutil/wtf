package gitlab

import (
	"fmt"

	"github.com/xanzy/go-gitlab"
)

func (widget *Widget) display() {
	widget.Redraw(widget.content)
}

func (widget *Widget) displayError() {
	widget.Redraw(widget.contentError)
}

func (widget *Widget) contentError() (string, string, bool) {

	title := fmt.Sprintf("%s - Error", widget.CommonSettings().Title)

	if widget.configError != nil {
		return title, fmt.Sprintf("Error: \n [red]%v[white]", widget.configError), false

	}
	return title, "Error", false
}

func (widget *Widget) content() (string, string, bool) {

	project := widget.currentGitlabProject()
	if project == nil {
		return widget.CommonSettings().Title, " Gitlab project data is unavailable ", true
	}

	// initial maxItems count
	widget.Items = make([]ContentItem, 0)
	widget.SetItemCount(0)

	title := fmt.Sprintf("%s - %s", widget.CommonSettings().Title, widget.title(project))

	_, _, width, _ := widget.View.GetRect()
	str := widget.settings.PaginationMarker(len(widget.GitlabProjects), widget.Idx, width) + "\n"
	str += fmt.Sprintf(" [%s]Stats[white]\n", widget.settings.Colors.Subheading)
	str += widget.displayStats(project)
	str += "\n"
	str += fmt.Sprintf(" [%s]Open Assigned Merge Requests[white]\n", widget.settings.Colors.Subheading)
	str += widget.displayMyAssignedMergeRequests(project, widget.settings.username)
	str += "\n"
	str += fmt.Sprintf(" [%s]My Merge Requests[white]\n", widget.settings.Colors.Subheading)
	str += widget.displayMyMergeRequests(project, widget.settings.username)
	str += "\n"
	str += fmt.Sprintf(" [%s]Open Assigned Issues[white]\n", widget.settings.Colors.Subheading)
	str += widget.displayMyAssignedIssues(project, widget.settings.username)
	str += "\n"
	str += fmt.Sprintf(" [%s]My Issues[white]\n", widget.settings.Colors.Subheading)
	str += widget.displayMyIssues(project, widget.settings.username)

	return title, str, false
}

func (widget *Widget) displayMyMergeRequests(project *GitlabProject, username string) string {
	mrs := project.myMergeRequests()
	return widget.renderMergeRequests(mrs)
}

func (widget *Widget) displayMyAssignedMergeRequests(project *GitlabProject, username string) string {
	mrs := project.myAssignedMergeRequests()
	return widget.renderMergeRequests(mrs)
}

func (widget *Widget) displayMyAssignedIssues(project *GitlabProject, username string) string {
	issues := project.myAssignedIssues()
	return widget.renderIssues(issues)
}

func (widget *Widget) displayMyIssues(project *GitlabProject, username string) string {
	issues := project.myIssues()
	return widget.renderIssues(issues)
}

func (widget *Widget) renderMergeRequests(mrs []*gitlab.MergeRequest) string {

	length := len(mrs)

	if length == 0 {
		return " [grey]none[white]\n"
	}
	maxItems := widget.GetItemCount()

	str := ""
	for idx, issue := range mrs {
		str += fmt.Sprintf(` [green]["%d"]%4d[""][white] %s`, maxItems+idx, issue.IID, issue.Title)
		str += "\n"
		widget.Items = append(widget.Items, ContentItem{Type: "MR", ID: issue.IID})
	}
	widget.SetItemCount(maxItems + length)

	return str
}

func (widget *Widget) renderIssues(issues []*gitlab.Issue) string {

	length := len(issues)

	if length == 0 {
		return " [grey]none[white]\n"
	}
	maxItems := widget.GetItemCount()

	str := ""
	for idx, issue := range issues {
		str += fmt.Sprintf(` [green]["%d"]%4d[""][white] %s`, maxItems+idx, issue.IID, issue.Title)
		str += "\n"
		widget.Items = append(widget.Items, ContentItem{Type: "ISSUE", ID: issue.IID})
	}
	widget.SetItemCount(maxItems + length)

	return str
}

func (widget *Widget) displayStats(project *GitlabProject) string {
	str := fmt.Sprintf(
		" MRs: %d  Issues: %d  Stars: %d\n",
		project.MergeRequestCount(),
		project.IssueCount(),
		project.StarCount(),
	)

	return str
}

func (widget *Widget) title(project *GitlabProject) string {
	return fmt.Sprintf("[green]%s [white]", project.path)
}
