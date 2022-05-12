package krisinformation

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

// Widget is the container for your module's data
type Widget struct {
	view.TextWidget

	app      *tview.Application
	settings *Settings
	err      error
	client   *Client
}

// NewWidget creates and returns an instance of Widget
func NewWidget(app *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, redrawChan, nil, settings.common),
		app:        app,
		settings:   settings,
		client: NewClient(
			settings.latitude,
			settings.longitude,
			settings.radius,
			settings.county,
			settings.country),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}
	// The last call should always be to the display function
	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	var title = defaultTitle
	if widget.CommonSettings().Title != "" {
		title = widget.CommonSettings().Title
	}
	now := time.Now()
	kriser, err := widget.client.getKrisinformation()
	if err != nil {
		handleError(widget, err)
	}

	var str string
	i := 0
	for k := range kriser {
		diff := now.Sub(kriser[k].Updated)
		if widget.settings.maxage != -1 {
			// Skip if message is too old
			if int(diff.Hours()) > widget.settings.maxage {
				//logger.Log(fmt.Sprintf("Article to old: (%s) Days: %d", kriser[k].HeadLine, int(diff.Hours())))
				continue
			}
		}
		i++
		if i > widget.settings.maxitems && widget.settings.maxitems != -1 {
			break
		}
		str += fmt.Sprintf("- %s\n", kriser[k].HeadLine)
	}
	return title, str, true
}

func (widget *Widget) display() {
	widget.Redraw(func() (string, string, bool) {
		return widget.content()
	})
}

func handleError(widget *Widget, err error) {
	widget.err = err
}
