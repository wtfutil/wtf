package trello

import (
	"sort"
	"time"

	"github.com/pkg/errors"
)

type ListDuration struct {
	ListID       string
	ListName     string
	Duration     time.Duration
	FirstEntered time.Time
	TimesInList  int
}

func (l *ListDuration) AddDuration(d time.Duration) {
	l.Duration = l.Duration + d
	l.TimesInList++
}

// Analytzes a Cards actions to figure out how long it was in each List
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
