package spotigopher

import (
	"errors"
	"log"

	"github.com/godbus/dbus"
)

type SpotifyClient struct {
}

/*
Info holds all available Information we get over the dbus
*/
type Info struct {
	Artist      []string
	Title       string
	Album       string
	Rating      string
	TrackID     string
	TrackNumber int32
	URL         string
	ArtworkURL  string
	Status      string
}

var sdbus dbus.Object

/*
NewClient returns an Instance of Spotify Cleint
*/
func NewClient() SpotifyClient {

	client := SpotifyClient{}

	return client
}

func getSpotifyBus() dbus.BusObject {
	con, err := dbus.SessionBus()
	if err != nil {
		log.Fatalf("Could not get SessionBus: %v", err)
	}
	//TODO: Test what happens when Spotify is Closed!
	spotifyBus := con.Object("org.mpris.MediaPlayer2.spotify", "/org/mpris/MediaPlayer2")
	return spotifyBus
}

/*
PlayPause sends a PlayPause Command on the DBus
*/
func (s *SpotifyClient) PlayPause() {
	sendAction("org.mpris.MediaPlayer2.Player.PlayPause")
}

/*
Next sends a Next Command on the DBus
*/
func (s *SpotifyClient) Next() {
	sendAction("org.mpris.MediaPlayer2.Player.Next")
}

/*
Previous sends a Previous Command on the DBus
*/
func (s *SpotifyClient) Previous() {
	sendAction("org.mpris.MediaPlayer2.Player.Previous")
}

/*
Stop sends a Stop Command on the DBus
*/
func (s *SpotifyClient) Stop() {
	sendAction("org.mpris.MediaPlayer2.Player.Stop")

}

/*
GetInfo returns all Spotify related Information, when Spotify is running
*/
func (s *SpotifyClient) GetInfo() (Info, error) {
	info := Info{}

	spotifyBus := getSpotifyBus()
	props, err := spotifyBus.GetProperty("org.mpris.MediaPlayer2.Player.Metadata")
	if err != nil {
		return Info{}, errors.New("Could not get any Info from Spotify. Are you sure Spotify is running?")
	}
	songData := props.Value().(map[string]dbus.Variant)
	info.TrackID = songData["mpris:trackid"].Value().(string)
	info.Artist = songData["xesam:artist"].Value().([]string)
	info.Title = songData["xesam:title"].Value().(string)
	info.Album = songData["xesam:album"].Value().(string)
	info.TrackNumber = songData["xesam:trackNumber"].Value().(int32)
	info.URL = songData["xesam:url"].Value().(string)
	info.ArtworkURL = songData["mpris:artUrl"].Value().(string)

	status, err := spotifyBus.GetProperty("org.mpris.MediaPlayer2.Player.PlaybackStatus")
	if err != nil {
		log.Fatalf("Could not get Playback Status : %v", err)

	}
	info.Status = status.Value().(string)
	return info, nil
}

func sendAction(method string) error {
	sdbus := getSpotifyBus()
	call := sdbus.Call(method, 0)
	return call.Err
}
