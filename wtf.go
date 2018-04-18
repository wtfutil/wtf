package main

import (
	"flag"
	"os"
	"time"
	//"math/rand"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/bamboohr"
	"github.com/senorprogrammer/wtf/clocks"
	"github.com/senorprogrammer/wtf/gcal"
	"github.com/senorprogrammer/wtf/git"
	"github.com/senorprogrammer/wtf/github"
	"github.com/senorprogrammer/wtf/jira"
	"github.com/senorprogrammer/wtf/newrelic"
	"github.com/senorprogrammer/wtf/opsgenie"
	"github.com/senorprogrammer/wtf/security"
	"github.com/senorprogrammer/wtf/status"
	"github.com/senorprogrammer/wtf/textfile"
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
		false,
	)
}

// Grid stores all the widgets onscreen (like an HTML table)
func buildGrid(modules []wtf.TextViewer) *tview.Grid {
	grid := tview.NewGrid()
	grid.SetColumns(wtf.ToInts(Config.UList("wtf.grid.columns"))...)
	grid.SetRows(wtf.ToInts(Config.UList("wtf.grid.rows"))...)
	grid.SetBorder(false)

	for _, module := range modules {
		addToGrid(grid, module)
		go wtf.Schedule(module)
	}

	return grid
}

func keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyCtrlR:
		refreshAllModules()
	case tcell.KeyTab:
		FocusTracker.Next()
	case tcell.KeyBacktab:
		FocusTracker.Prev()
	case tcell.KeyEsc:
		FocusTracker.None()
	default:
		return event
	}

	return event
}

// redrawApp redraws the rendered views to screen on a defined interval (set in config.yml)
// Use this because each textView widget can have it's own update interval, and I don't want to
// manage drawing co-ordination amongst them all. If you need to have a
// widget redraw on it's own schedule, use the view's SetChangedFunc() and pass it `app`.
func redrawApp(app *tview.Application) {
	tick := time.NewTicker(time.Duration(Config.UInt("wtf.refreshInterval", 2)) * time.Second)
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

func refreshAllModules() {
	for _, module := range Widgets {
		go module.Refresh()
	}
}

var result = wtf.CreateConfigDir()

var Config *config.Config
var FocusTracker wtf.FocusTracker
var Widgets []wtf.TextViewer

/* -------------------- Main -------------------- */

func main() {
	// Optional argument to accept path to a config file
	configFile := flag.String("config", "~/.wtf/config.yml", "Path to config file")
	flag.Parse()

	Config = wtf.LoadConfigFile(*configFile)

	wtf.Config = Config

	bamboohr.Config = Config
	clocks.Config = Config
	gcal.Config = Config
	git.Config = Config
	github.Config = Config
	jira.Config = Config
	newrelic.Config = Config
	opsgenie.Config = Config
	security.Config = Config
	status.Config = Config
	textfile.Config = Config
	weather.Config = Config

	Widgets = []wtf.TextViewer{
		bamboohr.NewWidget(),
		clocks.NewWidget(),
		gcal.NewWidget(),
		git.NewWidget(),
		github.NewWidget(),
		jira.NewWidget(),
		newrelic.NewWidget(),
		opsgenie.NewWidget(),
		security.NewWidget(),
		status.NewWidget(),
		textfile.NewWidget(),
		weather.NewWidget(),
	}

	app := tview.NewApplication()
	app.SetInputCapture(keyboardIntercept)

	FocusTracker = wtf.FocusTracker{
		App:     app,
		Idx:     0,
		Widgets: Widgets,
	}

	// Loop in a routine to redraw the screen
	go redrawApp(app)

	grid := buildGrid(Widgets)
	if err := app.SetRoot(grid, true).Run(); err != nil {
		os.Exit(1)
	}
}
