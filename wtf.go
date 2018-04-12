package main

import (
	"os"
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

	grid.AddItem(
		widget.TextView(),
		widget.Top(),
		widget.Left(),
		widget.Height(),
		widget.Width(),
		0,
		0,
		false, // has focus
	)
}

var result = wtf.CreateConfigDir()
var Config = wtf.LoadConfigFile()

/* -------------------- Main -------------------- */

func main() {
	wtf.Config = Config

	// Grid stores all the widgets onscreen (like an HTML table)
	grid := tview.NewGrid()
	grid.SetColumns(wtf.ToInts(Config.UList("wtf.grid.columns"))...)
	grid.SetRows(wtf.ToInts(Config.UList("wtf.grid.rows"))...)
	grid.SetBorder(false)

	// TODO: Really need to generalize all of these. This don't scale
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

	app := tview.NewApplication()

	// Loop in a routine to redraw the screen
	go refresher(stat, app)

	if err := app.SetRoot(grid, true).Run(); err != nil {
		os.Exit(1)
	}
}
