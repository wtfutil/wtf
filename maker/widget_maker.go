package maker

import (
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/logger"
	"github.com/wtfutil/wtf/modules/bamboohr"
	"github.com/wtfutil/wtf/modules/bargraph"
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
	// "github.com/wtfutil/wtf/modules/system"
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

func MakeWidget(
	app *tview.Application,
	pages *tview.Pages,
	widgetName string,
	widgetType string,
	moduleConfig *config.Config,
	globalConfig *config.Config,
) wtf.Wtfable {
	var widget wtf.Wtfable

	// Always in alphabetical order
	switch widgetType {
	case "bamboohr":
		settings := bamboohr.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = bamboohr.NewWidget(app, settings)
	case "bargraph":
		settings := bargraph.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = bargraph.NewWidget(app, settings)
	case "bittrex":
		settings := bittrex.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = bittrex.NewWidget(app, settings)
	case "blockfolio":
		settings := blockfolio.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = blockfolio.NewWidget(app, settings)
	case "circleci":
		settings := circleci.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = circleci.NewWidget(app, settings)
	case "clocks":
		settings := clocks.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = clocks.NewWidget(app, settings)
	case "cmdrunner":
		settings := cmdrunner.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = cmdrunner.NewWidget(app, settings)
	case "cryptolive":
		settings := cryptolive.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = cryptolive.NewWidget(app, settings)
	case "datadog":
		settings := datadog.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = datadog.NewWidget(app, pages, settings)
	case "gcal":
		settings := gcal.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = gcal.NewWidget(app, settings)
	case "gerrit":
		settings := gerrit.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = gerrit.NewWidget(app, pages, settings)
	case "git":
		settings := git.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = git.NewWidget(app, pages, settings)
	case "github":
		settings := github.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = github.NewWidget(app, pages, settings)
	case "gitlab":
		settings := gitlab.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = gitlab.NewWidget(app, pages, settings)
	case "gitter":
		settings := gitter.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = gitter.NewWidget(app, pages, settings)
	case "gspreadsheets":
		settings := gspreadsheets.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = gspreadsheets.NewWidget(app, settings)
	case "hackernews":
		settings := hackernews.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = hackernews.NewWidget(app, pages, settings)
	case "ipapi":
		settings := ipapi.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = ipapi.NewWidget(app, settings)
	case "ipinfo":
		settings := ipinfo.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = ipinfo.NewWidget(app, settings)
	case "jenkins":
		settings := jenkins.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = jenkins.NewWidget(app, pages, settings)
	case "jira":
		settings := jira.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = jira.NewWidget(app, pages, settings)
	case "logger":
		settings := logger.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = logger.NewWidget(app, settings)
	case "mercurial":
		settings := mercurial.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = mercurial.NewWidget(app, pages, settings)
	case "nbascore":
		settings := nbascore.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = nbascore.NewWidget(app, pages, settings)
	case "newrelic":
		settings := newrelic.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = newrelic.NewWidget(app, settings)
	case "opsgenie":
		settings := opsgenie.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = opsgenie.NewWidget(app, settings)
	case "pagerduty":
		settings := pagerduty.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = pagerduty.NewWidget(app, settings)
	case "power":
		settings := power.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = power.NewWidget(app, settings)
	case "prettyweather":
		settings := prettyweather.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = prettyweather.NewWidget(app, settings)
	case "resourceusage":
		settings := resourceusage.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = resourceusage.NewWidget(app, settings)
	case "rollbar":
		settings := rollbar.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = rollbar.NewWidget(app, pages, settings)
	case "security":
		settings := security.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = security.NewWidget(app, settings)
	case "spotify":
		settings := spotify.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = spotify.NewWidget(app, pages, settings)
	case "spotifyweb":
		settings := spotifyweb.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = spotifyweb.NewWidget(app, pages, settings)
	case "status":
		settings := status.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = status.NewWidget(app, settings)
	// case "system":
	// 	settings := system.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
	// 	widget = system.NewWidget(app, date, version, settings)
	case "textfile":
		settings := textfile.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = textfile.NewWidget(app, pages, settings)
	case "todo":
		settings := todo.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = todo.NewWidget(app, pages, settings)
	case "todoist":
		settings := todoist.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = todoist.NewWidget(app, pages, settings)
	case "travisci":
		settings := travisci.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = travisci.NewWidget(app, pages, settings)
	case "trello":
		settings := trello.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = trello.NewWidget(app, settings)
	case "twitter":
		settings := twitter.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = twitter.NewWidget(app, pages, settings)
	case "victorops":
		settings := victorops.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = victorops.NewWidget(app, settings)
	case "weather":
		settings := weather.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = weather.NewWidget(app, pages, settings)
	case "zendesk":
		settings := zendesk.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = zendesk.NewWidget(app, pages, settings)
	default:
		settings := unknown.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = unknown.NewWidget(app, settings)
	}

	return widget
}

func MakeWidgets(app *tview.Application, pages *tview.Pages, config *config.Config) []wtf.Wtfable {
	widgets := []wtf.Wtfable{}

	mods, _ := config.Map("wtf.mods")

	for mod := range mods {
		modConfig, _ := config.Get("wtf.mods." + mod)
		widgetType := modConfig.UString("type", mod)
		if enabled := modConfig.UBool("enabled", false); enabled {
			widget := MakeWidget(app, pages, mod, widgetType, modConfig, config)
			widgets = append(widgets, widget)
		}
	}

	return widgets
}
