package healthchecks

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
)

type Widget struct {
	view.ScrollableWidget
	checks   []Checks
	settings *Settings
	err      error
}

type Health struct {
	Checks []Checks `json:"checks"`
}

type Checks struct {
	Name         string    `json:"name"`
	Tags         string    `json:"tags"`
	Desc         string    `json:"desc"`
	Grace        int       `json:"grace"`
	NPings       int       `json:"n_pings"`
	Status       string    `json:"status"`
	LastPing     time.Time `json:"last_ping"`
	NextPing     time.Time `json:"next_ping"`
	ManualResume bool      `json:"manual_resume"`
	Methods      string    `json:"methods"`
	PingURL      string    `json:"ping_url"`
	UpdateURL    string    `json:"update_url"`
	PauseURL     string    `json:"pause_url"`
	Channels     string    `json:"channels"`
	Timeout      int       `json:"timeout,omitempty"`
	Schedule     string    `json:"schedule,omitempty"`
	Tz           string    `json:"tz,omitempty"`
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, redrawChan, pages, settings.Common),
		settings:         settings,
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

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	numUp := 0
	for _, check := range widget.checks {
		if check.Status == "up" {
			numUp++
		}
	}

	title := fmt.Sprintf("Healthchecks (%d/%d)", numUp, len(widget.checks))

	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	if widget.checks == nil {
		return title, "No checks to display", false
	}

	str := widget.contentFrom(widget.checks)

	return title, str, false
}

func (widget *Widget) contentFrom(checks []Checks) string {
	var str string

	for _, check := range checks {
		prefix := ""

		switch check.Status {
		case "up":
			prefix += "[green] + "
		case "down":
			prefix += "[red] - "
		default:
			prefix += "[yellow] ~ "
		}

		str += fmt.Sprintf(`%s%s [gray](%s|%d)[white]%s`,
			prefix,
			check.Name,
			timeSincePing(check.LastPing),
			check.NPings,
			"\n",
		)
	}

	return str
}

func timeSincePing(ts time.Time) string {
	dur := time.Since(ts)
	return dur.Truncate(time.Second).String()
}

func makeURL(baseurl string, path string, tags []string) (string, error) {
	u, err := url.Parse(baseurl)
	if err != nil {
		return "", err
	}
	u.Path = path
	q := u.Query()
	// If we have several tags
	if len(tags) > 0 {
		for _, tag := range tags {
			q.Add("tag", tag)
		}
		u.RawQuery = q.Encode()
	}
	return u.String(), nil
}

func (widget *Widget) getExistingChecks() ([]Checks, error) {
	// See: https://healthchecks.io/docs/api/#list-checks
	u, err := makeURL(widget.settings.apiURL, "/api/v1/checks/", widget.settings.tags)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", u, http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("X-Api-Key", widget.settings.apiKey)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}

	defer func() { _ = resp.Body.Close() }()

	var health Health
	err = utils.ParseJSON(&health, resp.Body)
	if err != nil {
		return nil, err
	}

	return health.Checks, nil
}
