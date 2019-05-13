package hackernews

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	*wtf.ScrollableWidget

	stories  []Story
	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		ScrollableWidget: wtf.NewScrollableWidget(app, pages, settings.common, true),

		settings: settings,
	}

	widget.SetRefreshFunction(widget.Refresh)
	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	storyIds, err := GetStories(widget.settings.storyType)
	if storyIds == nil {
		return
	}

	if err != nil {
		widget.Redraw(widget.CommonSettings.Title, err.Error(), true)
		return
	}
	var stories []Story
	for idx := 0; idx < widget.settings.numberOfStories; idx++ {
		story, e := GetStory(storyIds[idx])
		if e != nil {
			// panic(e)
		} else {
			stories = append(stories, story)
		}
	}

	widget.stories = stories
	widget.SetItemCount(len(stories))

	widget.Render()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) Render() {
	if widget.stories == nil {
		return
	}

	title := fmt.Sprintf("%s - %sstories", widget.CommonSettings.Title, widget.settings.storyType)
	widget.Redraw(title, widget.contentFrom(widget.stories), false)
}

func (widget *Widget) contentFrom(stories []Story) string {
	var str string
	for idx, story := range stories {

		u, _ := url.Parse(story.URL)

		str = str + fmt.Sprintf(
			`["%d"][""][%s]%2d. %s [lightblue](%s)[white][""]`,
			idx,
			widget.RowColor(idx),
			idx+1,
			story.Title,
			strings.TrimPrefix(u.Host, "www."),
		)

		str = str + "\n"
	}

	return str
}

func (widget *Widget) openStory() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.stories != nil && sel < len(widget.stories) {
		story := &widget.stories[sel]
		wtf.OpenFile(story.URL)
	}
}

func (widget *Widget) openComments() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.stories != nil && sel < len(widget.stories) {
		story := &widget.stories[sel]
		wtf.OpenFile(fmt.Sprintf("https://news.ycombinator.com/item?id=%d", story.ID))
	}
}
