package spotify

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
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

type Widget struct {
	wtf.HelpfulWidget
	wtf.TextWidget

	settings *Settings
	spotigopher.Info
	spotigopher.SpotifyClient
}

func NewWidget(app *tview.Application, refreshChan chan<- string, pages *tview.Pages, settings *Settings) *Widget {
	spotifyClient := spotigopher.NewClient()
	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget(refreshChan, settings.common, true),

		Info:          spotigopher.Info{},
		SpotifyClient: spotifyClient,
		settings:      settings,
	}

	widget.settings.common.RefreshInterval = 5

	widget.HelpfulWidget.SetView(widget.View)
	widget.View.SetInputCapture(widget.captureInput)
	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)
	widget.View.SetTitle(fmt.Sprint("[green]Spotify[white]"))
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
	w.View.Clear()
	if err != nil {
		w.TextWidget.View.SetText(err.Error())
	} else {
		w.TextWidget.View.SetText(w.createOutput())
	}
}

func (w *Widget) captureInput(event *tcell.EventKey) *tcell.EventKey {
	switch (string)(event.Rune()) {
	case "h":
		w.SpotifyClient.Previous()
		time.Sleep(time.Second * 1)
		w.Refresh()
		return nil
	case "l":
		w.SpotifyClient.Next()
		time.Sleep(time.Second * 1)
		w.Refresh()
		return nil
	case " ":
		w.SpotifyClient.PlayPause()
		time.Sleep(time.Second * 1)
		w.Refresh()
		return nil
	}
	return nil
}

func (w *Widget) createOutput() string {
	output := wtf.CenterText(fmt.Sprintf("[green]Now %v [white]\n", w.Info.Status), w.Width())
	output += wtf.CenterText(fmt.Sprintf("[green]Title:[white] %v\n ", w.Info.Title), w.Width())
	output += wtf.CenterText(fmt.Sprintf("[green]Artist:[white] %v\n", w.Info.Artist), w.Width())
	output += wtf.CenterText(fmt.Sprintf("[green]%v:[white] %v\n", w.Info.TrackNumber, w.Info.Album), w.Width())
	return output
}
