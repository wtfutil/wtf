package spotigopher

import (
	"errors"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

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
	if runtime.GOOS == "darwin" {
		sendDarwinAction("playpause")
	} else {
		sendAction("org.mpris.MediaPlayer2.Player.PlayPause")
	}
}

/*
Next sends a Next Command on the DBus
*/
func (s *SpotifyClient) Next() {
	if runtime.GOOS == "darwin" {
		sendDarwinAction("next")
	} else {
		sendAction("org.mpris.MediaPlayer2.Player.Next")
	}
}

/*
Previous sends a Previous Command on the DBus
*/
func (s *SpotifyClient) Previous() {
	if runtime.GOOS == "darwin" {
		sendDarwinAction("previous")
	} else {
		sendAction("org.mpris.MediaPlayer2.Player.Previous")
	}
}

/*
Stop sends a Stop Command on the DBus
*/
func (s *SpotifyClient) Stop() {
	if runtime.GOOS == "darwin" {
		sendDarwinAction("stop")
	} else {
		sendAction("org.mpris.MediaPlayer2.Player.Stop")
	}
}

/*
GetInfo returns all Spotify related Information, when Spotify is running
*/
func (s *SpotifyClient) GetInfo() (Info, error) {
	info := Info{}

	if runtime.GOOS == "linux" {
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
	} else if runtime.GOOS == "darwin" {
		info.TrackID = getDarwinInfo("trackid")
		info.Artist = strings.Split(getDarwinInfo("artist"), ",")
		info.Album = getDarwinInfo("album")
		info.Title = getDarwinInfo("track")
		tmpTrackNumber, err := strconv.ParseInt(getDarwinInfo("tracknumber"), 0, 32)
		if err != nil {
			log.Fatalf("Could not get track number: %s", err)
		}
		info.TrackNumber = int32(tmpTrackNumber)
		info.URL = getDarwinInfo("url")
		info.ArtworkURL = getDarwinInfo("artworkurl")
		info.Status = getDarwinInfo("status")
	}
	return info, nil
}

func getDarwinInfo(infoType string) string {
	args := []string{`-etell application "Spotify" to name of current track as string`}
	switch infoType {
	case "trackid":
		args = []string{`-etell application "Spotify" to id of current track as string`}
	case "artist":
		args = []string{`-etell application "Spotify" to artist of current track as string`}
	case "title":
		args = []string{`-etell application "Spotify" to name of current track as string`}
	case "album":
		args = []string{`-etell application "Spotify" to album of current track as string`}
	case "tracknumber":
		args = []string{`-etell application "Spotify" to track number of current track as string`}
	case "url":
		args = []string{`-etell application "Spotify" to spotify url of current track as string`}
	case "artworkurl":
		args = []string{`-etell application "Spotify" to artwork url of current track as string`}
	case "status":
		args = []string{`-etell application "Spotify" to player state as string`}
	default:
		args = []string{`-etell application "Spotify" to name of current track as string`}
	}
	info, err := exec.Command("osascript", args...).Output()
	if err != nil {
		log.Fatalf("Could not get info: %s", infoType)
	}
	return strings.Trim(string(info), "\n")
}

func sendDarwinAction(action string) {
	args := []string{`-etell application "Spotify" to pause`}
	switch action {
	case "play":
		args = []string{`-etell application "Spotify" to play`}
	case "pause":
		args = []string{`-etell application "Spotify" to pause`}
	case "stop":
		args = []string{`-etell application "Spotify" to pause`}
	case "playpause":
		args = []string{`-etell application "Spotify" to playpause`}
	case "next":
		args = []string{`-etell application "Spotify" to next track`}
	case "previous":
		args = []string{`-etell application "Spotify" to previous track`}
	default:
		args = []string{`-etell application "Spotify" to pause`}
	}
	err := exec.Command("osascript", args...).Run()
	if err != nil {
		log.Fatalf("Could not completed action: %s", action)
	}
}

func sendAction(method string) error {
	sdbus := getSpotifyBus()
	call := sdbus.Call(method, 0)
	return call.Err
}
