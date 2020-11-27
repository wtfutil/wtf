package hackernews

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.ScrollableWidget

	stories  []Story
	settings *Settings
	err      error
}

func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, pages, settings.Common),

		settings: settings,
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()

	return widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	storyIds, err := GetStories(widget.settings.storyType)
	if err != nil {
		widget.err = err
		widget.stories = nil
		widget.SetItemCount(0)
	} else {
		var stories []Story
		for idx := 0; idx < widget.settings.numberOfStories; idx++ {
			story, e := GetStory(storyIds[idx])
			if e == nil {
				stories = append(stories, story)
			}
		}
		widget.stories = stories
		widget.SetItemCount(len(stories))
	}

	widget.Render()
}

// Render sets up the widget data for redrawing to the screen
func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	title := fmt.Sprintf("%s - %s stories", widget.CommonSettings().Title, widget.settings.storyType)

	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	if len(widget.stories) == 0 {
		return title, "No stories to display", false
	}

	var str string
	for idx, story := range widget.stories {
		u, _ := url.Parse(story.URL)

		row := fmt.Sprintf(
			`[%s]%2d. %s [lightblue](%s)[white]`,
			widget.RowColor(idx),
			idx+1,
			story.Title,
			strings.TrimPrefix(u.Host, "www."),
		)

		str += utils.HighlightableHelper(widget.View, row, idx, len(story.Title))
	}

	return title, str, false
}

func (widget *Widget) openComments() {
	story := widget.selectedStory()
	if story != nil {
		utils.OpenFile(story.CommentLink())
	}
}

func (widget *Widget) openStory() {
	story := widget.selectedStory()
	if story != nil {
		utils.OpenFile(story.Link())
	}
}

func (widget *Widget) selectedStory() *Story {
	var story *Story

	sel := widget.GetSelected()
	if sel >= 0 && widget.stories != nil && sel < len(widget.stories) {
		story = &widget.stories[sel]
	}

	return story
}
