package victorops

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

// HelpText to display to users
const HelpText = `
	Keyboard commands for VictorOps

	/: Show/hide this help window
	arrow down: Scroll down the list
	arrow up: Scroll up the list
`

// Widget contains text info
type Widget struct {
	wtf.TextWidget

	app      *tview.Application
	teams    []OnCallTeam
	settings *Settings
}

// NewWidget creates a new widget
func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, settings.common, true),
	}

	widget.View.SetScrollable(true)
	widget.View.SetRegions(true)

	return &widget
}

// Refresh gets latest content for the widget
func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	teams, err := Fetch(widget.settings.apiID, widget.settings.apiKey)

	if err != nil {
		widget.View.SetWrap(true)

		widget.app.QueueUpdateDraw(func() {
			widget.View.SetText(err.Error())
		})
	} else {
		widget.teams = teams
	}

	widget.app.QueueUpdateDraw(func() {
		widget.View.SetTitle(widget.ContextualTitle(widget.CommonSettings.Title))
		widget.display()
	})
}

func (widget *Widget) display() {
	if widget.teams == nil {
		return
	}

	widget.View.SetWrap(false)
	widget.View.Clear()
	widget.View.SetText(widget.contentFrom(widget.teams))
}

func (widget *Widget) contentFrom(teams []OnCallTeam) string {
	var str string

	for _, team := range teams {
		if len(widget.settings.team) > 0 && widget.settings.team != team.Slug {
			continue
		}

		str = fmt.Sprintf("%s[green]%s\n", str, team.Name)
		if len(team.OnCall) == 0 {
			str = fmt.Sprintf("%s[grey]no one\n", str)
		}
		for _, onCall := range team.OnCall {
			str = fmt.Sprintf("%s[white]%s - %s\n", str, onCall.Policy, onCall.Userlist)
		}

		str = fmt.Sprintf("%s\n", str)
	}

	if len(str) == 0 {
		str = "Could not find any teams to display"
	}
	return str
}
