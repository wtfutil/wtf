package github

import (
	"fmt"

	"github.com/senorprogrammer/wtf/wtf"
)

func (widget *Widget) display() {
	widget.View.Clear()

	repo := widget.currentGithubRepo()
	if repo == nil {
		fmt.Fprintf(widget.View, "%s", " Github repo data is unavailable (1)")
		return
	}

	widget.View.SetTitle(fmt.Sprintf(" Github: %s ", widget.title(repo)))

	str := wtf.SigilStr(len(widget.GithubRepos), widget.Idx, widget.View) + "\n"
	str = str + " [red]Open Review Requests[white]\n"
	str = str + repo.pullRequetsForMeToReview(Config.UString("wtf.mods.github.username"))
	str = str + "\n"
	str = str + " [red]My Pull Requests[white]\n"
	str = str + repo.myPullRequests(Config.UString("wtf.mods.github.username"))

	fmt.Fprintf(widget.View, str)
}

func (widget *Widget) title(repo *GithubRepo) string {
	return fmt.Sprintf("[green]%s - %s[white]", repo.Owner, repo.Name)
}
