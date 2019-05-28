package wtf

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// A WidgetController is used to manipulate the onscreen display of widgets
type WidgetController struct {
	app       *tview.Application
	isVisible bool
	pages     *tview.Pages

	// prevFocused is the widget that was in focus before this modal is displayed
	// This is used to reset focus after this modal is removed
	prevFocused tview.Primitive
}

// NewWidgetController creates and returns an instance of a widget controller
func NewWidgetController(app *tview.Application, pages *tview.Pages) WidgetController {
	controller := WidgetController{
		app:       app,
		isVisible: false,
		pages:     pages,
	}

	return controller
}

// IsVisible returns whether or not this modal window is currently showing onscreen
func (cont *WidgetController) IsVisible() bool {
	return cont.isVisible
}

// ShowVisibilityModal displays an instance of VisibilityModal on the screen
func (cont *WidgetController) ShowVisibilityModal() {
	modalName := "visibility"

	closeFunc := func() {
		cont.pages.RemovePage(modalName)
		cont.app.SetFocus(cont.prevFocused)

		cont.isVisible = false
	}

	modal := NewVisibilityModal(closeFunc)

	cont.prevFocused = cont.app.GetFocus()

	cont.pages.AddPage(modalName, modal, false, true)
	cont.app.SetFocus(modal)

	cont.isVisible = true

	cont.app.QueueUpdate(func() {
		cont.app.Draw()
	})
}

/* -------------------- Visibility Modal -------------------- */

// NewVisibilityModal creates and returns a control modal
// This modal is used to toggle active modules on and off
func NewVisibilityModal(closeFunc func()) *tview.Frame {
	keyboardIntercept := func(event *tcell.EventKey) *tcell.EventKey {
		if string(event.Rune()) == "/" {
			closeFunc()
			return nil
		}

		switch event.Key() {
		case tcell.KeyCtrlE:
			closeFunc()
			return nil
		case tcell.KeyEsc:
			closeFunc()
			return nil
		case tcell.KeyTab:
			return nil
		default:
			return event
		}
	}

	textView := tview.NewTextView()
	textView.SetDynamicColors(true)
	textView.SetInputCapture(keyboardIntercept)
	textView.SetText(" Hello")
	textView.SetWrap(false)

	frame := tview.NewFrame(textView)
	frame.SetRect(offscreen, offscreen, modalWidth, modalHeight)

	drawFunc := func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		w, h := screen.Size()
		frame.SetRect((w/2)-(width/2), (h/2)-(height/2), width, height)
		return x, y, width, height
	}

	frame.SetBorder(true)
	frame.SetBorders(1, 1, 0, 0, 1, 1)
	frame.SetDrawFunc(drawFunc)

	return frame
}
