package spotifyweb

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
	"github.com/zmb3/spotify"
)

var (
	auth           spotify.Authenticator
	tempClientChan = make(chan *spotify.Client)
	state          = "wtfSpotifyWebStateString"
	authURL        string
	callbackPort   string
	redirectURI    string
)

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
	view.TextWidget

	Info

	client      *spotify.Client
	clientChan  chan *spotify.Client
	playerState *spotify.PlayerState
	settings    *Settings
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
	}
	// use the token to get an authenticated client
	client := auth.NewClient(tok)
	_, err = fmt.Fprintf(w, "Login Completed!")
	if err != nil {
		return
	}
	tempClientChan <- &client
}

// NewWidget creates a new widget for WTF
func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	redirectURI = "http://localhost:" + settings.callbackPort + "/callback"

	auth = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadCurrentlyPlaying, spotify.ScopeUserReadPlaybackState, spotify.ScopeUserModifyPlaybackState)
	auth.SetAuthInfo(settings.clientID, settings.secretKey)
	authURL = auth.AuthURL(state)

	var client *spotify.Client
	var playerState *spotify.PlayerState

	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, pages, settings.Common),

		Info: Info{},

		client:      client,
		clientChan:  tempClientChan,
		playerState: playerState,
		settings:    settings,
	}

	http.HandleFunc("/callback", authHandler)
	go func() {
		err := http.ListenAndServe(":"+callbackPort, nil)
		if err != nil {
			return
		}
	}()

	go func() {
		// wait for auth to complete
		client = <-tempClientChan

		// use the client to make calls that require authorization
		_, err := client.CurrentUser()
		if err != nil {
			panic(err)
		}

		playerState, err = client.PlayerState()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

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
	utils.OpenFile(`"` + authURL + `"`)

	widget.settings.RefreshInterval = 5

	widget.initializeKeyboardControls()

	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)

	return &widget
}

func (w *Widget) refreshSpotifyInfos() error {
	if w.client == nil || w.playerState == nil {
		return errors.New("authentication failed! Please log in to Spotify by visiting the following page in your browser: " + authURL)
	}
	var err error
	w.playerState, err = w.client.PlayerState()
	if err != nil {
		return errors.New("extracting player state failed! Please refresh or restart WTF")
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
	w.Redraw(w.createOutput)
}

func (w *Widget) createOutput() (string, string, bool) {
	var output string

	err := w.refreshSpotifyInfos()
	if err != nil {
		output = err.Error()
	} else {
		output += utils.CenterText(fmt.Sprintf("[green]Now %v [white]\n", w.Info.Status), w.CommonSettings().Width)
		output += utils.CenterText(fmt.Sprintf("[green]Title:[white] %v\n", w.Info.Title), w.CommonSettings().Width)
		output += utils.CenterText(fmt.Sprintf("[green]Artist:[white] %v\n", w.Info.Artists), w.CommonSettings().Width)
		output += utils.CenterText(fmt.Sprintf("[green]Album:[white] %v\n", w.Info.Album), w.CommonSettings().Width)
		if w.playerState.ShuffleState {
			output += utils.CenterText("[green]Shuffle:[white] on\n", w.CommonSettings().Width)
		} else {
			output += utils.CenterText("[green]Shuffle:[white] off\n", w.CommonSettings().Width)
		}
	}
	return w.CommonSettings().Title, output, true
}
