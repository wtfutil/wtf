package wtf

import (
	"github.com/olebedev/config"
	"github.com/rivo/tview"
)

type Display struct {
	Grid   *tview.Grid
	config *config.Config
}

func NewDisplay(widgets []Wtfable, config *config.Config) *Display {
	display := Display{
		Grid:   tview.NewGrid(),
		config: config,
	}

	display.build(widgets)
	display.Grid.SetBackgroundColor(ColorFor(config.UString("wtf.colors.background", "black")))

	return &display
}

/* -------------------- Unexported Functions -------------------- */

func (display *Display) add(widget Wtfable) {
	if widget.Disabled() {
		return
	}

	if !widget.IsPositionable() {
		return
	}

	display.Grid.AddItem(
		widget.TextView(),
		widget.Top(),
		widget.Left(),
		widget.Height(),
		widget.Width(),
		0,
		0,
		false,
	)
}

func (display *Display) build(widgets []Wtfable) *tview.Grid {
	display.Grid.SetColumns(ToInts(display.config.UList("wtf.grid.columns"))...)
	display.Grid.SetRows(ToInts(display.config.UList("wtf.grid.rows"))...)
	display.Grid.SetBorder(false)

	for _, widget := range widgets {
		display.add(widget)
		go Schedule(widget)
	}

	return display.Grid
}
