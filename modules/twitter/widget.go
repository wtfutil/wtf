package twitter

import (
	"fmt"
	"html"
	"regexp"

	"github.com/dustin/go-humanize"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

const HelpText = `
  Keyboard commands for Twitter:

    /: Show/hide this help window
    h: Previous Twitter name
    l: Next Twitter name
    o: Open the Twitter handle in a browser

    arrow left:  Previous Twitter name
    arrow right: Next Twitter name
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.KeyboardWidget
	wtf.MultiSourceWidget
	wtf.TextWidget

	client   *Client
	idx      int
	settings *Settings
	sources  []string
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		HelpfulWidget:     wtf.NewHelpfulWidget(app, pages, HelpText),
		KeyboardWidget:    wtf.NewKeyboardWidget(),
		MultiSourceWidget: wtf.NewMultiSourceWidget(settings.common, "screenName", "screenNames"),
		TextWidget:        wtf.NewTextWidget(app, settings.common, true),

		idx:      0,
		settings: settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.SetDisplayFunction(widget.display)

	widget.client = NewClient(settings)

	widget.View.SetBorderPadding(1, 1, 1, 1)
	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)

	widget.HelpfulWidget.SetView(widget.View)

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

	title := fmt.Sprintf("Twitter - [green]@%s[white]", widget.CurrentSource())

	if len(tweets) == 0 {
		str := fmt.Sprintf("\n\n\n%s", wtf.CenterText("[blue]No Tweets[white]", 50))
		widget.View.SetText(str)
		return
	}

	_, _, width, _ := widget.View.GetRect()
	str := widget.settings.common.SigilStr(len(widget.Sources), widget.Idx, width-2) + "\n"
	for _, tweet := range tweets {
		str = str + widget.format(tweet)
	}

	widget.Redraw(title, str, false)
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
func (widget *Widget) currentSourceURI() string {

	src := "https://twitter.com/" + widget.CurrentSource()
	return src
}
