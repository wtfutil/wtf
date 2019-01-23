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
	teams []OnCallTeam
}

// NewWidget creates a new widget
func NewWidget(app *tview.Application) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, "VictorOps - OnCall", "victorops", true),
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

	teams, err := Fetch()
	widget.View.SetTitle(widget.ContextualTitle(widget.Name))

	if err != nil {
		widget.View.SetWrap(true)
		widget.View.SetText(err.Error())
	} else {
		widget.teams = teams
	}

	widget.display()
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
	teamToDisplay := wtf.Config.UString("wtf.mods.victorops.team")
	var str string
	for _, team := range teams {
		if len(teamToDisplay) > 0 && teamToDisplay != team.Slug {
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
