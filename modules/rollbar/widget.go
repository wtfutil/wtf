package rollbar

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

const HelpText = `
 Keyboard commands for Rollbar:

   /: Show/hide this help window
   j: Select the next item in the list
   k: Select the previous item in the list
   r: Refresh the data
   u: unselect the current item(removes item being perma highlighted)

   arrow down: Select the next item in the list
   arrow up:   Select the previous item in the list

   return: Open the selected item in a browser
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.TextWidget

	items    *Result
	selected int
	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget(app, settings.common, true),

		settings: settings,
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

	items, err := CurrentActiveItems(
		widget.settings.accessToken,
		widget.settings.assignedToName,
		widget.settings.activeOnly,
	)

	if err != nil {
		widget.View.SetWrap(true)
		widget.View.SetTitle(widget.Name())
		widget.View.SetText(err.Error())
	} else {
		widget.items = &items.Results
	}

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	if widget.items == nil {
		return
	}

	widget.View.SetWrap(false)
	widget.View.SetTitle(widget.ContextualTitle(fmt.Sprintf("%s - %s", widget.Name(), widget.settings.projectName)))
	widget.View.SetText(widget.contentFrom(widget.items))
}

func (widget *Widget) contentFrom(result *Result) string {
	var str string
	if len(result.Items) > widget.settings.count {
		result.Items = result.Items[:widget.settings.count]
	}
	for idx, item := range result.Items {

		str = str + fmt.Sprintf(
			"[%s] [%s] %s [%s] %s [%s]count: %d [%s]%s\n",
			widget.rowColor(idx),
			levelColor(&item),
			item.Level,
			statusColor(&item),
			item.Title,
			widget.rowColor(idx),
			item.TotalOccurrences,
			"blue",
			item.Environment,
		)
	}

	return str
}

func (widget *Widget) rowColor(idx int) string {
	if widget.View.HasFocus() && (idx == widget.selected) {
		widget.settings.common.DefaultFocussedRowColor()
	}

	return widget.settings.common.RowColor(idx)
}

func statusColor(item *Item) string {
	switch item.Status {
	case "active":
		return "red"
	case "resolved":
		return "green"
	default:
		return "red"
	}
}
func levelColor(item *Item) string {
	switch item.Level {
	case "error":
		return "red"
	case "critical":
		return "green"
	case "warning":
		return "yellow"
	default:
		return "grey"
	}
}

func (widget *Widget) next() {
	widget.selected++
	if widget.items != nil && widget.selected >= len(widget.items.Items) {
		widget.selected = 0
	}

	widget.display()
}

func (widget *Widget) prev() {
	widget.selected--
	if widget.selected < 0 && widget.items.Items != nil {
		widget.selected = len(widget.items.Items) - 1
	}

	widget.display()
}

func (widget *Widget) openBuild() {
	if widget.selected >= 0 && widget.items != nil && widget.selected < len(widget.items.Items) {
		item := &widget.items.Items[widget.selected]

		wtf.OpenFile(
			fmt.Sprintf(
				"https://rollbar.com/%s/%s/%s/%d",
				widget.settings.projectOwner,
				widget.settings.projectName,
				"items",
				item.ID,
			),
		)
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
	case "u":
		widget.unselect()
		return nil
	}

	switch event.Key() {
	case tcell.KeyDown:
		widget.next()
		return nil
	case tcell.KeyEnter:
		widget.openBuild()
		return nil
	case tcell.KeyEsc:
		widget.unselect()
		return event
	case tcell.KeyUp:
		widget.prev()
		widget.display()
		return nil
	default:
		return event
	}
}
