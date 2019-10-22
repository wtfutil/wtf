package twitterstats

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	client *Client
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, settings.common),
		client:     NewClient(settings),
	}

	widget.View.SetBorderPadding(1, 1, 1, 1)
	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)

	return &widget
}

func (widget *Widget) Refresh() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	usernames := widget.client.screenNames
	stats := widget.client.GetStats()

	// Add header row
	str := fmt.Sprintf("%-16s %8s %8s\n", "Username", "Followers", "Tweets")

	// Add rows for each of the followed usernames
	for i, username := range usernames {
		followerCount := stats[i].followerCount
		tweetCount := stats[i].tweetCount

		str += fmt.Sprintf("%-16s %8d %8d\n", username, followerCount, tweetCount)
	}

	return "Twitter Stats", str, true
}
