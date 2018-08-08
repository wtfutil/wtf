package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/pkg/profile"
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
	"github.com/senorprogrammer/wtf/gerrit"
	"github.com/senorprogrammer/wtf/git"
	"github.com/senorprogrammer/wtf/github"
	"github.com/senorprogrammer/wtf/gitlab"
	"github.com/senorprogrammer/wtf/gitter"
	"github.com/senorprogrammer/wtf/gspreadsheets"
	"github.com/senorprogrammer/wtf/hackernews"
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
	"github.com/senorprogrammer/wtf/todoist"
	"github.com/senorprogrammer/wtf/travisci"
	"github.com/senorprogrammer/wtf/trello"
	"github.com/senorprogrammer/wtf/twitter"
	"github.com/senorprogrammer/wtf/weatherservices/prettyweather"
	"github.com/senorprogrammer/wtf/weatherservices/weather"
	"github.com/senorprogrammer/wtf/wtf"
	"github.com/senorprogrammer/wtf/zendesk"
)

var focusTracker wtf.FocusTracker
var widgets []wtf.Wtfable

// Config parses the config.yml file and makes available the settings within
var Config *config.Config

var (
	commit  = "dev"
	date    = "dev"
	version = "dev"
)

/* -------------------- Functions -------------------- */

func disableAllWidgets() {
	for _, widget := range widgets {
		widget.Disable()
	}
}

func initializeFocusTracker(app *tview.Application) {
	focusTracker = wtf.FocusTracker{
		App:     app,
		Idx:     -1,
		Widgets: widgets,
	}

	focusTracker.AssignHotKeys()
}

func keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyCtrlR:
		refreshAllWidgets()
	case tcell.KeyTab:
		focusTracker.Next()
	case tcell.KeyBacktab:
		focusTracker.Prev()
	case tcell.KeyEsc:
		focusTracker.None()
	}

	if focusTracker.FocusOn(string(event.Rune())) {
		return nil
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
	for _, widget := range widgets {
		go widget.Refresh()
	}
}

func setTerm() {
	err := os.Setenv("TERM", Config.UString("wtf.term", os.Getenv("TERM")))
	if err != nil {
		return
	}
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
				widgets = nil
				makeWidgets(app, pages)
				initializeFocusTracker(app)
				display := wtf.NewDisplay(widgets)
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
		widgets = append(widgets, bamboohr.NewWidget())
	case "bargraph":
		widgets = append(widgets, bargraph.NewWidget())
	case "bittrex":
		widgets = append(widgets, bittrex.NewWidget())
	case "blockfolio":
		widgets = append(widgets, blockfolio.NewWidget())
	case "circleci":
		widgets = append(widgets, circleci.NewWidget())
	case "clocks":
		widgets = append(widgets, clocks.NewWidget())
	case "cmdrunner":
		widgets = append(widgets, cmdrunner.NewWidget())
	case "cryptolive":
		widgets = append(widgets, cryptolive.NewWidget())
	case "gcal":
		widgets = append(widgets, gcal.NewWidget())
	case "gerrit":
		widgets = append(widgets, gerrit.NewWidget(app, pages))
	case "git":
		widgets = append(widgets, git.NewWidget(app, pages))
	case "github":
		widgets = append(widgets, github.NewWidget(app, pages))
	case "gitlab":
		widgets = append(widgets, gitlab.NewWidget(app, pages))
	case "gitter":
		widgets = append(widgets, gitter.NewWidget(app, pages))
	case "gspreadsheets":
		widgets = append(widgets, gspreadsheets.NewWidget())
	case "hackernews":
		widgets = append(widgets, hackernews.NewWidget(app, pages))
	case "ipapi":
		widgets = append(widgets, ipapi.NewWidget())
	case "ipinfo":
		widgets = append(widgets, ipinfo.NewWidget())
	case "jenkins":
		widgets = append(widgets, jenkins.NewWidget())
	case "jira":
		widgets = append(widgets, jira.NewWidget(app, pages))
	case "logger":
		widgets = append(widgets, logger.NewWidget())
	case "newrelic":
		widgets = append(widgets, newrelic.NewWidget())
	case "opsgenie":
		widgets = append(widgets, opsgenie.NewWidget())
	case "power":
		widgets = append(widgets, power.NewWidget())
	case "prettyweather":
		widgets = append(widgets, prettyweather.NewWidget())
	case "security":
		widgets = append(widgets, security.NewWidget())
	case "status":
		widgets = append(widgets, status.NewWidget())
	case "system":
		widgets = append(widgets, system.NewWidget(date, version))
	case "textfile":
		widgets = append(widgets, textfile.NewWidget(app, pages))
	case "todo":
		widgets = append(widgets, todo.NewWidget(app, pages))
	case "todoist":
		widgets = append(widgets, todoist.NewWidget(app, pages))
	case "travisci":
		widgets = append(widgets, travisci.NewWidget(app, pages))
	case "trello":
		widgets = append(widgets, trello.NewWidget())
	case "twitter":
		widgets = append(widgets, twitter.NewWidget())
	case "weather":
		widgets = append(widgets, weather.NewWidget(app, pages))
	case "zendesk":
		widgets = append(widgets, zendesk.NewWidget())
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

	if flags.Profile {
		defer profile.Start(profile.MemProfile).Stop()
	}

	setTerm()

	app := tview.NewApplication()
	pages := tview.NewPages()

	makeWidgets(app, pages)
	initializeFocusTracker(app)

	display := wtf.NewDisplay(widgets)
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
