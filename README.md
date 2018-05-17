<p align="center">
<img src="./docs/img/wtf.jpg?raw=true" title="WTF" width="852" height="240" />
</p>

A personal terminal-based dashboard utility, designed for
displaying infrequently-updating, but very important, daily data.

<p align="center">
<img src="./docs/img/screenshot_sm.png" title="screenshot" width="800" height="507" />
</p>

## Prerequisites

Ensure you have [Go](https://golang.org/doc/install) installed and
operational.

Depending on which modules you want to run, you'll need the appropriate API and
configuration credentials. For each of the following supported services
that you want to use, create an ENV var named as below with the
described value.

|    | Service         | ENV var                        | Value                         |
|----|-----------------|--------------------------------|-------------------------------|
| 1  | [BambooHR](https://www.bamboohr.com/api/documentation/)        | `WTF_BAMBOO_HR_TOKEN`          | BambooHR API token            |
| 2  |                 | `WTF_BAMBOO_HR_SUBDOMAIN`      | BambooHR subdomain            |
| 3  | [Github](https://developer.github.com/v3/)          | `WTF_GITHUB_TOKEN`             | Github API token              |
| 4  | [Google Calendar](https://developers.google.com/calendar/) | `WTF_GOOGLE_CAL_CLIENT_ID`     | Google Calendar client ID     |
| 5  |                 | `WTF_GOOGLE_CAL_CLIENT_SECRET` | Google Calendar client secret |
| 6  | [JIRA](https://developer.atlassian.com/server/jira/platform/rest-apis/)            | `WTF_JIRA_API_KEY`            | You JIRA API key             |
| 8  | [New Relic](https://docs.newrelic.com/docs/apis/rest-api-v2/getting-started/introduction-new-relic-rest-api-v2)       | `WTF_NEW_RELIC_API_KEY`        | New Relic API key             |
| 9  | [OpsGenie](https://docs.opsgenie.com/docs/api-overview)        | `WTF_OPS_GENIE_API_KEY`        | OpsGenie API key              |
| 10 | [OpenWeatherMap](https://openweathermap.org/api)  | `WTF_OWM_API_KEY`              | OpenWeatherMap API key        |

## Installation

1. Clone this directory and install all the dependencies.
2. Create a directory called `.wtf` in your `home` directory (ie:
   `~/.wtf/`)
3. In that directory copy the `config.yml` file (ie: `~/.wtf/config.yml`)
4. Disable all the modules for which you need an API key, by setting
   `enabled: false` in the config file
5. `go run wtf.go`

It'll probably run.

## Modules

Modules are the heart of `WTF`. Each one connects to a specific services
or renders data about a specific concept. The following modules are
included in `WTF`.

#### BambooHR

Displays the following Away information for your organization:

* Names of the people away today
* Datespan for when they'll be away

#### Git

Specify a local git repository to watch for the following:

* Current branch name
* List all changed files
* List last n commits against that branch

#### Github

Specify a Github repository to watch for the following:

* Lists all open code review requests assigned to you
* Lists all open pull requests created by you

#### Google Calendar

Displays the following information about your upcoming calendar events:

* Event title
* Date and time
* Hours/minutes/seconds until event

#### OpsGenie

Displays the following on-call information for all your active schedules:

* Schedule name
* Who's currently on call

#### Security

Displays the following security/network related information about your
local machine:

* Wifi network name
* Wifi network encryption
* Firewall enabled/disabled
* Firewall stealth mode enabled/disabled
* DNS entries

#### Weather

Displays the following current weather information for the specified city:

* weather description
* current temperature
* today's high
* today's low
* sunrise
* sunset

In the configuration, use a city code from the OpenWeatherMap [city
list](http://openweathermap.org/help/city_list.txt).

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests.

## Authors

* Chris Cummer, [senorprogrammer](https://github.com/senorprogrammer)

## License

See [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments

The inspiration for `WTF` came from Monica Dinculescu's
[tiny-care-terminal](https://github.com/notwaldorf/tiny-care-terminal).

The following open-source libraries were used in the creation of `WTF`.
Many thanks to all these developers.

* [calendar](https://google.golang.org/api/calendar/v3)
* [config](https://github.com/olebedev/config)
* [go-github](https://github.com/google/go-github)
* [goreleaser](https://github.com/goreleaser/goreleaser)
* [newrelic](https://github.com/yfronto/newrelic)
* [openweathremap](https://github.com/briandowns/openweathermap)
* [tcell](https://github.com/gdamore/tcell)
* [tview](https://github.com/rivo/tview)
