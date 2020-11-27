package pagerduty

import (
	"fmt"
	"sort"
	"time"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

const (
	onCallTimeAPILayout     = "2006-01-02T15:04:05Z"
	onCallTimeDisplayLayout = "Jan 2, 2006"
)

type Widget struct {
	view.TextWidget

	settings *Settings
}

// NewWidget creates and returns an instance of PagerDuty widget
func NewWidget(tviewApp *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	var onCalls []pagerduty.OnCall
	var incidents []pagerduty.Incident

	var err1 error
	var err2 error

	if widget.settings.showIncidents {
		teamIDs := utils.ToStrs(widget.settings.teamIDs)
		userIDs := utils.ToStrs(widget.settings.userIDs)
		incidents, err2 = GetIncidents(widget.settings.apiKey, teamIDs, userIDs)
	}

	if widget.settings.showSchedules {
		scheduleIDs := utils.ToStrs(widget.settings.scheduleIDs)
		onCalls, err1 = GetOnCalls(widget.settings.apiKey, scheduleIDs)
	}

	var content string
	wrap := false
	if err1 != nil || err2 != nil {
		wrap = true
		if err1 != nil {
			content += err1.Error()
		}
		if err2 != nil {
			content += err2.Error()
		}
	} else {
		widget.View.SetWrap(false)
		content = widget.contentFrom(onCalls, incidents)
	}

	widget.Redraw(func() (string, string, bool) { return widget.CommonSettings().Title, content, wrap })
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(onCalls []pagerduty.OnCall, incidents []pagerduty.Incident) string {
	var str string

	// Incidents

	if widget.settings.showIncidents {
		str += fmt.Sprintf("[%s] Incidents[white]\n", widget.settings.Colors.Subheading)

		if len(incidents) > 0 {
			for _, incident := range incidents {
				str += fmt.Sprintf("\n [%s]%s[white]\n", widget.settings.Colors.Label, tview.Escape(incident.Summary))
				str += fmt.Sprintf("     Status: %s\n", incident.Status)
				str += fmt.Sprintf("    Service: %s\n", incident.Service.Summary)
				str += fmt.Sprintf(" Escalation: %s\n", incident.EscalationPolicy.Summary)
				str += fmt.Sprintf("       Link: %s\n", incident.HTMLURL)
			}
		} else {
			str += "\n No open incidents\n"
		}

		str += "\n"
	}

	onCallTree := make(map[string][]pagerduty.OnCall)

	filter := make(map[string]bool)
	for _, item := range widget.settings.escalationFilter {
		filter[item.(string)] = true
	}

	// OnCalls

	for _, onCall := range onCalls {
		summary := onCall.EscalationPolicy.Summary
		if len(widget.settings.escalationFilter) == 0 || filter[summary] {
			onCallTree[summary] = append(onCallTree[summary], onCall)
		}
	}

	// We want to sort our escalation policies for predictability/ease of finding
	keys := make([]string, 0, len(onCallTree))
	for k := range onCallTree {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	if len(keys) > 0 {
		str += fmt.Sprintf("[%s] Schedules[white]\n", widget.settings.Colors.Subheading)

		// Print out policies, and escalation order of users
		for _, key := range keys {
			str += fmt.Sprintf(
				"\n [%s]%s\n",
				widget.settings.Colors.Label,
				key,
			)

			values := onCallTree[key]
			sort.Sort(ByEscalationLevel(values))

			for _, onCall := range values {
				str += fmt.Sprintf(
					" [%s]%d - %s\n",
					widget.settings.Colors.Text,
					onCall.EscalationLevel,
					widget.userSummary(onCall),
				)

				onCallEnd := widget.onCallEndSummary(onCall)
				if onCallEnd != "" {
					str += fmt.Sprintf(
						"     %s\n",
						onCallEnd,
					)
				}
			}
		}
	}

	return str
}

// onCallEndSummary may or may not return the date that the specified onCall schedule ends
func (widget *Widget) onCallEndSummary(onCall pagerduty.OnCall) string {
	if !widget.settings.showOnCallEnd {
		return ""
	}

	if onCall.End == "" {
		return ""
	}

	end, err := time.Parse(onCallTimeAPILayout, onCall.End)
	if err != nil {
		return ""
	}

	return end.Format(onCallTimeDisplayLayout)
}

// userSummary returns the name of the person assigned to the specified onCall schedule
func (widget *Widget) userSummary(onCall pagerduty.OnCall) string {
	summary := onCall.User.Summary

	if summary == widget.settings.myName {
		summary = fmt.Sprintf("[::b]%s", summary)
	}

	return summary
}
