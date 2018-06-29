/*
* This butt-ugly code is direct from Google itself
* https://developers.google.com/calendar/quickstart/go
*
* With some changes by me to improve things a bit.
 */

package gcal

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"time"

	"github.com/senorprogrammer/wtf/wtf"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

/* -------------------- Exported Functions -------------------- */

func Fetch() ([]*CalEvent, error) {
	ctx := context.Background()

	secretPath, _ := wtf.ExpandHomeDir(wtf.Config.UString("wtf.mods.gcal.secretFile"))

	b, err := ioutil.ReadFile(secretPath)
	if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		return nil, err
	}
	client := getClient(ctx, config)

	srv, err := calendar.New(client)
	if err != nil {
		return nil, err
	}

	calendarIds, err := getCalendarIdList(srv)

	// Get calendar events
	var events calendar.Events

	startTime := fromMidnight().Format(time.RFC3339)
	eventLimit := int64(wtf.Config.UInt("wtf.mods.gcal.eventCount", 10))

	for _, calendarId := range calendarIds {
		calendarEvents, err := srv.Events.List(calendarId).ShowDeleted(false).TimeMin(startTime).MaxResults(eventLimit).SingleEvents(true).OrderBy("startTime").Do()
		if err != nil {
			break
		}
		events.Items = append(events.Items, calendarEvents.Items...)
	}
	if err != nil {
		return nil, err
	}

	// Sort events
	timeDateChooser := func(event *calendar.Event) (time.Time, error) {
		if len(event.Start.Date) > 0 {
			return time.Parse("2006-01-02", event.Start.Date)
		} else {
			return time.Parse(time.RFC3339, event.Start.DateTime)
		}
	}

	sort.Slice(events.Items, func(i, j int) bool {
		dateA, _ := timeDateChooser(events.Items[i])
		dateB, _ := timeDateChooser(events.Items[j])
		return dateA.Before(dateB)
	})

	// Wrap the calendar events in our custom CalEvent
	calEvents := []*CalEvent{}
	for _, event := range events.Items {
		calEvents = append(calEvents, NewCalEvent(event))
	}

	return calEvents, err
}

/* -------------------- Unexported Functions -------------------- */

func fromMidnight() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}
	return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("calendar-go-quickstart.json")), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()

	json.NewEncoder(f).Encode(token)
}

func getCalendarIdList(srv *calendar.Service) ([]string, error) {
	// Return single calendar if settings specify we should
	if !wtf.Config.UBool("wtf.mods.gcal.multiCalendar", false) {
		id, err := srv.CalendarList.Get("primary").Do()
		if err != nil {
			return nil, err
		}
		return []string{id.Id}, nil
	}

	// Get all user calendars with at the least writing access
	var calendarIds []string
	var pageToken string
	for {
		calendarList, err := srv.CalendarList.List().ShowHidden(false).MinAccessRole("writer").PageToken(pageToken).Do()
		if err != nil {
			return nil, err
		}
		for _, calendarListItem := range calendarList.Items {
			calendarIds = append(calendarIds, calendarListItem.Id)
		}

		pageToken = calendarList.NextPageToken
		if pageToken == "" {
			break
		}
	}
	return calendarIds, nil
}
