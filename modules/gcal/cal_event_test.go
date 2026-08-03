package gcal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/api/calendar/v3"
)

func allDayEvent(startDate, endDate string) *CalEvent {
	event := &calendar.Event{
		Start: &calendar.EventDateTime{Date: startDate},
		End:   &calendar.EventDateTime{Date: endDate},
	}
	return NewCalEvent(event)
}

func Test_Past_AllDayEvent(t *testing.T) {
	yesterday := time.Now().AddDate(0, 0, -1)
	today := time.Now()
	tomorrow := time.Now().AddDate(0, 0, 1)

	dateOnly := func(d time.Time) string {
		return d.Format("2006-01-02")
	}

	testCases := []struct {
		name       string
		startDate  string
		endDate    string
		pastWanted bool
	}{
		{
			name:       "single-day all-day event yesterday is past",
			startDate:  dateOnly(yesterday),
			endDate:    dateOnly(today),
			pastWanted: true,
		},
		{
			name:       "single-day all-day event today is not past",
			startDate:  dateOnly(today),
			endDate:    dateOnly(tomorrow),
			pastWanted: false,
		},
		{
			name:       "single-day all-day event tomorrow is not past",
			startDate:  dateOnly(tomorrow),
			endDate:    dateOnly(tomorrow.AddDate(0, 0, 1)),
			pastWanted: false,
		},
		{
			name:       "multi-day all-day event ending yesterday is past",
			startDate:  dateOnly(yesterday.AddDate(0, 0, -2)),
			endDate:    dateOnly(yesterday),
			pastWanted: true,
		},
		{
			name:       "multi-day all-day event spanning today is not past",
			startDate:  dateOnly(yesterday),
			endDate:    dateOnly(tomorrow),
			pastWanted: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			event := allDayEvent(tt.startDate, tt.endDate)
			assert.Equal(t, tt.pastWanted, event.Past())
		})
	}
}
