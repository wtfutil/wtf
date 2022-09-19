package twitch

import (
	"errors"
	"fmt"
	"os/exec"

	helix "github.com/nicklaw5/helix/v2"
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

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	clientOpts := &ClientOpts{
		ClientID:         settings.clientId,
		ClientSecret:     settings.clientSecret,
		AppAccessToken:   settings.appAccessToken,
		UserAccessToken:  settings.userAccessToken,
		UserRefreshToken: settings.userRefreshToken,
		RedirectURI:      settings.redirectURI,
		Streams:          settings.streams,
		UserID:           settings.userId,
	}

	twitchClient, err := NewClient(clientOpts)
	if err != nil {
		fmt.Println(err)
	}

	widget := &Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, redrawChan, pages, settings.Common),
		settings:         settings,
		twitch:           twitchClient,
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()

	return widget
}

func (widget *Widget) Refresh() {
	var err error
	var response *helix.StreamsResponse
	// Refresh the auth token on each refresh to be sure we aren't using an expired one.
	if err = widget.twitch.RefreshOAuthToken(); err != nil {
		handleError(widget, err)
	}

	if widget.twitch.Streams == "followed" {
		response, err = widget.twitch.FollowedStreams(&helix.FollowedStreamsParams{
			UserID: widget.twitch.UserID,
		})
	} else if widget.twitch.Streams == "top" {
		response, err = widget.twitch.TopStreams(&helix.StreamsParams{
			First:      widget.settings.numberOfResults,
			GameIDs:    widget.settings.gameIds,
			Language:   widget.settings.languages,
			Type:       widget.settings.streamType,
			UserIDs:    widget.settings.userIds,
			UserLogins: widget.settings.userLogins,
		})
	}

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
	streams := make([]*Stream, len(response.Data.Streams))
	for i, b := range response.Data.Streams {
		streams[i] = &Stream{
			b.UserName,
			b.ViewerCount,
			b.Language,
			b.GameID,
			b.Title,
		}
	}
	return streams
}

func handleError(widget *Widget, err error) {
	widget.err = err
	widget.topStreams = nil
	widget.SetItemCount(0)
}

func (widget *Widget) content() (string, string, bool) {
	var title = "Twitch Streams"
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

	locPrinter, _ := widget.settings.LocalizedPrinter()

	for idx, stream := range widget.topStreams {
		row := fmt.Sprintf(
			"[%s]%2d. [red]%s [white]%s - %s",
			widget.RowColor(idx),
			idx+1,
			utils.PrettyNumber(locPrinter, float64(stream.ViewerCount)),
			stream.Streamer,
			stream.Title,
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

func (widget *Widget) openStreamlink() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.topStreams != nil && sel < len(widget.topStreams) {
		stream := widget.topStreams[sel]
		fullLink := "https://twitch.tv/" + stream.Streamer
		cmd := exec.Command("streamlink", fullLink, "best")
		err := cmd.Start()
		if err != nil {
			handleError(widget, err)
		}
	}
}
