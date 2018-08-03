package hackernews

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/logger"
	"github.com/senorprogrammer/wtf/wtf"
	"net/url"
	"strings"
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
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.TextWidget

	stories  []Story
	selected int
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget("Hacker News", "hackernews", true),
	}

	widget.HelpfulWidget.SetView(widget.View)
	widget.unselect()

	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	storyIds, err := GetStories(wtf.Config.UString("wtf.mods.hackernews.storyType", "top"))

	widget.UpdateRefreshedAt()

	if err != nil {
		widget.View.SetWrap(true)
		widget.View.SetTitle(widget.Name)
		widget.View.SetText(err.Error())
	} else {
		var stories []Story
		numberOfStoriesToDisplay := wtf.Config.UInt("wtf.mods.hackernews.numberOfStories", 10)
		for idx := 0; idx <= numberOfStoriesToDisplay; idx++ {
			story, e := GetStory(storyIds[idx])
			logger.Log(fmt.Sprintf("stories %s", story.Title))

			if e != nil {
				panic(e)
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

	widget.View.SetWrap(false)

	widget.View.SetTitle(widget.ContextualTitle(fmt.Sprintf("%s - Stories", widget.Name)))
	widget.View.SetText(widget.contentFrom(widget.stories))
}

func (widget *Widget) contentFrom(stories []Story) string {
	var str string
	for idx, story := range stories {
		u, _ := url.Parse(story.URL)
		str = str + fmt.Sprintf(
			"[%s] [yellow]%d. [%s]%s [blue](%s)\n",
			widget.rowColor(idx),
			idx+1,
			widget.rowColor(idx),
			story.Title,
			strings.TrimPrefix(u.Host, "www."),
		)
	}

	return str
}

func (widget *Widget) rowColor(idx int) string {
	if widget.View.HasFocus() && (idx == widget.selected) {
		return wtf.DefaultFocussedRowColor()
	}

	return wtf.RowColor("hackernews", idx)
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

func (widget *Widget) unselect() {
	widget.selected = -1
	widget.display()
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "/":
		widget.ShowHelp()
	case "j":
		widget.next()
		return nil
	case "k":
		widget.prev()
		return nil
	case "r":
		widget.Refresh()
		return nil
	}

	switch event.Key() {
	case tcell.KeyDown:
		widget.next()
		return nil
	case tcell.KeyEnter:
		widget.openStory()
		return nil
	case tcell.KeyEsc:
		widget.unselect()
		return event
	case tcell.KeyUp:
		widget.prev()
		return nil
	default:
		return event
	}
}
