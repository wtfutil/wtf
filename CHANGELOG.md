# Changelog

## Unreleased

### ⚡️ Added

* Jira widget navigable via up/down arrow keys, by [@jdenoy](https://github.com/jdenoy)
* Windows security module improved, by [E3V3A](https://github.com/E3V3A)

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
