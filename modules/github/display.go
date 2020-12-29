package github

import (
	"fmt"

	ghb "github.com/google/go-github/v32/github"
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
	str := widget.settings.PaginationMarker(len(widget.GithubRepos), widget.Idx, width)
	if widget.settings.showStats {
		str += fmt.Sprintf("\n [%s]Stats[white]\n", widget.settings.Colors.Subheading)
		str += widget.displayStats(repo)
	}
	if widget.settings.showOpenReviewRequests {
		str += fmt.Sprintf("\n [%s]Open Review Requests[white]\n", widget.settings.Colors.Subheading)
		str += widget.displayMyReviewRequests(repo, username)
	}
	if widget.settings.showMyPullRequests {
		str += fmt.Sprintf("\n [%s]My Pull Requests[white]\n", widget.settings.Colors.Subheading)
		str += widget.displayMyPullRequests(repo, username)
	}
	for _, customQuery := range widget.settings.customQueries {
		str += fmt.Sprintf("\n [%s]%s[white]\n", widget.settings.Colors.Subheading, customQuery.title)
		str += widget.displayCustomQuery(repo, customQuery.filter, customQuery.perPage)
	}

	return title, str, false
}

func (widget *Widget) displayMyPullRequests(repo *Repo, username string) string {
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

func (widget *Widget) displayCustomQuery(repo *Repo, filter string, perPage int) string {
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

func (widget *Widget) displayMyReviewRequests(repo *Repo, username string) string {
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

func (widget *Widget) displayStats(repo *Repo) string {
	locPrinter, err := widget.settings.LocalizedPrinter()
	if err != nil {
		return err.Error()
	}

	str := fmt.Sprintf(
		" PRs: %s  Issues: %s  Stars: %s\n",
		locPrinter.Sprintf("%d", repo.PullRequestCount()),
		locPrinter.Sprintf("%d", repo.IssueCount()),
		locPrinter.Sprintf("%d", repo.StarCount()),
	)

	return str
}

func (widget *Widget) title(repo *Repo) string {
	return fmt.Sprintf(
		"[%s]%s - %s[white]",
		widget.settings.Colors.TextTheme.Title,
		repo.Owner,
		repo.Name,
	)
}

var mergeIcons = map[string]string{
	"dirty":    "[red]\u0021[white] ",
	"clean":    "[green]\u2713[white] ",
	"unstable": "[red]\u2717[white] ",
	"blocked":  "[red]\u2717[white] ",
}

func (widget *Widget) mergeString(pr *ghb.PullRequest) string {
	if !widget.settings.enableStatus {
		return ""
	}
	if str, ok := mergeIcons[pr.GetMergeableState()]; ok {
		return str
	}
	return "? "
}
