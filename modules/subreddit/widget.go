package subreddit

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.KeyboardWidget
	view.ScrollableWidget

	settings *Settings
	err      error
	links    []Link
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		KeyboardWidget:   view.NewKeyboardWidget(app, pages, settings.common),
		ScrollableWidget: view.NewScrollableWidget(app, settings.common),

		settings: settings,
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.KeyboardWidget.SetView(widget.View)

	return widget

}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	links, err := GetLinks(widget.settings.subreddit, widget.settings.sortOrder, widget.settings.topTimePeriod)
	if err != nil {
		widget.err = err
		widget.links = nil
		widget.SetItemCount(0)
	} else {
		widget.links = links[:widget.settings.numberOfPosts]
		widget.SetItemCount(len(widget.links))
		widget.err = nil
	}
	widget.Render()
}

func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	title := "/r/" + widget.settings.subreddit + " - " + widget.settings.sortOrder
	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	var content string
	for idx, link := range widget.links {
		row := fmt.Sprintf(
			`[%s]%2d. %s`,
			widget.RowColor(idx),
			idx+1,
			tview.Escape(link.Title),
		)
		content += utils.HighlightableHelper(widget.View, row, idx, len(link.Title))
	}

	return title, content, false
}

func (widget *Widget) openLink() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.links != nil && sel < len(widget.links) {
		story := &widget.links[sel]
		utils.OpenFile(story.ItemURL)
	}
}

func (widget *Widget) openReddit() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.links != nil && sel < len(widget.links) {
		story := &widget.links[sel]
		fullLink := "http://reddit.com" + story.Permalink
		utils.OpenFile(fullLink)
	}
}
