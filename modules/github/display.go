package github

import (
	"fmt"

	"github.com/google/go-github/v25/github"
)

func (widget *Widget) display() {
	repo := widget.currentGithubRepo()
	if repo == nil {
		widget.View.SetText(" GitHub repo data is unavailable ")
		return
	}

	widget.View.SetTitle(widget.ContextualTitle(fmt.Sprintf("%s - %s", widget.Name(), widget.title(repo))))

	_, _, width, _ := widget.View.GetRect()
	str := widget.settings.common.SigilStr(len(widget.GithubRepos), widget.Idx, width) + "\n"
	str = str + " [red]Stats[white]\n"
	str = str + widget.displayStats(repo)
	str = str + "\n"
	str = str + " [red]Open Review Requests[white]\n"
	str = str + widget.displayMyReviewRequests(repo, widget.settings.username)
	str = str + "\n"
	str = str + " [red]My Pull Requests[white]\n"
	str = str + widget.displayMyPullRequests(repo, widget.settings.username)

	widget.View.SetText(str)
}

func (widget *Widget) displayMyPullRequests(repo *GithubRepo, username string) string {
	prs := repo.myPullRequests(username, widget.settings.enableStatus)

	if len(prs) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for _, pr := range prs {
		str = str + fmt.Sprintf(" %s[green]%4d[white] %s\n", widget.mergeString(pr), *pr.Number, *pr.Title)
	}

	return str
}

func (widget *Widget) displayMyReviewRequests(repo *GithubRepo, username string) string {
	prs := repo.myReviewRequests(username)

	if len(prs) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for _, pr := range prs {
		str = str + fmt.Sprintf(" [green]%4d[white] %s\n", *pr.Number, *pr.Title)
	}

	return str
}

func (widget *Widget) displayStats(repo *GithubRepo) string {
	str := fmt.Sprintf(
		" PRs: %d  Issues: %d  Stars: %d\n",
		repo.PullRequestCount(),
		repo.IssueCount(),
		repo.StarCount(),
	)

	return str
}

func (widget *Widget) title(repo *GithubRepo) string {
	return fmt.Sprintf("[green]%s - %s[white]", repo.Owner, repo.Name)
}

var mergeIcons = map[string]string{
	"dirty":    "[red]![white] ",
	"clean":    "[green]✔[white] ",
	"unstable": "[red]✖[white] ",
	"blocked":  "[red]✖[white] ",
}

func (widget *Widget) mergeString(pr *github.PullRequest) string {
	if !widget.settings.enableStatus {
		return ""
	}
	if str, ok := mergeIcons[pr.GetMergeableState()]; ok {
		return str
	}
	return "? "
}
