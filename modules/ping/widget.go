package ping

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
	"github.com/prometheus-community/pro-bing"
)

// Widget is the container for your module's data
type Widget struct {
	view.TextWidget
	targets []Target

	settings *Settings
}

// NewWidget creates and returns an instance of Widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.common),

		settings: settings,
	}
	widget.targets = widget.settings.targets

	return &widget
}

/* -------------------- Exported Functions -------------------- */


func (widget *Widget) doPings() {
	var wg sync.WaitGroup
	for i := range widget.targets {
		idx := i
		target := widget.targets[idx]
		widget.targets[idx].Up = false // reset to false each time
		wg.Add(1)
		go func() {
			defer wg.Done()
			pinger, err := probing.NewPinger(target.Host)
			if err == nil {
				pinger.Count = 1
				pinger.Timeout = 10*time.Second
				err = pinger.Run() // Blocks until finished.
				if err == nil {
					stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
					if stats.PacketsRecv > 0 {
						widget.targets[idx].Up = true
					} else {
						widget.targets[idx].Up = false
					}
				} else {
					log.Fatalf("error sending ping: %v", err)
				}
			}

		}()
	}
	wg.Wait()
}
func (widget *Widget) Refresh() {

	widget.doPings()
    widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() string {
	nameWidth := 12
	for _,t := range widget.targets {
		if len(t.Name) > nameWidth {
			nameWidth = len(t.Name) + 2
		}
	}

	s := []string{}
	for _,t := range widget.targets {
		var status string
		if t.Up == true {
			status = "[green]Up"
		} else {
			status = "[red]DOWN"
		}
		statusLine := fmt.Sprintf("[white]%-*s: %s", nameWidth, t.Name, status)
		s = append(s, statusLine)
	}

	return strings.Join(s, "\n")
}

func (widget *Widget) display() {
	widget.Redraw(func() (string, string, bool) {
		return widget.CommonSettings().Title, widget.content(), false
	})
}
