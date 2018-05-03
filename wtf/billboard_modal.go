package wtf

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type BillboardModal struct{}

func NewBillboardModal(text string, cancelFn func()) *tview.Frame {
	modal := BillboardModal{}

	frame := modal.build(cancelFn)
	frame.AddText(text, false, tview.AlignLeft, tcell.ColorWhite)

	return frame
}

/* -------------------- Unexported Functions -------------------- */

func (modal *BillboardModal) build(cancelFn func()) *tview.Frame {
	form := modal.buildForm(cancelFn)
	frame := modal.buildFrame(form)

	return frame
}

func (modal *BillboardModal) buildForm(cancelFn func()) *tview.Form {
	form := tview.NewForm().
		SetButtonsAlign(tview.AlignCenter).
		SetButtonTextColor(tview.Styles.PrimaryTextColor)

	form.AddButton("Cancel", cancelFn)

	return form
}

func (modal *BillboardModal) buildFrame(form *tview.Form) *tview.Frame {

	frame := tview.NewFrame(form).SetBorders(0, 0, 0, 0, 0, 0)
	frame.SetBorder(true)
	frame.SetRect(20, 20, 80, 20)

	return frame
}
