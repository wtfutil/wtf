package main

import (
	"time"

	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/bamboohr"
	"github.com/senorprogrammer/wtf/gcal"
	"github.com/senorprogrammer/wtf/git"
	"github.com/senorprogrammer/wtf/github"
	"github.com/senorprogrammer/wtf/jira"
	"github.com/senorprogrammer/wtf/newrelic"
	"github.com/senorprogrammer/wtf/opsgenie"
	"github.com/senorprogrammer/wtf/security"
	"github.com/senorprogrammer/wtf/status"
	"github.com/senorprogrammer/wtf/weather"
	"github.com/senorprogrammer/wtf/wtf"
)

func refresher(stat *status.Widget, app *tview.Application) {
	tick := time.NewTicker(time.Duration(Config.UInt("wtf.refreshInterval", 1)) * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-tick.C:
			app.Draw()
		case <-quit:
			tick.Stop()
			return
		}
	}
}

func addToApp(grid *tview.Grid, widget wtf.TextViewer) {
	if widget.Disabled() {
		return
	}

	grid.AddItem(widget.TextView(), 0, 0, 2, 1, 0, 0, false)
}

var result = wtf.CreateConfigDir()
var Config = wtf.LoadConfigFile()

/* -------------------- Main -------------------- */

func main() {
	wtf.Config = Config

	bamboohr.Config = Config
	bamboo := bamboohr.NewWidget()
	go wtf.Schedule(bamboo)

	gcal.Config = Config
	cal := gcal.NewWidget()
	go wtf.Schedule(cal)

	git.Config = Config
	git := git.NewWidget()
	go wtf.Schedule(git)

	github.Config = Config
	github := github.NewWidget()
	go wtf.Schedule(github)

	jira.Config = Config
	jira := jira.NewWidget()
	go wtf.Schedule(jira)

	newrelic.Config = Config
	newrelic := newrelic.NewWidget()
	go wtf.Schedule(newrelic)

	opsgenie.Config = Config
	opsgenie := opsgenie.NewWidget()
	go wtf.Schedule(opsgenie)

	security.Config = Config
	sec := security.NewWidget()
	go wtf.Schedule(sec)

	status.Config = Config
	stat := status.NewWidget()
	go wtf.Schedule(stat)

	weather.Config = Config
	weather := weather.NewWidget()
	go wtf.Schedule(weather)

	grid := tview.NewGrid()
	grid.SetRows(10, 10, 10, 10, 10, 4) // How _high_ the row is, in terminal rows
	grid.SetColumns(37, 37, 37, 37, 37) // How _wide_ the column is, in terminal columns
	grid.SetBorder(false)

	addToApp(grid, bamboo)
	addToApp(grid, cal)
	addToApp(grid, git)
	addToApp(grid, github)
	addToApp(grid, newrelic)
	addToApp(grid, weather)
	addToApp(grid, sec)
	addToApp(grid, opsgenie)
	addToApp(grid, jira)
	addToApp(grid, stat)

	//grid.AddItem(bamboo.View, 0, 0, 2, 1, 0, 0, false)
	//grid.AddItem(cal.View, 2, 1, 4, 1, 0, 0, false)
	//grid.AddItem(git.View, 0, 2, 2, 3, 0, 0, false)
	//grid.AddItem(github.View, 2, 2, 2, 3, 0, 0, false)
	//grid.AddItem(newrelic.View, 4, 2, 1, 3, 0, 0, false)
	//grid.AddItem(weather.View, 0, 1, 1, 1, 0, 0, false)
	//grid.AddItem(sec.View, 5, 0, 1, 1, 0, 0, false)
	//grid.AddItem(opsgenie.View, 2, 0, 2, 1, 0, 0, false)
	//grid.AddItem(jira.View, 1, 1, 1, 1, 0, 0, false)
	//grid.AddItem(stat.View, 5, 2, 3, 3, 0, 0, false)

	app := tview.NewApplication()

	// Loop in a routine to redraw the screen
	go refresher(stat, app)

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
