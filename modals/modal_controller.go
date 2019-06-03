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

// Modallike defines the interface that all modals must adhere to
type Modallike interface {
	Frame() *tview.Frame
	Name() string
	SetCloseFunction(fn func())
}

// A ModalController is used to manipulate the onscreen display of modal overlays
type ModalController struct {
	app            *tview.Application
	currentModal   Modallike
	modalIsVisible bool
	pages          *tview.Pages

	// prevFocused is the widget that was in focus before the modal is displayed
	// This is used to reset focus after the modal is removed
	prevFocused tview.Primitive
}

// NewModalController creates and returns an instance of a ModalController
func NewModalController(app *tview.Application, pages *tview.Pages) ModalController {
	controller := ModalController{
		app:            app,
		modalIsVisible: false,
		pages:          pages,
	}

	return controller
}

// CloseModal closes the currently-displayed modal
func (cont *ModalController) CloseModal() {
	if !cont.modalIsVisible {
		return
	}

	if cont.currentModal == nil {
		return
	}

	cont.pages.RemovePage(cont.currentModal.Name())
	cont.app.SetFocus(cont.prevFocused)

	cont.modalIsVisible = false
}

// ModalIsVisible returns whether or not a modal window is currently showing
func (cont *ModalController) ModalIsVisible() bool {
	return cont.modalIsVisible
}

// ShowModal displays the specified Modallike instance onscreen
func (cont *ModalController) ShowModal(modal Modallike) {
	cont.currentModal = modal
	cont.prevFocused = cont.app.GetFocus()

	cont.currentModal.SetCloseFunction(cont.CloseModal)

	cont.pages.AddPage(cont.currentModal.Name(), cont.currentModal.Frame(), false, true)
	cont.app.SetFocus(cont.currentModal.Frame())

	cont.modalIsVisible = true

	cont.app.QueueUpdate(func() {
		cont.app.Draw()
	})
}
