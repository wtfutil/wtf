package updown

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

const (
	userAgent = "WTFUtil"

	apiURLBase = "https://updown.io"
)

type Widget struct {
	view.ScrollableWidget
	checks   []Check
	settings *Settings
	tokenSet map[string]struct{}
	err      error
}

// Taken from https://github.com/AntoineAugusti/updown/blob/d590ab97f115302c73ecf21647909d8fd06ed6ac/checks.go#L17
type Check struct {
	Token             string            `json:"token,omitempty"`
	URL               string            `json:"url,omitempty"`
	Alias             string            `json:"alias,omitempty"`
	LastStatus        int               `json:"last_status,omitempty"`
	Uptime            float64           `json:"uptime,omitempty"`
	Down              bool              `json:"down"`
	DownSince         string            `json:"down_since,omitempty"`
	Error             string            `json:"error,omitempty"`
	Period            int               `json:"period,omitempty"`
	Apdex             float64           `json:"apdex_t,omitempty"`
	Enabled           bool              `json:"enabled"`
	Published         bool              `json:"published"`
	LastCheckAt       time.Time         `json:"last_check_at,omitempty"`
	NextCheckAt       time.Time         `json:"next_check_at,omitempty"`
	FaviconURL        string            `json:"favicon_url,omitempty"`
	SSL               SSL               `json:"ssl,omitempty"`
	StringMatch       string            `json:"string_match,omitempty"`
	MuteUntil         string            `json:"mute_until,omitempty"`
	DisabledLocations []string          `json:"disabled_locations,omitempty"`
	CustomHeaders     map[string]string `json:"custom_headers,omitempty"`
}

// Taken from https://github.com/AntoineAugusti/updown/blob/d590ab97f115302c73ecf21647909d8fd06ed6ac/checks.go#L10
type SSL struct {
	TestedAt string `json:"tested_at,omitempty"`
	Valid    bool   `json:"valid,omitempty"`
	Error    string `json:"error,omitempty"`
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, redrawChan, pages, settings.Common),
		settings:         settings,
		tokenSet:         make(map[string]struct{}),
	}

	for _, t := range settings.tokens {
		widget.tokenSet[t] = struct{}{}
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()

	return widget
}

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	checks, err := widget.getExistingChecks()
	widget.checks = checks
	widget.err = err
	widget.SetItemCount(len(checks))
	widget.Render()
}

// Render sets up the widget data for redrawing to the screen
func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	numUp := 0
	for _, check := range widget.checks {
		if !check.Down {
			numUp++
		}
	}

	title := fmt.Sprintf("Updown (%d/%d)", numUp, len(widget.checks))

	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	if widget.checks == nil {
		return title, "No checks to display", false
	}

	str := widget.contentFrom(widget.checks)

	return title, str, false
}

func (widget *Widget) contentFrom(checks []Check) string {
	var str string

	for _, check := range checks {
		prefix := ""

		if !check.Enabled {
			prefix += "[yellow] ~ "
		} else if check.Down {
			prefix += "[red] - "
		} else {
			prefix += "[green] + "
		}

		str += fmt.Sprintf(`%s%s [gray](%0.2f|%s)[white]%s`,
			prefix,
			check.Alias,
			check.Uptime,
			timeSincePing(check.LastCheckAt),
			"\n",
		)
	}

	return str
}

func timeSincePing(ts time.Time) string {
	dur := time.Since(ts)
	return dur.Truncate(time.Second).String()
}

func makeURL(baseurl string, path string) (string, error) {
	u, err := url.Parse(baseurl)
	if err != nil {
		return "", err
	}
	u.Path = path
	return u.String(), nil
}

func filterChecks(checks []Check, tokenSet map[string]struct{}) []Check {
	j := 0
	for i := 0; i < len(checks); i++ {
		if _, ok := tokenSet[checks[i].Token]; ok {
			checks[j] = checks[i]
			j++
		}
	}
	return checks[:j]
}

func (widget *Widget) getExistingChecks() ([]Check, error) {
	// See: https://updown.io/api#rest
	u, err := makeURL(apiURLBase, "/api/checks")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", u, http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("X-API-KEY", widget.settings.apiKey)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}

	defer func() { _ = resp.Body.Close() }()

	var checks []Check
	err = utils.ParseJSON(&checks, resp.Body)
	if err != nil {
		return nil, err
	}

	if len(widget.tokenSet) > 0 {
		checks = filterChecks(checks, widget.tokenSet)
	}

	return checks, nil
}
