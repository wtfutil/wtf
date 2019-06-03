package modals

import (
	"github.com/rivo/tview"
)

const (
	// ModalWidth is the default width for a modal overlay, in characters
	ModalWidth = 80

	// ModalHeight is the default height for a modal overlay, in screen lines
	ModalHeight = 22

	// OffscreenPos is the default starting position of a modal, used to build it before
	// blitting it onscreen
	OffscreenPos = -1000
)

// A ModalController is used to manipulate the onscreen display of widgets
type ModalController struct {
	app            *tview.Application
	modalIsVisible bool
	pages          *tview.Pages

	// prevFocused is the widget that was in focus before this modal is displayed
	// This is used to reset focus after this modal is removed
	prevFocused tview.Primitive
}

// NewModalController creates and returns an instance of a widget controller
func NewModalController(app *tview.Application, pages *tview.Pages) ModalController {
	controller := ModalController{
		app:            app,
		modalIsVisible: false,
		pages:          pages,
	}

	return controller
}

// ModalIsVisible returns whether or not this modal window is currently showing onscreen
func (cont *ModalController) ModalIsVisible() bool {
	return cont.modalIsVisible
}

// ShowWidgetVisibilityModal displays an instance of VisibilityModal on the screen
func (cont *ModalController) ShowWidgetVisibilityModal() {
	modalName := "visibility"

	closeFunc := func() {
		cont.pages.RemovePage(modalName)
		cont.app.SetFocus(cont.prevFocused)

		cont.modalIsVisible = false
	}

	modal := NewWidgetVisibilityModal(closeFunc)

	cont.prevFocused = cont.app.GetFocus()

	cont.pages.AddPage(modalName, modal, false, true)
	cont.app.SetFocus(modal)

	cont.modalIsVisible = true

	cont.app.QueueUpdate(func() {
		cont.app.Draw()
	})
}
