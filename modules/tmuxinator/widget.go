package tmuxinator

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	settings *Settings
	Selected int
	maxItems int
	Items    []string
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, pages, settings.common),

		settings: settings,
	}

	widget.initializeKeyboardControls()

	widget.Items = fetchProjectList()

	widget.Unselect()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) GetSelected() int {
	if widget.Selected < 0 {
		return 0
	}

	return widget.Selected
}

func (widget *Widget) MaxItems() int {
	return len(widget.Items)
}

func (widget *Widget) Refresh() {
	widget.Items = fetchProjectList()
	widget.Unselect()
	widget.display()
}

func (widget *Widget) RowColor(idx int) string {
	if widget.View.HasFocus() && (idx == widget.GetSelected()) {
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
