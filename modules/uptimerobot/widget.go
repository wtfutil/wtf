package uptimerobot

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.KeyboardWidget
	view.ScrollableWidget

	monitors []Monitor
	settings *Settings
	err      error
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		KeyboardWidget:   view.NewKeyboardWidget(app, pages, settings.common),
		ScrollableWidget: view.NewScrollableWidget(app, settings.common),

		settings: settings,
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.KeyboardWidget.SetView(widget.View)

	return widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	monitors, err := widget.getMonitors()
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

	title := fmt.Sprintf("UptimeRobot (%d/%d)", numUp, len(widget.monitors))

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
			monitor.Uptime,
		)
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
	resp, err_h := http.PostForm("https://api.uptimerobot.com/v2/getMonitors",
		url.Values{
			"api_key":              {widget.settings.apiKey},
			"format":               {"json"},
			"custom_uptime_ratios": {widget.settings.uptimePeriods},
		},
	)

	if err_h != nil {
		return nil, err_h
	}

	body, _ := ioutil.ReadAll(resp.Body)

	// First pass to read the status
	c := make(map[string]json.RawMessage)
	err_j1 := json.Unmarshal([]byte(body), &c)

	if err_j1 != nil {
		return nil, err_j1
	}

	if string(c["stat"]) != `"ok"` {
		return nil, errors.New(string(body))
	}

	// Second pass to get the actual info
	var monitors []Monitor
	err_j2 := json.Unmarshal(c["monitors"], &monitors)

	if err_j2 != nil {
		return nil, err_j2
	}

	return monitors, nil
}
