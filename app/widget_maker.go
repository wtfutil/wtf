package app

import (
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/modules/azuredevops"
	"github.com/wtfutil/wtf/modules/bamboohr"
	"github.com/wtfutil/wtf/modules/bargraph"
	"github.com/wtfutil/wtf/modules/buildkite"
	"github.com/wtfutil/wtf/modules/circleci"
	"github.com/wtfutil/wtf/modules/clocks"
	"github.com/wtfutil/wtf/modules/cmdrunner"
	"github.com/wtfutil/wtf/modules/cryptoexchanges/bittrex"
	"github.com/wtfutil/wtf/modules/cryptoexchanges/blockfolio"
	"github.com/wtfutil/wtf/modules/cryptoexchanges/cryptolive"
	"github.com/wtfutil/wtf/modules/datadog"
	"github.com/wtfutil/wtf/modules/devto"
	"github.com/wtfutil/wtf/modules/digitalclock"
	"github.com/wtfutil/wtf/modules/docker"
	"github.com/wtfutil/wtf/modules/exchangerates"
	"github.com/wtfutil/wtf/modules/feedreader"
	"github.com/wtfutil/wtf/modules/football"
	"github.com/wtfutil/wtf/modules/gcal"
	"github.com/wtfutil/wtf/modules/gerrit"
	"github.com/wtfutil/wtf/modules/git"
	"github.com/wtfutil/wtf/modules/github"
	"github.com/wtfutil/wtf/modules/gitlab"
	"github.com/wtfutil/wtf/modules/gitter"
	"github.com/wtfutil/wtf/modules/googleanalytics"
	"github.com/wtfutil/wtf/modules/gspreadsheets"
	"github.com/wtfutil/wtf/modules/hackernews"
	"github.com/wtfutil/wtf/modules/hibp"
	"github.com/wtfutil/wtf/modules/ipaddresses/ipapi"
	"github.com/wtfutil/wtf/modules/ipaddresses/ipinfo"
	"github.com/wtfutil/wtf/modules/jenkins"
	"github.com/wtfutil/wtf/modules/jira"
	"github.com/wtfutil/wtf/modules/kubernetes"
	"github.com/wtfutil/wtf/modules/logger"
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
	"github.com/wtfutil/wtf/modules/subreddit"
	"github.com/wtfutil/wtf/modules/textfile"
	"github.com/wtfutil/wtf/modules/todo"
	"github.com/wtfutil/wtf/modules/todoist"
	"github.com/wtfutil/wtf/modules/transmission"
	"github.com/wtfutil/wtf/modules/travisci"
	"github.com/wtfutil/wtf/modules/trello"
	"github.com/wtfutil/wtf/modules/twitter"
	"github.com/wtfutil/wtf/modules/twitterstats"
	"github.com/wtfutil/wtf/modules/unknown"
	"github.com/wtfutil/wtf/modules/victorops"
	"github.com/wtfutil/wtf/modules/weatherservices/arpansagovau"
	"github.com/wtfutil/wtf/modules/weatherservices/prettyweather"
	"github.com/wtfutil/wtf/modules/weatherservices/weather"
	"github.com/wtfutil/wtf/modules/zendesk"
	"github.com/wtfutil/wtf/wtf"
)

