package ping

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"sync"
	"time"

	probing "github.com/prometheus-community/pro-bing"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

// pinger is the subset of *probing.Pinger behavior doPings relies on. It
// exists so tests can substitute a fake implementation instead of sending
// real ICMP packets.
type pinger interface {
	Run() error
	Statistics() *probing.Statistics
}

// newRealPinger builds a *probing.Pinger configured the way doPings expects
// (single packet, 10s timeout) and returns it as a pinger.
func newRealPinger(hostname string) (pinger, error) {
	p, err := probing.NewPinger(hostname)
	if err != nil {
		return nil, err
	}
	p.Count = 1
	p.Timeout = 10 * time.Second
	// Unprivileged (UDP) ICMP isn't supported at all on Windows ("socket:
	// the requested protocol has not been configured into the system"), so
	// use privileged (raw socket) mode there. Other platforms keep the
	// unprivileged default so they don't need elevated/root privileges.
	if runtime.GOOS == "windows" {
		p.SetPrivileged(true)
	}

	return p, nil
}

// Widget is the container for your module's data
type Widget struct {
	view.TextWidget
	hosts []Host

	settings *Settings

	// pingerFactory creates the pinger used for each host. Defaults to
	// newRealPinger; tests override it to avoid real network calls.
	pingerFactory func(hostname string) (pinger, error)
}

// NewWidget creates and returns an instance of Widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.common),

		settings:      settings,
		pingerFactory: newRealPinger,
	}
	widget.hosts = widget.settings.hosts

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// hostUpFromStatistics interprets ping statistics as up/down status. A host
// is considered up if at least one reply packet was received.
func hostUpFromStatistics(stats *probing.Statistics) bool {
	return stats.PacketsRecv > 0
}

func (widget *Widget) doPings() {
	var wg sync.WaitGroup
	for i := range widget.hosts {
		idx := i
		host := widget.hosts[idx]
		widget.hosts[idx].Up = false // reset to false each time
		wg.Add(1)
		go func() {
			defer wg.Done()
			pinger, err := widget.pingerFactory(host.Hostname)
			if err == nil {
				err = pinger.Run() // Blocks until finished.
				if err == nil {
					stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
					widget.hosts[idx].Up = hostUpFromStatistics(stats)
				} else {
					log.Printf("error sending ping: %v", err)
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
	for _, t := range widget.hosts {
		if len(t.Label) > nameWidth {
			nameWidth = len(t.Label) + 2
		}
	}

	s := []string{}
	for _, t := range widget.hosts {
		var status string
		if t.Up {
			status = "[green]Up"
		} else {
			status = "[red]DOWN"
		}
		statusLine := fmt.Sprintf("[white]%-*s: %s", nameWidth, t.Label, status)
		s = append(s, statusLine)
	}

	return strings.Join(s, "\n")
}

func (widget *Widget) display() {
	widget.Redraw(func() (string, string, bool) {
		return widget.CommonSettings().Title, widget.content(), false
	})
}
