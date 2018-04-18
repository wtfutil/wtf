package github

import (
	"fmt"
	"time"

	ghb "github.com/google/go-github/github"
	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Github ", "github"),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.display()
	widget.RefreshedAt = time.Now()
}

/* -------------------- Unexported Functions -------------------- */

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
