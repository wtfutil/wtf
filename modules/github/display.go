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
	username := widget.settings.username

	// Choses the correct place to scroll to when changing sources
	if len(widget.View.GetHighlights()) > 0 {
		widget.View.ScrollToHighlight()
	} else {
		widget.View.ScrollToBeginning()
	}

	// initial maxItems count
	widget.Items = make([]int, 0)
	widget.SetItemCount(len(repo.myReviewRequests((username))))

	title := fmt.Sprintf("%s - %s", widget.CommonSettings().Title, widget.title(repo))
	if repo == nil {
		return title, " GitHub repo data is unavailable ", false
	} else if repo.Err != nil {
		return title, repo.Err.Error(), true
	}

	_, _, width, _ := widget.View.GetRect()
	str := widget.settings.common.SigilStr(len(widget.GithubRepos), widget.Idx, width) + "\n"
	str += " [red]Stats[white]\n"
	str += widget.displayStats(repo)
	str += "\n [red]Open Review Requests[white]\n"
	str += widget.displayMyReviewRequests(repo, username)
	str += "\n [red]My Pull Requests[white]\n"
	str += widget.displayMyPullRequests(repo, username)
	for _, customQuery := range widget.settings.customQueries {
		str += fmt.Sprintf("\n [red]%s[white]\n", customQuery.title)
		str += widget.displayCustomQuery(repo, customQuery.filter, customQuery.perPage)
	}

	return title, str, false
}

func (widget *Widget) displayMyPullRequests(repo *GithubRepo, username string) string {
	prs := repo.myPullRequests(username, widget.settings.enableStatus)

	prLength := len(prs)

	if prLength == 0 {
		return " [grey]none[white]\n"
	}

	maxItems := widget.GetItemCount()

	str := ""
	for idx, pr := range prs {
		str += fmt.Sprintf(` %s[green]["%d"]%4d[""][white] %s`, widget.mergeString(pr), maxItems+idx, *pr.Number, *pr.Title)
		str += "\n"
		widget.Items = append(widget.Items, *pr.Number)
	}

	widget.SetItemCount(maxItems + prLength)

	return str
}

func (widget *Widget) displayCustomQuery(repo *GithubRepo, filter string, perPage int) string {
	res := repo.customIssueQuery(filter, perPage)

	if res == nil {
		return " [grey]Invalid Query[white]\n"
	}

	issuesLength := len(res.Issues)

	if issuesLength == 0 {
		return " [grey]none[white]\n"
	}

	maxItems := widget.GetItemCount()

	str := ""
	for idx, issue := range res.Issues {
		str += fmt.Sprintf(` [green]["%d"]%4d[""][white] %s`, maxItems+idx, *issue.Number, *issue.Title)
		str += "\n"
		widget.Items = append(widget.Items, *issue.Number)
	}

	widget.SetItemCount(maxItems + issuesLength)

	return str
}

func (widget *Widget) displayMyReviewRequests(repo *GithubRepo, username string) string {
	prs := repo.myReviewRequests(username)

	if len(prs) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for idx, pr := range prs {
		str += fmt.Sprintf(` [green]["%d"]%4d[""][white] %s`, idx, *pr.Number, *pr.Title)
		str += "\n"
		widget.Items = append(widget.Items, *pr.Number)
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
	"dirty":    "[red]\u0021[white] ",
	"clean":    "[green]\u2713[white] ",
	"unstable": "[red]\u2717[white] ",
	"blocked":  "[red]\u2717[white] ",
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
