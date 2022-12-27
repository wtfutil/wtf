package lunarphase

import (
	"io"
	"net/http"
	"strings"
    "time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	view.ScrollableWidget

	day       string
	date      time.Time
	result    string
	settings  *Settings
	timeout   time.Duration
	titleBase string
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, redrawChan, pages, settings.Common),
		settings: settings,
	}

    widget.titleBase = widget.settings.Title
	widget.timeout = time.Duration(widget.settings.requestTimeout) * time.Second
    widget.date = time.Now()
    widget.day = widget.date.Format(dateFormat)

	widget.SetRenderFunction(widget.Refresh)
	widget.initializeKeyboardControls()

	return widget
}

func (widget *Widget) Refresh() {
	widget.lunarPhase()

	if !widget.settings.Enabled {
		widget.View.Clear()
		return
	}
	widget.settings.Common.Title = widget.titleBase + " " + widget.day

	widget.Redraw(func() (string, string, bool) { return widget.CommonSettings().Title, widget.result, false })
}

// this method reads the config and calls wttr.in for lunar phase
func (widget *Widget) lunarPhase() {
	client := &http.Client{
		Timeout: widget.timeout,
	}

	language := widget.settings.language

	req, err := http.NewRequest("GET", "https://wttr.in/Moon@" + widget.day + "?AF&lang=" + language, http.NoBody)
	if err != nil {
		widget.result = err.Error()
		return
	}

	req.Header.Set("Accept-Language", widget.settings.language)
	req.Header.Set("User-Agent", "curl")
	response, err := client.Do(req)
	if err != nil {
		widget.result = err.Error()
		return
	}
	defer func() { _ = response.Body.Close() }()

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		widget.result = err.Error()
		return
	}

	widget.result = strings.TrimSpace(wtf.ASCIItoTviewColors(string(contents)))
}

// NextDay shows the next day's lunar phase (KeyRight / 'n')
func (widget *Widget) NextDay() {
	tomorrow := widget.date.AddDate(0, 0, 1)
	widget.setDay(tomorrow)
}

// NextWeek shows the next week's lunar phase (KeyUp / 'N')
func (widget *Widget) NextWeek() {
	nextweek := widget.date.AddDate(0, 0, 7)
	widget.setDay(nextweek)
}

// PrevDay shows the previous day's lunar phase (KeyLeft / 'p')
func (widget *Widget) PrevDay() {
	yesterday := widget.date.AddDate(0, 0, -1)
	widget.setDay(yesterday)
}

// PrevWeek shows the previous week's lunar phase (KeyDown / 'P')
func (widget *Widget) PrevWeek() {
	lastweek := widget.date.AddDate(0, 0, -7)
	widget.setDay(lastweek)
}

func (widget *Widget) setDay(ts time.Time) {
	widget.date = ts
    widget.day = widget.date.Format(dateFormat)
	widget.Refresh()
}

// Open nineplanets.org in a browser (Enter / 'o')
func (widget *Widget) OpenMoonPhase() {
	phasedate := widget.date.Format(phaseFormat)
	utils.OpenFile("https://nineplanets.org/moon/phase/" + phasedate + "/")
}

// Disable/Enable the widget (Ctrl-D)
func (widget *Widget) DisableWidget() {

	if widget.settings.Enabled {
		widget.settings.Enabled = false
	} else {
		widget.settings.Enabled = true
	}

	widget.Refresh()
}
