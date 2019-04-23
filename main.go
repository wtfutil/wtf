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
var runningWidgets []wtf.Wtfable

// Config parses the config.yml file and makes available the settings within
var Config *config.Config

var (
	commit  = "dev"
	date    = "dev"
	version = "dev"
)

/* -------------------- Functions -------------------- */

func disableAllWidgets(widgets []wtf.Wtfable) {
	for _, widget := range widgets {
		widget.Disable()
	}
}

func initializeFocusTracker(app *tview.Application, widgets []wtf.Wtfable) {
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
		refreshAllWidgets(runningWidgets)
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

func refreshAllWidgets(widgets []wtf.Wtfable) {
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

	// Notify write events
	watch.FilterOps(watcher.Write)

	go func() {
		for {
			select {
			case <-watch.Event:
				// Disable all widgets to stop scheduler goroutines and rmeove widgets from memory.
				disableAllWidgets(runningWidgets)

				loadConfigFile(absPath)

				widgets := makeWidgets(app, pages)
				wtf.ValidateWidgets(widgets)

				initializeFocusTracker(app, widgets)

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

func makeWidget(app *tview.Application, pages *tview.Pages, widgetName string) wtf.Wtfable {
	var widget wtf.Wtfable

	// Always in alphabetical order
	switch widgetName {
	case "bamboohr":
		settings := bamboohr.NewSettingsFromYAML("BambooHR", wtf.Config)
		widget = bamboohr.NewWidget(app, settings)
	case "bargraph":
		widget = bargraph.NewWidget(app)
	case "bittrex":
		settings := bittrex.NewSettingsFromYAML("Bittrex", wtf.Config)
		widget = bittrex.NewWidget(app, settings)
	case "blockfolio":
		settings := blockfolio.NewSettingsFromYAML("Blockfolio", wtf.Config)
		widget = blockfolio.NewWidget(app, settings)
	case "circleci":
		settings := circleci.NewSettingsFromYAML("CircleCI", wtf.Config)
		widget = circleci.NewWidget(app, settings)
	case "clocks":
		settings := clocks.NewSettingsFromYAML("Clocks", wtf.Config)
		widget = clocks.NewWidget(app, settings)
	case "cmdrunner":
		settings := cmdrunner.NewSettingsFromYAML("CmdRunner", wtf.Config)
		widget = cmdrunner.NewWidget(app, settings)
	case "cryptolive":
		settings := cryptolive.NewSettingsFromYAML("CryptoLive", wtf.Config)
		widget = cryptolive.NewWidget(app, settings)
	case "datadog":
		settings := datadog.NewSettingsFromYAML("DataDog", wtf.Config)
		widget = datadog.NewWidget(app, settings)
	case "gcal":
		settings := gcal.NewSettingsFromYAML("Calendar", wtf.Config)
		widget = gcal.NewWidget(app, settings)
	case "gerrit":
		settings := gerrit.NewSettingsFromYAML("Gerrit", wtf.Config)
		widget = gerrit.NewWidget(app, pages, settings)
	case "git":
		settings := git.NewSettingsFromYAML("Git", wtf.Config)
		widget = git.NewWidget(app, pages, settings)
	case "github":
		settings := github.NewSettingsFromYAML("GitHub", wtf.Config)
		widget = github.NewWidget(app, pages, settings)
	case "gitlab":
		settings := gitlab.NewSettingsFromYAML("GitLab", wtf.Config)
		widget = gitlab.NewWidget(app, pages, settings)
	case "gitter":
		settings := gitter.NewSettingsFromYAML("Gitter", wtf.Config)
		widget = gitter.NewWidget(app, pages, settings)
	case "gspreadsheets":
		settings := gspreadsheets.NewSettingsFromYAML("Google Spreadsheets", wtf.Config)
		widget = gspreadsheets.NewWidget(app, settings)
	case "hackernews":
		settings := hackernews.NewSettingsFromYAML("HackerNews", wtf.Config)
		widget = hackernews.NewWidget(app, pages, settings)
	case "ipapi":
		settings := ipapi.NewSettingsFromYAML("IPAPI", wtf.Config)
		widget = ipapi.NewWidget(app, settings)
	case "ipinfo":
		settings := ipinfo.NewSettingsFromYAML("IPInfo", wtf.Config)
		widget = ipinfo.NewWidget(app, settings)
	case "jenkins":
		settings := jenkins.NewSettingsFromYAML("Jenkins", wtf.Config)
		widget = jenkins.NewWidget(app, pages, settings)
	case "jira":
		settings := jira.NewSettingsFromYAML("Jira", wtf.Config)
		widget = jira.NewWidget(app, pages, settings)
	case "logger":
		settings := logger.NewSettingsFromYAML("Log", wtf.Config)
		widget = logger.NewWidget(app, settings)
	case "mercurial":
		settings := mercurial.NewSettingsFromYAML("Mercurial", wtf.Config)
		widget = mercurial.NewWidget(app, pages, settings)
	case "nbascore":
		settings := nbascore.NewSettingsFromYAML("NBA Score", wtf.Config)
		widget = nbascore.NewWidget(app, pages, settings)
	case "newrelic":
		settings := newrelic.NewSettingsFromYAML("NewRelic", wtf.Config)
		widget = newrelic.NewWidget(app, settings)
	case "opsgenie":
		settings := opsgenie.NewSettingsFromYAML("OpsGenie", wtf.Config)
		widget = opsgenie.NewWidget(app, settings)
	case "pagerduty":
		settings := pagerduty.NewSettingsFromYAML("PagerDuty", wtf.Config)
		widget = pagerduty.NewWidget(app, settings)
	case "power":
		settings := power.NewSettingsFromYAML("Power", wtf.Config)
		widget = power.NewWidget(app, settings)
	case "prettyweather":
		settings := prettyweather.NewSettingsFromYAML("Pretty Weather", wtf.Config)
		widget = prettyweather.NewWidget(app, settings)
	case "resourceusage":
		settings := resourceusage.NewSettingsFromYAML("Resource Usage", wtf.Config)
		widget = resourceusage.NewWidget(app, settings)
	case "rollbar":
		settings := rollbar.NewSettingsFromYAML("Rollbar", wtf.Config)
		widget = rollbar.NewWidget(app, pages, settings)
	case "security":
		settings := security.NewSettingsFromYAML("Security", wtf.Config)
		widget = security.NewWidget(app, settings)
	case "spotify":
		settings := spotify.NewSettingsFromYAML("Spotify", wtf.Config)
		widget = spotify.NewWidget(app, pages, settings)
	case "spotifyweb":
		settings := spotifyweb.NewSettingsFromYAML("Spotify Web", wtf.Config)
		widget = spotifyweb.NewWidget(app, pages, settings)
	case "status":
		settings := status.NewSettingsFromYAML("Status", wtf.Config)
		widget = status.NewWidget(app, settings)
	case "system":
		settings := system.NewSettingsFromYAML("System", wtf.Config)
		widget = system.NewWidget(app, date, version, settings)
	case "textfile":
		settings := textfile.NewSettingsFromYAML("Textfile", wtf.Config)
		widget = textfile.NewWidget(app, pages, settings)
	case "todo":
		settings := todo.NewSettingsFromYAML("Todo", wtf.Config)
		widget = todo.NewWidget(app, pages, settings)
	case "todoist":
		settings := todoist.NewSettingsFromYAML("Todoist", wtf.Config)
		widget = todoist.NewWidget(app, pages, settings)
	case "travisci":
		settings := travisci.NewSettingsFromYAML("TravisCI", wtf.Config)
		widget = travisci.NewWidget(app, pages, settings)
	case "trello":
		settings := trello.NewSettingsFromYAML("Trello", wtf.Config)
		widget = trello.NewWidget(app, settings)
	case "twitter":
		settings := twitter.NewSettingsFromYAML("Twitter", wtf.Config)
		widget = twitter.NewWidget(app, pages, settings)
	case "victorops":
		settings := victorops.NewSettingsFromYAML("VictorOps - OnCall", wtf.Config)
		widget = victorops.NewWidget(app, settings)
	case "weather":
		settings := weather.NewSettingsFromYAML("Weather", wtf.Config)
		widget = weather.NewWidget(app, pages, settings)
	case "zendesk":
		settings := zendesk.NewSettingsFromYAML("Zendesk", wtf.Config)
		widget = zendesk.NewWidget(app, settings)
	default:
		settings := unknown.NewSettingsFromYAML(widgetName, wtf.Config)
		widget = unknown.NewWidget(app, widgetName, settings)
	}

	return widget
}

func makeWidgets(app *tview.Application, pages *tview.Pages) []wtf.Wtfable {
	widgets := []wtf.Wtfable{}

	mods, _ := Config.Map("wtf.mods")

	for mod := range mods {
		if enabled := Config.UBool("wtf.mods."+mod+".enabled", false); enabled {
			widget := makeWidget(app, pages, mod)
			widgets = append(widgets, widget)
		}
	}

	// This is a hack to allow refreshAllWidgets and disableAllWidgets to work
	// Need to implement a non-global way to track these
	runningWidgets = widgets

	return widgets
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

	widgets := makeWidgets(app, pages)
	wtf.ValidateWidgets(widgets)

	initializeFocusTracker(app, widgets)

	display := wtf.NewDisplay(widgets)
	pages.AddPage("grid", display.Grid, true, true)

	app.SetInputCapture(keyboardIntercept)

	go watchForConfigChanges(app, flags.Config, display.Grid, pages)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
