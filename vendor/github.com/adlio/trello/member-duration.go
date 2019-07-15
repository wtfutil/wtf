package trello

import (
	"sort"
	"time"

	"github.com/pkg/errors"
)

// MemberDuration is used to track the periods of time which a user (member) is attached to a card.
type MemberDuration struct {
	MemberID   string
	MemberName string
	FirstAdded time.Time
	Duration   time.Duration
	active     bool
	lastAdded  time.Time
}

// ByLongestDuration is a slice of *MemberDuration
type ByLongestDuration []*MemberDuration

// Len returns the length of the ByLongestDuration slice.
func (d ByLongestDuration) Len() int { return len(d) }

// Less takes two indexes i and j and returns true exactly if the Duration
// at i is larger than the Duration at j.
func (d ByLongestDuration) Less(i, j int) bool { return d[i].Duration > d[j].Duration }
func (d ByLongestDuration) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

func (d *MemberDuration) addAsOf(t time.Time) {
	d.active = true
	if d.FirstAdded.IsZero() {
		d.FirstAdded = t
	}
	d.startTimerAsOf(t)
}

func (d *MemberDuration) startTimerAsOf(t time.Time) {
	if d.active {
		d.lastAdded = t
	}
}

func (d *MemberDuration) removeAsOf(t time.Time) {
	d.stopTimerAsOf(t)
	d.active = false
	d.lastAdded = time.Time{}
}

func (d *MemberDuration) stopTimerAsOf(t time.Time) {
	if d.active {
		d.Duration = d.Duration + t.Sub(d.lastAdded)
	}
}

// GetMemberDurations returns a slice containing all durations of a card.
func (c *Card) GetMemberDurations() (durations []*MemberDuration, err error) {
	var actions ActionCollection
	if len(c.Actions) == 0 {
		c.client.log("[trello] GetMemberDurations() called on card '%s' without any Card.Actions. Fetching fresh.", c.ID)
		actions, err = c.GetMembershipChangeActions()
		if err != nil {
			err = errors.Wrap(err, "GetMembershipChangeActions() call failed.")
			return
		}
	} else {
		actions = c.Actions.FilterToCardMembershipChangeActions()
	}

	return actions.GetMemberDurations()
}

// GetMemberDurations is similar to GetListDurations. It returns a slice of MemberDuration objects,
// which describes the length of time each member was attached to this card. Durations are
// calculated such that being added to a card starts a timer for that member, and being removed
// starts it again (so that if a person is added and removed multiple times, the duration
// captures only the times which they were attached). Archiving the card also stops the timer.
func (actions ActionCollection) GetMemberDurations() (durations []*MemberDuration, err error) {
	sort.Sort(actions)
	durs := make(map[string]*MemberDuration)
	for _, action := range actions {
		if action.DidChangeCardMembership() {
			_, durExists := durs[action.Member.ID]
			if !durExists {
				switch action.Type {
				case "addMemberToCard":
					durs[action.Member.ID] = &MemberDuration{MemberID: action.Member.ID, MemberName: action.Member.FullName}
					durs[action.Member.ID].addAsOf(action.Date)
				case "removeMemberFromCard":
					// Surprisingly, this is possible. If a card was copied, and members were preserved, those
					// members exist on the card without a corresponding addMemberToCard action.
					t, _ := IDToTime(action.Data.Card.ID)
					durs[action.Member.ID] = &MemberDuration{MemberID: action.Member.ID, MemberName: action.Member.FullName, lastAdded: t}
					durs[action.Member.ID].removeAsOf(action.Date)
				}
			} else {
				switch action.Type {
				case "addMemberToCard":
					durs[action.Member.ID].addAsOf(action.Date)
				case "removeMemberFromCard":
					durs[action.Member.ID].removeAsOf(action.Date)
				}
			}
		} else if action.DidArchiveCard() {
			for id := range durs {
				durs[id].stopTimerAsOf(action.Date)
			}
		} else if action.DidUnarchiveCard() {
			for id := range durs {
				durs[id].startTimerAsOf(action.Date)
			}
		}
	}

	durations = make([]*MemberDuration, 0, len(durs))
	for _, md := range durs {
		durations = append(durations, md)
	}
	// sort.Sort(ByLongestDuration(durations))
	return durations, nil
}
