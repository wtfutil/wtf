package pagerduty

import "github.com/PagerDuty/go-pagerduty"

type ByEscalationLevel []pagerduty.OnCall

func (s ByEscalationLevel) Len() int      { return len(s) }
func (s ByEscalationLevel) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s ByEscalationLevel) Less(i, j int) bool {
	return s[i].EscalationLevel < s[j].EscalationLevel
}
