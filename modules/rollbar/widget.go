package rollbar

import (
	"fmt"

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

// A Widget represents a Rollbar widget
type Widget struct {
	wtf.HelpfulWidget
	wtf.KeyboardWidget
	wtf.ScrollableWidget

	items    *Result
	settings *Settings
}

// NewWidget creates a new instance of a widget
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

	items, err := CurrentActiveItems(
		widget.settings.accessToken,
		widget.settings.assignedToName,
		widget.settings.activeOnly,
	)

	if err != nil {
		widget.Redraw(widget.CommonSettings.Title, err.Error(), true)
		return
	}
	widget.items = &items.Results
	widget.SetItemCount(len(widget.items.Items))

	widget.Refresh()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) Render() {
	if widget.items == nil {
		return
	}

	title := fmt.Sprintf("%s - %s", widget.CommonSettings.Title, widget.settings.projectName)
	widget.Redraw(title, widget.contentFrom(widget.items), false)
}

func (widget *Widget) contentFrom(result *Result) string {
	if result == nil {
		return "No results"
	}
	var str string
	if len(result.Items) > widget.settings.count {
		result.Items = result.Items[:widget.settings.count]
	}
	for idx, item := range result.Items {

		str = str + fmt.Sprintf(
			"[%s] [%s] %s [%s] %s [%s]count: %d [%s]%s\n",
			widget.RowColor(idx),
			levelColor(&item),
			item.Level,
			statusColor(&item),
			item.Title,
			widget.RowColor(idx),
			item.TotalOccurrences,
			"blue",
			item.Environment,
		)
	}

	return str
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

func (widget *Widget) openBuild() {
	if widget.GetSelected() >= 0 && widget.items != nil && widget.GetSelected() < len(widget.items.Items) {
		item := &widget.items.Items[widget.GetSelected()]

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
