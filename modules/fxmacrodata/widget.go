package fxmacrodata

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

// Widget is the container for this module's data
type Widget struct {
	view.TextWidget

	settings *Settings
	client   *http.Client
	err      error
	releases []Release
}

// NewWidget creates and returns an instance of Widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.common),

		settings: settings,
		client:   &http.Client{Timeout: 10 * time.Second},
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {
	widget.releases, widget.err = upcoming(widget.settings, widget.client, time.Now())

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() string {
	if widget.err != nil {
		return fmt.Sprintf("[red]%s[white]", widget.err.Error())
	}

	if len(widget.releases) == 0 {
		return " no upcoming releases"
	}

	out := ""
	for _, release := range widget.releases {
		when := time.Unix(release.AnnouncementDatetime, 0).Local()

		// An unconfirmed date is the authority's provisional slot rather than a
		// published time, and acting on it as though it were fixed is the whole
		// problem this module exists to avoid.
		marker := " "
		if !release.Confirmed {
			marker = "~"
		}

		name := release.Name
		if release.TopTier {
			name = fmt.Sprintf("[::b]%s[::-]", name)
		}

		out += fmt.Sprintf(
			" %s [yellow]%s[white] %s  %s\n",
			marker,
			when.Format("Mon 02 Jan 15:04"),
			release.Currency,
			name,
		)
	}

	return out
}

func (widget *Widget) display() {
	widget.Redraw(func() (string, string, bool) {
		return widget.CommonSettings().Title, widget.content(), false
	})
}
