package feedreader

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mmcdole/gofeed"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
	"jaytaylor.com/html2text"
)

type ShowType int

const (
	SHOW_TITLE ShowType = iota
	SHOW_LINK
	SHOW_CONTENT
)

// FeedItem represents an item returned from an RSS or Atom feed
type FeedItem struct {
	item   *gofeed.Item
	viewed bool
}

// Widget is the container for RSS and Atom data
type Widget struct {
	view.ScrollableWidget

	stories  []*FeedItem
	parser   *gofeed.Parser
	settings *Settings
	err      error
	showType ShowType
}

func rotateShowType(showtype ShowType) ShowType {
	returnValue := SHOW_TITLE
	switch showtype {
	case SHOW_TITLE:
		returnValue = SHOW_LINK
	case SHOW_LINK:
		returnValue = SHOW_CONTENT
	case SHOW_CONTENT:
		returnValue = SHOW_TITLE
	}
	return returnValue
}

func getShowText(feedItem *FeedItem, showType ShowType) string {
	returnValue := feedItem.item.Title
	switch showType {
	case SHOW_LINK:
		returnValue = feedItem.item.Link
	case SHOW_CONTENT:
		text, _ := html2text.FromString(feedItem.item.Content, html2text.Options{PrettyTables: true})
		returnValue = strings.TrimSpace(feedItem.item.Title + "\n" + strings.TrimSpace(text))
	}
	return returnValue
}

// NewWidget creates a new instance of a widget
func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, pages, settings.Common),

		parser:   gofeed.NewParser(),
		settings: settings,
		showType: SHOW_TITLE,
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()

	return widget
}

/* -------------------- Exported Functions -------------------- */

// Fetch retrieves RSS and Atom feed data
func (widget *Widget) Fetch(feedURLs []string) ([]*FeedItem, error) {
	data := []*FeedItem{}

	for _, feedURL := range feedURLs {
		feedItems, err := widget.fetchForFeed(feedURL)
		if err != nil {
			return nil, err
		}

		data = append(data, feedItems...)
	}

	data = widget.sort(data)

	return data, nil
}

// Refresh updates the data in the widget
func (widget *Widget) Refresh() {
	feedItems, err := widget.Fetch(widget.settings.feeds)
	if err != nil {
		widget.err = err
		widget.stories = nil
		widget.SetItemCount(0)
	} else {
		widget.err = nil
		widget.stories = feedItems
		widget.SetItemCount(len(feedItems))
	}

	widget.Render()
}

// Render sets up the widget data for redrawing to the screen
func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) fetchForFeed(feedURL string) ([]*FeedItem, error) {
	feed, err := widget.parser.ParseURL(feedURL)
	if err != nil {
		return nil, err
	}

	feedItems := []*FeedItem{}

	for idx, gofeedItem := range feed.Items {
		if widget.settings.feedLimit >= 1 && idx >= widget.settings.feedLimit {
			// We only want to get the widget.settings.feedLimit latest articles,
			// not all of them. To get all, set feedLimit to < 1
			break
		}

		feedItem := &FeedItem{
			item:   gofeedItem,
			viewed: false,
		}

		feedItems = append(feedItems, feedItem)
	}

	return feedItems, nil
}

func (widget *Widget) content() (string, string, bool) {
	title := widget.CommonSettings().Title
	if widget.err != nil {
		return title, widget.err.Error(), true
	}
	data := widget.stories
	if len(data) == 0 {
		return title, "No data", false
	}
	var str string

	for idx, feedItem := range data {
		rowColor := widget.RowColor(idx)

		if feedItem.viewed {
			// Grays out viewed items in the list, while preserving background highlighting when selected
			rowColor = "gray"
			if idx == widget.Selected {
				rowColor = fmt.Sprintf("gray:%s", widget.settings.Colors.RowTheme.HighlightedBackground)
			}
		}

		displayText := getShowText(feedItem, widget.showType)

		row := fmt.Sprintf(
			"[%s]%2d. %s[white]",
			rowColor,
			idx+1,
			displayText,
		)

		str += utils.HighlightableHelper(widget.View, row, idx, len(feedItem.item.Title))
	}

	return title, str, false
}

// feedItems are sorted by published date
func (widget *Widget) sort(feedItems []*FeedItem) []*FeedItem {
	sort.Slice(feedItems, func(i, j int) bool {
		return feedItems[i].item.PublishedParsed != nil &&
			feedItems[j].item.PublishedParsed != nil &&
			feedItems[i].item.PublishedParsed.After(*feedItems[j].item.PublishedParsed)
	})

	return feedItems
}

func (widget *Widget) openStory() {
	sel := widget.GetSelected()

	if sel >= 0 && widget.stories != nil && sel < len(widget.stories) {
		story := widget.stories[sel]
		story.viewed = true

		utils.OpenFile(story.item.Link)
	}
}

func (widget *Widget) toggleDisplayText() {
	widget.showType = rotateShowType(widget.showType)
	widget.Render()
}
