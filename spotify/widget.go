package spotify

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
	"github.com/sticreations/spotigopher/spotigopher"
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
	spotigopher.SpotifyClient
	spotigopher.Info
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	spotifyClient := spotigopher.NewClient()
	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget(app, "Spotify", "spotify", true),
		SpotifyClient: spotifyClient,
		Info:          spotigopher.Info{},
	}
	widget.HelpfulWidget.SetView(widget.View)
	widget.TextWidget.RefreshInt = 5
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
