# Changelog

## Unreleased

### 🐞 Fixed

* FeedReader story sorting bug fixed

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
