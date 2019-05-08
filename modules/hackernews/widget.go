package hackernews

import (
	"fmt"
	"net/url"
	"strconv"
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
	wtf.TextWidget

	app *tview.Application

	stories  []Story
	selected int
	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		HelpfulWidget:  wtf.NewHelpfulWidget(app, pages, HelpText),
		KeyboardWidget: wtf.NewKeyboardWidget(),
		TextWidget:     wtf.NewTextWidget(app, settings.common, true),

		app:      app,
		settings: settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.unselect()

	widget.View.SetScrollable(true)
	widget.View.SetRegions(true)

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
		widget.View.SetWrap(true)
		widget.View.SetTitle(widget.ContextualTitle(widget.CommonSettings.Title))
		widget.View.SetText(err.Error())
	} else {
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
	}

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	if widget.stories == nil {
		return
	}

	title := fmt.Sprintf("%s - %sstories", widget.CommonSettings.Title, widget.settings.storyType)
	widget.Redraw(title, widget.contentFrom(widget.stories), false)
	widget.app.QueueUpdateDraw(func() {
		widget.View.Highlight(strconv.Itoa(widget.selected)).ScrollToHighlight()
	})
}

func (widget *Widget) contentFrom(stories []Story) string {
	var str string
	for idx, story := range stories {
		u, _ := url.Parse(story.URL)
		str = str + fmt.Sprintf(
			`["%d"][""][%s] [yellow]%d. [%s]%s [blue](%s)`,
			idx,
			widget.rowColor(idx),
			idx+1,
			widget.rowColor(idx),
			story.Title,
			strings.TrimPrefix(u.Host, "www."),
		)

		str = str + "\n"
	}

	return str
}

func (widget *Widget) rowColor(idx int) string {
	if widget.View.HasFocus() && (idx == widget.selected) {
		return widget.settings.common.DefaultFocussedRowColor()
	}

	return widget.settings.common.RowColor(idx)
}

func (widget *Widget) next() {
	widget.selected++
	if widget.stories != nil && widget.selected >= len(widget.stories) {
		widget.selected = 0
	}

	widget.display()
}

func (widget *Widget) prev() {
	widget.selected--
	if widget.selected < 0 && widget.stories != nil {
		widget.selected = len(widget.stories) - 1
	}

	widget.display()
}

func (widget *Widget) openStory() {
	sel := widget.selected
	if sel >= 0 && widget.stories != nil && sel < len(widget.stories) {
		story := &widget.stories[widget.selected]
		wtf.OpenFile(story.URL)
	}
}

func (widget *Widget) openComments() {
	sel := widget.selected
	if sel >= 0 && widget.stories != nil && sel < len(widget.stories) {
		story := &widget.stories[widget.selected]
		wtf.OpenFile(fmt.Sprintf("https://news.ycombinator.com/item?id=%d", story.ID))
	}
}

func (widget *Widget) unselect() {
	widget.selected = -1
	widget.display()
}