// MakeWidget creates and returns instances of widgets
func MakeWidget(
	app *tview.Application,
	pages *tview.Pages,
	moduleName string,
	config *config.Config,
) wtf.Wtfable {
	var widget wtf.Wtfable

	moduleConfig, _ := config.Get("wtf.mods." + moduleName)
	if enabled := moduleConfig.UBool("enabled", false); !enabled {
		// Don't initialize modules that aren't enabled
		return nil
	}

	// Always in alphabetical order
	switch moduleConfig.UString("type", moduleName) {
	case "arpansagovau":
		settings := arpansagovau.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = arpansagovau.NewWidget(app, settings)
	case "azuredevops":
		settings := azuredevops.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = azuredevops.NewWidget(app, pages, settings)
	case "bamboohr":
		settings := bamboohr.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = bamboohr.NewWidget(app, settings)
	case "bargraph":
		settings := bargraph.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = bargraph.NewWidget(app, settings)
	case "bittrex":
		settings := bittrex.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = bittrex.NewWidget(app, settings)
	case "blockfolio":
		settings := blockfolio.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = blockfolio.NewWidget(app, settings)
	case "buildkite":
		settings := buildkite.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = buildkite.NewWidget(app, pages, settings)
	case "circleci":
		settings := circleci.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = circleci.NewWidget(app, settings)
	case "clocks":
		settings := clocks.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = clocks.NewWidget(app, settings)
	case "digitalclock":
		settings := digitalclock.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = digitalclock.NewWidget(app, settings)
	case "cmdrunner":
		settings := cmdrunner.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = cmdrunner.NewWidget(app, settings)
	case "cryptolive":
		settings := cryptolive.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = cryptolive.NewWidget(app, settings)
	case "datadog":
		settings := datadog.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = datadog.NewWidget(app, pages, settings)
	case "devto":
		settings := devto.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = devto.NewWidget(app, pages, settings)
	case "docker":
		settings := docker.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = docker.NewWidget(app, pages, settings)
	case "feedreader":
		settings := feedreader.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = feedreader.NewWidget(app, pages, settings)
	case "football":
		settings := football.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = football.NewWidget(app, pages, settings)
	case "gcal":
		settings := gcal.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = gcal.NewWidget(app, settings)
	case "gerrit":
		settings := gerrit.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = gerrit.NewWidget(app, pages, settings)
	case "git":
		settings := git.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = git.NewWidget(app, pages, settings)
	case "github":
		settings := github.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = github.NewWidget(app, pages, settings)
	case "gitlab":
		settings := gitlab.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = gitlab.NewWidget(app, pages, settings)
	case "gitter":
		settings := gitter.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = gitter.NewWidget(app, pages, settings)
	case "googleanalytics":
		settings := googleanalytics.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = googleanalytics.NewWidget(app, settings)
	case "gspreadsheets":
		settings := gspreadsheets.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = gspreadsheets.NewWidget(app, settings)
	case "hackernews":
		settings := hackernews.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = hackernews.NewWidget(app, pages, settings)
	case "hibp":
		settings := hibp.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = hibp.NewWidget(app, settings)
	case "ipapi":
		settings := ipapi.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = ipapi.NewWidget(app, settings)
	case "ipinfo":
		settings := ipinfo.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = ipinfo.NewWidget(app, settings)
	case "jenkins":
		settings := jenkins.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = jenkins.NewWidget(app, pages, settings)
	case "jira":
		settings := jira.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = jira.NewWidget(app, pages, settings)
	case "kubernetes":
		settings := kubernetes.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = kubernetes.NewWidget(app, settings)
	case "logger":
		settings := logger.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = logger.NewWidget(app, settings)
	case "mercurial":
		settings := mercurial.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = mercurial.NewWidget(app, pages, settings)
	case "nbascore":
		settings := nbascore.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = nbascore.NewWidget(app, pages, settings)
	case "newrelic":
		settings := newrelic.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = newrelic.NewWidget(app, pages, settings)
	case "opsgenie":
		settings := opsgenie.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = opsgenie.NewWidget(app, settings)
	case "pagerduty":
		settings := pagerduty.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = pagerduty.NewWidget(app, settings)
	case "power":
		settings := power.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = power.NewWidget(app, settings)
	case "prettyweather":
		settings := prettyweather.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = prettyweather.NewWidget(app, settings)
	case "resourceusage":
		settings := resourceusage.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = resourceusage.NewWidget(app, settings)
	case "rollbar":
		settings := rollbar.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = rollbar.NewWidget(app, pages, settings)
	case "security":
		settings := security.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = security.NewWidget(app, settings)
	case "spotify":
		settings := spotify.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = spotify.NewWidget(app, pages, settings)
	case "spotifyweb":
		settings := spotifyweb.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = spotifyweb.NewWidget(app, pages, settings)
	case "status":
		settings := status.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = status.NewWidget(app, settings)
	case "subreddit":
		settings := subreddit.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = subreddit.NewWidget(app, pages, settings)
	case "textfile":
		settings := textfile.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = textfile.NewWidget(app, pages, settings)
	case "todo":
		settings := todo.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = todo.NewWidget(app, pages, settings)
	case "todoist":
		settings := todoist.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = todoist.NewWidget(app, pages, settings)
	case "transmission":
		settings := transmission.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = transmission.NewWidget(app, pages, settings)
	case "travisci":
		settings := travisci.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = travisci.NewWidget(app, pages, settings)
	case "trello":
		settings := trello.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = trello.NewWidget(app, settings)
	case "twitter":
		settings := twitter.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = twitter.NewWidget(app, pages, settings)
	case "twitterstats":
		settings := twitterstats.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = twitterstats.NewWidget(app, pages, settings)
	case "victorops":
		settings := victorops.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = victorops.NewWidget(app, settings)
	case "weather":
		settings := weather.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = weather.NewWidget(app, pages, settings)
	case "zendesk":
		settings := zendesk.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = zendesk.NewWidget(app, pages, settings)
	case "exchangerates":
		settings := exchangerates.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = exchangerates.NewWidget(app, pages, settings)
	default:
		settings := unknown.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = unknown.NewWidget(app, settings)
	}

	return widget
}

// MakeWidgets creates and returns a collection of enabled widgets
func MakeWidgets(app *tview.Application, pages *tview.Pages, config *config.Config) []wtf.Wtfable {
	widgets := []wtf.Wtfable{}

	moduleNames, _ := config.Map("wtf.mods")

	for moduleName := range moduleNames {
		widget := MakeWidget(app, pages, moduleName, config)

		if widget != nil {
			widgets = append(widgets, widget)
		}
	}

	return widgets
}
