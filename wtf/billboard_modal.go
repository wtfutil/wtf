package wtf

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func NewBillboardModal(text string) *tview.Frame {
	textView := tview.NewTextView()
	textView.SetWrap(true)
	textView.SetText(text)

	textView.SetBackgroundColor(tview.Styles.ContrastBackgroundColor)

	thing := tview.NewFrame(textView)

	drawFunc := func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		w, h := screen.Size()
		thing.SetRect((w/2)-40, (h/2)-11, 80, 22)
		return x, y, width, height
	}

	thing.SetBackgroundColor(tview.Styles.ContrastBackgroundColor)
	thing.SetBorder(true)
	thing.SetBorders(1, 1, 0, 0, 1, 1)
	thing.SetDrawFunc(drawFunc)

	return thing
}
