package gitlab

import (
	"fmt"
)

func (widget *Widget) display() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {

	project := widget.currentGitlabProject()
	if project == nil {
		return widget.CommonSettings().Title, " Gitlab project data is unavailable ", true
	}

	title := fmt.Sprintf("%s- %s", widget.CommonSettings().Title, widget.title(project))

	_, _, width, _ := widget.View.GetRect()
	str := widget.settings.common.SigilStr(len(widget.GitlabProjects), widget.Idx, width) + "\n"
	str += " [red]Stats[white]\n"
	str += widget.displayStats(project)
	str += "\n"
	str += " [red]Open Assigned Merge Requests[white]\n"
	str += widget.displayMyAssignedMergeRequests(project, widget.settings.username)
	str += "\n"
	str += " [red]My Merge Requests[white]\n"
	str += widget.displayMyMergeRequests(project, widget.settings.username)
	str += "\n"
	str += " [red]Open Assigned Issues[white]\n"
	str += widget.displayMyAssignedIssues(project, widget.settings.username)
	str += "\n"
	str += " [red]My Issues[white]\n"
	str += widget.displayMyIssues(project, widget.settings.username)

	return title, str, false
}

func (widget *Widget) displayMyMergeRequests(project *GitlabProject, username string) string {
	mrs := project.myMergeRequests(username)

	if len(mrs) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for _, mr := range mrs {
		str += fmt.Sprintf(" [green]%4d[white] %s\n", mr.IID, mr.Title)
	}

	return str
}

func (widget *Widget) displayMyAssignedMergeRequests(project *GitlabProject, username string) string {
	mrs := project.myAssignedMergeRequests(username)

	if len(mrs) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for _, mr := range mrs {
		str += fmt.Sprintf(" [green]%4d[white] %s\n", mr.IID, mr.Title)
	}

	return str
}

func (widget *Widget) displayMyAssignedIssues(project *GitlabProject, username string) string {
	issues := project.myAssignedIssues(username)

	if len(issues) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for _, issue := range issues {
		str += fmt.Sprintf(" [green]%4d[white] %s\n", issue.IID, issue.Title)
	}

	return str
}

func (widget *Widget) displayMyIssues(project *GitlabProject, username string) string {
	issues := project.myIssues(username)

	if len(issues) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for _, issue := range issues {
		str += fmt.Sprintf(" [green]%4d[white] %s\n", issue.IID, issue.Title)
	}

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
