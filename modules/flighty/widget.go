package flighty

import (
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.MultiSourceWidget
	view.TextWidget

	Data []*Flight

	client   *OpenSkyClient
	pages    *tview.Pages
	settings *Settings
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		MultiSourceWidget: view.NewMultiSourceWidget(settings.Common, "", "aircraft"),
		TextWidget:        view.NewTextWidget(tviewApp, redrawChan, pages, settings.Common),

		client:   NewOpenSkyClient("", ""),
		pages:    pages,
		settings: settings,
	}

	widget.initializeKeyboardControls()

	widget.SetDisplayFunction(widget.display)

	return widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Fetch(aircraft []string) ([]*Flight, error) {
	data := []*Flight{}

	yesterday := time.Now().Add(-24 * time.Hour)
	today := time.Now()

	for _, plane := range aircraft {
		result, _ := widget.client.Flight(plane, yesterday, today)
		for _, d := range result {
			data = append(data, d)
		}
	}

	return data, nil
}

func (widget *Widget) Refresh() {
	// if widget.authenticationCredentialsValid() {
	data, err := widget.Fetch(widget.settings.aircraft)
	if err == nil {
		widget.Data = data
	}
	// }

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

// func (widget *Widget) authenticationCredentialsValid() bool {
// 	if widget.settings.username == "" {
// 		return false
// 	}

// 	if widget.settings.password == "" {
// 		return false
// 	}

// 	return true
// }
