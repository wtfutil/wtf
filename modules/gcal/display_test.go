package gcal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/cfg"
	"google.golang.org/api/calendar/v3"
)

func Test_display_content(t *testing.T) {
	startTime := &calendar.EventDateTime{DateTime: "1986-04-19T01:00:00.00Z"}
	endTime := &calendar.EventDateTime{DateTime: "1986-04-19T02:00:00.00Z"}
	event := &calendar.Event{Summary: "Foo", Start: startTime, End: endTime}

	testCases := []struct {
		descriptionWanted string
		events            []*CalEvent
		name              string
		settings          *Settings
	}{
		{
			name:              "Event content without any events",
			settings:          &Settings{Common: &cfg.Common{}},
			events:            nil,
			descriptionWanted: "No calendar events",
		},
		{
			name:              "Event content with a single event, without end times displayed",
			settings:          &Settings{Common: &cfg.Common{}, showEndTime: false},
			events:            []*CalEvent{NewCalEvent(event)},
			descriptionWanted: "[]Saturday, Apr 19\n  []01:00 []Foo[white]\n   \n",
		},
		{
			name:              "Event content with a single event without showEndTime explicitly set in settings",
			settings:          &Settings{Common: &cfg.Common{}},
			events:            []*CalEvent{NewCalEvent(event)},
			descriptionWanted: "[]Saturday, Apr 19\n  []01:00 []Foo[white]\n   \n",
		},
		{
			name:              "Event content with a single event with end times displayed",
			settings:          &Settings{Common: &cfg.Common{}, showEndTime: true},
			events:            []*CalEvent{NewCalEvent(event)},
			descriptionWanted: "[]Saturday, Apr 19\n  []01:00-02:00 []Foo[white]\n   \n",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			w := &Widget{calEvents: tt.events, settings: tt.settings, err: nil}
			_, description, err := w.content()

			assert.Equal(t, false, err, tt.name)
			assert.Equal(t, tt.descriptionWanted, description, tt.name)
		})
	}

}
