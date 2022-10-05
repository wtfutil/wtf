package app

import (
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/modules/airbrake"
	"github.com/wtfutil/wtf/modules/asana"
	"github.com/wtfutil/wtf/modules/azuredevops"
	"github.com/wtfutil/wtf/modules/bamboohr"
	"github.com/wtfutil/wtf/modules/bargraph"
	"github.com/wtfutil/wtf/modules/buildkite"
	cdsfavorites "github.com/wtfutil/wtf/modules/cds/favorites"
	cdsqueue "github.com/wtfutil/wtf/modules/cds/queue"
	cdsstatus "github.com/wtfutil/wtf/modules/cds/status"
	"github.com/wtfutil/wtf/modules/circleci"
	"github.com/wtfutil/wtf/modules/clocks"
	"github.com/wtfutil/wtf/modules/cmdrunner"
	"github.com/wtfutil/wtf/modules/covid"
	"github.com/wtfutil/wtf/modules/cryptocurrency/bittrex"
	"github.com/wtfutil/wtf/modules/cryptocurrency/blockfolio"
	"github.com/wtfutil/wtf/modules/cryptocurrency/cryptolive"
	"github.com/wtfutil/wtf/modules/cryptocurrency/mempool"
	"github.com/wtfutil/wtf/modules/datadog"
	"github.com/wtfutil/wtf/modules/devto"
	"github.com/wtfutil/wtf/modules/digitalclock"
	"github.com/wtfutil/wtf/modules/digitalocean"
	"github.com/wtfutil/wtf/modules/docker"
	"github.com/wtfutil/wtf/modules/feedreader"
	"github.com/wtfutil/wtf/modules/football"
	"github.com/wtfutil/wtf/modules/gcal"
	"github.com/wtfutil/wtf/modules/gerrit"
	"github.com/wtfutil/wtf/modules/git"
	"github.com/wtfutil/wtf/modules/github"
	"github.com/wtfutil/wtf/modules/gitlab"
	"github.com/wtfutil/wtf/modules/gitlabtodo"
	"github.com/wtfutil/wtf/modules/gitter"
	"github.com/wtfutil/wtf/modules/googleanalytics"
	"github.com/wtfutil/wtf/modules/grafana"
	"github.com/wtfutil/wtf/modules/gspreadsheets"
	"github.com/wtfutil/wtf/modules/hackernews"
	"github.com/wtfutil/wtf/modules/healthchecks"
	"github.com/wtfutil/wtf/modules/hibp"
	"github.com/wtfutil/wtf/modules/ipaddresses/ipapi"
	"github.com/wtfutil/wtf/modules/ipaddresses/ipinfo"
	"github.com/wtfutil/wtf/modules/jenkins"
	"github.com/wtfutil/wtf/modules/jira"
	"github.com/wtfutil/wtf/modules/krisinformation"
	"github.com/wtfutil/wtf/modules/kubernetes"
	"github.com/wtfutil/wtf/modules/logger"
	"github.com/wtfutil/wtf/modules/mercurial"
	"github.com/wtfutil/wtf/modules/nbascore"
	"github.com/wtfutil/wtf/modules/newrelic"
	"github.com/wtfutil/wtf/modules/nextbus"
	"github.com/wtfutil/wtf/modules/opsgenie"
	"github.com/wtfutil/wtf/modules/pagerduty"
	"github.com/wtfutil/wtf/modules/pihole"
	"github.com/wtfutil/wtf/modules/pocket"
	"github.com/wtfutil/wtf/modules/power"
	"github.com/wtfutil/wtf/modules/resourceusage"
	"github.com/wtfutil/wtf/modules/rollbar"
	"github.com/wtfutil/wtf/modules/security"
	"github.com/wtfutil/wtf/modules/spacex"
	"github.com/wtfutil/wtf/modules/spotify"
	"github.com/wtfutil/wtf/modules/spotifyweb"
	"github.com/wtfutil/wtf/modules/status"
	"github.com/wtfutil/wtf/modules/steam"
	"github.com/wtfutil/wtf/modules/stocks/finnhub"
	"github.com/wtfutil/wtf/modules/stocks/yfinance"
	"github.com/wtfutil/wtf/modules/subreddit"
	"github.com/wtfutil/wtf/modules/textfile"
	"github.com/wtfutil/wtf/modules/todo"
	"github.com/wtfutil/wtf/modules/todo_plus"
	"github.com/wtfutil/wtf/modules/transmission"
	"github.com/wtfutil/wtf/modules/travisci"
	"github.com/wtfutil/wtf/modules/twitch"
	"github.com/wtfutil/wtf/modules/twitter"
	"github.com/wtfutil/wtf/modules/twitterstats"
	"github.com/wtfutil/wtf/modules/unknown"
	"github.com/wtfutil/wtf/modules/updown"
	"github.com/wtfutil/wtf/modules/uptimerobot"
	"github.com/wtfutil/wtf/modules/victorops"
	"github.com/wtfutil/wtf/modules/weatherservices/arpansagovau"
	"github.com/wtfutil/wtf/modules/weatherservices/prettyweather"
	"github.com/wtfutil/wtf/modules/weatherservices/weather"
	"github.com/wtfutil/wtf/modules/zendesk"
	"github.com/wtfutil/wtf/wtf"
)

