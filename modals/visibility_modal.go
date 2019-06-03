package modals

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// VisibilityModal is a modal that controls the onscreen visibility of widgets
type VisibilityModal struct {
	closeFunc         func()
	frame             *tview.Frame
	keyboardIntercept func(event *tcell.EventKey) *tcell.EventKey
	name              string
}

// NewVisibilityModal creates and returns a control modal
// This modal is used to toggle active modules on and off
func NewVisibilityModal() *VisibilityModal {
	modal := &VisibilityModal{
		name: "visibility",
	}

	keyboardIntercept := func(event *tcell.EventKey) *tcell.EventKey {
		if string(event.Rune()) == "/" {
			modal.execCloseFunc()
			return nil
		}

		switch event.Key() {
		case tcell.KeyCtrlE:
			modal.execCloseFunc()
			return nil
		case tcell.KeyEsc:
			modal.execCloseFunc()
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
	frame.SetRect(OffscreenPos, OffscreenPos, ModalWidth, ModalHeight)

	drawFunc := func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		w, h := screen.Size()
		frame.SetRect((w/2)-(width/2), (h/2)-(height/2), width, height)
		return x, y, width, height
	}

	frame.SetBorder(true)
	frame.SetBorders(1, 1, 0, 0, 1, 1)
	frame.SetDrawFunc(drawFunc)

	modal.frame = frame
	modal.keyboardIntercept = keyboardIntercept

	return modal
}

/* -------------------- Exported Functions -------------------- */

// Frame returns the frame used to display the modal
func (modal *VisibilityModal) Frame() *tview.Frame {
	return modal.frame
}

// Name returns the name of the modal
func (modal *VisibilityModal) Name() string {
	return modal.name
}

// SetCloseFunction assigns the function that should be called to close this modal
func (modal *VisibilityModal) SetCloseFunction(fn func()) {
	modal.closeFunc = fn
}

/* -------------------- Unexported Functions -------------------- */

func (modal *VisibilityModal) execCloseFunc() {
	modal.closeFunc()
}
