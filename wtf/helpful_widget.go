package wtf

import (
	"github.com/rivo/tview"
)

type HelpfulWidget struct {
	app      *tview.Application
	helpText string
	pages    *tview.Pages
	view     *tview.TextView
}

func NewHelpfulWidget(app *tview.Application, pages *tview.Pages, helpText string) HelpfulWidget {
	widget := HelpfulWidget{
		app:      app,
		helpText: helpText,
		pages:    pages,
	}

	return widget
}

func (widget *HelpfulWidget) SetView(view *tview.TextView) {
	widget.view = view
}

func (widget *HelpfulWidget) ShowHelp() {
	closeFunc := func() {
		widget.pages.RemovePage("help")
		widget.app.SetFocus(widget.view)
	}

	modal := NewBillboardModal(widget.helpText, closeFunc)

	widget.pages.AddPage("help", modal, false, true)
	widget.app.SetFocus(modal)
	widget.app.Draw()
}
