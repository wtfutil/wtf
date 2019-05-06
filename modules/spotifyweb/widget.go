package spotifyweb

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/logger"
	"github.com/wtfutil/wtf/wtf"
	"github.com/zmb3/spotify"
)

// HelpText contains the help text for the Spotify Web API widget.
const HelpText = `
Keyboard commands for Spotify Web:

	Before any of these commands are used, you should authenticate using the
	URL provided by the widget.

	The widget should automatically open a browser window for you, otherwise
	you should check out the logs for the URL.

	/: Show/hide this help window
	h: Switch to previous song in Spotify queue
	l: Switch to next song in Spotify queue
	s: Toggle shuffle

	[space]: Pause/play current song

	esc: Unselect the Spotify Web module
`

// Info is the struct that contains all the information the Spotify player displays to the user
type Info struct {
	Artists     string
	Title       string
	Album       string
	TrackNumber int
	Status      string
}

// Widget is the struct used by all WTF widgets to transfer to the main widget controller
type Widget struct {
	wtf.HelpfulWidget
	wtf.KeyboardWidget
	wtf.TextWidget

	Info

	app         *tview.Application
	client      *spotify.Client
	clientChan  chan *spotify.Client
	playerState *spotify.PlayerState
	settings    *Settings
}

var (
	auth           spotify.Authenticator
	tempClientChan = make(chan *spotify.Client)
	state          = "wtfSpotifyWebStateString"
	authURL        string
	callbackPort   string
	redirectURI    string
)

func authHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log("[SpotifyWeb] Got an authentication hit!")
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		logger.Log(err.Error())
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		logger.Log(fmt.Sprintf("State mismatch: %s != %s\n", st, state))
	}
	// use the token to get an authenticated client
	client := auth.NewClient(tok)
	fmt.Fprintf(w, "Login Completed!")
	tempClientChan <- &client
}

// NewWidget creates a new widget for WTF
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	redirectURI = "http://localhost:" + settings.callbackPort + "/callback"

	auth = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadCurrentlyPlaying, spotify.ScopeUserReadPlaybackState, spotify.ScopeUserModifyPlaybackState)
	auth.SetAuthInfo(settings.clientID, settings.secretKey)
	authURL = auth.AuthURL(state)

	var client *spotify.Client
	var playerState *spotify.PlayerState

	widget := Widget{
		HelpfulWidget:  wtf.NewHelpfulWidget(app, pages, HelpText),
		KeyboardWidget: wtf.NewKeyboardWidget(),
		TextWidget:     wtf.NewTextWidget(app, settings.common, true),

		Info: Info{},

		app:         app,
		client:      client,
		clientChan:  tempClientChan,
		playerState: playerState,
		settings:    settings,
	}

	http.HandleFunc("/callback", authHandler)
	go http.ListenAndServe(":"+callbackPort, nil)

	go func() {
		// wait for auth to complete
		logger.Log("[SpotifyWeb] Waiting for authentication... URL: " + authURL)
		client = <-tempClientChan

		// use the client to make calls that require authorization
		_, err := client.CurrentUser()
		if err != nil {
			panic(err)
		}

		playerState, err = client.PlayerState()
		if err != nil {
			panic(err)
		}
		logger.Log("[SpotifyWeb] Authentication complete.")
		widget.client = client
		widget.playerState = playerState
		widget.Refresh()
	}()

	// While I wish I could find the reason this doesn't work, I can't.
	//
	// Normally, this should open the URL to the browser, however it opens the Explorer window in Windows.
	// This mostly likely has to do with the fact that the URL includes some very special characters that no terminal likes.
	// The only solution would be to include quotes in the command, which is why I do here, but it doesn't work.
	//
	// If inconvenient, I'll remove this option and save the URL in a file or some other method.
	wtf.OpenFile(`"` + authURL + `"`)

	widget.settings.common.RefreshInterval = 5

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)
	widget.View.SetTitle("[green]Spotify Web[white]")

	widget.HelpfulWidget.SetView(widget.View)

	return &widget
}

func (w *Widget) refreshSpotifyInfos() error {
	if w.client == nil || w.playerState == nil {
		return errors.New("Authentication failed! Please log in to Spotify by visiting the following page in your browser: " + authURL)
	}
	var err error
	w.playerState, err = w.client.PlayerState()
	if err != nil {
		return errors.New("Extracting player state failed! Please refresh or restart WTF")
	}
	w.Info.Album = fmt.Sprint(w.playerState.CurrentlyPlaying.Item.Album.Name)
	artists := ""
	for _, artist := range w.playerState.CurrentlyPlaying.Item.Artists {
		artists += artist.Name + ", "
	}
	artists = artists[:len(artists)-2]
	w.Info.Artists = artists
	w.Info.Title = fmt.Sprint(w.playerState.CurrentlyPlaying.Item.Name)
	w.Info.TrackNumber = w.playerState.CurrentlyPlaying.Item.TrackNumber
	if w.playerState.CurrentlyPlaying.Playing {
		w.Info.Status = "Playing"
	} else {
		w.Info.Status = "Paused"
	}
	return nil
}

// Refresh refreshes the current view of the widget
func (w *Widget) Refresh() {
	w.app.QueueUpdateDraw(func() {
		w.render()
	})
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

func (w *Widget) createOutput() string {
	output := wtf.CenterText(fmt.Sprintf("[green]Now %v [white]\n", w.Info.Status), w.Width())
	output += wtf.CenterText(fmt.Sprintf("[green]Title:[white] %v\n", w.Info.Title), w.Width())
	output += wtf.CenterText(fmt.Sprintf("[green]Artist:[white] %v\n", w.Info.Artists), w.Width())
	output += wtf.CenterText(fmt.Sprintf("[green]Album:[white] %v\n", w.Info.Album), w.Width())
	if w.playerState.ShuffleState {
		output += wtf.CenterText(fmt.Sprintf("[green]Shuffle:[white] on\n"), w.Width())
	} else {
		output += wtf.CenterText(fmt.Sprintf("[green]Shuffle:[white] off\n"), w.Width())
	}
	return output
}
