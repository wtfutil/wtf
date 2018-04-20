package github

import (
	"fmt"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

func (widget *Widget) display() {
	widget.View.Clear()

	repo := widget.currentData()
	if repo == nil {
		fmt.Fprintf(widget.View, "%s", " Github repo data is unavailable (1)")
		return
	}

	widget.View.SetTitle(fmt.Sprintf(" Github: %s ", widget.title(repo)))

	str := widget.tickMarks(widget.Data) + "\n"
	str = str + " [red]Open Review Requests[white]\n"
	str = str + widget.prsForReview(repo, Config.UString("wtf.mods.github.username"))
	str = str + "\n"
	str = str + " [red]Open Pull Requests[white]\n"
	str = str + widget.openPRs(repo, Config.UString("wtf.mods.github.username"))

	fmt.Fprintf(widget.View, str)
}

func (widget *Widget) openPRs(repo *GithubRepo, username string) string {
	if len(repo.PullRequests) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""

	for _, pr := range repo.PullRequests {
		user := *pr.User

		if *user.Login == username {
			str = str + fmt.Sprintf(" [green]%d[white] %s\n", *pr.Number, *pr.Title)
		}
	}

	if str == "" {
		return " [grey]none[white]\n"
	}

	return str
}

func (widget *Widget) prsForReview(repo *GithubRepo, username string) string {
	if len(repo.PullRequests) > 0 {
		return " [grey]none[white]\n"
	}

	str := ""

	for _, pr := range repo.PullRequests {
		for _, reviewer := range pr.RequestedReviewers {
			if *reviewer.Login == username {
				str = str + fmt.Sprintf(" [green]%d[white] %s\n", *pr.Number, *pr.Title)
			}
		}
	}

	if str == "" {
		return " [grey]none[white]\n"
	}

	return str
}

func (widget *Widget) tickMarks(data []*GithubRepo) string {
	str := ""

	if len(data) > 1 {
		marks := strings.Repeat("*", len(data))
		marks = marks[:widget.Idx] + "_" + marks[widget.Idx+1:]

		str = "[lightblue]" + fmt.Sprintf(wtf.RightAlignFormat(widget.View), marks) + "[white]"
	}

	return str
}

func (widget *Widget) title(repo *GithubRepo) string {
	return fmt.Sprintf("[green]%s - %s[white]", repo.Owner, repo.Name)
}
