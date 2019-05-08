package spotify

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/sticreations/spotigopher/spotigopher"
	"github.com/wtfutil/wtf/wtf"
)

const HelpText = `
	To control Spotify use:
		[Spacebar] for Play & Pause
		[h] for Previous Song
		[l] for Next Song
`

// A Widget represents a Spotify widget
type Widget struct {
	wtf.HelpfulWidget
	wtf.KeyboardWidget
	wtf.TextWidget

	settings *Settings
	spotigopher.Info
	spotigopher.SpotifyClient
}

// NewWidget creates a new instance of a widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	spotifyClient := spotigopher.NewClient()
	widget := Widget{
		HelpfulWidget:  wtf.NewHelpfulWidget(app, pages, HelpText),
		KeyboardWidget: wtf.NewKeyboardWidget(),
		TextWidget:     wtf.NewTextWidget(app, settings.common, true),

		Info:          spotigopher.Info{},
		SpotifyClient: spotifyClient,

		settings: settings,
	}

	widget.settings.common.RefreshInterval = 5

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)
	widget.View.SetTitle(fmt.Sprint("[green]Spotify[white]"))

	widget.HelpfulWidget.SetView(widget.View)

	return &widget
}

func (w *Widget) refreshSpotifyInfos() error {
	info, err := w.SpotifyClient.GetInfo()
	w.Info = info
	return err
}

func (w *Widget) Refresh() {
	w.render()
}

func (w *Widget) render() {
	err := w.refreshSpotifyInfos()
	var content string
	if err != nil {
		content = err.Error()
	} else {
		content = w.createOutput()
	}
	w.Redraw(w.CommonSettings.Title, content, true)
}

func (w *Widget) createOutput() string {
	output := wtf.CenterText(fmt.Sprintf("[green]Now %v [white]\n", w.Info.Status), w.Width())
	output += wtf.CenterText(fmt.Sprintf("[green]Title:[white] %v\n ", w.Info.Title), w.Width())
	output += wtf.CenterText(fmt.Sprintf("[green]Artist:[white] %v\n", w.Info.Artist), w.Width())
	output += wtf.CenterText(fmt.Sprintf("[green]%v:[white] %v\n", w.Info.TrackNumber, w.Info.Album), w.Width())
	return output
}
