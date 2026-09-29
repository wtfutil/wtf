package tennis

import (
	"context"
	"fmt"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

// Widget displays tennis matches from the Live Tennis API
type Widget struct {
	view.TextWidget

	client   *Client
	settings *Settings
	matches  []Match
	err      error
}

// NewWidget creates and returns an instance of Widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, pages, settings.Common),

		client:   NewClient(settings.apiKey, nil, settings.baseURL),
		settings: settings,
	}

	widget.initializeKeyboardControls()

	widget.View.SetScrollable(true)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh fetches the latest matches and redraws the widget
func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	if widget.settings.apiKey == "" {
		widget.matches = nil
		widget.err = nil
		widget.Redraw(widget.content)
		return
	}

	matches, err := widget.client.FetchMatches(
		context.Background(),
		widget.settings.status,
		widget.settings.tour,
		widget.settings.matchLimit,
	)
	if err != nil {
		widget.err = err
		widget.matches = nil
	} else {
		widget.err = nil
		if limit := widget.settings.matchLimit; limit > 0 && len(matches) > limit {
			matches = matches[:limit]
		}
		widget.matches = matches
	}

	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	title := widget.title()

	if widget.settings.apiKey == "" {
		return title, missingKeyText(), true
	}

	if widget.err != nil {
		return title, errorText(widget.err), true
	}

	if len(widget.matches) == 0 {
		return title, fmt.Sprintf("No %s matches", widget.settings.status), false
	}

	lines := make([]string, 0, len(widget.matches))
	for _, match := range widget.matches {
		lines = append(lines, renderMatchLine(match))
	}

	return title, strings.Join(lines, "\n"), false
}

// title renders the widget title. It deliberately avoids square brackets:
// tview parses those as style tags in titles and silently swallows them.
func (widget *Widget) title() string {
	title := widget.CommonSettings().Title

	if widget.settings.tour != "" {
		title = fmt.Sprintf("%s %s", title, strings.ToUpper(widget.settings.tour))
	}

	return fmt.Sprintf("%s (%s)", title, widget.settings.status)
}
