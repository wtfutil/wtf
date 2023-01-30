package pivotal

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"regexp"
)

const (
	hasPullFailIcon = '💥'
	hasPullIcon     = "🌱"
)

var statusMapEmoji = map[string]string{
	"started":     "🚧",
	"unstarted":   "  ",
	"finished":    "🚀",
	"delivered":   "🚢",
	"rejected":    "❌",
	"accepted":    "✅",
	"planned":     "📅",
	"unscheduled": "❓",
}

func (widget *Widget) display() {
	widget.SetItemCount(widget.CurrentSource().getItemCount())
	widget.ScrollableWidget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	proj := widget.CurrentSource()

	if proj == nil {
		return widget.CommonSettings().Title, "No sources", false
	}

	if proj.Err != nil {
		return widget.CommonSettings().Title, proj.Err.Error(), true
	}

	title := fmt.Sprintf(
		"[%s]%s[white] - %d ",
		widget.settings.Colors.TextTheme.Title,
		proj.name, proj.getItemCount())

	str := ""
	for idx, item := range proj.stories {
		rowColor := widget.RowColor(idx)
		displayText := getShowText(&item)

		row := fmt.Sprintf(
			`[%s]|%s%s| %s[%s]`,
			widget.RowColor(idx),
			getStatusIcon(&item),
			getPullStatusIcon(&item),
			tview.Escape(displayText),
			rowColor,
		)

		str += utils.HighlightableHelper(widget.View, row, idx, len(item.Name))
	}

	return title, str, false
}

func getStatusIcon(story *Story) string {
	state := story.CurrentState
	val, ok := statusMapEmoji[state]
	if ok {
		state = val
	}
	return state
}

func getPullStatusIcon(story *Story) string {
	//prs := len(story.PullRequests)
	var prs string
	prs = "  "
	if len(story.PullRequests) > 0 {
		prs = hasPullIcon
	}
	return prs
}

func getShowText(story *Story) string {
	if story == nil {
		return ""
	}

	space := regexp.MustCompile(`\s+`)
	title := space.ReplaceAllString(story.Name, " ")
	//html.UnescapeString("[" + rowColor + "]" + title)
	return title
}
