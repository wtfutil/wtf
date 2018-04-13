package main

import (
	"os"
	"time"

	"github.com/gdamore/tcell"
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

func addToGrid(grid *tview.Grid, widget wtf.TextViewer) {
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

// Grid stores all the widgets onscreen (like an HTML table)
func buildGrid() *tview.Grid {
	grid := tview.NewGrid()
	grid.SetColumns(wtf.ToInts(Config.UList("wtf.grid.columns"))...)
	grid.SetRows(wtf.ToInts(Config.UList("wtf.grid.rows"))...)
	grid.SetBorder(false)

	return grid
}

func keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	// Ctrl-R: force-refreshes every widget
	if event.Key() == tcell.KeyCtrlR {
		for _, module := range Modules {
			go module.Refresh()
		}
	}

	return event
}

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

var result = wtf.CreateConfigDir()
var Config = wtf.LoadConfigFile()

var Modules []wtf.TextViewer

/* -------------------- Main -------------------- */

func main() {
	wtf.Config = Config

	// TODO: Really need to generalize all of these. This don't scale
	bamboohr.Config = Config
	bamboo := bamboohr.NewWidget()

	gcal.Config = Config
	cal := gcal.NewWidget()

	git.Config = Config
	git := git.NewWidget()

	github.Config = Config
	github := github.NewWidget()

	jira.Config = Config
	jira := jira.NewWidget()

	newrelic.Config = Config
	newrelic := newrelic.NewWidget()

	opsgenie.Config = Config
	opsgenie := opsgenie.NewWidget()

	security.Config = Config
	sec := security.NewWidget()

	status.Config = Config
	stat := status.NewWidget()

	weather.Config = Config
	weather := weather.NewWidget()

	Modules = []wtf.TextViewer{
		bamboo,
		cal,
		git,
		github,
		jira,
		newrelic,
		opsgenie,
		sec,
		stat,
		weather,
	}

	grid := buildGrid()

	for _, module := range Modules {
		addToGrid(grid, module)
		go wtf.Schedule(module)
	}

	app := tview.NewApplication()
	app.SetInputCapture(keyboardIntercept)

	// Loop in a routine to redraw the screen
	go refresher(stat, app)

	if err := app.SetRoot(grid, true).Run(); err != nil {
		os.Exit(1)
	}
}
