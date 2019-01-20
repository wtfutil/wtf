// Package {{(Lower .Name)}}
package {{(Lower .Name)}}

import (
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

const HelpText = `
 Keyboard commands for {{(Title .Name)}}:
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.TextWidget
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget(app, "{{(Title .Name)}}", "{{(Lower .Name)}}", true),
	}

	widget.HelpfulWidget.SetView(widget.View)
	widget.unselect()

	widget.View.SetScrollable(true)
	widget.View.SetRegions(true)
	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {

	// The last call should always be to the display function
	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	widget.View.SetWrap(false)

	widget.View.Clear()
	widget.View.SetTitle(widget.Name)
	widget.View.SetText("Some Text")
	widget.View.Highlight(strconv.Itoa(widget.selected)).ScrollToHighlight()
}

func (widget *Widget) unselect() {
	widget.selected = -1
	widget.display()
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	// This switch statement could handle alphanumeric keys
	switch string(event.Rune()) {
	}

	// This switch statement could handle events like the "enter" key
	switch event.Key() {
	}
}
