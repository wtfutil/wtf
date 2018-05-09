package wtf

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const offscreen = -1000
const modalWidth = 80
const modalHeight = 22

func NewBillboardModal(text string, closeFunc func()) *tview.Frame {
	keyboardIntercept := func(event *tcell.EventKey) *tcell.EventKey {
		switch string(event.Rune()) {
		case "/":
			closeFunc()
			return nil
		}

		switch event.Key() {
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
	textView.SetWrap(true)
	textView.SetText(text)

	textView.SetBackgroundColor(tview.Styles.ContrastBackgroundColor)
	textView.SetInputCapture(keyboardIntercept)

	thing := tview.NewFrame(textView)
	thing.SetRect(offscreen, offscreen, modalWidth, modalHeight)

	drawFunc := func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		w, h := screen.Size()
		thing.SetRect((w/2)-(width/2), (h/2)-(height/2), width, height)
		return x, y, width, height
	}

	thing.SetBackgroundColor(tview.Styles.ContrastBackgroundColor)
	thing.SetBorder(true)
	thing.SetBorders(1, 1, 0, 0, 1, 1)
	thing.SetDrawFunc(drawFunc)

	return thing
}
