package twitter

import (
	"fmt"
	"regexp"

	"github.com/senorprogrammer/wtf/wtf"
)

const apiURL = "https://api.twitter.com/1.1/"

type Widget struct {
	wtf.TextWidget

	screenName string
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget("Twitter", "twitter", false),

		screenName: wtf.Config.UString("wtf.mods.twitter.screenName", "wtfutil"),
	}

	widget.View.SetBorderPadding(1, 1, 1, 1)
	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	client := NewClient(widget.screenName, apiURL)
	userTweets := client.Tweets()

	widget.UpdateRefreshedAt()
	widget.View.SetTitle(widget.ContextualTitle(fmt.Sprintf("Twitter - [green]@%s[white]", client.screenName)))

	widget.View.SetText(widget.contentFrom(userTweets))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(tweets []Tweet) string {
	if len(tweets) == 0 {
		return fmt.Sprintf("\n\n\n%s", wtf.CenterText("[blue]No Tweets[white]", 50))
	}

	str := ""
	for _, tweet := range tweets {
		str = str + widget.format(tweet)
	}

	return str
}

// If the tweet's Username is the same as the account we're watching, no
// need to display the username
func (widget *Widget) displayName(tweet Tweet) string {
	if widget.screenName == tweet.User.ScreenName {
		return ""
	}
	return tweet.User.ScreenName
}

func (widget *Widget) highlightText(text string) string {
	result := text

	// RT indicator
	rtRegExp := regexp.MustCompile(`^RT`)
	result = rtRegExp.ReplaceAllString(result, "[olive]${0}[white::-]")

	// @name mentions
	atRegExp := regexp.MustCompile(`@[0-9A-Za-z_]*`)
	result = atRegExp.ReplaceAllString(result, "[blue]${0}[white]")

	// HTTP(S) links
	linkRegExp := regexp.MustCompile(`http[s:\/.0-9A-Za-z]*`)
	result = linkRegExp.ReplaceAllString(result, "[lightblue::u]${0}[white::-]")

	// Hash tags
	hashRegExp := regexp.MustCompile(`#[0-9A-Za-z_]*`)
	result = hashRegExp.ReplaceAllString(result, "[yellow]${0}[white]")

	return result
}

func (widget *Widget) format(tweet Tweet) string {
	body := widget.highlightText(tweet.Text)
	name := widget.displayName(tweet)

	var attribution string
	if name == "" {
		attribution = tweet.PrettyCreatedAt()
	} else {
		attribution = fmt.Sprintf("%s, %s", name, tweet.PrettyCreatedAt())
	}

	return fmt.Sprintf("%s\n[grey]%s[white]\n\n", body, attribution)
}
