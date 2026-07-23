package devto

import (
	"context"
	"fmt"

	"github.com/rivo/tview"

	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.ScrollableWidget

	articles []Article
	client   *Client
	settings *Settings
	err      error
	openURL  func(string)
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, redrawChan, pages, settings.Common),

		client:   NewClient(nil, ""),
		settings: settings,
		openURL:  utils.OpenFile,
	}

	widget.SetRenderFunction(widget.Render)
	widget.View.SetScrollable(true)
	widget.initializeKeyboardControls()

	return widget
}

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	ctx := context.Background()

	articles, err := widget.client.FetchArticles(
		ctx,
		widget.settings.contentTag,
		widget.settings.contentUsername,
		widget.settings.contentState,
		widget.settings.numberOfArticles,
	)
	if err != nil {
		widget.err = err
		widget.articles = nil
		widget.SetItemCount(0)
	} else {
		limit := widget.settings.numberOfArticles
		if len(articles) < limit {
			limit = len(articles)
		}
		widget.articles = articles[:limit]
		widget.SetItemCount(len(widget.articles))
	}

	widget.Render()
}

// Render sets up the widget data for redrawing to the screen
func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	title := fmt.Sprintf("%s - %s stories", widget.CommonSettings().Title, widget.settings.contentTag)

	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	articles := widget.articles
	if len(articles) == 0 {
		return title, "No stories to display", false
	}

	var str string
	for idx, article := range articles {
		row := formatArticleRow(idx, article.Title, article.User.Username, widget.RowColor(idx))
		str += utils.HighlightableHelper(widget.View, row, idx, len(article.Title))
	}

	return title, str, false
}

func (widget *Widget) openStory() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.articles != nil && sel < len(widget.articles) {
		article := &widget.articles[sel]
		widget.openURL(article.URL)
	}
}

// formatArticleRow formats a single article row for display.
func formatArticleRow(idx int, title, username, rowColor string) string {
	return fmt.Sprintf(
		`[%s]%2d. %s [lightblue](%s)[white]`,
		rowColor,
		idx+1,
		title,
		username,
	)
}
