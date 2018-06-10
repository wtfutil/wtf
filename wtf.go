package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/radovskyb/watcher"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/bamboohr"
	"github.com/senorprogrammer/wtf/bargraph"
	"github.com/senorprogrammer/wtf/cfg"
	"github.com/senorprogrammer/wtf/circleci"
	"github.com/senorprogrammer/wtf/clocks"
	"github.com/senorprogrammer/wtf/cmdrunner"
	"github.com/senorprogrammer/wtf/cryptoexchanges/bittrex"
	"github.com/senorprogrammer/wtf/cryptoexchanges/cryptolive"
	"github.com/senorprogrammer/wtf/gcal"
	"github.com/senorprogrammer/wtf/git"
	"github.com/senorprogrammer/wtf/github"
	"github.com/senorprogrammer/wtf/gspreadsheets"
	"github.com/senorprogrammer/wtf/help"
	"github.com/senorprogrammer/wtf/ipinfo"
	"github.com/senorprogrammer/wtf/ipapi"
	"github.com/senorprogrammer/wtf/jira"
	"github.com/senorprogrammer/wtf/newrelic"
	"github.com/senorprogrammer/wtf/opsgenie"
	"github.com/senorprogrammer/wtf/power"
	"github.com/senorprogrammer/wtf/prettyweather"
	"github.com/senorprogrammer/wtf/security"
	"github.com/senorprogrammer/wtf/status"
	"github.com/senorprogrammer/wtf/system"
	"github.com/senorprogrammer/wtf/textfile"
	"github.com/senorprogrammer/wtf/todo"
	"github.com/senorprogrammer/wtf/weather"
	"github.com/senorprogrammer/wtf/wtf"
)

/* -------------------- Functions -------------------- */

