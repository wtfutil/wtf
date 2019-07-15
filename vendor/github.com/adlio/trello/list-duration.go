package trello

import (
	"sort"
	"time"

	"github.com/pkg/errors"
)

// ListDuration represents the time a Card has been or was in list.
type ListDuration struct {
	ListID       string
	ListName     string
	Duration     time.Duration
	FirstEntered time.Time
	TimesInList  int
}

// AddDuration takes a duration and adds it to the ListDuration's Duration.
// Also increments TimesInList.
func (l *ListDuration) AddDuration(d time.Duration) {
	l.Duration = l.Duration + d
	l.TimesInList++
}

// GetListDurations analyses a Card's actions to figure out how long it was in each List.
// It returns a slice of the ListDurations, one Duration per list, or an error.
func (c *Card) GetListDurations() (durations []*ListDuration, err error) {

	var actions ActionCollection
	if len(c.Actions) == 0 {
		// Get all actions which affected the Card's List
		c.client.log("[trello] GetListDurations() called on card '%s' without any Card.Actions. Fetching fresh.", c.ID)
		actions, err = c.GetListChangeActions()
		if err != nil {
			err = errors.Wrap(err, "GetListChangeActions() call failed.")
			return
		}
	} else {
		actions = c.Actions.FilterToListChangeActions()
	}

	return actions.GetListDurations()
}

// GetListDurations returns a slice of ListDurations based on the receiver Actions.
func (actions ActionCollection) GetListDurations() (durations []*ListDuration, err error) {
	sort.Sort(actions)

	var prevTime time.Time
	var prevList *List

	durs := make(map[string]*ListDuration)
	for _, action := range actions {
		if action.DidChangeListForCard() {
			if prevList != nil {
				duration := action.Date.Sub(prevTime)
				_, durExists := durs[prevList.ID]
				if !durExists {
					durs[prevList.ID] = &ListDuration{ListID: prevList.ID, ListName: prevList.Name, Duration: duration, TimesInList: 1, FirstEntered: prevTime}
				} else {
					durs[prevList.ID].AddDuration(duration)
				}
			}
			prevList = ListAfterAction(action)
			prevTime = action.Date
		}
	}

	if prevList != nil {
		duration := time.Now().Sub(prevTime)
		_, durExists := durs[prevList.ID]
		if !durExists {
			durs[prevList.ID] = &ListDuration{ListID: prevList.ID, ListName: prevList.Name, Duration: duration, TimesInList: 1, FirstEntered: prevTime}
		} else {
			durs[prevList.ID].AddDuration(duration)
		}
	}

	durations = make([]*ListDuration, 0, len(durs))
	for _, ld := range durs {
		durations = append(durations, ld)
	}
	sort.Sort(ByFirstEntered(durations))

	return durations, nil
}
