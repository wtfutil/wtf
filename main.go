package main

// Generators
// To generate the skeleton for a new TextWidget use 'WTF_WIDGET_NAME=MySuperAwesomeWidget go generate -run=text
//go:generate -command text go run generator/textwidget.go
//go:generate text

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
	"github.com/wtfutil/wtf/bargraph"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/flags"
	"github.com/wtfutil/wtf/logger"
	"github.com/wtfutil/wtf/modules/bamboohr"
	"github.com/wtfutil/wtf/modules/circleci"
	"github.com/wtfutil/wtf/modules/clocks"
	"github.com/wtfutil/wtf/modules/cmdrunner"
	"github.com/wtfutil/wtf/modules/cryptoexchanges/bittrex"
	"github.com/wtfutil/wtf/modules/cryptoexchanges/blockfolio"
	"github.com/wtfutil/wtf/modules/cryptoexchanges/cryptolive"
	"github.com/wtfutil/wtf/modules/datadog"
	"github.com/wtfutil/wtf/modules/gcal"
	"github.com/wtfutil/wtf/modules/gerrit"
	"github.com/wtfutil/wtf/modules/git"
	"github.com/wtfutil/wtf/modules/github"
	"github.com/wtfutil/wtf/modules/gitlab"
	"github.com/wtfutil/wtf/modules/gitter"
	"github.com/wtfutil/wtf/modules/gspreadsheets"
	"github.com/wtfutil/wtf/modules/hackernews"
	"github.com/wtfutil/wtf/modules/ipaddresses/ipapi"
	"github.com/wtfutil/wtf/modules/ipaddresses/ipinfo"
	"github.com/wtfutil/wtf/modules/jenkins"
	"github.com/wtfutil/wtf/modules/jira"
	"github.com/wtfutil/wtf/modules/mercurial"
	"github.com/wtfutil/wtf/modules/nbascore"
	"github.com/wtfutil/wtf/modules/newrelic"
	"github.com/wtfutil/wtf/modules/opsgenie"
	"github.com/wtfutil/wtf/modules/pagerduty"
	"github.com/wtfutil/wtf/modules/power"
	"github.com/wtfutil/wtf/modules/resourceusage"
	"github.com/wtfutil/wtf/modules/rollbar"
	"github.com/wtfutil/wtf/modules/security"
	"github.com/wtfutil/wtf/modules/spotify"
	"github.com/wtfutil/wtf/modules/spotifyweb"
	"github.com/wtfutil/wtf/modules/status"
	"github.com/wtfutil/wtf/modules/system"
	"github.com/wtfutil/wtf/modules/textfile"
	"github.com/wtfutil/wtf/modules/todo"
	"github.com/wtfutil/wtf/modules/todoist"
	"github.com/wtfutil/wtf/modules/travisci"
	"github.com/wtfutil/wtf/modules/trello"
	"github.com/wtfutil/wtf/modules/twitter"
	"github.com/wtfutil/wtf/modules/unknown"
	"github.com/wtfutil/wtf/modules/victorops"
	"github.com/wtfutil/wtf/modules/weatherservices/prettyweather"
	"github.com/wtfutil/wtf/modules/weatherservices/weather"
	"github.com/wtfutil/wtf/modules/zendesk"
	"github.com/wtfutil/wtf/wtf"
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
	absPath, _ := wtf.ExpandHomeDir(configFilePath)

	// notify write events.
	watch.FilterOps(watcher.Write)

	go func() {
		for {
			select {
			case <-watch.Event:
				loadConfigFile(absPath)
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
	if err := watch.Add(absPath); err != nil {
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
		widgets = append(widgets, bamboohr.NewWidget(app))
	case "bargraph":
		widgets = append(widgets, bargraph.NewWidget(app))
	case "bittrex":
		widgets = append(widgets, bittrex.NewWidget(app))
	case "blockfolio":
		widgets = append(widgets, blockfolio.NewWidget(app))
	case "circleci":
		widgets = append(widgets, circleci.NewWidget(app))
	case "clocks":
		widgets = append(widgets, clocks.NewWidget(app))
	case "cmdrunner":
		widgets = append(widgets, cmdrunner.NewWidget(app))
	case "resourceusage":
		widgets = append(widgets, resourceusage.NewWidget(app))
	case "cryptolive":
		widgets = append(widgets, cryptolive.NewWidget(app))
	case "datadog":
		widgets = append(widgets, datadog.NewWidget(app))
	case "gcal":
		widgets = append(widgets, gcal.NewWidget(app))
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
		widgets = append(widgets, gspreadsheets.NewWidget(app))
	case "hackernews":
		widgets = append(widgets, hackernews.NewWidget(app, pages))
	case "ipapi":
		widgets = append(widgets, ipapi.NewWidget(app))
	case "ipinfo":
		widgets = append(widgets, ipinfo.NewWidget(app))
	case "jenkins":
		widgets = append(widgets, jenkins.NewWidget(app, pages))
	case "jira":
		widgets = append(widgets, jira.NewWidget(app, pages))
	case "logger":
		widgets = append(widgets, logger.NewWidget(app))
	case "mercurial":
		widgets = append(widgets, mercurial.NewWidget(app, pages))
	case "nbascore":
		widgets = append(widgets, nbascore.NewWidget(app, pages))
	case "newrelic":
		widgets = append(widgets, newrelic.NewWidget(app))
	case "opsgenie":
		widgets = append(widgets, opsgenie.NewWidget(app))
	case "pagerduty":
		widgets = append(widgets, pagerduty.NewWidget(app))
	case "power":
		widgets = append(widgets, power.NewWidget(app))
	case "prettyweather":
		widgets = append(widgets, prettyweather.NewWidget(app))
	case "security":
		widgets = append(widgets, security.NewWidget(app))
	case "status":
		widgets = append(widgets, status.NewWidget(app))
	case "system":
		widgets = append(widgets, system.NewWidget(app, date, version))
	case "spotify":
		widgets = append(widgets, spotify.NewWidget(app, pages))
	case "spotifyweb":
		widgets = append(widgets, spotifyweb.NewWidget(app, pages))
	case "textfile":
		widgets = append(widgets, textfile.NewWidget(app, pages))
	case "todo":
		widgets = append(widgets, todo.NewWidget(app, pages))
	case "todoist":
		widgets = append(widgets, todoist.NewWidget(app, pages))
	case "travisci":
		widgets = append(widgets, travisci.NewWidget(app, pages))
	case "rollbar":
		widgets = append(widgets, rollbar.NewWidget(app, pages))
	case "trello":
		widgets = append(widgets, trello.NewWidget(app))
	case "twitter":
		widgets = append(widgets, twitter.NewWidget(app, pages))
	case "victorops":
		widgets = append(widgets, victorops.NewWidget(app))
	case "weather":
		widgets = append(widgets, weather.NewWidget(app, pages))
	case "zendesk":
		widgets = append(widgets, zendesk.NewWidget(app))
	default:
		widgets = append(widgets, unknown.NewWidget(app, widgetName))
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

	go watchForConfigChanges(app, flags.Config, display.Grid, pages)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