func addToGrid(grid *tview.Grid, widget wtf.Wtfable) {
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
func buildGrid(modules []wtf.Wtfable) *tview.Grid {
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
		refreshAllWidgets()
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

func refreshAllWidgets() {
	for _, widget := range Widgets {
		go widget.Refresh()
	}
}

func watchForConfigChanges(app *tview.Application, configFlag string, grid *tview.Grid, pages *tview.Pages) {
	watch := watcher.New()

	// notify write events.
	watch.FilterOps(watcher.Write)

	go func() {
		for {
			select {
			case <-watch.Event:
				loadConfig(configFlag)
				// Disable all widgets to stop scheduler goroutines and rmeove widgets from memory.
				disableAllWidgets()
				makeWidgets(app, pages)
				grid = buildGrid(Widgets)
				pages.AddPage("grid", grid, true, true)
			case err := <-watch.Error:
				log.Fatalln(err)
			case <-watch.Closed:
				return
			}
		}
	}()

	// Watch config file for changes.
	if err := watch.Add(configFlag); err != nil {
		log.Fatalln(err)
	}

	// Start the watching process - it'll check for changes every 100ms.
	if err := watch.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}

/* -------------------- Main -------------------- */

var Config *config.Config
var FocusTracker wtf.FocusTracker
var Widgets []wtf.Wtfable

var (
	commit  = "dev"
	date    = "dev"
	version = "dev"
)

func disableAllWidgets() {
	for _, widget := range Widgets {
		widget.Disable()
	}
}

func addWidget(app *tview.Application, pages *tview.Pages, widgetName string) {
	// Always in alphabetical order
	switch widgetName {
	case "bamboohr":
		Widgets = append(Widgets, bamboohr.NewWidget())
	case "bargraph":
		Widgets = append(Widgets, bargraph.NewWidget())
	case "bittrex":
		Widgets = append(Widgets, bittrex.NewWidget())
	case "circleci":
		Widgets = append(Widgets, circleci.NewWidget())
	case "clocks":
		Widgets = append(Widgets, clocks.NewWidget())
	case "cmdrunner":
		Widgets = append(Widgets, cmdrunner.NewWidget())
	case "cryptolive":
		Widgets = append(Widgets, cryptolive.NewWidget())
	case "gcal":
		Widgets = append(Widgets, gcal.NewWidget())
	case "git":
		Widgets = append(Widgets, git.NewWidget(app, pages))
	case "github":
		Widgets = append(Widgets, github.NewWidget(app, pages))
	case "gspreadsheets":
		Widgets = append(Widgets, gspreadsheets.NewWidget())
	case "ipinfo":
		Widgets = append(Widgets, ipinfo.NewWidget())
	case "ipapi":
		Widgets = append(Widgets, ipapi.NewWidget())
	case "jira":
		Widgets = append(Widgets, jira.NewWidget())
	case "newrelic":
		Widgets = append(Widgets, newrelic.NewWidget())
	case "opsgenie":
		Widgets = append(Widgets, opsgenie.NewWidget())
	case "power":
		Widgets = append(Widgets, power.NewWidget())
	case "prettyweather":
		Widgets = append(Widgets, prettyweather.NewWidget())
	case "security":
		Widgets = append(Widgets, security.NewWidget())
	case "status":
		Widgets = append(Widgets, status.NewWidget())
	case "system":
		Widgets = append(Widgets, system.NewWidget(date, version))
	case "textfile":
		Widgets = append(Widgets, textfile.NewWidget(app, pages))
	case "todo":
		Widgets = append(Widgets, todo.NewWidget(app, pages))
	case "weather":
		Widgets = append(Widgets, weather.NewWidget(app, pages))
	default:
	}
}

func makeWidgets(app *tview.Application, pages *tview.Pages) {
	Widgets = []wtf.Wtfable{}

	// Always in alphabetical order
	bamboohr.Config = Config
	bargraph.Config = Config
	bittrex.Config = Config
	circleci.Config = Config
	clocks.Config = Config
	cmdrunner.Config = Config
	cryptolive.Config = Config
	gcal.Config = Config
	git.Config = Config
	github.Config = Config
	gspreadsheets.Config = Config
	ipinfo.Config = Config
	ipapi.Config = Config
	jira.Config = Config
	newrelic.Config = Config
	opsgenie.Config = Config
	power.Config = Config
	prettyweather.Config = Config
	security.Config = Config
	status.Config = Config
	system.Config = Config
	textfile.Config = Config
	todo.Config = Config
	weather.Config = Config
	wtf.Config = Config

	mods, _ := Config.Map("wtf.mods")
	for mod := range mods {
		if enabled, _ := Config.Bool("wtf.mods." + mod + ".enabled"); enabled {
			addWidget(app, pages, mod)
		}

	}

	FocusTracker = wtf.FocusTracker{
		App:     app,
		Idx:     -1,
		Widgets: Widgets,
	}
}

func loadConfig(configFlag string) {
	Config = cfg.LoadConfigFile(configFlag)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cmdFlags := wtf.NewCommandFlags()
	cmdFlags.Parse(version)

	if cmdFlags.HasModule() {
		help.DisplayModuleInfo(cmdFlags.Module)
	}

	/* -------------------- end flag parsing and handling -------------------- */

	// Responsible for creating the configuration directory and default
	// configuration file if they don't already exist
	cfg.CreateConfigDir()
	cfg.WriteConfigFile()

	loadConfig(cmdFlags.Config)
	os.Setenv("TERM", Config.UString("wtf.term", os.Getenv("TERM")))

	app := tview.NewApplication()
	pages := tview.NewPages()

	makeWidgets(app, pages)

	grid := buildGrid(Widgets)
	pages.AddPage("grid", grid, true, true)
	app.SetInputCapture(keyboardIntercept)

	grid.SetBackgroundColor(wtf.ColorFor(Config.UString("wtf.colors.background", "black")))

	// Loop in a routine to redraw the screen
	go redrawApp(app)
	go watchForConfigChanges(app, cmdFlags.Config, grid, pages)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		os.Exit(1)
	}

	wtf.Log("running!")
}
