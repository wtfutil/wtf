package twitch

import (
	"errors"
	"fmt"

	"github.com/nicklaw5/helix"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.ScrollableWidget

	settings   *Settings
	err        error
	twitch     *Twitch
	topStreams []*Stream
}

type Stream struct {
	Streamer    string
	ViewerCount int
	Language    string
	GameID      string
	Title       string
}

func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, pages, settings.Common),
		settings:         settings,
		twitch:           NewClient(settings.clientId),
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()

	return widget
}

func (widget *Widget) Refresh() {
	response, err := widget.twitch.TopStreams(&helix.StreamsParams{
		First:      widget.settings.numberOfResults,
		GameIDs:    widget.settings.gameIds,
		Language:   widget.settings.languages,
		Type:       widget.settings.streamType,
		UserIDs:    widget.settings.gameIds,
		UserLogins: widget.settings.userLogins,
	})

	if err != nil {
		handleError(widget, err)
	} else if response.ErrorMessage != "" {
		handleError(widget, errors.New(response.ErrorMessage))
	} else {
		streams := makeStreams(response)
		widget.topStreams = streams
		widget.err = nil
		if len(streams) <= widget.settings.numberOfResults {
			widget.SetItemCount(len(widget.topStreams))
		} else {
			widget.topStreams = streams[:widget.settings.numberOfResults]
			widget.SetItemCount(len(widget.topStreams))
		}
	}
	widget.Render()
}

func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

func makeStreams(response *helix.StreamsResponse) []*Stream {
	streams := make([]*Stream, 0)
	for _, b := range response.Data.Streams {
		streams = append(streams, &Stream{
			b.UserName,
			b.ViewerCount,
			b.Language,
			b.GameID,
			b.Title,
		})
	}
	return streams
}

func handleError(widget *Widget, err error) {
	widget.err = err
	widget.topStreams = nil
	widget.SetItemCount(0)
}

func (widget *Widget) content() (string, string, bool) {
	var title = "Top Streams"
	if widget.CommonSettings().Title != "" {
		title = widget.CommonSettings().Title
	}
	if widget.err != nil {
		return title, widget.err.Error(), true
	}
	if len(widget.topStreams) == 0 {
		return title, "No data", false
	}
	var str string

	for idx, stream := range widget.topStreams {
		row := fmt.Sprintf(
			"[%s]%2d. [red]%s [white]%s",
			widget.RowColor(idx),
			idx+1,
			utils.PrettyNumber(float64(stream.ViewerCount)),
			stream.Streamer,
		)
		str += utils.HighlightableHelper(widget.View, row, idx, len(stream.Streamer))
	}

	return title, str, false
}

// Opens stream in the browser
func (widget *Widget) openTwitch() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.topStreams != nil && sel < len(widget.topStreams) {
		stream := widget.topStreams[sel]
		fullLink := "https://twitch.com/" + stream.Streamer
		utils.OpenFile(fullLink)
	}
}
