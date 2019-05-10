package hackernews

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

const HelpText = `
 Keyboard commands for Hacker News:

   /: Show/hide this help window
   j: Select the next story in the list
   k: Select the previous story in the list
   r: Refresh the data

   arrow down: Select the next story in the list
   arrow up:   Select the previous story in the list

   return: Open the selected story in a browser
   c: Open the comments of the article
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.KeyboardWidget
	wtf.ScrollableWidget

	stories  []Story
	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		HelpfulWidget:    wtf.NewHelpfulWidget(app, pages, HelpText),
		KeyboardWidget:   wtf.NewKeyboardWidget(),
		ScrollableWidget: wtf.NewScrollableWidget(app, settings.common, true),

		settings: settings,
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.HelpfulWidget.SetView(widget.View)

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
			`["%d"][""][%s] [yellow]%d. [%s]%s [blue](%s)[""]`,
			idx,
			widget.RowColor(idx),
			idx+1,
			widget.RowColor(idx),
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
