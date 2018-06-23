package gitlab

import (
	"fmt"

	"github.com/senorprogrammer/wtf/wtf"
)

func (widget *Widget) display() {

	project := widget.currentGitlabProject()
	if project == nil {
		widget.View.SetText(" Gitlab project data is unavailable ")
		return
	}

	widget.View.SetTitle(fmt.Sprintf("%s- %s", widget.Name, widget.title(project)))

	str := wtf.SigilStr(len(widget.GitlabProjects), widget.Idx, widget.View) + "\n"
	str = str + " [red]Stats[white]\n"
	str = str + widget.displayStats(project)
	str = str + "\n"
	str = str + " [red]Open Approval Requests[white]\n"
	str = str + widget.displayMyApprovalRequests(project, wtf.Config.UString("wtf.mods.gitlab.username"))
	str = str + "\n"
	str = str + " [red]My Merge Requests[white]\n"
	str = str + widget.displayMyMergeRequests(project, wtf.Config.UString("wtf.mods.gitlab.username"))

	widget.View.SetText(str)
}

func (widget *Widget) displayMyMergeRequests(project *GitlabProject, username string) string {
	mrs := project.myMergeRequests(username)

	if len(mrs) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for _, mr := range mrs {
		str = str + fmt.Sprintf(" [green]%4d[white] %s\n", mr.IID, mr.Title)
	}

	return str
}

func (widget *Widget) displayMyApprovalRequests(project *GitlabProject, username string) string {
	mrs := project.myApprovalRequests(username)

	if len(mrs) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for _, mr := range mrs {
		str = str + fmt.Sprintf(" [green]%4d[white] %s\n", mr.IID, mr.Title)
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
	return fmt.Sprintf("[green]%s [white]", project.Path)
}
