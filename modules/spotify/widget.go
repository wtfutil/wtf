package spotify

import (
	"fmt"
	"strings"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/spotigopher/spotigopher"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

// A Widget represents a Spotify widget
type Widget struct {
	view.TextWidget

	client   spotigopher.SpotifyClient
	settings *Settings
	spotigopher.Info
}

// NewWidget creates a new instance of a widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, pages, settings.Common),

		Info:   spotigopher.Info{},
		client: spotigopher.NewClient(),

		settings: settings,
	}

	widget.settings.RefreshInterval = 5 * time.Second

	widget.initializeKeyboardControls()

	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)

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

func (w *Widget) createOutput() (string, string, bool) {
	var content string
	err := w.refreshSpotifyInfos()
	if err != nil {
		content = err.Error()
	} else {
		labelColor := w.settings.label
		textColor := w.settings.text

		artist := strings.Join(w.Artist, ", ")

		content = utils.CenterText(fmt.Sprintf("[%s]Now %v [%s]\n", labelColor, w.Status, textColor), w.CommonSettings().Width)
		content += utils.CenterText(fmt.Sprintf("[%s]Title:[%s] %v\n ", labelColor, textColor, w.Title), w.CommonSettings().Width)
		content += utils.CenterText(fmt.Sprintf("[%s]Artist:[%s] %v\n", labelColor, textColor, artist), w.CommonSettings().Width)
		content += utils.CenterText(fmt.Sprintf("[%s]%v:[%s] %v\n", labelColor, w.TrackNumber, textColor, w.Album), w.CommonSettings().Width)
	}
	return w.CommonSettings().Title, content, true
}
