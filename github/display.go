package github

import (
	"fmt"
)

func (widget *Widget) display() {
	client := NewClient()
	prs, _ := client.PullRequests(Config.UString("wtf.mods.github.owner"), Config.UString("wtf.mods.github.repo"))

	widget.View.Clear()

	widget.View.SetTitle(fmt.Sprintf(" Github: %s ", widget.title()))

	str := " [red]Open Review Requests[white]\n"
	str = str + widget.prsForReview(prs, Config.UString("wtf.mods.github.username"))
	str = str + "\n"
	str = str + " [red]Open Pull Requests[white]\n"
	str = str + widget.openPRs(prs, Config.UString("wtf.mods.github.username"))

	fmt.Fprintf(widget.View, str)
}
