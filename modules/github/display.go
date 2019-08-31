package github

import (
	"fmt"

	"github.com/google/go-github/v26/github"
)

func (widget *Widget) display() {
	widget.TextWidget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	repo := widget.currentGithubRepo()
	title := fmt.Sprintf("%s - %s", widget.CommonSettings().Title, widget.title(repo))
	if repo == nil {
		return title, " GitHub repo data is unavailable ", false
	}

	_, _, width, _ := widget.View.GetRect()
	str := widget.settings.common.SigilStr(len(widget.GithubRepos), widget.Idx, width) + "\n"
	str += " [red]Stats[white]\n"
	str += widget.displayStats(repo)
	str += "\n [red]Open Review Requests[white]\n"
	str += widget.displayMyReviewRequests(repo, widget.settings.username)
	str += "\n [red]My Pull Requests[white]\n"
	str += widget.displayMyPullRequests(repo, widget.settings.username)
	for _, customQuery := range widget.settings.customQueries {
		str += fmt.Sprintf("\n [red]%s[white]\n", customQuery.title)
		str += widget.displayCustomQuery(repo, customQuery.filter, customQuery.perPage)
	}

	return title, str, false
}

func (widget *Widget) displayMyPullRequests(repo *GithubRepo, username string) string {
	prs := repo.myPullRequests(username, widget.settings.enableStatus)

	numSelections := widget.GetSelected()

	if len(prs) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for _, pr := range prs {
		str += fmt.Sprintf(`[green]["%d"]%4d[""][white] %s`, numSelections, *pr.Number, *pr.Title)
		str += "\n"
		numSelections++
	}

	widget.SetItemCount(numSelections)

	return str
}

func (widget *Widget) displayCustomQuery(repo *GithubRepo, filter string, perPage int) string {
	res := repo.customIssueQuery(filter, perPage)
	
	numSelections := widget.GetSelected()
	
	if res == nil {
		return " [grey]Invalid Query[white]\n"
	}

	if len(res.Issues) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for _, issue := range res.Issues {
		str += fmt.Sprintf(`[green]["%d"]%4d[""][white] %s`, numSelections, *issue.Number, *issue.Title)
		str += "\n"
		numSelections++
	}

	widget.SetItemCount(numSelections)

	return str
}

func (widget *Widget) displayMyReviewRequests(repo*GithubRepo, username string) string {
	prs := repo.myReviewRequests(username)

	numSelections := widget.GetSelected()

	if len(prs) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for _, pr := range prs {
		str += fmt.Sprintf(`[green]["%d"]%4d[""][white] %s`, numSelections, *pr.Number, *pr.Title)
		str += "\n"
		numSelections++
	}

	widget.SetItemCount(numSelections)

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
