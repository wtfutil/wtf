package status

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rivo/tview"
)

type Widget struct {
	Name            string
	RefreshedAt     time.Time
	RefreshInterval int
	View            *tview.TextView
}

func NewWidget() *Widget {
	widget := Widget{
		Name:            "Status",
		RefreshedAt:     time.Now(),
		RefreshInterval: 1,
	}

	widget.addView()
	go widget.refresher()

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
	view.SetTitle(widget.Name)

	widget.View = view
}

func (widget *Widget) contentFrom() string {
	return fmt.Sprint(rand.Intn(100))
}

func (widget *Widget) refresher() {
	tick := time.NewTicker(time.Duration(widget.RefreshInterval) * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-tick.C:
			widget.Refresh()
		case <-quit:
			tick.Stop()
			return
		}
	}
}
