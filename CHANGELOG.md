# Changelog

## Unreleased

## v0.42.0

### ⚡️ Added

* [FeedReader module](https://wtfutil.com/modules/feedreader/) now displays the name of the parent feed in front of the feed item, [#1138](https://github.com/wtfutil/wtf/issues/1138) by [@igbanam](https://github.com/igbanam)
* Steam Status module added by [@despreston](https://github.com/despreston)
* Working directory option for CmdRunner, [#1293](https://github.com/wtfutil/wtf/issues/1293) by [@Seanstoppable](https://github.com/Seanstoppable)
* Updated to Go 1.18, by [@Seanstoppable](https://github.com/Seanstoppable)
* "compact" option added to [Weather module](https://wtfutil.com/modules/weather_services/weather/) by [@Bizzaro](https://github.com/Bizzaro)
* Mempool cryptocurrency module by [@Bizzaro](https://github.com/Bizzaro)
* [Nextbus module](http://nextbus.com) added by [@Bizzaro](https://github.com/Bizzaro)
* URLCheck module added by [@pliski](https://github.com/pliski)
* Twitch modules added by [@zyrre](https://github.com/zyrre)

### 🐞 Fixed

* [Spotify](https://wtfutil.com/modules/spotify/) module now properly displays the artist list, by [@inetAnt](https://github.com/inetAnt)
* [Football](https://wtfutil.com/modules/sports/football/) module now uses HTTPS to connect to API, [#1183](https://github.com/wtfutil/wtf/issues/1183) by [@despreston](https://github.com/despreston)
* [Jira module](https://wtfutil.com/modules/jira/) now properly escapes brackets in issue summary, by [@Eraac](https://github.com/Eraac)
* [Pocket module](https://wtfutil.com/modules/pocket/) now properly handles empty lists, by [@asadali](https://github.com/asadali)
* Twitch module OAuth token support added, UI improved, and underlying library upgraded, by [@despreston](https://github.com/despreston)
* Typing "q" no longer quits the entire app every time, [#1257](https://github.com/wtfutil/wtf/issues/1257) by [@senorprogrammer](https://github.com/senorprogrammer)
* Row highlighting properly terminates the highlight color, [#1280](https://github.com/wtfutil/wtf/issues/1280) by [@Seanstoppable](https://github.com/Seanstoppable)
* Textfile widget properly displays text inside square brackets
* CmdRunner doesn't continue to keep around processes in pty mode [#1225](https://github.com/wtfutil/wtf/issues/1225) by [@Seanstoppable](https://github.com/Seanstoppable)
* Fix tag filtering in Todo module [#1172](https://github.com/wtfutil/wtf/issues/1225) by [@Seanstoppable](https://github.com/Seanstoppable)

## 0.41.0

### 🐞 Fixed

* [Kubernetes module](https://wtfutil.com/modules/kubernetes/) now uses a client per widget, rather than a global client, [#1081](https://github.com/wtfutil/wtf/issues/1081) by [@hypnoglow](https://github.com/hypnoglow)
* [FeedReader module]() now supports basic auth, [#957](https://github.com/wtfutil/wtf/issues/957) by [@ResamVi](https://github.com/ResamVi)
* Single words inside square brackets are displayed in the [CmdRunner module](https://wtfutil.com/modules/cmdrunner/), [#1142](https://github.com/wtfutil/wtf/issues/1142) by [@yosmoc](https://github.com/yosmoc)
* `text_widget.tpl` generator fixed and documentation improved, by [@Quentinchampenois](https://github.com/Quentinchampenois)

### 👍 Updated

* ExchangeRates module has been removed. Exchange rate data is supported by the [Yahoo Finance module](https://wtfutil.com/modules/stocks/yfinance/), by [@yosmoc](https://github.com/yosmoc)
* Replace usage of `ioutil` with `io` and `os` functions, by [@Juneezee](https://github.com/Juneezee)

## v0.40.0

### ⚡️ Added

* Numerous improvements to the display of the Git module, [#1102](https://github.com/wtfutil/wtf/pull/1102) by [@tadeaspaule](https://github.com/tadeaspaule)
* Updated to Go 1.17, by [nontw](https://github.com/nontw)
* [Airbrake module](https://github.com/wtfutil/wtf/tree/master/modules/airbrake) added, by [@kyrylo](https://github.com/kyrylo)
* Displays the battery level of connected devices, if the level can be read by `ioreg`, on MacOS, [#1108](https://github.com/wtfutil/wtf/issues/1108) by [@senorprogrammer](https://github.com/senorprogrammer)
* Adds an `args` setting to [IP API](https://wtfutil.com/modules/ipaddress/ipapi/) for defining which fields are displayed, by [@Choromanski](https://github.com/Choromanski)
* [Security module](https://wtfutil.com/modules/security/) no longer displays wifi info if wifi is missing or disabled, [#1124](https://github.com/wtfutil/wtf/issues/1124) by [@yosmoc](https://github.com/yosmoc)
* `q` now quits WTF. Added as a new global hotkey, by [@thuan1412](https://github.com/thuan1412)
* [DigitalOcean module](https://wtfutil.com/modules/digitalocean/) now properly displays the droplet region, [#1053](https://github.com/wtfutil/wtf/issues/1053) by [@yosmoc](https://github.com/yosmoc)
* [gCal module](https://wtfutil.com/modules/google/gcal/) now supports multiple configurations with different email addresses, [#1016](https://github.com/wtfutil/wtf/issues/1016) by [@Seanstoppable](https://github.com/Seanstoppable)
* All modules now support sub-1 second refresh intervals by specifying an optional time unit (ie: `50ms`, `20s`), [#1103](https://github.com/wtfutil/wtf/issues/1103) by [@ResamVi](https://github.com/ResamVi)

### 🐞 Fixed

* [GitHub module](https://wtfutil.com/modules/github/) no longer crashes when selecting different items across multiple repos, by [@Seanstoppable](https://github.com/Seanstoppable)
* `make lint` errors fixed, [#1127](https://github.com/wtfutil/wtf/issues/1127) by [@yosmoc](https://github.com/yosmoc)
* `jwt-go` replaced by `golang-jwt`, because security vulnerabilities, [#1121](https://github.com/wtfutil/wtf/issues/1121) by [@siddhant94](https://github.com/siddhant94)

## v0.39.2

Fixes a critical crashing bug that prevented WTF from launching on Intel Macs. The solution was to `go get -u golang.org/x/sys` to force a compatibility update with Go 1.17.

## v0.39.0

### ⚡️ Added

* [Updown module](https://github.com/wtfutil/wtf/tree/master/modules/updown) added, by [@hjr265](https://github.com/hjr265)
* Considerable improvements to the [Todo module](https://github.com/wtfutil/wtf/tree/master/modules/todo). See PR [#1100](https://github.com/wtfutil/wtf/pull/1100) for details. By [@tadeaspaule](https://github.com/tadeaspaule)

### 🐞 Fixed

* [Healthchecks module](https://github.com/wtfutil/wtf/tree/master/modules/healthchecks) now respects the `apiURL` setting in the config file, by [@hjr265](https://github.com/hjr265)
* Example FeedReader RSS feed in the default config changed to a working feed, [#1096](https://github.com/wtfutil/wtf/issues/1096) by [@senorprogrammer](https://github.com/senorprogrammer)

## 0.38.0

### 🐞 Fixed

* [FeedReader module](https://github.com/wtfutil/wtf/tree/master/modules/feedreader) handles line breaks in item titles gracefully, [#1088](https://github.com/wtfutil/wtf/issues/1088) by [@GibranHL0](https://github.com/GibranHL0)

## 0.37.0

### ⚡️ Added

* [Yahoo Finance](https://github.com/wtfutil/wtf/tree/master/modules/stocks/yfinance) module added, by [@inetAnt](https://github.com/inetAnt)

### 🐞 Fixed

* Fixed error closing the Jira module response body, [#1070](https://github.com/wtfutil/wtf/issues/1070) by [@senorprogrammer](https://github.com/senorprogrammer)

## v0.36.0

### ⚡️ Added

* [Healthchecks module](https://github.com/wtfutil/wtf/tree/master/modules/healthchecks) added, by [@stone](https://github.com/stone)
* [Asana module](https://github.com/wtfutil/wtf/tree/master/modules/asana) added, by [@xntrik](https://github.com/xntrik)
* [Kubernetes module](https://github.com/wtfutil/wtf/tree/master/modules/kubernetes) now displays the Ready Replicas and Total Replicas count, by [@LKaemmerling](https://github.com/LKaemmerling)
* Now builds using Go 1.16, which provides support for Apple's M1 ARM hardware, by [@LKaemmerling](https://github.com/LKaemmerling)

### 🐞 Fixed

* Numerous dangling response bodies closed, by [@dabcoder](https://github.com/dabcoder)
* HackerNews read-on-response error fixed, by [@LKaemmerling](https://github.com/LKaemmerling)
* CircleCI read-on-response error fixed, by [@dabcoder](https://github.com/dabcoder)

## v0.35.0

### ⚡️ Added

* CircleCI module now supports `numberOfBuilds` setting defining how many builds to display, [#1020](https://github.com/wtfutil/wtf/issues/1020) by [@dabcoder](https://github.com/dabcoder)
* The <kbd>\\</kbd> key now opens module documentation in the browser for the currently-focused module, by [@senorprogrammer](https://github.com/senorprogrammer)
* Krisinformation module added, by [@stone](https://github.com/stone)
* Covid module has been added, by [@dabcoder](https://github.com/dabcoder)
* Support for a "Language" configuration key added. Supports any BCP 47 tag for identifying a langauge, by [@senorprogrammer](https://github.com/senorprogrammer)

### 🐞 Fixed

* CryptoLive prices update again, [#1041](https://github.com/wtfutil/wtf/issues/1041) by [@zye1996](https://github.com/zye1996)

## v0.34.0
2020-11-08

### ⚡️ Added

* DigitalOcean module now supports custom column definitions in the default layout. See [the docs](https://wtfutil.com/modules/digitalocean/) for details, [#986](https://github.com/wtfutil/wtf/issues/986) by [@senorprogrammer](https://github.com/senorprogrammer)
* Finnhub stock quotes module added, [#930](https://github.com/wtfutil/wtf/issues/930) by [@applegreengrape](https://github.com/applegreengrape)
* Grafana Alerts module added, by [@schoentoon](https://github.com/schoentoon)

### 👍 Updated

* Updated `google.golang.org/api` from 0.30.0 to 0.33.0
* Updated `github.com/digitalocean/godo` from 1.46.0 to 1.52.0
* Updated `github.com/xanzy/go-gitlab` from 0.38.1 to 0.39.0
* Updated `github.com/zorkian/go-datadog-api` from 2.29.0+incompatible to 2.30.0+incompatible
* Updated `golang.org/x/text` from 0.3.3 to 0.3.4

## v0.33.0
2020-10-13

### ⚡️ Added

* FeedReader module supports [display mode switching](https://github.com/wtfutil/wtf/pull/976) via the `t` key, by [@cyingfan](https://github.com/cyingfan)
* Clocks module now supports `natural` ordering option, [#896](https://github.com/wtfutil/wtf/issues/896) by [@seanstoppable](https://github.com/Seanstoppable)
* Exit banner message added, by [@senorprogrammer](https://github.com/senorprogrammer)
* [UptimeRobot module added](https://github.com/wtfutil/wtf/pull/983), by [@franga2000](https://github.com/franga2000)
* Gitlab module now defaults the domain to `https://gitlab.com`, by [@jmks](https://github.com/jmks)

### 🐞 Fixed

* Exchange Rates module now displays rates in alphabetical order, sorted by the base rate, by [@senorprogrammer](https://github.com/senorprogrammer)
* Feed Reader module no longer crashes on feeds that don't have a published date, [#958](https://github.com/wtfutil/wtf/issues/958) by [@cyingfan](https://github.com/cyingfan)
* Stray blank lines no longer appear at the end of highlightable lists, [#977](https://github.com/wtfutil/wtf/pull/977) by [@cyingfan](https://github.com/cyingfan)
* `config.yml` now properly set to `0600` instead of `0666`, by [@cyingfan](https://github.com/cyingfan)
* Now blank-imports the `time/tzinfo` package to provide timezone support for Windows users, [#994](https://github.com/wtfutil/wtf/issues/994) by [@trimble](https://github.com/trimble)

### 👍 Updated

* Updated `github.com/gdamore/tcell` from 1.3.0 to 1.4.0
* Updated `github.com/shirou/gopsutil` from 2.20.7+incompatible to 2.20.9+incompatible
* Updated `github.com/mmcdole/gofeed` from 1.0.0 to 1.1.0
* Updated `github.com/digitalocean/godo` from 1.42.1 to 1.46.0
* Updated `github.com/xanzy/go-gitlab` from 0.33.0 to 0.38.1
* Updated `github.com/adlio/trello` from 1.7.0 to 1.8.0
* Updated `github.com/alecthomas/chroma` from 0.8.0 to 0.8.1
* Updated `github.com/nicklaw5/helix` from 0.5.9 to 0.7.0

* Now requires Go 1.15 for compilation in Dockerfiles and go.mod, by [@seanstoppable](https://github.com/Seanstoppable)

## v0.32.0

### ⚡️ Added

* PagerDuty module now supports `showOncallEnd` setting, which determines whether or not to display the oncall rotation end date, by [@senorprogrammer](https://github.com/senorprogrammer)
* Exchange Rates module now supports `precision` setting, which determines how many decimal places to display, [#897](https://github.com/wtfutil/wtf/issues/897) by [@senorprogrammer](https://github.com/senorprogrammer)

### 🐞 Fixed

* Todo module's edit modal now shows up quickly, [#876](https://github.com/wtfutil/wtf/issues/876) by [@aeter](https://github.com/aeter)
* Jira module title now displays properly, [#942](https://github.com/wtfutil/wtf/issues/942) by [@timhwang21](https://github.com/timhwang21)

### 👍 Updated

* Updated `alecthomas/chroma` from 0.7.3 to 0.8.0
* Updated `digitalocean/godo` from 1.38.0 to 1.42.1
* Updated `google.golang.org/api` from 0.29.0 to 0.30.0
* Updated `shirou/gopsutil` from 2.20.6+incompatible to 2.20.7+incompatible

## v0.31.0

### ⚡️ Added

* [docker-credential-helpers](https://github.com/docker/docker-credential-helpers) support added to provide optional securing of module secrets, [#517](https://github.com/wtfutil/wtf/issues/517) by [@sam-github](https://github.com/sam-github)
* `make lint` now uses [golangci-lint](https://github.com/golangci/golangci-lint-action), by [@sam-github](https://github.com/sam-github)
* GitHub module now supports per-section configuration, [#864](https://github.com/wtfutil/wtf/issues/864) by [@mryanmurphy](https://github.com/mryanmurphy)
* PagerDuty module now accepts Team ID and User ID as filters, by [@mryanmurphy](https://github.com/mryanmurphy)
* Weather module now supports a `useEmoji` setting, toggling the use of emoji characters in the title, by [@jonhadfield](https://github.com/jonhadfield)
* Pocket documentation added, by [@3mard](https://github.com/3mard)


### 🐞 Fixed

* GitLab Todo module now allows domain to be set, defaults to `https://gitlab.com`, by [@gerchardon](https://github.com/gerchardon)
* Pi-hole module now works with Pi-hole v4 and v5, by [@jonhadfield](https://github.com/jonhadfield)
* Null pointer exception fixed in VictorOps module, by [@ripienaar](https://github.com/ripienaar)

### 👍 Updated

* Updated `google.golang.org/api` from 0.25.0 to 0.29.0
* Updated `pkg/profile` from 1.4.0 to 1.5.0
* Updated `nicklaw5/helix` from 0.5.8 to 0.5.9
* Updated `digitalocean/godo` from 1.37.0 to 1.38.0
* Updated `gopkg.in/yaml.v2` from 2.2.8 to 2.3.0
* Updated `xanzy/go-gitlab` from 0.31.0 to 0.33.0
* Updated `stretchr/testify` from 1.5.1 to 1.6.1
* Updated `golang.org/x/text` from 0.3.2 to 0.3.3
* Updated `shirou/gopsutil` from 2.20.4+incompatible to 2.20.6+incompatible

## v0.30.0

### ⚡️ Added

* Expansive keyboard controls for the [GitLab](https://wtfutil.com/modules/gitlab/) module, by [@mogensen](https://github.com/mogensen)
* [IPInfo](https://wtfutil.com/modules/ipaddress/ipinfo/) module now supports IPv6 lookup, by [@acaloiaro](https://github.com/acaloiaro)
* The GitHub repo now runs a set of actions that lint PRs for code quality and conformity and comment in-line on the PRs. See here for action details https://github.com/wtfutil/wtf/actions

### 🐞 Fixed

* Random capital letters around display, [#811](https://github.com/wtfutil/wtf/issues/811) by [@jonhadfield](https://github.com/jonhadfield)
* [Feed Reader](https://wtfutil.com/modules/feedreader/) module now properly sorts multi-feed items by date, by [@rsaarelm](https://github.com/rsaarelm)

### 👍 Updated

* Updated `zorkian/go-datadog-api` from 2.28.0+incompatible to 2.29.0+incompatible
* Updated `shirou/gopsutil` from 2.20.3+incompatible to 2.20.4+incompatible
* Updated `google.golang.org/api` from 0.22.0 to 0.23.0

## 0.29.0

### ⚡️ Added

* gCal module now has a `showEndTime` boolean option for displaying meeting end times, by [@acaloiaro](https://github.com/acaloiaro)
* [Pi-hole](https://wtfutil.com/modules/pihole/) module added, by [@jonhadfield](https://github.com/jonhadfield)

### 🐞 Fixed

* Docker module subheading display, [#847](https://github.com/wtfutil/wtf/issues/847) by [@senorprogrammer](https://github.com/senorprogrammer)
* Improved display of currencies for the Exchange Rate module, by [@indradhanush](https://github.com/indradhanush)

### 👍 Updated

* Updated `nicklaw5/helix` from 0.5.7 to 0.5.8
* Updated `digitalocean/godo` from 1.34.0 to 1.35.1
* Updated `xanzy/go-gitlab` from 0.28.0 to 0.31.0
* Updated `shirou/gopsutil` from 2.20.2+incompatible to 2.20.3+incompatible
* Updated `alecthomas/chroma` from 0.7.1 to 0.7.2
* Updated `google.golang.org/api` from 0.21.0 to 0.22.0

## v0.28.0

### ⚡️ Added

* Support for customizing CPU, Mem, and Swap display in ResourceUsage, by [@leterio](https://github.com/leterio)

### 👍 Updated

* Now prefers Go 1.14 for compilation (should still work under 1.13 however)
* Updated `shirou/gopsutil` from 2.20.1+incompatible to 2.20.2+incompatible
* Updated `google.golang.org/api` from 0.17.0 to 0.20.0
* Updated `digitalocean/godo` from 1.30.0 to 1.32.0
* Updated `xanzy/go-gitlab` from 0.26.0 to 0.28.0
* Updated `adlio/trello` from 1.6.0 to 1.7.0
* Updated `zorkian/go-datadog-api` from 2.27.0+incompatible to 2.28.0+incompatible

## 0.27.0

### ⚡️ Added

* GitLab Todo module added, by [@elliotrushton](https://github.com/elliotrushton)
* [CDS](https://wtfutil.com/modules/cds/) module added, by [@yesnault](https://github.com/yesnault)

### 🐞 Fixed

* The `vendor` directory has been removed, [#792](https://github.com/wtfutil/wtf/issues/792) by [@bjoernw](https://github.com/bjoernw)

### 👍 Updated

* Updated `zorkian/go-datadog-api` from 2.26.0+incompatible to 2.27.0+incompatible
* Updated `google.golang.org/api` from 0.15.0 to 0.17.0
* Updated `github.com/nicklaw5/helix` from 0.5.5 to 0.5.7
* Updated `xanzy/go-gitlab` from 0.23.0 to 0.26.0
* Updated `stretchr/testify` from o.4.0 to 0.5.1

## 0.26.0

### ⚡️ Added

* `myName` config setting added to PagerDuty module, by [@senorprogrammer](https://github.com/senorprogrammer)
* `withDate` config setting added to Digital Clock module, by [@senorprogrammer](https://github.com/senorprogrammer)
* Twitch module added, by [@bjoernw](https://github.com/bjoernw)
* HackerNews module now opens HN comments when there is no alternative external link, [#758](https://github.com/wtfutil/wtf/issues/758) by [@senorprogrammer](https://github.com/senorprogrammer)
* gCal module now now allows users to hide all-day events, [#733](https://github.com/wtfutil/wtf/issues/733) by [@senorprogrammer](https://github.com/senorprogrammer)
* SpaceX module added, by [@bjoernw](https://github.com/bjoernw)
* Support for obeying `XDG_CONFIG_HOME` when set, [#699](https://github.com/wtfutil/wtf/issues/699) by [@Seanstoppable](https://github.com/Seanstoppable)

### 🐞 Fixed

* Module templating working again, [#748](https://github.com/wtfutil/wtf/issues/748) by [@senorprogrammer](https://github.com/senorprogrammer)
* CmdRunner title spacing issue fixed, [#784](https://github.com/wtfutil/wtf/issues/784) by [@senorprogrammer](https://github.com/senorprogrammer)
* Colors in cmdrunner fixed when using nodejs chalk et al., [#618](https://github.com/wtfutil/wtf/issues/618) by [@Seanstoppable](https://github.com/Seanstoppable)
* Docker buid instructions updated and improved, by [@firecat53](https://github.com/firecat53)
* Kubernetes module can now be used with multiple contexts, [#809](https://github.com/wtfutil/wtf/issues/809) by [@davidsbond](https://github.com/davidsbond)

### 👍 Updated

* Updated `digitalocean/godo` from 1.22.0 to 1.30.0
* Updated `google.golang.org/api` from 0.14.0 to 0.15.0
* Updated `alecthomas/chroma` from 0.7.0 to 0.7.1
* Updated `pkg/errors` from 0.8.1 to 0.9.1
* Updated `xanzy/go-gitlab` from 0.22.2 to 0.23.0
* Updated `shirou/gopsutil` from 2.19.11 to 2.20.1
* Updated `zorkian/go-datadog-api` from 2.25.0 to 2.26.0
* Updated `gopkg.in/yaml.v2` from 2.2.7 to 2.2.8
* Updated `nicklaw5/helix` from 0.5.4 to 0.5.5

## 0.25.0

### ⚡️ Added

* [DigitalOcean](https://wtfutil.com/modules/digitalocean/) module added, by [@senorprogrammer](https://github.com/senorprogrammer)
* [Transmission](https://wtfutil.com/modules/transmission/) module now supports a `hideComplete` configuration setting, by [@schoentoon](https://github.com/schoentoon)
* Pocket module added, [#742] by [@3mard](https://github.com/3mard)
* [Exchange Rates](https://wtfutil.com/modules/exchange_rates/) module added, by [@schoentoon](https://github.com/schoentoon)
* [GitHub](https://wtfutil.com/modules/github/) modules supports 'p' keyboard command to open **p**ull requests, by [@NickyMateev](https://github.com/NickyMateev)
* [GitHub](https://wtfutil.com/modules/github/) modules supports 'i' keyboard command to open **i**ssues, by [@NickyMateev](https://github.com/NickyMateev)
* [Jenkins](https://wtfutil.com/modules/jenkins/) module now supports multi-configuration projects, by [@NickyMateev](https://github.com/NickyMateev)

### 🐞 Fixed

* Subreddit out of bounds error fixed, [#753](https://github.com/wtfutil/wtf/issues/753) by [@TDHTTTT](https://github.com/TDHTTTT)
* Homebrew builds now contain version information, [#557](https://github.com/wtfutil/wtf/issues/557) by [@jottr](https://github.com/jottr)
* CmdRunner flicker problem, [#732](https://github.com/wtfutil/wtf/issues/732) by [@Gibstick](https://github.com/Gibstick)

### 👍 Updated

* Switched from `gocenter.io` as the Go proxy to `proxy.golang.org`, by [@chenrui333](https://github.com/chenrui333)
* Updated `go-datadog-api` to version 2.25.0+incompatible
* Updated `adlio/trello` to version 1.6.0
* Updated `alecthomas/chroma` to version 0.7.0
* Updated `olekukonko/tablewriter` to version 0.0.3
* Updated `pkg/profile` to version 1.4.0
* Updated `yaml.v2` to 2.2.7
* Updated `google.golang.org/api` to 0.14.0
* Updated `xanzy/go-gitlab` to 0.22.2
* Uodated `shirou/gopsutil` to 2.19.11+incompatible


## v0.24.0

### ⚡️ Added

* Proper, usable [Docker file](https://github.com/wtfutil/wtf/blob/master/Dockerfile) added, by [@Boot-Error](https://github.com/Boot-Error)
* [GitLab](https://wtfutil.com/modules/gitlab/) module displays issues assigned to, and opened by, the user, by [@caalberts](https://github.com/caalberts)
* [TravisCI](https://wtfutil.com/modules/travisci/) now checks for uncommitted vendor changes, by [@indradhanush](https://github.com/indradhanush)
* Football module added, by [@C123R](https://github.com/C123R)
* [resourceuseage](https://wtfutil.com/modules/resourceusage/) now supports a `cpuCombined` setting, by [@madepolli](https://github.com/madepolli)
* [Twitter Stats](https://wtfutil.com/modules/twitter/twittertweets/) module added, by [@Ameobea](https://github.com/Ameobea)

### 🐞 Fixed

* Github PRs do not count against issues, by [@alexfornuto](https://github.com/alexfornuto)
* Todo scrolling now works properly, [#707](https://github.com/wtfutil/wtf/issues/707) by [3mard](https://github.com/3mard)
* Configs with a missing `color` key now load properly, [#718](https://github.com/wtfutil/wtf/issues/718) and [#730](https://github.com/wtfutil/wtf/issues/730) by [@senorprogrammer](https://github.com/senorprogrammer)

## 0.23.0

### ⚡️ Added

* [Azure DevOps](https://wtfutil.com/modules/azure-devops/) module added, by [@v-braun](https://github.com/v-braun)
* [Dev.to](https://wtfutil.com/modules/devto/) module added, by [@VictorAvelar](https://github.com/VictorAvelar)
* [TravisCI]() module now supports enterprise endpoints, [#652](https://github.com/wtfutil/wtf/issues/652) by [@scw007](https://github.com/scw007)
* [Subreddit](https://wtfutil.com/modules/subreddit/) module added, by [@lawrencecraft](https://github.com/lawrencecraft)
* [gCal](https://wtfutil.com/modules/google/gcal/) module now supports a `hourFormat` setting for defining whether to display 12 or 24-hour times, [#665](https://github.com/wtfutil/wtf/issues/665) by [@senorprogrammer](https://github.com/senorprogrammer)
* [Scarf](https://scarf.sh) installation instructions added to README, by [@aviaviavi](https://github.com/aviaviavi)
* Spotify widget now supports colour themes, [#659](https://github.com/wtfutil/wtf/issues/659) by [@Tardog](https://github.com/Tardog)
* [Buildkite](https://wtfutil.com/modules/buildkite/) module added, by [@jmks](https://github.com/jmks)
* [Improvements](https://github.com/wtfutil/wtf/pull/680) to the [CmdRunner](https://wtfutil.com/modules/cmdrunner/) module, by [@noxer](https://github.com/noxer)

### 🐞 Fixed

* gCal calendar event time colour can now be changed by setting the `eventTime` configuration setting, [#638](https://github.com/wtfutil/wtf/issues/638) by [@indradhanush](https://github.com/indradhanush)
* [Clocks](https://wtfutil.com/modules/clocks/) now obeys global row colour settings, [#658](https://github.com/wtfutil/wtf/issues/658) by [@senorprogrammer](https://github.com/senorprogrammer)
* [Transmission](https://wtfutil.com/modules/transmission/) module no longer blocks rendering when a Transmission daemon cannot be found, [#661](https://github.com/wtfutil/wtf/issues/661) by [@senorprogrammer](https://github.com/senorprogrammer)
* [Trello](https://wtfutil.com/modules/trello/) module now respects project list order, [#664](https://github.com/wtfutil/wtf/issues/664) by [@Seanstoppable](https://github.com/Seanstoppable)
* [Todo](https://wtfutil.com/modules/todo/) module now respects checkbox settings, [#616](https://github.com/wtfutil/wtf/issues/616) by [@Seanstoppable](https://github.com/Seanstoppable)
* [Todoist](https://wtfutil.com/modules/todoist/) module now properly handles todo items with due date and times, [#645](https://github.com/wtfutil/wtf/issues/645) by [@massa1240](https://github.com/massa1240)
* Invalid pointer error in [Azure DevOps](https://wtfutil.com/modules/azure-devops/) fixed by [@Boot-Error](https://github.com/Boot-Error)
* Renamed slice error in [Dev](https://wtfutil.com/modules/devto/) fixed by [@Boot-Error](https://github.com/Boot-Error)

### 👍 Updated

* Updated `go-datadog-api` to version v2.24.0
* Updated `go-github` to version 26.13
* Updated `watcher` to version 1.0.7
* Updated `google-api-go-client` to version 0.10.0
* Updated `chroma` to version 0.6.7
* Updated `go-gitlab` to version 0.20.1
* Updated `trello` to version 1.4.0
* Updated `tcell` to version 1.3.0
* Updated `gopsutil` to version 2.19.9+incompatible
* Updated `yaml` to version 2.2.4

## v0.22.0

### ⚡️ Added

* [Arpansagovau](https://wtfutil.com/modules/weather_services/arpansagovau/) (arpansa.gov.au) module added, by [@jeffz](https://github.com/jeffz)
* 'calendarReadLevel' setting added to gCal module, by [@mikkeljuhl](https://github.com/mikkeljuhl)
* Todoist module now catches and displays API errors, by [@Seanstoppable](https://github.com/Seanstoppable)
* TravisCI sort orders now configurable,  by [@nyourchuck](https://github.com/nyourchuck)
* Google Analytics module now supports real-time metrics, [#581](https://github.com/wtfutil/wtf/issues/581) by [@Ameobea](https://github.com/Ameobea)
* Colors in configuration can now be defined using long-form hex, i.e.: #ff0000, by [@Seanstoppable](https://github.com/Seanstoppable)
* GitHub module pull requests are now selectable and openable via keyboard, [#547](https://github.com/wtfutil/wtf/issues/547) by [@Midnight-Conqueror](https://github.com/Midnight-Conqueror)
* [Docker](https://wtfutil.com/modules/docker/) module added, [#594](https://github.com/wtfutil/wtf/issues/594) by [@v-braun](https://github.com/v-braun)
* NewRelic module now supports displaying data from multiple apps, [#471](https://github.com/wtfutil/wtf/issues/471) by [@ChrisDBrown](https://github.com/ChrisDBrown) and [@Seanstoppable](https://github.com/Seanstoppable)
* [Digital Clock](https://wtfutil.com/modules/digitalclock/) module added, by [@Narengowda](https://github.com/Narengowda)

### 🐞 Fixed

* ScrollableWidget bounds checking error fixed, [#578](https://github.com/wtfutil/wtf/issues/578) by [@Seanstoppable](https://github.com/Seanstoppable)
* Now properly URL-decodes Jenkins branch names, [#575](https://github.com/wtfutil/wtf/issues/575) by [@lesteenman](https://github.com/lesteenman)
* Jira column sizes render properly, [#574](https://github.com/wtfutil/wtf/issues/574) by [@lesteenman](https://github.com/lesteenman)
* Todoist module updated to latest API version, by [@Seanstoppable](https://github.com/Seanstoppable)
* gCal colour highlighting working again, [#611](https://github.com/wtfutil/wtf/issues/611) by [@senorprogrammer](https://github.com/senorprogrammer)
* Per-module background and text colour settings working again, [#568](https://github.com/wtfutil/wtf/issues/568) by [@Seanstoppable](https://github.com/Seanstoppable)
* Git module no longer forces sorting of repositories, [#608](https://github.com/wtfutil/wtf/pull/608) by [@Seanstoppable](https://github.com/Seanstoppable)
* GitHub PR icons render properly without phantom characters, by [@Midnight-Conqueror](https://github.com/Midnight-Conqueror)
* GitLab configuration now takes a list of project paths, [#566](https://github.com/wtfutil/wtf/issues/566) by [@senorprogrammer](https://github.com/senorprogrammer)
* Kubernetes configuration segfault fixed, [#549](https://github.com/wtfutil/wtf/issues/549) by [@ibaum](https://github.com/ibaum)

## v0.21.0

### ⚡️ Added

* Power Soure module support added for FreeBSD, by [@hxw](https://github.com/hxw)

### 🐞 Fixed

* Power indicator displays ∞ on Linux when fully-charged and on AC power, [#534](https://github.com/wtfutil/wtf/issues/534) by [@Seanstoppable](https://github.com/Seanstoppable)
* Default background color is now the terminal background color, making transparency support possible in MacOS and Linux,  by [@Seanstoppable](https://github.com/Seanstoppable)
* `xdg-open` now used as the opener for HTTP/HTTPS by default, by [@hxw](https://github.com/hxw)
* Transmission port over-ride now working, [#565](https://github.com/wtfutil/wtf/issues/565) by [@Seanstoppable](https://github.com/Seanstoppable)
* Default config is now created on first run, [#553](https://github.com/wtfutil/wtf/issues/553) by [@senorprogrammer](https://github.com/senorprogrammer)

## v0.20.0

### ⚡️ Added

* Kubernetes module added, [#142](https://github.com/wtfutil/wtf/issues/142) by [@sudermanjr](https://github.com/sudermanjr)

### 🐞 Fixed

* Tab and Esc keys work properly in modal dialogs, [#520](https://github.com/wtfutil/wtf/issues/520) by [@Seanstoppable](https://github.com/Seanstoppable)
* `wtfutil -m` flag now works with non-enabled modules, [#529](https://github.com/wtfutil/wtf/issues/529) by [@Seanstoppable](https://github.com/Seanstoppable)
* Jenkins job filtering preserved across redraws, [#532](https://github.com/wtfutil/wtf/issues/532) by [@Seanstoppable](https://github.com/Seanstoppable)

## v0.19.1

### ⚡️ Added

* Dockerfile, by [@senorprogrammer](https://github.com/senorprogrammer)
* Add build targets for arm and arm64 architectures, by [@senorprogrammer](https://github.com/senorprogrammer)

## v0.19.0

### ☠️ Breaking Change

* HIBP module now requires an API key to operate. See [Authentication and the Have I Been Pwned API](https://www.troyhunt.com/authentication-and-the-have-i-been-pwned-api/) for more details, [#508](https://github.com/wtfutil/wtf/issues/508) by [@senorprogrammer](https://github.com/senorprogrammer)

### ⚡️ Added

* OpsGenie module now supports "region" configuration option ("us" or "eu"), by [@l13t](https://github.com/l13t)

### 🐞 Fixed

* Fixes the error message shown when an explicitly-specified custom config file cannot be found or cannot be read, by [@senorprogrammer](https://github.com/senorprogrammer)
* Rollbar module works again, [#507](https://github.com/wtfutil/wtf/issues/507) by [@Seanstoppable](https://github.com/Seanstoppable)
* The default config that gets installed on first run is much improved, [#504](https://github.com/wtfutil/wtf/issues/504) by [@senorprogrammer](https://github.com/senorprogrammer)
* Default config file is now `chmod 0600` to ensure only the owning user can read it, by [@senorprogrammer](https://github.com/senorprogrammer)

## v0.18.0

### ⚡️ Added

* Google Analytics module, by [@DylanBartels](https://github.com/DylanBartels)

### 🐞 Fixed

* Now created ~/.config if that directory is missing, [#510](https://github.com/wtfutil/wtf/issues/510) by [@senorprogrammer](https://github.com/senorprogrammer)

## v0.17.1

### 🐞 Fixed

* Fixes an issue in which the default config file was not being created on first run

## v0.17.0

### 🐞 Fixed

* FeedReader story sorting bug fixed
* NewRelic dependency vendored

## v0.16.1
## v0.16.0

### ⚡️ Added

* Config and recipe added for installing via Homebrew

## v0.15.0

### ❗️Changed

* The installed binary has been renamed from `wtf` to `wtfutil`. [Read more about it here](https://wtfutil.com/blog/2019-07-10-wtfutil-release/).

## v0.14.0

### ⚡️ Added

* CmdRunner module now supports custom titles, by [@Seanstoppable](https://github.com/Seanstoppable)
* FeedReader module added (https://wtfutil.com/modules/feedreader/), a rudimentary Rss & Atom feed reader

### 🐞 Fixed

* Cryptolive module works again, [#481](https://github.com/wtfutil/wtf/issues/481) by [@Seanstoppable](https://github.com/Seanstoppable)
* gCal module now supports setting an explicit time zone via the "timezone" config attribute, [#382](https://github.com/wtfutil/wtf/issues/382) by [@jeangovil](https://github.com/jeangovil)
* Misconfigured module positions in `config.yaml` now attempt to provide informative error messages on launch, [#482](https://github.com/wtfutil/wtf/issues/482)

## v0.13.0

### ⚡️ Added

* Transmission module addedd (https://wtfutil.com/modules/transmission/)

## v0.12.0

### ⚡️ Added

* Textfile module's text wrapping is configurable via the 'wrapText' boolean setting
* Have I Been Pwned (HIBP) module added (https://wtfutil.com/modules/hibp/)

## v0.11.0

### ⚡️ Added

* GitHub module now supports custom queries for issues and pull requests, by [@Sean-Der](https://github.com/Sean-Der)

### 🐞 Fixed

* Todoist now properly updates list items when Refresh is called
* Keyboard modal displays properly when tabbing between widgets, [#467](https://github.com/wtfutil/wtf/issues/467)

## v0.10.3

### ❗️Changed

* Invalid glog dependency removed, by [@bosr](https://github.com/bosr)

## v0.10.2

### 🐞 Fixed

* Weather module no longer crashes if there's no weather data or no internet connection
* Gitlab no longer prevents installing with missing param, [#455](https://github.com/wtfutil/wtf/issues/455)

## v0.10.1

### 🐞 Fixed

* Trello now displays multiple lists properly, [#454](https://github.com/wtfutil/wtf/issues/454)

## v0.10.0

### ⚡️ Added

* DataDog module is now scrollable and interactive, by [@Seanstoppable](https://github.com/Seanstoppable)
* Focusable hot key numbers are now assigned in a stable order, [#435](https://github.com/wtfutil/wtf/issues/435) by [@Seanstoppable](https://github.com/Seanstoppable)
* Zendesk widget now has help text, by [@Seanstoppable](https://github.com/Seanstoppable)
* Scrollable widget added to provide common keyboard-navigation list functionality, by [@Seanstoppable](https://github.com/Seanstoppable)
* Logging functionality extracted out from Log module, by [@Seanstoppable](https://github.com/Seanstoppable)
* Improved sample configs with up-to-date attributes and examples, by [@retgits](https://github.com/retgits)
* PagerDuty config now supports schedule filtering using the `scheduleIDs` config option, by [@senporprogrammer](https://github.com/senporprogrammer)

## v0.9.2

### ⚡️ Added

* Keyboard management system for modules reworked. Now has a KeyboardWidget to simplify adding keyboard commands

### Fixed

* WTF versions are now prefixed with `v` again so module systems pick up the latest versions

## 0.9.1

### ⚡️ Added

* Increased the pagination limit for GitHub to 100, by [@Seanstoppable](https://github.com/Seanstoppable)
* Support for multiple instances of the same widget added, by [@Seanstoppable](https://github.com/Seanstoppable)

## 0.9.0

* Null release

## 0.8.0

### ⚡️ Added

* Dependencies are now managed and installed using Go modules. See README.md for details, [#406](https://github.com/wtfutil/wtf/issues/406) by [@retgits](https://github.com/retgits)

## 0.7.2

### ⚡️ Added

* NBA Scores now navigable via arrow keys, [#415](https://github.com/wtfutil/wtf/issues/415)

### 🐞 Fixed

* Multi-page sigils off-by-one error fixed, [#413](https://github.com/wtfutil/wtf/issues/413)
* Many points of potential and probable race conditions have been improved to not have race conditions. WTF should be quite a bit more stable now
* In the Twitter module, the following have been fixed:
  * Help text says "Twitter" instead of "TextFile"
  * Keyboard-command "o" opens the current Twitter handle in the browser
  * Keyboard-command "o" is documented in the help text

## 0.7.1

### 🐞 Fixed

* HackerNews row selections are visible again, [#411](https://github.com/wtfutil/wtf/issues/411)

## 0.7.0

### ⚡️ Added

* Jenkins now supports coloured balls, [#358](https://github.com/wtfutil/wtf/issues/358) by [@rudolphjacksonm](https://github.com/rudolphjacksonm)
* Jenkins now supports regular expressions, [#359](https://github.com/wtfutil/wtf/issues/359) by [@rudolphjacksonm](https://github.com/rudolphjacksonm)
* Complete refactoring of the module settings system, reducing the dependency on `config` and making it possible to configure modules by other means, by [@senporprogrammer](https://github.com/senporprogrammer)

## 0.6.0

### ⚡️ Added

* Jira widget navigable via up/down arrow keys, by [@jdenoy](https://github.com/jdenoy)
* Windows security module improved, by [@E3V3A](https://github.com/E3V3A)
* Function modules moved into the `/modules` directory, by [@Seanstoppable](https://github.com/Seanstoppable)
* NBA Score module added by [@FriedCosey](https://github.com/FriedCosey)

### 🐞 Fixed

* Now displays an error on start-up if a widget has mis-configured `position` attributes ([#389](https://github.com/wtfutil/wtf/issues/389) by @senporprogrammer)

## 0.5.0

### ⚡️ Added

* Resource Usage module added by [@nicholas-eden](https://github.com/nicholas-eden)
* Recursive repo search in Git module ([#126](https://github.com/wtfutil/wtf/issues/126) by [@anandsudhir](http://github.com/anandsudhir))
* HTTP/HTTPS handling in OpenFile() util function by [@jdenoy](https://github.com/jdenoy)
* Honor system http proxies when using non-default transport by [@skymeyer](https://github.com/skymeyer)
* VictorOps module added by [ImDevinC](https://github.com/ImDevinC)
* Module templates added by [retgits](https://github.com/retgits)

## 0.4.0

### ⚡️ Added

* Mecurial module added ([@mweb](https://github.com/mweb))
* Can now define numeric hotkeys in config ([@mweb](https://github.com/mweb))
* Linux firewall support added ([@TheRedSpy15](https://github.com/TheRedSpy15))
* Spotify Web module added ([@StormFireFox1](https://github.com/StormFireFox1))

### 🐞 Fixed

* Google Calendar module now displays all-day events ([#306](https://github.com/wtfutil/wtf/issues/306) by [@nicholas-eden](https://github.com/nicholas-eden))
* Google Calendar configuration much improved ([#326](https://github.com/wtfutil/wtf/issues/326) by [@dvdmssmnn](https://github.com/dvdmssmnn))

## 0.3.0

### ⚡️ Added

* Spotify module added (@sticreations)
* Clocks module now supports configurable datetime formats (@danrabinowitz)
* Twitter module now supports subscribing to multiple screen names

### 🐞 Fixed

* Textfile module now watches files for changes ([#276](https://github.com/wtfutil/wtf/issues/276) by @senporprogrammer)
* Nav shortcuts now use numbers rather than letters to allow the use of letters in widget menus
* Twitter widget no longer crashes app when closing the help modal

## 0.2.2
#### Aug 25, 2018

### ⚡️ Added

* Twitter tweets are now colourized (@senorprogrammer)
* Twitter number of tweets to fetch is now customizable via config (@senorprogrammer)
* Google Calendar: widget is now focusable (@anandsudhir)
* [DataDog module](https://wtfutil.com/modules/datadog/) added (@Seanstoppable)

### 🐞 Fixed

* Textfile syntax highlighting now included in stand-alone binary ([#261](https://github.com/wtfutil/wtf/issues/261) by @senporprogrammer)
* Config param now supports relative paths starting with `~` ([#295](https://github.com/wtfutil/wtf/issues/295) by @anandsudhir)

## 0.2.1
#### Aug 17, 2018

### ⚡️ Added

* HackerNews widget is now scrollable (@anandsudhir)

### 🐞 Fixed

* Twitter screen name now configurable in configuration file (@senorprogrammer)
* Gerrit module no longer dies if it can't connect to the server (@anandsudhir)
* Pretty Weather properly displays colours again (([#298](https://github.com/wtfutil/wtf/issues/298) by @bertl4398)
* Clocks row colour configuration fixed (([#282](https://github.com/wtfutil/wtf/issues/282) by @anandsudhir)
* Sigils no longer display when there's only one option (([#291](https://github.com/wtfutil/wtf/issues/291) by @anandsudhir)
* Jira module now responds to the "/" key (([#268](https://github.com/wtfutil/wtf/issues/268)) by @senorprogrammer)

## 0.2.0
#### Aug 3, 2018

### ⚡️ Added

* [HackerNews module](https://wtfutil.com/modules/hackernews/) added (@anandsudhir)
* [Twitter module](https://wtfutil.com/modules/twitter/) added (@Trinergy)

### 🐞 Fixed

* TravisCI module now works with Pro version thanks to @ruggi
* Sensitive credentials can now be stored in config.yml instead of ENV vars
* GCal.showDeclined config added (@baustinanki)
* Gerrit widget is now interactive, added (@anandsudhir)

---

This file attempts to adhere to the principles of [keep a changelog](https://keepachangelog.com/en/1.0.0/).
