package twitterstats

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	client   *Client
	settings *Settings
}

func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),

		client:   NewClient(settings),
		settings: settings,
	}

	widget.View.SetBorderPadding(1, 1, 1, 1)
	widget.View.SetWrap(false)
	widget.View.SetWordWrap(true)

	return &widget
}

func (widget *Widget) Refresh() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	// Add header row
	str := fmt.Sprintf(
		"[%s]%-12s %10s %8s[white]\n",
		widget.settings.Colors.Subheading,
		"Username",
		"Followers",
		"Tweets",
	)

	stats := widget.client.GetStats()

	// Add rows for each of the followed usernames
	for i, username := range widget.client.screenNames {
		str += fmt.Sprintf(
			"%-12s %10d %8d\n",
			username,
			stats[i].FollowerCount,
			stats[i].TweetCount,
		)
	}

	return "Twitter Stats", str, false
}
