/*
* This butt-ugly code is direct from Google itself
* https://developers.google.com/calendar/quickstart/go
*
* With some changes by me to improve things a bit.
 */

package gcal

import (
	"sort"
	"time"

	"github.com/senorprogrammer/wtf/wtf"
	"google.golang.org/api/calendar/v3"
)

/* -------------------- Exported Functions -------------------- */

func Fetch(oauthHandler OAuthUIHandler) ([]*CalEvent, error) {
	secretFile, _ := wtf.ExpandHomeDir(wtf.Config.UString("wtf.mods.gcal.secretFile"))
	tokenFile, _ := wtf.ExpandHomeDir(wtf.Config.UString("wtf.mods.gcal.tokenFile", "~/.gcal.token"))
	config := OAuthClientConfig{
		SecretFile: secretFile,
		TokenFile:  tokenFile,
		Scope:      calendar.CalendarReadonlyScope,
	}
	client, err := BuildOAuthClient(config, oauthHandler)
	if err != nil {
		return nil, err
	}

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
