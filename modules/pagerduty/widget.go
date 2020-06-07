package pagerduty

import (
	"fmt"
	"sort"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	settings *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, settings.common),

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

	if widget.settings.showIncidents {
		str += "[yellow] Incidents[white]"
		if len(incidents) > 0 {
			for _, incident := range incidents {
				str += fmt.Sprintf("\n [%s]%s[white]\n", widget.settings.common.Colors.Subheading, incident.Summary)
				str += fmt.Sprintf(" Status: %s\n", incident.Status)
				str += fmt.Sprintf(" Service: %s\n", incident.Service.Summary)
				str += fmt.Sprintf(" Escalation: %s\n", incident.EscalationPolicy.Summary)
			}
		} else {
			str += "\n No open incidents\n"
		}
	}

	tree := make(map[string][]pagerduty.OnCall)

	filter := make(map[string]bool)
	for _, item := range widget.settings.escalationFilter {
		filter[item.(string)] = true
	}

	for _, onCall := range onCalls {
		key := onCall.EscalationPolicy.Summary
		if len(widget.settings.escalationFilter) == 0 || filter[key] {
			tree[key] = append(tree[key], onCall)
		}
	}

	// We want to sort our escalation policies for predictability/ease of finding
	keys := make([]string, 0, len(tree))
	for k := range tree {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	if len(keys) > 0 {
		str += fmt.Sprintf("\n[%s] Schedules[white]\n", widget.settings.common.Colors.Subheading)

		// Print out policies, and escalation order of users
		for _, key := range keys {
			str += fmt.Sprintf(
				"\n [%s]%s\n",
				widget.settings.common.Colors.Subheading,
				key,
			)

			values := tree[key]
			sort.Sort(ByEscalationLevel(values))

			for _, item := range values {
				str += fmt.Sprintf(
					" [%s]%d - %s\n",
					widget.settings.common.Colors.Text,
					item.EscalationLevel,
					widget.userSummary(item),
				)
			}
		}
	}

	return str
}

func (widget *Widget) userSummary(item pagerduty.OnCall) string {
	summary := item.User.Summary

	if summary == widget.settings.myName {
		summary = fmt.Sprintf("[::b]%s", summary)
	}

	return summary
}
