package rollbar

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

// A Widget represents a Rollbar widget
type Widget struct {
	view.ScrollableWidget

	items    *Result
	settings *Settings
	err      error
}

// NewWidget creates a new instance of a widget
func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, pages, settings.Common),

		settings: settings,
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()

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
		widget.err = err
		widget.items = nil
		widget.SetItemCount(0)
	} else {
		widget.items = &items.Results
		widget.SetItemCount(len(widget.items.Items))
	}

	widget.Render()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	title := fmt.Sprintf("%s - %s", widget.CommonSettings().Title, widget.settings.projectName)
	if widget.err != nil {
		return widget.CommonSettings().Title, widget.err.Error(), true
	}
	result := widget.items
	if result == nil || len(result.Items) == 0 {
		return title, "No results", false
	}
	var str string
	if len(result.Items) > widget.settings.count {
		result.Items = result.Items[:widget.settings.count]
	}
	for idx, item := range result.Items {

		row := fmt.Sprintf(
			"[%s] [%s] %s [%s] %s [%s]count: %d [%s]%s",
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
		str += utils.HighlightableHelper(widget.View, row, idx, len(item.Title))
	}

	return title, str, false
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

		utils.OpenFile(
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
