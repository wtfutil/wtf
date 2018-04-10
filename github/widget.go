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
	if widget.Disabled() {
		return
	}

	client := NewClient()
	prs, _ := client.PullRequests(Config.UString("wtf.mods.github.owner"), Config.UString("wtf.mods.github.repo"))

	widget.View.SetTitle(fmt.Sprintf(" Github: %s ", widget.title()))
	widget.RefreshedAt = time.Now()

	str := " [red]Open Review Requests[white]\n"
	str = str + widget.prsForReview(prs, Config.UString("wtf.mods.github.username"))
	str = str + "\n"
	str = str + " [red]Open Pull Requests[white]\n"
	str = str + widget.openPRs(prs, Config.UString("wtf.mods.github.username"))

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

func (widget *Widget) prsForReview(prs []*ghb.PullRequest, username string) string {
	if len(prs) > 0 {
		str := ""

		for _, pr := range prs {
			for _, reviewer := range pr.RequestedReviewers {
				if *reviewer.Login == username {
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

func (widget *Widget) openPRs(prs []*ghb.PullRequest, username string) string {
	if len(prs) > 0 {
		str := ""

		for _, pr := range prs {
			user := *pr.User

			if *user.Login == username {
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

func (widget *Widget) title() string {
	return fmt.Sprintf("[green]%s - %s[white]", Config.UString("wtf.mods.github.owner"), Config.UString("wtf.mods.github.repo"))
}
