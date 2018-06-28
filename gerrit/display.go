package gerrit

import (
	"fmt"

	"github.com/senorprogrammer/wtf/wtf"
)

func (widget *Widget) display() {

	project := widget.currentGerritProject()
	if project == nil {
		fmt.Fprintf(widget.View, "%s", " Gerrit project data is unavailable (1)")
		return
	}

	widget.View.SetTitle(fmt.Sprintf("%s- %s", widget.Name, widget.title(project)))

	str := wtf.SigilStr(len(widget.GerritProjects), widget.Idx, widget.View) + "\n"
	str = str + " [red]Stats[white]\n"
	str = str + widget.displayStats(project)
	str = str + "\n"
	str = str + " [red]Open Incoming Reviews[white]\n"
	str = str + widget.displayMyIncomingReviews(project, wtf.Config.UString("wtf.mods.gerrit.username"))
	str = str + "\n"
	str = str + " [red]My Outgoing Reviews[white]\n"
	str = str + widget.displayMyOutgoingReviews(project, wtf.Config.UString("wtf.mods.gerrit.username"))

	widget.View.SetText(str)
}

func (widget *Widget) displayMyOutgoingReviews(project *GerritProject, username string) string {
	ors := project.myOutgoingReviews(username)

	if len(ors) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for _, r := range ors {
		str = str + fmt.Sprintf(" [green]%4s[white] %s\n", r.ChangeID, r.Subject)
	}

	return str
}

func (widget *Widget) displayMyIncomingReviews(project *GerritProject, username string) string {
	irs := project.myIncomingReviews(username)

	if len(irs) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for _, r := range irs {
		str = str + fmt.Sprintf(" [green]%4s[white] %s\n", r.ChangeID, r.Subject)
	}

	return str
}

func (widget *Widget) displayStats(project *GerritProject) string {
	str := fmt.Sprintf(
		" Reviews: %d\n",
		project.ReviewCount(),
	)

	return str
}

func (widget *Widget) title(project *GerritProject) string {
	return fmt.Sprintf("[green]%s [white]", project.Path)
}
