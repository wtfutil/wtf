# WTF

WTF is a personal terminal-based dashboard utility, designed for
displaying infrequently-updating, but very important, daily data.

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
