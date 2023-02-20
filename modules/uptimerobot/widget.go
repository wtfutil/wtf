package uptimerobot

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.ScrollableWidget

	monitors []Monitor
	settings *Settings
	err      error
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, redrawChan, pages, settings.Common),

		settings: settings,
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()

	return widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	monitors, err := widget.getMonitors()

	if widget.settings.offlineFirst {
		var tmp Monitor
		var next int
		for i := 0; i < len(monitors); i++ {
			if monitors[i].State != 2 {
				tmp = monitors[i]
				for j := i; j > next; j-- {
					monitors[j] = monitors[j-1]
				}
				monitors[next] = tmp
				next++
			}
		}
	}

	widget.monitors = monitors
	widget.err = err
	widget.SetItemCount(len(monitors))

	widget.Render()
}

// Render sets up the widget data for redrawing to the screen
func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	numUp := 0
	for _, monitor := range widget.monitors {
		if monitor.State == 2 {
			numUp++
		}
	}

	title := fmt.Sprintf("%s (%d/%d)", widget.CommonSettings().Title, numUp, len(widget.monitors))

	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	if widget.monitors == nil {
		return title, "No monitors to display", false
	}

	str := widget.contentFrom(widget.monitors)

	return title, str, false
}

func (widget *Widget) contentFrom(monitors []Monitor) string {
	var str string

	for _, monitor := range monitors {
		prefix := ""

		switch monitor.State {
		case 2:
			prefix += "[green] + "
		case 8:
		case 9:
			prefix += "[red] - "
		default:
			prefix += "[yellow] ~ "
		}

		str += fmt.Sprintf(`%s%s [gray](%s)[white]
`,
			prefix,
			monitor.Name,
			formatUptimes(monitor.Uptime),
		)
	}

	return str
}

func formatUptimes(str string) string {
	splits := strings.Split(str, "-")
	str = ""
	for i, s := range splits {
		if i != 0 {
			str += "|"
		}
		s = s[:5]
		s = strings.TrimRight(s, "0")
		s = strings.TrimRight(s, ".") + "%"
		str += s
	}
	return str
}

type Monitor struct {
	Name string `json:"friendly_name"`
	// Monitor state, see: https://uptimerobot.com/api/#parameters
	State int8 `json:"status"`
	// Uptime ratio, preformatted, e.g.: 100.000-97.233-96.975
	Uptime string `json:"custom_uptime_ratio"`
}

func (widget *Widget) getMonitors() ([]Monitor, error) {
	// See: https://uptimerobot.com/api/#getMonitorsWrap
	resp, errh := http.PostForm("https://api.uptimerobot.com/v2/getMonitors",
		url.Values{
			"api_key":              {widget.settings.apiKey},
			"format":               {"json"},
			"custom_uptime_ratios": {widget.settings.uptimePeriods},
		},
	)

	if errh != nil {
		return nil, errh
	}
	defer func() { _ = resp.Body.Close() }()

	body, _ := io.ReadAll(resp.Body)

	// First pass to read the status
	c := make(map[string]json.RawMessage)
	errj1 := json.Unmarshal(body, &c)

	if errj1 != nil {
		return nil, errj1
	}

	if string(c["stat"]) != `"ok"` {
		return nil, errors.New(string(body))
	}

	// Second pass to get the actual info
	var monitors []Monitor
	errj2 := json.Unmarshal(c["monitors"], &monitors)

	if errj2 != nil {
		return nil, errj2
	}

	return monitors, nil
}
