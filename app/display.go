package app

import (
	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/wtf"
)

// Display is the container for the onscreen representation of a WtfApp
type Display struct {
	Grid *tview.Grid

	bgColor tcell.Color
	config  *config.Config
}

// NewDisplay creates and returns a Display
func NewDisplay(widgets []wtf.Wtfable, config *config.Config) *Display {
	display := Display{
		Grid: tview.NewGrid(),

		bgColor: wtf.ColorFor("black"),
		config:  config,
	}

	// Make sure we have some widgets before trying to get the first one
	if len(widgets) > 0 {
		firstWidget := widgets[0]
		display.bgColor = wtf.ColorFor(firstWidget.CommonSettings().Colors.WidgetTheme.Background)
	}

	display.Grid.SetBackgroundColor(display.bgColor)
	display.build(widgets)

	return &display
}

/* -------------------- Unexported Functions -------------------- */

func (display *Display) add(widget wtf.Wtfable) {
	if widget.Disabled() {
		return
	}

	display.Grid.AddItem(
		widget.TextView(),
		widget.CommonSettings().Top,
		widget.CommonSettings().Left,
		widget.CommonSettings().Height,
		widget.CommonSettings().Width,
		0,
		0,
		false,
	)
}

func (display *Display) build(widgets []wtf.Wtfable) *tview.Grid {
	cols := utils.ToInts(display.config.UList("wtf.grid.columns"))
	rows := utils.ToInts(display.config.UList("wtf.grid.rows"))

	display.Grid.SetColumns(cols...)
	display.Grid.SetRows(rows...)
	display.Grid.SetBorder(false)

	for _, widget := range widgets {
		display.add(widget)
	}

	return display.Grid
}
