package status

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rivo/tview"
)

type Widget struct {
	RefreshedAt time.Time
	View        *tview.TextView
}

func NewWidget() *Widget {
	widget := Widget{
		RefreshedAt: time.Now(),
	}

	widget.addView()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.View.SetTitle(" 🦊 Status ")
	widget.RefreshedAt = time.Now()

	widget.View.Clear()
	fmt.Fprintf(widget.View, "%s", widget.contentFrom())
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) addView() {
	view := tview.NewTextView()

	view.SetBorder(true)
	view.SetDynamicColors(true)
	view.SetTitle(" BambooHR ")

	widget.View = view
}

func (widget *Widget) contentFrom() string {
	//return "cats and gods\ndogs and tacs"
	return fmt.Sprint(rand.Intn(100))
}
