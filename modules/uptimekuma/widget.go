package uptimekuma

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

// HeartbeatStatus represents the status of a heartbeat
// Matches JS: DOWN=0, UP=1, PENDING=2, MAINTENANCE=3
type HeartbeatStatus int

const (
	DOWN HeartbeatStatus = iota
	UP
	PENDING
	MAINTENANCE
)

// StatusPageData represents the data from the /api/status-page/<slug> endpoint
type StatusPageData struct {
	Incident *Incident `json:"incident"`
}

// Incident represents an incident in Uptime Kuma
type Incident struct {
	CreatedDate string `json:"createdDate"`
}

// HeartbeatData represents the data from the /api/status-page/heartbeat/<slug> endpoint
type HeartbeatData struct {
	HeartbeatList map[string][]*Heartbeat `json:"heartbeatList"`
	UptimeList    map[string]float64      `json:"uptimeList"`
}

// Heartbeat represents a single heartbeat event
type Heartbeat struct {
	Status int `json:"status"`
}

// Widget is the container for your module's data
type Widget struct {
	view.TextWidget

	settings      *Settings
	statusData    *StatusPageData
	heartbeatData *HeartbeatData
	err           error
}

// NewWidget creates and returns an instance of Widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, pages, settings.common),
		settings:   settings,
	}

	widget.initializeKeyboardControls()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {
	widget.err = nil

	baseURL, slug, err := parseURL(widget.settings.url)
	if err != nil {
		widget.err = err
		widget.display()
		return
	}

	statusData, err := widget.fetchStatusData(baseURL, slug)
	if err != nil {
		widget.err = err
		widget.display()
		return
	}
	widget.statusData = statusData

	heartbeatData, err := widget.fetchHeartbeatData(baseURL, slug)
	if err != nil {
		widget.err = err
		widget.display()
		return
	}
	widget.heartbeatData = heartbeatData

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() string {
	if widget.err != nil {
		return fmt.Sprintf("[red]Error: %v", widget.err)
	}

	if widget.statusData == nil || widget.heartbeatData == nil {
		return "Loading..."
	}

	// Use a single indexed variable for status counts
	statusCounts := [4]int{}
	for _, siteList := range widget.heartbeatData.HeartbeatList {
		if len(siteList) > 0 {
			lastHeartbeat := siteList[len(siteList)-1]
			status := HeartbeatStatus(lastHeartbeat.Status)
			if status >= 0 && int(status) < len(statusCounts) {
				statusCounts[status]++
			}
		}
	}

	var totalUptime float64
	numMonitors := len(widget.heartbeatData.UptimeList)
	if numMonitors > 0 {
		for _, uptime := range widget.heartbeatData.UptimeList {
			totalUptime += uptime
		}
	}

	var avgUptime float64
	if numMonitors > 0 {
		avgUptime = (totalUptime / float64(numMonitors)) * 100
	}

	// Adapted from https://github.com/gethomepage/homepage/blob/00bb1a3f37940a0c3c681c3eef0a10d3e1fa0053/src/widgets/uptimekuma/component.jsx#L41C1-L48C1
	var builder strings.Builder
	var textColor = widget.settings.common.Colors.Text
	downColor := "red"
	if statusCounts[DOWN] == 0 {
		downColor = "green"
	}
	builder.WriteString(fmt.Sprintf("[%s] Up: [green]%d", textColor, statusCounts[UP]))
	builder.WriteString(fmt.Sprintf("[%s] (%.1f%%)", textColor, avgUptime))
	builder.WriteString(fmt.Sprintf("[%s], Down: [%s]%d", textColor, downColor, statusCounts[DOWN]))
	if statusCounts[MAINTENANCE] > 0 {
		builder.WriteString(fmt.Sprintf("[%s], Maint: [%s]%d", textColor, "blue", statusCounts[MAINTENANCE]))
	}
	if statusCounts[PENDING] > 0 {
		builder.WriteString(fmt.Sprintf("[%s], Pend: [%s]%d", textColor, "orange", statusCounts[PENDING]))
	}

	if widget.statusData.Incident != nil {
		// Uptime Kuma's API returns dates like "2023-10-27 10:30:00.123"
		layout := "2006-01-02 15:04:05.999"
		created, err := time.Parse(layout, widget.statusData.Incident.CreatedDate)
		if err == nil {
			hoursAgo := time.Since(created).Hours()
			builder.WriteString(fmt.Sprintf("[%s]\n Incident: %.0fh ago", textColor, hoursAgo))
		} else {
			builder.WriteString(fmt.Sprintf("[%s]\n Incident [unparsable date]", textColor))
		}
	}

	return builder.String()
}

func (widget *Widget) display() {
	widget.Redraw(func() (string, string, bool) {
		return widget.CommonSettings().Title, widget.content(), false
	})
}

func (*Widget) fetchStatusData(baseURL, slug string) (*StatusPageData, error) {
	apiURL := fmt.Sprintf("%s/api/status-page/%s", baseURL, slug)

	resp, err := http.Get(apiURL)
	if resp != nil && resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", resp.Status)
	}
	if resp == nil || err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	var data StatusPageData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (*Widget) fetchHeartbeatData(baseURL, slug string) (*HeartbeatData, error) {
	apiURL := fmt.Sprintf("%s/api/status-page/heartbeat/%s", baseURL, slug)

	resp, err := http.Get(apiURL)
	if resp != nil && resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", resp.Status)
	}
	if resp == nil || err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	var data HeartbeatData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func parseURL(rawURL string) (string, string, error) {
	if rawURL == "" {
		return "", "", fmt.Errorf("URL is not defined")
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return "", "", fmt.Errorf("invalid URL: %w", err)
	}

	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(parts) < 2 || parts[0] != "status" {
		return "", "", fmt.Errorf("invalid status page URL format. Expected '.../status/<slug>'")
	}

	slug := parts[1]
	baseURL := fmt.Sprintf("%s://%s", u.Scheme, u.Host)

	return baseURL, slug, nil
}
