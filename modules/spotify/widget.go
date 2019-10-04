package spotify

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/sticreations/spotigopher/spotigopher"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

// A Widget represents a Spotify widget
type Widget struct {
	view.KeyboardWidget
	view.TextWidget

	client   spotigopher.SpotifyClient
	settings *Settings
	spotigopher.Info
}

// NewWidget creates a new instance of a widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget: view.NewKeyboardWidget(app, pages, settings.common),
		TextWidget:     view.NewTextWidget(app, settings.common),

		Info:   spotigopher.Info{},
		client: spotigopher.NewClient(),

		settings: settings,
	}

	widget.settings.common.RefreshInterval = 5

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)

	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

func (w *Widget) refreshSpotifyInfos() error {
	info, err := w.client.GetInfo()
	w.Info = info
	return err
}

func (w *Widget) Refresh() {
	w.Redraw(w.createOutput)
}

func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
}

func (w *Widget) createOutput() (string, string, bool) {
	var content string
	err := w.refreshSpotifyInfos()
	if err != nil {
		content = err.Error()
	} else {
		labelColor := w.settings.colors.label
		textColor := w.settings.colors.text

		content = utils.CenterText(fmt.Sprintf("[%s]Now %v [%s]\n", labelColor, w.Info.Status, textColor), w.CommonSettings().Width)
		content += utils.CenterText(fmt.Sprintf("[%s]Title:[%s] %v\n ", labelColor, textColor, w.Info.Title), w.CommonSettings().Width)
		content += utils.CenterText(fmt.Sprintf("[%s]Artist:[%s] %v\n", labelColor, textColor, w.Info.Artist), w.CommonSettings().Width)
		content += utils.CenterText(fmt.Sprintf("[%s]%v:[%s] %v\n", labelColor, w.Info.TrackNumber, textColor, w.Info.Album), w.CommonSettings().Width)
	}
	return w.CommonSettings().Title, content, true
}
