package maker

import (
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/bargraph"
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

func MakeWidget(app *tview.Application, pages *tview.Pages, widgetName string) wtf.Wtfable {
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
	// case "system":
	// 	settings := system.NewSettingsFromYAML("System", wtf.Config)
	// 	widget = system.NewWidget(app, date, version, settings)
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

func MakeWidgets(app *tview.Application, pages *tview.Pages, config *config.Config) []wtf.Wtfable {
	widgets := []wtf.Wtfable{}

	mods, _ := config.Map("wtf.mods")

	for mod := range mods {
		if enabled := config.UBool("wtf.mods."+mod+".enabled", false); enabled {
			widget := MakeWidget(app, pages, mod)
			widgets = append(widgets, widget)
		}
	}

	return widgets
}
