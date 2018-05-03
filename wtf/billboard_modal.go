package wtf

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const offscreen = -1000
const modalWidth = 80
const modalHeight = 22

func NewBillboardModal(text string) *tview.Frame {
	textView := tview.NewTextView()
	textView.SetWrap(true)
	textView.SetText(text)

	textView.SetBackgroundColor(tview.Styles.ContrastBackgroundColor)

	thing := tview.NewFrame(textView)
	thing.SetRect(offscreen, offscreen, modalWidth, modalHeight) // First draw it offscreen and then reposition below

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
