package main

import (
	"flag"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
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
		false, // has focus
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

// FIXME: Not a fan of how this function has to reach outside itself, grab
// Modules, and then operate on them. Should be able to pass that in instead
func keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	// Ctrl-R: force-refreshes every widget
	if event.Key() == tcell.KeyCtrlR {
		for _, module := range Widgets {
			go module.Refresh()
		}
	} else if event.Key() == tcell.KeyTab {
	}

	return event
}

func refresher(app *tview.Application) {
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

var Config *config.Config
var Widgets []wtf.TextViewer

/* -------------------- Main -------------------- */

func main() {
	// Optional argument to accept path to a config file
	configFile := flag.String("config", "~/.wtf/config.yml", "Path to config file")
	flag.Parse()

	Config = wtf.LoadConfigFile(*configFile)

	wtf.Config = Config

	bamboohr.Config = Config
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

	// Loop in a routine to redraw the screen
	go refresher(app)

	grid := buildGrid(Widgets)
	if err := app.SetRoot(grid, true).Run(); err != nil {
		os.Exit(1)
	}
}
