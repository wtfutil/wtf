package github

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
	ghb "github.com/google/go-github/github"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

var Config *config.Config

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget("Github", "github"),
	}

	widget.addView()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	client := NewClient()

	repo, _ := client.Repository(Config.UString("wtf.github.organization"), Config.UString("wtf.github.repo"))
	org := *repo.Organization

	prs, _ := client.PullRequests(Config.UString("wtf.github.organization"), Config.UString("wtf.github.repo"))

	title := fmt.Sprintf("[green]%s - %s[white]", *org.Login, *repo.Name)

	widget.View.SetTitle(fmt.Sprintf(" Github: %s ", title))
	widget.RefreshedAt = time.Now()

	str := " [red]Open Review Requests[white]\n"
	str = str + widget.prsForReview(prs)
	str = str + "\n"
	str = str + " [red]Open Pull Requests[white]\n"
	str = str + widget.openPRs(prs)

	widget.View.Clear()
	fmt.Fprintf(widget.View, str)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) addView() {
	view := tview.NewTextView()

	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorGray)
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name)
	view.SetWrap(false)

	widget.View = view
}

func (widget *Widget) prsForReview(prs []*ghb.PullRequest) string {
	if len(prs) > 0 {
		str := ""

		for _, pr := range prs {
			for _, reviewer := range pr.RequestedReviewers {
				if *reviewer.Login == Config.UString("wtf.github.username") {
					str = str + fmt.Sprintf(" [green]%d[white] %s\n", *pr.Number, *pr.Title)
				}
			}
		}

		if str == "" {
			str = " [grey]none[white]\n"
		}

		return str
	}

	return " [grey]none[white]\n"
}

func (widget *Widget) openPRs(prs []*ghb.PullRequest) string {
	if len(prs) > 0 {
		str := ""

		for _, pr := range prs {
			user := *pr.User

			if *user.Login == Config.UString("wtf.github.username") {
				str = str + fmt.Sprintf(" [green]%d[white] %s\n", *pr.Number, *pr.Title)
			}
		}

		if str == "" {
			str = " [grey]none[white]\n"
		}

		return str
	}

	return " [grey]none[white]\n"
}