// MakeWidget creates and returns instances of widgets
func MakeWidget(
	tviewApp *tview.Application,
	pages *tview.Pages,
	moduleName string,
	config *config.Config,
	redrawChan chan bool,
) wtf.Wtfable {
	var widget wtf.Wtfable

	moduleConfig, _ := config.Get("wtf.mods." + moduleName)

	// Don' try to initialize modules that don't exist
	if moduleConfig == nil {
		return nil
	}

	// Don't try to initialize modules that aren't enabled
	if enabled := moduleConfig.UBool("enabled", false); !enabled {
		return nil
	}

	// Always in alphabetical order
	switch moduleConfig.UString("type", moduleName) {
	case "airbrake":
		settings := airbrake.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = airbrake.NewWidget(tviewApp, redrawChan, pages, settings)
	case "arpansagovau":
		settings := arpansagovau.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = arpansagovau.NewWidget(tviewApp, redrawChan, settings)
	case "asana":
		settings := asana.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = asana.NewWidget(tviewApp, redrawChan, pages, settings)
	case "azuredevops":
		settings := azuredevops.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = azuredevops.NewWidget(tviewApp, redrawChan, pages, settings)
	case "bamboohr":
		settings := bamboohr.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = bamboohr.NewWidget(tviewApp, redrawChan, settings)
	case "bargraph":
		settings := bargraph.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = bargraph.NewWidget(tviewApp, redrawChan, settings)
	case "bittrex":
		settings := bittrex.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = bittrex.NewWidget(tviewApp, redrawChan, settings)
	case "blockfolio":
		settings := blockfolio.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = blockfolio.NewWidget(tviewApp, redrawChan, settings)
	case "buildkite":
		settings := buildkite.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = buildkite.NewWidget(tviewApp, redrawChan, pages, settings)
	case "cdsFavorites":
		settings := cdsfavorites.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = cdsfavorites.NewWidget(tviewApp, redrawChan, pages, settings)
	case "cdsQueue":
		settings := cdsqueue.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = cdsqueue.NewWidget(tviewApp, redrawChan, pages, settings)
	case "cdsStatus":
		settings := cdsstatus.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = cdsstatus.NewWidget(tviewApp, redrawChan, pages, settings)
	case "circleci":
		settings := circleci.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = circleci.NewWidget(tviewApp, redrawChan, settings)
	case "clocks":
		settings := clocks.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = clocks.NewWidget(tviewApp, redrawChan, settings)
	case "covid":
		settings := covid.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = covid.NewWidget(tviewApp, redrawChan, settings)
	case "cmdrunner":
		settings := cmdrunner.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = cmdrunner.NewWidget(tviewApp, redrawChan, settings)
	case "cryptolive":
		settings := cryptolive.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = cryptolive.NewWidget(tviewApp, redrawChan, settings)
	case "datadog":
		settings := datadog.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = datadog.NewWidget(tviewApp, redrawChan, pages, settings)
	case "devto":
		settings := devto.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = devto.NewWidget(tviewApp, redrawChan, pages, settings)
	case "digitalclock":
		settings := digitalclock.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = digitalclock.NewWidget(tviewApp, redrawChan, settings)
	case "digitalocean":
		settings := digitalocean.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = digitalocean.NewWidget(tviewApp, redrawChan, pages, settings)
	case "docker":
		settings := docker.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = docker.NewWidget(tviewApp, redrawChan, pages, settings)
	case "feedreader":
		settings := feedreader.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = feedreader.NewWidget(tviewApp, redrawChan, pages, settings)
	case "football":
		settings := football.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = football.NewWidget(tviewApp, redrawChan, pages, settings)
	case "gcal":
		settings := gcal.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = gcal.NewWidget(tviewApp, redrawChan, settings)
	case "gerrit":
		settings := gerrit.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = gerrit.NewWidget(tviewApp, redrawChan, pages, settings)
	case "git":
		settings := git.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = git.NewWidget(tviewApp, redrawChan, pages, settings)
	case "github":
		settings := github.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = github.NewWidget(tviewApp, redrawChan, pages, settings)
	case "gitlab":
		settings := gitlab.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = gitlab.NewWidget(tviewApp, redrawChan, pages, settings)
	case "gitlabtodo":
		settings := gitlabtodo.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = gitlabtodo.NewWidget(tviewApp, redrawChan, pages, settings)
	case "gitter":
		settings := gitter.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = gitter.NewWidget(tviewApp, redrawChan, pages, settings)
	case "googleanalytics":
		settings := googleanalytics.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = googleanalytics.NewWidget(tviewApp, redrawChan, settings)
	case "gspreadsheets":
		settings := gspreadsheets.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = gspreadsheets.NewWidget(tviewApp, redrawChan, settings)
	case "grafana":
		settings := grafana.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = grafana.NewWidget(tviewApp, redrawChan, pages, settings)
	case "hackernews":
		settings := hackernews.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = hackernews.NewWidget(tviewApp, redrawChan, pages, settings)
	case "healthchecks":
		settings := healthchecks.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = healthchecks.NewWidget(tviewApp, redrawChan, pages, settings)
	case "hibp":
		settings := hibp.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = hibp.NewWidget(tviewApp, redrawChan, settings)
	case "ipapi":
		settings := ipapi.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = ipapi.NewWidget(tviewApp, redrawChan, settings)
	case "ipinfo":
		settings := ipinfo.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = ipinfo.NewWidget(tviewApp, redrawChan, settings)
	case "jenkins":
		settings := jenkins.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = jenkins.NewWidget(tviewApp, redrawChan, pages, settings)
	case "jira":
		settings := jira.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = jira.NewWidget(tviewApp, redrawChan, pages, settings)
	case "kubernetes":
		settings := kubernetes.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = kubernetes.NewWidget(tviewApp, redrawChan, settings)
	case "krisinformation":
		settings := krisinformation.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = krisinformation.NewWidget(tviewApp, redrawChan, settings)
	case "logger":
		settings := logger.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = logger.NewWidget(tviewApp, redrawChan, settings)
	case "mercurial":
		settings := mercurial.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = mercurial.NewWidget(tviewApp, redrawChan, pages, settings)
	case "mempool":
		settings := mempool.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = mempool.NewWidget(tviewApp, redrawChan, pages, settings)
	case "nbascore":
		settings := nbascore.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = nbascore.NewWidget(tviewApp, redrawChan, pages, settings)
	case "newrelic":
		settings := newrelic.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = newrelic.NewWidget(tviewApp, redrawChan, pages, settings)
	case "nextbus":
		settings := nextbus.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = nextbus.NewWidget(tviewApp, redrawChan, pages, settings)
	case "opsgenie":
		settings := opsgenie.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = opsgenie.NewWidget(tviewApp, redrawChan, settings)
	case "pagerduty":
		settings := pagerduty.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = pagerduty.NewWidget(tviewApp, redrawChan, settings)
	case "pihole":
		settings := pihole.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = pihole.NewWidget(tviewApp, redrawChan, pages, settings)
	case "power":
		settings := power.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = power.NewWidget(tviewApp, redrawChan, settings)
	case "prettyweather":
		settings := prettyweather.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = prettyweather.NewWidget(tviewApp, redrawChan, settings)
	case "pocket":
		settings := pocket.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = pocket.NewWidget(tviewApp, redrawChan, pages, settings)
	case "resourceusage":
		settings := resourceusage.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = resourceusage.NewWidget(tviewApp, redrawChan, settings)
	case "rollbar":
		settings := rollbar.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = rollbar.NewWidget(tviewApp, redrawChan, pages, settings)
	case "security":
		settings := security.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = security.NewWidget(tviewApp, redrawChan, settings)
	case "spacex":
		settings := spacex.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = spacex.NewWidget(tviewApp, redrawChan, settings)
	case "spotify":
		settings := spotify.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = spotify.NewWidget(tviewApp, redrawChan, pages, settings)
	case "spotifyweb":
		settings := spotifyweb.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = spotifyweb.NewWidget(tviewApp, redrawChan, pages, settings)
	case "status":
		settings := status.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = status.NewWidget(tviewApp, redrawChan, settings)
	case "steam":
		settings := steam.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = steam.NewWidget(tviewApp, redrawChan, pages, settings)
	case "subreddit":
		settings := subreddit.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = subreddit.NewWidget(tviewApp, redrawChan, pages, settings)
	case "textfile":
		settings := textfile.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = textfile.NewWidget(tviewApp, redrawChan, pages, settings)
	case "todo":
		settings := todo.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = todo.NewWidget(tviewApp, redrawChan, pages, settings)
	case "todo_plus":
		settings := todo_plus.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = todo_plus.NewWidget(tviewApp, redrawChan, pages, settings)
	case "todoist":
		settings := todo_plus.FromTodoist(moduleName, moduleConfig, config)
		widget = todo_plus.NewWidget(tviewApp, redrawChan, pages, settings)
	case "transmission":
		settings := transmission.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = transmission.NewWidget(tviewApp, redrawChan, pages, settings)
	case "travisci":
		settings := travisci.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = travisci.NewWidget(tviewApp, redrawChan, pages, settings)
	case "trello":
		settings := todo_plus.FromTrello(moduleName, moduleConfig, config)
		widget = todo_plus.NewWidget(tviewApp, redrawChan, pages, settings)
	case "twitch":
		settings := twitch.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = twitch.NewWidget(tviewApp, redrawChan, pages, settings)
	case "twitter":
		settings := twitter.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = twitter.NewWidget(tviewApp, redrawChan, pages, settings)
	case "twitterstats":
		settings := twitterstats.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = twitterstats.NewWidget(tviewApp, redrawChan, pages, settings)
	case "updown":
		settings := updown.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = updown.NewWidget(tviewApp, redrawChan, pages, settings)
	case "uptimerobot":
		settings := uptimerobot.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = uptimerobot.NewWidget(tviewApp, redrawChan, pages, settings)
	case "victorops":
		settings := victorops.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = victorops.NewWidget(tviewApp, redrawChan, settings)
	case "weather":
		settings := weather.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = weather.NewWidget(tviewApp, redrawChan, pages, settings)
	case "zendesk":
		settings := zendesk.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = zendesk.NewWidget(tviewApp, redrawChan, pages, settings)
	case "finnhub":
		settings := finnhub.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = finnhub.NewWidget(tviewApp, redrawChan, settings)
	case "yfinance":
		settings := yfinance.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = yfinance.NewWidget(tviewApp, redrawChan, settings)
	default:
		settings := unknown.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = unknown.NewWidget(tviewApp, redrawChan, settings)
	}

	return widget
}

// MakeWidgets creates and returns a collection of enabled widgets
func MakeWidgets(tviewApp *tview.Application, pages *tview.Pages, config *config.Config, redrawChan chan bool) []wtf.Wtfable {
	widgets := []wtf.Wtfable{}

	moduleNames, _ := config.Map("wtf.mods")

	for moduleName := range moduleNames {
		widget := MakeWidget(tviewApp, pages, moduleName, config, redrawChan)

		if widget != nil {
			widgets = append(widgets, widget)
		}
	}

	return widgets
}
