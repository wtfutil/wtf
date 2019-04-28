package gerrit

import (
	"fmt"
)

func (widget *Widget) display() {

	project := widget.currentGerritProject()
	if project == nil {
		widget.View.SetText(fmt.Sprintf("%s", " Gerrit project data is unavailable (1)"))
		return
	}

	widget.View.SetTitle(widget.ContextualTitle(fmt.Sprintf("%s- %s", widget.CommonSettings.Title, widget.title(project))))

	_, _, width, _ := widget.View.GetRect()
	str := widget.settings.common.SigilStr(len(widget.GerritProjects), widget.Idx, width) + "\n"
	str = str + " [red]Stats[white]\n"
	str = str + widget.displayStats(project)
	str = str + "\n"
	str = str + " [red]Open Incoming Reviews[white]\n"
	str = str + widget.displayMyIncomingReviews(project, widget.settings.username)
	str = str + "\n"
	str = str + " [red]My Outgoing Reviews[white]\n"
	str = str + widget.displayMyOutgoingReviews(project, widget.settings.username)

	widget.View.SetText(str)
}

func (widget *Widget) displayMyIncomingReviews(project *GerritProject, username string) string {
	if len(project.IncomingReviews) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for idx, r := range project.IncomingReviews {
		str = str + fmt.Sprintf(" [%s] [green]%d[white] [%s] %s\n", widget.rowColor(idx), r.Number, widget.rowColor(idx), r.Subject)
	}

	return str
}

func (widget *Widget) displayMyOutgoingReviews(project *GerritProject, username string) string {
	if len(project.OutgoingReviews) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for idx, r := range project.OutgoingReviews {
		str = str + fmt.Sprintf(" [%s] [green]%d[white] [%s] %s\n", widget.rowColor(idx+len(project.IncomingReviews)), r.Number, widget.rowColor(idx+len(project.IncomingReviews)), r.Subject)
	}

	return str
}

func (widget *Widget) displayStats(project *GerritProject) string {
	str := fmt.Sprintf(
		" Reviews: %d\n",
		project.ReviewCount,
	)

	return str
}

func (widget *Widget) rowColor(idx int) string {
	if widget.View.HasFocus() && (idx == widget.selected) {
		return widget.settings.common.DefaultFocussedRowColor()
	}

	return widget.settings.common.RowColor(idx)
}

func (widget *Widget) title(project *GerritProject) string {
	return fmt.Sprintf("[green]%s [white]", project.Path)
}
