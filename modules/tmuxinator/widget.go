package tmuxinator

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

// Widget is the container for your module's data
type Widget struct {
	view.TextWidget

	settings *Settings
	Selected int
	maxItems int
	Items    []string
}

// NewWidget creates and returns an instance of Widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, pages, settings.common),

		settings: settings,
	}

	widget.View.SetRegions(true)

	widget.initializeKeyboardControls()

	widget.Unselect()

	widget.Items = fetchProjectList()
	widget.SetItemCount(len(widget.Items))

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// SetItemCount sets the amount of PRs RRs and other PRs throughout the widgets display creation
func (widget *Widget) SetItemCount(items int) {
	widget.maxItems = items
}

// GetItemCount returns the amount of PRs RRs and other PRs calculated so far as an int
func (widget *Widget) GetItemCount() int {
	return widget.maxItems
}

// GetSelected returns the index of the currently highlighted item as an int
func (widget *Widget) GetSelected() int {
	if widget.Selected < 0 {
		return 0
	}
	return widget.Selected
}


// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {

    // The last call should always be to the display function
    widget.display()
}

func (widget *Widget) RowColor(idx int) string {
	if widget.View.HasFocus() && (idx == widget.Selected) {
		foreground := widget.CommonSettings().Colors.RowTheme.HighlightedForeground

		return fmt.Sprintf(
			"%s:%s",
			foreground,
			widget.CommonSettings().Colors.RowTheme.HighlightedBackground,
		)
	}

	if idx%2 == 0 {
		return fmt.Sprintf(
			"%s:%s",
			widget.settings.common.Colors.RowTheme.EvenForeground,
			widget.settings.common.Colors.RowTheme.EvenBackground,
		)
	}

	return fmt.Sprintf(
		"%s:%s",
		widget.settings.common.Colors.RowTheme.OddForeground,
		widget.settings.common.Colors.RowTheme.OddBackground,
	)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	widget.Redraw(func() (string, string, bool) {
		return widget.CommonSettings().Title, widget.content(), false
	})
}

func (widget *Widget) content() string {
	cnt := ""

	if len(widget.Items) <= 0 {
		cnt += " [grey]No projects found[white]\n"
	} 

	for idx, projectName := range widget.Items {
		cnt += fmt.Sprintf("[%s] %s \n",
																	widget.RowColor(idx),
																	projectName)
	}

	return cnt
}
