package twitter

import (
	"fmt"
	"html"
	"regexp"

	"github.com/dustin/go-humanize"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

const HelpText = `
  Keyboard commands for Textfile:

    /: Show/hide this help window
    h: Previous Twitter name
    l: Next Twitter name

    arrow left:  Previous Twitter name
    arrow right: Next Twitter name
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.MultiSourceWidget
	wtf.TextWidget

	client  *Client
	idx     int
	sources []string
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		HelpfulWidget:     wtf.NewHelpfulWidget(app, pages, HelpText),
		MultiSourceWidget: wtf.NewMultiSourceWidget("twitter", "screenName", "screenNames"),
		TextWidget:        wtf.NewTextWidget(app, "Twitter", "twitter", true),

		idx: 0,
	}

	widget.HelpfulWidget.SetView(widget.View)

	widget.LoadSources()
	widget.SetDisplayFunction(widget.display)

	widget.client = NewClient()

	widget.View.SetBorderPadding(1, 1, 1, 1)
	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)
	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh is called on the interval and refreshes the data
func (widget *Widget) Refresh() {
	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	widget.client.screenName = widget.CurrentSource()
	tweets := widget.client.Tweets()

	widget.View.SetTitle(widget.ContextualTitle(fmt.Sprintf("Twitter - [green]@%s[white]", widget.CurrentSource())))

	if len(tweets) == 0 {
		str := fmt.Sprintf("\n\n\n%s", wtf.CenterText("[blue]No Tweets[white]", 50))
		widget.View.SetText(str)
		return
	}

	str := wtf.SigilStr(len(widget.Sources), widget.Idx, widget.View) + "\n"
	for _, tweet := range tweets {
		str = str + widget.format(tweet)
	}

	widget.View.SetText(str)
}

// If the tweet's Username is the same as the account we're watching, no
// need to display the username
func (widget *Widget) displayName(tweet Tweet) string {
	if widget.CurrentSource() == tweet.User.ScreenName {
		return ""
	}
	return tweet.User.ScreenName
}

func (widget *Widget) formatText(text string) string {
	result := text

	// Convert HTML entities
	result = html.UnescapeString(result)

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
	body := widget.formatText(tweet.Text)
	name := widget.displayName(tweet)

	var attribution string
	if name == "" {
		attribution = humanize.Time(tweet.Created())
	} else {
		attribution = fmt.Sprintf(
			"%s, %s",
			name,
			humanize.Time(tweet.Created()),
		)
	}

	return fmt.Sprintf("%s\n[grey]%s[white]\n\n", body, attribution)
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "/":
		widget.ShowHelp()
		return nil
	case "h":
		widget.Prev()
		return nil
	case "l":
		widget.Next()
		return nil
	case "o":
		wtf.OpenFile(widget.CurrentSource())
		return nil
	}

	switch event.Key() {
	case tcell.KeyLeft:
		widget.Prev()
		return nil
	case tcell.KeyRight:
		widget.Next()
		return nil
	default:
		return event
	}

	return event
}
