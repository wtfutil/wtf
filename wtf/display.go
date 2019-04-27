package wtf

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
)

type Display struct {
	Grid *tview.Grid
}

func NewDisplay(widgets []Wtfable) *Display {
	display := Display{
		Grid: tview.NewGrid(),
	}

	display.build(widgets)
	display.Grid.SetBackgroundColor(ColorFor(cfg.Config.UString("wtf.colors.background", "black")))

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
	display.Grid.SetColumns(ToInts(cfg.Config.UList("wtf.grid.columns"))...)
	display.Grid.SetRows(ToInts(cfg.Config.UList("wtf.grid.rows"))...)
	display.Grid.SetBorder(false)

	for _, widget := range widgets {
		display.add(widget)
		go Schedule(widget)
	}

	return display.Grid
}
