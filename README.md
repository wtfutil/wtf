
[![All Contributors](https://img.shields.io/badge/all_contributors-87-orange.svg?style=flat-square)](#contributors)
[![Build Status](https://travis-ci.com/wtfutil/wtf.svg?branch=master)](https://travis-ci.com/wtfutil/wtf)
[![Twitter](https://img.shields.io/badge/follow-on%20twitter-blue.svg)](https://twitter.com/wtfutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/wtfutil/wtf)](https://goreportcard.com/report/github.com/wtfutil/wtf)

# WTF

A personal terminal-based dashboard utility, designed for
displaying infrequently-needed, but very important, daily data.

* [Screenshot](#screenshot)
* [Installing](#installing)
    * [Install via Homebrew](#install-via-homebrew)
    * [Install via MacPorts](#install-via-macports)
    * [Install a Binary](#install-a-binary)
    * [Install from Source](#install-from-source)
* [Communication](#communication)
    * [Slack](#slack)
    * [Twitter](#twitter)
* [Documentation](#documentation)
* [Modules](#modules)
* [Contributing to the Source Code](#contributing-to-the-source-code)
    * [Adding Dependencies](#adding-dependencies)
* [Contributing to the Documentation](#contributing-to-the-documentation)
* [Contributors](#contributors)
* [Acknowledgements](#acknowledgments)

## Screenshot

<p align="center">
<img src="./images/screenshot.jpg" title="screenshot" width="720" height="420" />
</p>

## Installing

### Install via Homebrew

The simplest way from Homebrew:

```console
brew install wtfutil

wtfutil
```

That version can sometimes lag a bit, as recipe updates take time to get accepted into `homebrew-core`. If you always want the bleeding edge of releases, you can tap it:

```console
brew tap wtfutil/wtfutil
brew install wtfutil

wtfutil
```

### Install via MacPorts

You can also install via [MacPorts](https://www.macports.org/):

```console
sudo port selfupdate
sudo port install wtfutil

wtfutil
```

### Install a Binary

[Download the latest binary](https://github.com/wtfutil/wtf/releases) from GitHub.

WTF is a stand-alone binary. Once downloaded, copy it to a location you can run executables from (ie: `/usr/local/bin/`), and set the permissions accordingly:

```bash
chmod a+x /usr/local/bin/wtfutil
```

and you should be good to go.

### Install from Source

If you want to run the build command from within your `$GOPATH`:

```bash
# Set the Go proxy variable to GoCenter
export GOPROXY="https://gocenter.io"

# Enable Go modules
export GO111MODULE=on

go get -u github.com/wtfutil/wtf
cd $GOPATH/src/github.com/wtfutil/wtf
make install
make run
```

If you want to run the build command from a folder that is not in your `$GOPATH`:

```bash
# Set the Go proxy variable to GoCenter
export GOPROXY="https://gocenter.io"

go get -u github.com/wtfutil/wtf
cd $GOPATH/src/github.com/wtfutil/wtf
make install
make run
```

**Note:** WTF is _only_ compatible with Go versions **1.11.0** or later (due to the use of Go modules). If you would like to use `gccgo` to compile, you _must_ use `gccgo-9` or later which introduces support for Go modules.

## Communication

### Slack

If you’re a member of the Gophers Slack community (https://invite.slack.golangbridge.org) there’s a WTFUtil channel you should join for all your WTF questions, development conversations, etc.

Find #wtfutil on https://gophers.slack.com/ and join us.

### Twitter

Also, follow [on Twitter](https://twitter.com/wtfutil) for news and latest updates. 

## Documentation

See [https://wtfutil.com](https://wtfutil.com) for the definitive
documentation. Here's some short-cuts:

* [Installation](https://wtfutil.com/getting_started/installation/)
* [Configuration](https://wtfutil.com/configuration/)
* [Module Documentation](https://wtfutil.com/modules/)

## Modules

Modules are the chunks of functionality that make WTF useful. Modules are added and configured by including their configuration values in your `config.yml` file. The documentation for each module describes how to configure them.

Some interesting modules you might consider adding to get you started:

* [GitHub](https://wtfutil.com/modules/github/)
* [Google Calendar](https://wtfutil.com/modules/google/gcal/)
* [HackerNews](https://wtfutil.com/modules/hackernews/)
* [Have I Been Pwned](https://wtfutil.com/modules/hibp/)
* [NewRelic](https://wtfutil.com/modules/newrelic/)
* [OpsGenie](https://wtfutil.com/modules/opsgenie/)
* [Security](https://wtfutil.com/modules/security/)
* [Transmission](https://wtfutil.com/modules/transmission/)
* [Trello](https://wtfutil.com/modules/trello/)

## Contributing to the Source Code

First, please read [Talk, then code](https://dave.cheney.net/2019/02/18/talk-then-code) by Dave Cheney. It's great advice and will often save a lot of time and effort. 

Next, please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests.

Then create your branch, write your code, submit your PR, and join the rest of the awesome people who've contributed their time and effort towards WTF. Without their contributors, WTF wouldn't be possible.

And don't worry if you've never written Go before, or never contributed to an open source project before, or that your code won't be good enough. For a surprising number of people WTF has been their first Go project, or first open source contribution. If you're here, and you've read this far, you're the right stuff.

## Contributing to the Documentation

Documentation now lives in its own repository here: [https://github.com/wtfutil/wtfdocs](https://github.com/wtfutil/wtfdocs).

Please make all additions and updates to documentation in that repository.

### Adding Dependencies

Dependency management in WTF is handled by [Go modules](https://github.com/golang/go/wiki/Modules). Please check out that page for more details on how Go modules work.

## Contributors

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore -->
| [<img src="https://avatars0.githubusercontent.com/u/6413?v=4" width="48px;" alt="Chris Cummer"/><br /><sub><b>Chris Cummer</b></sub>](https://twitter.com/senorprogrammer)<br /> | [<img src="https://avatars2.githubusercontent.com/u/3252403?v=4" width="48px;" alt="Anand Sudhir Prayaga"/><br /><sub><b>Anand Sudhir Prayaga</b></sub>](https://github.com/anandsudhir)<br /> | [<img src="https://avatars1.githubusercontent.com/u/34973359?v=4" width="48px;" alt="Hossein Mehrabi"/><br /><sub><b>Hossein Mehrabi</b></sub>](https://github.com/jeangovil)<br /> | [<img src="https://avatars0.githubusercontent.com/u/11779018?v=4" width="48px;" alt="FengYa"/><br /><sub><b>FengYa</b></sub>](https://github.com/Fengyalv)<br /> | [<img src="https://avatars2.githubusercontent.com/u/17337753?v=4" width="48px;" alt="deltax"/><br /><sub><b>deltax</b></sub>](https://fluxionnetwork.github.io/fluxion/)<br /> | [<img src="https://avatars0.githubusercontent.com/u/1319630?v=4" width="48px;" alt="Bill Keenan"/><br /><sub><b>Bill Keenan</b></sub>](https://github.com/BillKeenan)<br /> | [<img src="https://avatars2.githubusercontent.com/u/118081?v=4" width="48px;" alt="June S"/><br /><sub><b>June S</b></sub>](http://blog.sapara.com)<br /> |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| [<img src="https://avatars3.githubusercontent.com/u/16461061?v=4" width="48px;" alt="liyiheng"/><br /><sub><b>liyiheng</b></sub>](https://github.com/XanthusL)<br /> | [<img src="https://avatars2.githubusercontent.com/u/9014288?v=4" width="48px;" alt="baustinanki"/><br /><sub><b>baustinanki</b></sub>](https://github.com/baustinanki)<br /> | [<img src="https://avatars0.githubusercontent.com/u/371475?v=4" width="48px;" alt="lucus lee"/><br /><sub><b>lucus lee</b></sub>](https://github.com/lixin9311)<br /> | [<img src="https://avatars1.githubusercontent.com/u/7537841?v=4" width="48px;" alt="Mike Lloyd"/><br /><sub><b>Mike Lloyd</b></sub>](https://github.com/mxplusb)<br /> | [<img src="https://avatars3.githubusercontent.com/u/10998?v=4" width="48px;" alt="Sergio Rubio"/><br /><sub><b>Sergio Rubio</b></sub>](http://rubiojr.rbel.co)<br /> | [<img src="https://avatars3.githubusercontent.com/u/17374492?v=4" width="48px;" alt="Farhad Farahi"/><br /><sub><b>Farhad Farahi</b></sub>](https://github.com/FarhadF)<br /> | [<img src="https://avatars1.githubusercontent.com/u/634604?v=4" width="48px;" alt="Lasantha Kularatne"/><br /><sub><b>Lasantha Kularatne</b></sub>](http://lasantha.blogspot.com/)<br /> |
| [<img src="https://avatars1.githubusercontent.com/u/823331?v=4" width="48px;" alt="Mark Old"/><br /><sub><b>Mark Old</b></sub>](https://github.com/dlom)<br /> | [<img src="https://avatars0.githubusercontent.com/u/5546718?v=4" width="48px;" alt="flw"/><br /><sub><b>flw</b></sub>](http://flw.tools/)<br /> | [<img src="https://avatars0.githubusercontent.com/u/6024927?v=4" width="48px;" alt="David Barda"/><br /><sub><b>David Barda</b></sub>](https://github.com/davebarda)<br /> | [<img src="https://avatars2.githubusercontent.com/u/4261980?v=4" width="48px;" alt="Geoff Lee"/><br /><sub><b>Geoff Lee</b></sub>](https://github.com/matrinox)<br /> | [<img src="https://avatars3.githubusercontent.com/u/1022918?v=4" width="48px;" alt="George Opritescu"/><br /><sub><b>George Opritescu</b></sub>](http://international.github.io)<br /> | [<img src="https://avatars3.githubusercontent.com/u/497310?v=4" width="48px;" alt="Grazfather"/><br /><sub><b>Grazfather</b></sub>](https://twitter.com/Grazfather)<br /> | [<img src="https://avatars2.githubusercontent.com/u/1691120?v=4" width="48px;" alt="Michael Cordell"/><br /><sub><b>Michael Cordell</b></sub>](http://www.mikecordell.com/)<br /> |
| [<img src="https://avatars2.githubusercontent.com/u/1215497?v=4" width="48px;" alt="Patrick José Pereira"/><br /><sub><b>Patrick José Pereira</b></sub>](http://patrick.ibexcps.com)<br /> | [<img src="https://avatars2.githubusercontent.com/u/1483092?v=4" width="48px;" alt="sherod taylor"/><br /><sub><b>sherod taylor</b></sub>](https://github.com/sherodtaylor)<br /> | [<img src="https://avatars2.githubusercontent.com/u/3062663?v=4" width="48px;" alt="Andrew Scott"/><br /><sub><b>Andrew Scott</b></sub>](http://cogentia.io)<br /> | [<img src="https://avatars1.githubusercontent.com/u/12018440?v=4" width="48px;" alt="Lassi Piironen"/><br /><sub><b>Lassi Piironen</b></sub>](https://github.com/lsipii)<br /> | [<img src="https://avatars0.githubusercontent.com/u/14799210?v=4" width="48px;" alt="BlackWebWolf"/><br /><sub><b>BlackWebWolf</b></sub>](https://github.com/BlackWebWolf)<br /> | [<img src="https://avatars0.githubusercontent.com/u/1894885?v=4" width="48px;" alt="andrewzolotukhin"/><br /><sub><b>andrewzolotukhin</b></sub>](https://github.com/andrewzolotukhin)<br /> | [<img src="https://avatars1.githubusercontent.com/u/8568280?v=4" width="48px;" alt="Leon Stigter"/><br /><sub><b>Leon Stigter</b></sub>](https://retgits.github.io)<br /> |
| [<img src="https://avatars3.githubusercontent.com/u/21756?v=4" width="48px;" alt="Amr Tamimi"/><br /><sub><b>Amr Tamimi</b></sub>](https://tamimi.se)<br /> | [<img src="https://avatars3.githubusercontent.com/u/3717137?v=4" width="48px;" alt="Jagdeep Singh"/><br /><sub><b>Jagdeep Singh</b></sub>](https://jagdeep.me)<br /> | [<img src="https://avatars0.githubusercontent.com/u/889171?v=4" width="48px;" alt="Lineu Felipe"/><br /><sub><b>Lineu Felipe</b></sub>](https://github.com/darkSasori)<br /> | [<img src="https://avatars2.githubusercontent.com/u/159124?v=4" width="48px;" alt="Konstantin"/><br /><sub><b>Konstantin</b></sub>](https://github.com/kvj)<br /> | [<img src="https://avatars2.githubusercontent.com/u/6044920?v=4" width="48px;" alt="Brendan O'Leary"/><br /><sub><b>Brendan O'Leary</b></sub>](http://www.brendanoleary.com)<br /> | [<img src="https://avatars2.githubusercontent.com/u/1226441?v=4" width="48px;" alt="bertl4398"/><br /><sub><b>bertl4398</b></sub>](https://github.com/bertl4398)<br /> | [<img src="https://avatars2.githubusercontent.com/u/6553695?v=4" width="48px;" alt="Ferenc-"/><br /><sub><b>Ferenc-</b></sub>](https://github.com/Ferenc-)<br /> |
| [<img src="https://avatars1.githubusercontent.com/u/952036?v=4" width="48px;" alt="Rohan Verma"/><br /><sub><b>Rohan Verma</b></sub>](http://rohanverma.net)<br /> | [<img src="https://avatars1.githubusercontent.com/u/19293566?v=4" width="48px;" alt="Tim Fitzgerald"/><br /><sub><b>Tim Fitzgerald</b></sub>](https://github.com/fimtitzgerald)<br /> | [<img src="https://avatars2.githubusercontent.com/u/1081051?v=4" width="48px;" alt="Federico Ruggi"/><br /><sub><b>Federico Ruggi</b></sub>](https://github.com/ruggi)<br /> | [<img src="https://avatars2.githubusercontent.com/u/7293328?v=4" width="48px;" alt="Craig Woodward"/><br /><sub><b>Craig Woodward</b></sub>](https://github.com/ctwoodward)<br /> | [<img src="https://avatars3.githubusercontent.com/u/15367484?v=4" width="48px;" alt="ReadmeCritic"/><br /><sub><b>ReadmeCritic</b></sub>](https://twitter.com/ReadmeCritic)<br /> | [<img src="https://avatars0.githubusercontent.com/u/141402?v=4" width="48px;" alt="Eugene"/><br /><sub><b>Eugene</b></sub>](https://github.com/jdevelop)<br /> | [<img src="https://avatars1.githubusercontent.com/u/12983705?s=460&v=4" width="48px;" alt="Kenny Wu"/><br /><sub><b>Kenny Wu</b></sub>](https://github.com/Trinergy)<br /> |
| [<img src="https://avatars0.githubusercontent.com/u/538234?v=4" width="48px;" alt="Renán Romero"/><br /><sub><b>Renán Romero</b></sub>](http://www.romeroruiz.com)<br /> | [<img src="https://avatars1.githubusercontent.com/u/5031240?v=4" width="48px;" alt="Bastian Groß"/><br /><sub><b>Bastian Groß</b></sub>](https://github.com/sticreations)<br /> | [<img src="https://avatars1.githubusercontent.com/u/2496835?v=4" width="48px;" alt="nicholas-eden"/><br /><sub><b>nicholas-eden</b></sub>](https://github.com/nicholas-eden)<br /> | [<img src="https://avatars1.githubusercontent.com/u/279390?v=4" width="48px;" alt="Dan Rabinowitz"/><br /><sub><b>Dan Rabinowitz</b></sub>](https://github.com/danrabinowitz)<br /> | [<img src="https://avatars1.githubusercontent.com/u/6897575?v=4" width="48px;" alt="David Missmann"/><br /><sub><b>David Missmann</b></sub>](https://github.com/dvdmssmnn)<br /> | [<img src="https://avatars2.githubusercontent.com/u/882006?v=4" width="48px;" alt="Mathias Weber"/><br /><sub><b>Mathias Weber</b></sub>](https://github.com/mweb)<br /> | [<img src="https://avatars1.githubusercontent.com/u/32081703?v=4" width="48px;" alt="TheRedSpy15"/><br /><sub><b>TheRedSpy15</b></sub>](https://github.com/TheRedSpy15)<br /> |
| [<img src="https://avatars0.githubusercontent.com/u/9569897?v=4" width="48px;" alt="Harald Nordgren"/><br /><sub><b>Harald Nordgren</b></sub>](https://www.linkedin.com/in/harald-nordgren-44778192)<br /> | [<img src="https://avatars0.githubusercontent.com/u/11583824?v=4" width="48px;" alt="Matei Alexandru Gardus"/><br /><sub><b>Matei Alexandru Gardus</b></sub>](http://stormfirefox1.github.io)<br /> | [<img src="https://avatars2.githubusercontent.com/u/1523955?v=4" width="48px;" alt="Sean Smith"/><br /><sub><b>Sean Smith</b></sub>](https://github.com/Seanstoppable)<br /> | [<img src="https://avatars1.githubusercontent.com/u/1646238?v=4" width="48px;" alt="Halil Kaskavalci"/><br /><sub><b>Halil Kaskavalci</b></sub>](http://kaskavalci.com)<br /> | [<img src="https://avatars2.githubusercontent.com/u/246715?v=4" width="48px;" alt="Johan Denoyer"/><br /><sub><b>Johan Denoyer</b></sub>](http://www.johandenoyer.fr)<br /> | [<img src="https://avatars1.githubusercontent.com/u/593516?v=4" width="48px;" alt="Jelle Vink"/><br /><sub><b>Jelle Vink</b></sub>](https://skymeyer.be)<br /> | [<img src="https://avatars1.githubusercontent.com/u/3997333?v=4" width="48px;" alt="Devin Collins"/><br /><sub><b>Devin Collins</b></sub>](http://imdevinc.com)<br /> |
| [<img src="https://avatars3.githubusercontent.com/u/18333?v=4" width="48px;" alt="Danne Stayskal"/><br /><sub><b>Danne Stayskal</b></sub>](http://danne.stayskal.com/)<br /> | [<img src="https://avatars1.githubusercontent.com/u/2006658?v=4" width="48px;" alt="Max Beizer"/><br /><sub><b>Max Beizer</b></sub>](https://www.maxbeizer.com)<br /> | [<img src="https://avatars1.githubusercontent.com/u/194392?v=4" width="48px;" alt="E:V:A"/><br /><sub><b>E:V:A</b></sub>](http://tinyurl.com/nwmj4as)<br /> | [<img src="https://avatars0.githubusercontent.com/u/1425500?v=4" width="48px;" alt="Gabriel"/><br /><sub><b>Gabriel</b></sub>](https://github.com/GaboFDC)<br /> | [<img src="https://avatars2.githubusercontent.com/u/10111411?v=4" width="48px;" alt="Andrew Scibek"/><br /><sub><b>Andrew Scibek</b></sub>](https://github.com/AndrewScibek)<br /> | [<img src="https://avatars0.githubusercontent.com/u/29709822?v=4" width="48px;" alt="FriedCosey"/><br /><sub><b>FriedCosey</b></sub>](https://github.com/FriedCosey)<br /> | [<img src="https://avatars1.githubusercontent.com/u/3891?v=4" width="48px;" alt="Michele Gerarduzzi"/><br /><sub><b>Michele Gerarduzzi</b></sub>](https://michelegera.dev/)<br /> |
| [<img src="https://avatars3.githubusercontent.com/u/13438569?v=4" width="48px;" alt="Jack Morris"/><br /><sub><b>Jack Morris</b></sub>](https://github.com/rudolphjacksonm)<br /> | [<img src="https://avatars0.githubusercontent.com/u/14993807?v=4" width="48px;" alt="foorb"/><br /><sub><b>foorb</b></sub>](https://github.com/foorb)<br /> | [<img src="https://avatars0.githubusercontent.com/u/5819098?v=4" width="48px;" alt="Levi Baber"/><br /><sub><b>Levi Baber</b></sub>](http://researchit.las.iastate.edu)<br /> | [<img src="https://avatars0.githubusercontent.com/u/38514?v=4" width="48px;" alt="Graham Anderson"/><br /><sub><b>Graham Anderson</b></sub>](https://github.com/gnanderson)<br /> | [<img src="https://avatars2.githubusercontent.com/u/1936828?v=4" width="48px;" alt="Romain Bossart"/><br /><sub><b>Romain Bossart</b></sub>](https://github.com/bosr)<br /> | [<img src="https://avatars0.githubusercontent.com/u/969838?v=4" width="48px;" alt="Kirill Motkov"/><br /><sub><b>Kirill Motkov</b></sub>](http://eonix.ru)<br /> | [<img src="https://avatars1.githubusercontent.com/u/3665694?v=4" width="48px;" alt="Brian Choromanski"/><br /><sub><b>Brian Choromanski</b></sub>](http://www.BrianChoromanski.com)<br /> |
| [<img src="https://avatars0.githubusercontent.com/u/1302304?v=4" width="48px;" alt="Sean DuBois"/><br /><sub><b>Sean DuBois</b></sub>](http://siobud.com)<br /> | [<img src="https://avatars1.githubusercontent.com/u/47195730?v=4" width="48px;" alt="Gary Kim"/><br /><sub><b>Gary Kim</b></sub>](https://github.com/gary-kim)<br /> | [<img src="https://avatars1.githubusercontent.com/u/6660171?v=4" width="48px;" alt="Dylan"/><br /><sub><b>Dylan</b></sub>](https://dylanbartels.com)<br /> | [<img src="https://avatars0.githubusercontent.com/u/1990354?v=4" width="48px;" alt="Dmytro Prokhorenkov"/><br /><sub><b>Dmytro Prokhorenkov</b></sub>](http://liet.me)<br /> | [<img src="https://avatars1.githubusercontent.com/u/590442?v=4" width="48px;" alt="Elliot"/><br /><sub><b>Elliot</b></sub>](https://github.com/elliotrushton)<br /> | [<img src="https://avatars3.githubusercontent.com/u/1580956?v=4" width="48px;" alt="chenrui"/><br /><sub><b>chenrui</b></sub>](http://chenrui.dev)<br /> | [<img src="https://avatars0.githubusercontent.com/u/7624765?v=4" width="48px;" alt="Andrew Suderman"/><br /><sub><b>Andrew Suderman</b></sub>](https://github.com/sudermanjr)<br /> |
| [<img src="https://avatars3.githubusercontent.com/u/2373856?v=4" width="48px;" alt="Bob 'Wombat' Hogg"/><br /><sub><b>Bob 'Wombat' Hogg</b></sub>](https://github.com/rwhogg)<br /> | [<img src="https://avatars0.githubusercontent.com/u/143462?v=4" width="48px;" alt="Christopher Hall"/><br /><sub><b>Christopher Hall</b></sub>](https://github.com/hxw)<br /> | [<img src="https://avatars1.githubusercontent.com/u/3451557?v=4" width="48px;" alt="Heitor Neiva"/><br /><sub><b>Heitor Neiva</b></sub>](https://github.com/hneiva)<br /> | [<img src="https://avatars3.githubusercontent.com/u/618376?v=4" width="48px;" alt="Herby Gillot"/><br /><sub><b>Herby Gillot</b></sub>](https://github.com/herbygillot)<br /> | [<img src="https://avatars3.githubusercontent.com/u/382352?v=4" width="48px;" alt="James Canning"/><br /><sub><b>James Canning</b></sub>](http://brudil.com)<br /> | [<img src="https://avatars1.githubusercontent.com/u/45892?v=4" width="48px;" alt="jeffz"/><br /><sub><b>jeffz</b></sub>](https://twitter.com/jeffz4000)<br /> | [<img src="https://avatars0.githubusercontent.com/u/1764035?v=4" width="48px;" alt="Mikkel Jeppesen Juhl"/><br /><sub><b>Mikkel Jeppesen Juhl</b></sub>](https://mikkeljuhl.com)<br /> |
| [<img src="https://avatars1.githubusercontent.com/u/963290?v=4" width="48px;" alt="Erik"/><br /><sub><b>Erik</b></sub>](https://github.com/lesteenman)<br /> | [<img src="https://avatars1.githubusercontent.com/u/155574?v=4" width="48px;" alt="Nate Yourchuck"/><br /><sub><b>Nate Yourchuck</b></sub>](https://github.com/nyourchuck)<br /> | [<img src="https://avatars3.githubusercontent.com/u/4335849?v=4" width="48px;" alt="Casey Primozic"/><br /><sub><b>Casey Primozic</b></sub>](https://cprimozic.net/)<br /> |
<!-- ALL-CONTRIBUTORS-LIST:END -->

## Acknowledgments

The inspiration for `WTF` came from Monica Dinculescu's
[tiny-care-terminal](https://github.com/notwaldorf/tiny-care-terminal).

WTF is built atop [tcell](https://github.com/gdamore/tcell) and [tview](https://github.com/rivo/tview), fantastic projects both.

Many thanks to <a href="https://lendesk.com">Lendesk</a> for supporting this project by
providing time to develop it.

<p align="center">
<img src="./images/dude_wtf.png?raw=true" title="Dude WTF" width="251" height="201" />
</p>
