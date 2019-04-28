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
	moduleConfig *config.Config,
	globalConfig *config.Config,
) wtf.Wtfable {
	var widget wtf.Wtfable

	// Always in alphabetical order
	switch widgetName {
	case "bamboohr":
		settings := bamboohr.NewSettingsFromYAML("BambooHR", moduleConfig, globalConfig)
		widget = bamboohr.NewWidget(app, settings)
	case "bargraph":
		settings := bargraph.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = bargraph.NewWidget(app, settings)
	case "bittrex":
		settings := bittrex.NewSettingsFromYAML("Bittrex", moduleConfig, globalConfig)
		widget = bittrex.NewWidget(app, settings)
	case "blockfolio":
		settings := blockfolio.NewSettingsFromYAML("Blockfolio", moduleConfig, globalConfig)
		widget = blockfolio.NewWidget(app, settings)
	case "circleci":
		settings := circleci.NewSettingsFromYAML("CircleCI", moduleConfig, globalConfig)
		widget = circleci.NewWidget(app, settings)
	case "clocks":
		settings := clocks.NewSettingsFromYAML("Clocks", moduleConfig, globalConfig)
		widget = clocks.NewWidget(app, settings)
	case "cmdrunner":
		settings := cmdrunner.NewSettingsFromYAML("CmdRunner", moduleConfig, globalConfig)
		widget = cmdrunner.NewWidget(app, settings)
	case "cryptolive":
		settings := cryptolive.NewSettingsFromYAML("CryptoLive", moduleConfig, globalConfig)
		widget = cryptolive.NewWidget(app, settings)
	case "datadog":
		settings := datadog.NewSettingsFromYAML("DataDog", moduleConfig, globalConfig)
		widget = datadog.NewWidget(app, settings)
	case "gcal":
		settings := gcal.NewSettingsFromYAML("Calendar", moduleConfig, globalConfig)
		widget = gcal.NewWidget(app, settings)
	case "gerrit":
		settings := gerrit.NewSettingsFromYAML("Gerrit", moduleConfig, globalConfig)
		widget = gerrit.NewWidget(app, pages, settings)
	case "git":
		settings := git.NewSettingsFromYAML("Git", moduleConfig, globalConfig)
		widget = git.NewWidget(app, pages, settings)
	case "github":
		settings := github.NewSettingsFromYAML("GitHub", moduleConfig, globalConfig)
		widget = github.NewWidget(app, pages, settings)
	case "gitlab":
		settings := gitlab.NewSettingsFromYAML("GitLab", moduleConfig, globalConfig)
		widget = gitlab.NewWidget(app, pages, settings)
	case "gitter":
		settings := gitter.NewSettingsFromYAML("Gitter", moduleConfig, globalConfig)
		widget = gitter.NewWidget(app, pages, settings)
	case "gspreadsheets":
		settings := gspreadsheets.NewSettingsFromYAML("Google Spreadsheets", moduleConfig, globalConfig)
		widget = gspreadsheets.NewWidget(app, settings)
	case "hackernews":
		settings := hackernews.NewSettingsFromYAML("HackerNews", moduleConfig, globalConfig)
		widget = hackernews.NewWidget(app, pages, settings)
	case "ipapi":
		settings := ipapi.NewSettingsFromYAML("IPAPI", moduleConfig, globalConfig)
		widget = ipapi.NewWidget(app, settings)
	case "ipinfo":
		settings := ipinfo.NewSettingsFromYAML("IPInfo", moduleConfig, globalConfig)
		widget = ipinfo.NewWidget(app, settings)
	case "jenkins":
		settings := jenkins.NewSettingsFromYAML("Jenkins", moduleConfig, globalConfig)
		widget = jenkins.NewWidget(app, pages, settings)
	case "jira":
		settings := jira.NewSettingsFromYAML("Jira", moduleConfig, globalConfig)
		widget = jira.NewWidget(app, pages, settings)
	case "logger":
		settings := logger.NewSettingsFromYAML("Log", moduleConfig, globalConfig)
		widget = logger.NewWidget(app, settings)
	case "mercurial":
		settings := mercurial.NewSettingsFromYAML("Mercurial", moduleConfig, globalConfig)
		widget = mercurial.NewWidget(app, pages, settings)
	case "nbascore":
		settings := nbascore.NewSettingsFromYAML("NBA Score", moduleConfig, globalConfig)
		widget = nbascore.NewWidget(app, pages, settings)
	case "newrelic":
		settings := newrelic.NewSettingsFromYAML("NewRelic", moduleConfig, globalConfig)
		widget = newrelic.NewWidget(app, settings)
	case "opsgenie":
		settings := opsgenie.NewSettingsFromYAML("OpsGenie", moduleConfig, globalConfig)
		widget = opsgenie.NewWidget(app, settings)
	case "pagerduty":
		settings := pagerduty.NewSettingsFromYAML("PagerDuty", moduleConfig, globalConfig)
		widget = pagerduty.NewWidget(app, settings)
	case "power":
		settings := power.NewSettingsFromYAML("Power", moduleConfig, globalConfig)
		widget = power.NewWidget(app, settings)
	case "prettyweather":
		settings := prettyweather.NewSettingsFromYAML("Pretty Weather", moduleConfig, globalConfig)
		widget = prettyweather.NewWidget(app, settings)
	case "resourceusage":
		settings := resourceusage.NewSettingsFromYAML("Resource Usage", moduleConfig, globalConfig)
		widget = resourceusage.NewWidget(app, settings)
	case "rollbar":
		settings := rollbar.NewSettingsFromYAML("Rollbar", moduleConfig, globalConfig)
		widget = rollbar.NewWidget(app, pages, settings)
	case "security":
		settings := security.NewSettingsFromYAML("Security", moduleConfig, globalConfig)
		widget = security.NewWidget(app, settings)
	case "spotify":
		settings := spotify.NewSettingsFromYAML("Spotify", moduleConfig, globalConfig)
		widget = spotify.NewWidget(app, pages, settings)
	case "spotifyweb":
		settings := spotifyweb.NewSettingsFromYAML("Spotify Web", moduleConfig, globalConfig)
		widget = spotifyweb.NewWidget(app, pages, settings)
	case "status":
		settings := status.NewSettingsFromYAML("Status", moduleConfig, globalConfig)
		widget = status.NewWidget(app, settings)
	// case "system":
	// 	settings := system.NewSettingsFromYAML("System", moduleConfig, globalConfig)
	// 	widget = system.NewWidget(app, date, version, settings)
	case "textfile":
		settings := textfile.NewSettingsFromYAML("Textfile", moduleConfig, globalConfig)
		widget = textfile.NewWidget(app, pages, settings)
	case "todo":
		settings := todo.NewSettingsFromYAML("Todo", moduleConfig, globalConfig)
		widget = todo.NewWidget(app, pages, settings)
	case "todoist":
		settings := todoist.NewSettingsFromYAML("Todoist", moduleConfig, globalConfig)
		widget = todoist.NewWidget(app, pages, settings)
	case "travisci":
		settings := travisci.NewSettingsFromYAML("TravisCI", moduleConfig, globalConfig)
		widget = travisci.NewWidget(app, pages, settings)
	case "trello":
		settings := trello.NewSettingsFromYAML("Trello", moduleConfig, globalConfig)
		widget = trello.NewWidget(app, settings)
	case "twitter":
		settings := twitter.NewSettingsFromYAML("Twitter", moduleConfig, globalConfig)
		widget = twitter.NewWidget(app, pages, settings)
	case "victorops":
		settings := victorops.NewSettingsFromYAML("VictorOps - OnCall", moduleConfig, globalConfig)
		widget = victorops.NewWidget(app, settings)
	case "weather":
		settings := weather.NewSettingsFromYAML("Weather", moduleConfig, globalConfig)
		widget = weather.NewWidget(app, pages, settings)
	case "zendesk":
		settings := zendesk.NewSettingsFromYAML("Zendesk", moduleConfig, globalConfig)
		widget = zendesk.NewWidget(app, settings)
	default:
		settings := unknown.NewSettingsFromYAML(widgetName, moduleConfig, globalConfig)
		widget = unknown.NewWidget(app, widgetName, settings)
	}

	return widget
}

func MakeWidgets(app *tview.Application, pages *tview.Pages, config *config.Config) []wtf.Wtfable {
	widgets := []wtf.Wtfable{}

	mods, _ := config.Map("wtf.mods")

	for mod := range mods {
		modConfig, _ := config.Get("wtf.mods." + mod)
		if enabled := modConfig.UBool("enabled", false); enabled {
			widget := MakeWidget(app, pages, mod, modConfig, config)
			widgets = append(widgets, widget)
		}
	}

	return widgets
}
