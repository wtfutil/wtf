package gerrit

import (
	"fmt"
)

func (widget *Widget) display() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	title := widget.CommonSettings().Title
	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	project := widget.currentGerritProject()
	if project == nil {
		return title, "Gerrit project data is unavailable", true
	}

	title = fmt.Sprintf("%s- %s", widget.CommonSettings().Title, widget.title(project))

	_, _, width, _ := widget.View.GetRect()
	str := widget.settings.common.SigilStr(len(widget.GerritProjects), widget.Idx, width) + "\n"
	str += " [red]Stats[white]\n"
	str += widget.displayStats(project)
	str += "\n"
	str += " [red]Open Incoming Reviews[white]\n"
	str += widget.displayMyIncomingReviews(project, widget.settings.username)
	str += "\n"
	str += " [red]My Outgoing Reviews[white]\n"
	str += widget.displayMyOutgoingReviews(project, widget.settings.username)

	return title, str, false
}

func (widget *Widget) displayMyIncomingReviews(project *GerritProject, username string) string {
	if len(project.IncomingReviews) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for idx, r := range project.IncomingReviews {
		str += fmt.Sprintf(" [%s] [green]%d[white] [%s] %s\n", widget.rowColor(idx), r.Number, widget.rowColor(idx), r.Subject)
	}

	return str
}

func (widget *Widget) displayMyOutgoingReviews(project *GerritProject, username string) string {
	if len(project.OutgoingReviews) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for idx, r := range project.OutgoingReviews {
		str += fmt.Sprintf(" [%s] [green]%d[white] [%s] %s\n", widget.rowColor(idx+len(project.IncomingReviews)), r.Number, widget.rowColor(idx+len(project.IncomingReviews)), r.Subject)
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
		return widget.settings.common.DefaultFocusedRowColor()
	}

	return widget.settings.common.RowColor(idx)
}

func (widget *Widget) title(project *GerritProject) string {
	return fmt.Sprintf("[green]%s [white]", project.Path)
}
