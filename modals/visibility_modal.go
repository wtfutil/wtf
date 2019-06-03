package modals

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// NewWidgetVisibilityModal creates and returns a control modal
// This modal is used to toggle active modules on and off
func NewWidgetVisibilityModal(closeFunc func()) *tview.Frame {
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
	frame.SetRect(OffscreenPos, OffscreenPos, ModalWidth, ModalHeight)

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
