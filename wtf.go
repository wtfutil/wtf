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
	"github.com/senorprogrammer/wtf/cryptoexchanges/blockfolio"
	"github.com/senorprogrammer/wtf/cryptoexchanges/cryptolive"
	"github.com/senorprogrammer/wtf/flags"
	"github.com/senorprogrammer/wtf/gcal"
	"github.com/senorprogrammer/wtf/git"
	"github.com/senorprogrammer/wtf/github"
	"github.com/senorprogrammer/wtf/gitlab"
	"github.com/senorprogrammer/wtf/gspreadsheets"
	"github.com/senorprogrammer/wtf/ipaddresses/ipapi"
	"github.com/senorprogrammer/wtf/ipaddresses/ipinfo"
	"github.com/senorprogrammer/wtf/jenkins"
	"github.com/senorprogrammer/wtf/jira"
	"github.com/senorprogrammer/wtf/logger"
	"github.com/senorprogrammer/wtf/newrelic"
	"github.com/senorprogrammer/wtf/opsgenie"
	"github.com/senorprogrammer/wtf/power"
	"github.com/senorprogrammer/wtf/security"
	"github.com/senorprogrammer/wtf/status"
	"github.com/senorprogrammer/wtf/system"
	"github.com/senorprogrammer/wtf/textfile"
	"github.com/senorprogrammer/wtf/todo"
	"github.com/senorprogrammer/wtf/trello"
	"github.com/senorprogrammer/wtf/weatherservices/prettyweather"
	"github.com/senorprogrammer/wtf/weatherservices/weather"
	"github.com/senorprogrammer/wtf/wtf"
)

var Config *config.Config
var FocusTracker wtf.FocusTracker
var Widgets []wtf.Wtfable

var (
	commit  = "dev"
	date    = "dev"
	version = "dev"
)

/* -------------------- Functions -------------------- */

func disableAllWidgets() {
	for _, widget := range Widgets {
		widget.Disable()
	}
}

func initializeFocusTracker(app *tview.Application) {
	FocusTracker = wtf.FocusTracker{
		App:     app,
		Idx:     -1,
		Widgets: Widgets,
	}
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

func loadConfigFile(filePath string) {
	Config = cfg.LoadConfigFile(filePath)
	wtf.Config = Config
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

func setTerm() {
	os.Setenv("TERM", Config.UString("wtf.term", os.Getenv("TERM")))
}

func watchForConfigChanges(app *tview.Application, configFilePath string, grid *tview.Grid, pages *tview.Pages) {
	watch := watcher.New()

	// notify write events.
	watch.FilterOps(watcher.Write)

	go func() {
		for {
			select {
			case <-watch.Event:
				loadConfigFile(configFilePath)
				// Disable all widgets to stop scheduler goroutines and rmeove widgets from memory.
				disableAllWidgets()
				Widgets = nil
				makeWidgets(app, pages)
				initializeFocusTracker(app)
				display := wtf.NewDisplay(Widgets)
				pages.AddPage("grid", display.Grid, true, true)
			case err := <-watch.Error:
				log.Fatalln(err)
			case <-watch.Closed:
				return
			}
		}
	}()

	// Watch config file for changes.
	if err := watch.Add(configFilePath); err != nil {
		log.Fatalln(err)
	}

	// Start the watching process - it'll check for changes every 100ms.
	if err := watch.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
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
	case "blockfolio":
		Widgets = append(Widgets, blockfolio.NewWidget(app, pages))
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
	case "gitlab":
		Widgets = append(Widgets, gitlab.NewWidget(app, pages))
	case "gspreadsheets":
		Widgets = append(Widgets, gspreadsheets.NewWidget())
	case "ipapi":
		Widgets = append(Widgets, ipapi.NewWidget())
	case "ipinfo":
		Widgets = append(Widgets, ipinfo.NewWidget())
	case "jenkins":
		Widgets = append(Widgets, jenkins.NewWidget())
	case "jira":
		Widgets = append(Widgets, jira.NewWidget())
	case "logger":
		Widgets = append(Widgets, logger.NewWidget())
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
	case "trello":
		Widgets = append(Widgets, trello.NewWidget())
	case "weather":
		Widgets = append(Widgets, weather.NewWidget(app, pages))
	default:
	}
}

func makeWidgets(app *tview.Application, pages *tview.Pages) {
	mods, _ := Config.Map("wtf.mods")

	for mod := range mods {
		if enabled := Config.UBool("wtf.mods."+mod+".enabled", false); enabled {
			addWidget(app, pages, mod)
		}

	}
}

/* -------------------- Main -------------------- */

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flags := flags.NewFlags()
	flags.Parse()
	flags.Display(version)

	cfg.MigrateOldConfig()
	cfg.CreateConfigDir()
	cfg.CreateConfigFile()
	loadConfigFile(flags.ConfigFilePath())

	setTerm()

	app := tview.NewApplication()
	pages := tview.NewPages()

	makeWidgets(app, pages)
	initializeFocusTracker(app)

	display := wtf.NewDisplay(Widgets)
	pages.AddPage("grid", display.Grid, true, true)
	app.SetInputCapture(keyboardIntercept)

	// Loop in a routine to redraw the screen
	go redrawApp(app)
	go watchForConfigChanges(app, flags.Config, display.Grid, pages)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
