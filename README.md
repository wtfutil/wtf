<p align="center">
<img src="docs/wtf.jpg?raw=true" title="WTF" width="852" height="240")
</p>

A personal terminal-based dashboard utility, designed for
displaying infrequently-updating, but very important, daily data.

## Required Prerequisites

1. [Git](https://git-scm.com/downloads)
1. [Go](https://golang.org/doc/install)

## Optional Prerequisites

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
| 6  | [JIRA](https://developer.atlassian.com/server/jira/platform/rest-apis/)            | `WTF_JIRA_USERNAME`            | You JIRA username             |
| 7  |                 | `WTF_JIRA_PASSWORD`            | Your JIRA password            |
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

Modules are the heart of WTF. Each one connects to a specific services
or renders data about a specific concept. The following modules are
included in WTF.

### BambooHR

Displays who's out on vacation or off sick today.

### Git

Displays the current branch, a list of changed files, and the last n
commits for a specified repository.

### Github

Displays your open pull requests and any code review requests assigned
to you.

### Google Calendar

Displays the next n calendar events.

### OpsGenie

Displays who's currently on call for all your schedules.

### Security

Displays whether or not your firewall is on and configured in 'stealth'
mode. Also displays the name of the current Wifi network and whether
or not it's encrypted.

### Weather

Displays the temperatures for the day, weather description, and the
sunrise and sunset times.
