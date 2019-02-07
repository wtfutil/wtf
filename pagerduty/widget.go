package pagerduty

import (
	"fmt"
	"sort"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget
}

func NewWidget(app *tview.Application) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, "PagerDuty", "pagerduty", false),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	onCalls, err := GetOnCalls()

	widget.View.SetTitle(widget.ContextualTitle(fmt.Sprintf("%s", widget.Name)))
	widget.View.Clear()

	var content string
	if err != nil {
		widget.View.SetWrap(true)
		content = err.Error()
	} else {
		widget.View.SetWrap(false)
		content = widget.contentFrom(onCalls)
	}

	widget.View.SetText(content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(onCalls []pagerduty.OnCall) string {
	var str string

	tree := make(map[string][]pagerduty.OnCall)

	filtering := wtf.Config.UList("wtf.mods.pagerduty.escalationFilter")
	filter := make(map[string]bool)
	for _, item := range filtering {
		filter[item.(string)] = true
	}

	for _, onCall := range onCalls {
		key := onCall.EscalationPolicy.Summary
		if len(filtering) == 0 || filter[key] {
			tree[key] = append(tree[key], onCall)
		}
	}

	// We want to sort our escalation policies for predictability/ease of finding
	keys := make([]string, 0, len(tree))
	for k := range tree {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	// Print out policies, and escalation order of users
	for _, key := range keys {
		str = str + fmt.Sprintf("[red]%s\n", key)
		values := tree[key]
		sort.Sort(ByEscalationLevel(values))
		for _, item := range values {
			str = str + fmt.Sprintf("[white]%d - %s\n", item.EscalationLevel, item.User.Summary)
		}
	}

	return str
}
