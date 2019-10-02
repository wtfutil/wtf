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
		incidents, err2 = GetIncidents(widget.settings.apiKey)
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

	if len(incidents) > 0 {
		str += "[yellow]Incidents[white]\n"
		for _, incident := range incidents {
			str += fmt.Sprintf("[red]%s[white]\n", incident.Summary)
			str += fmt.Sprintf("Status: %s\n", incident.Status)
			str += fmt.Sprintf("Service: %s\n", incident.Service.Summary)
			str += fmt.Sprintf("Escalation: %s\n", incident.EscalationPolicy.Summary)
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
		str += "[red] Schedules[white]\n"
		// Print out policies, and escalation order of users
		for _, key := range keys {
			str += fmt.Sprintf("\n [green::b]%s\n", key)
			values := tree[key]
			sort.Sort(ByEscalationLevel(values))
			for _, item := range values {
				str += fmt.Sprintf(" [white]%d - %s\n", item.EscalationLevel, item.User.Summary)
			}
		}
	}

	return str
}
