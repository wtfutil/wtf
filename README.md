<p align="center">
    <img src="./images/logo_transparent.png?raw=true" title="WTF" alt="WTF" width="560" height="560" />
</p>

<p align="left">
    <a href="https://goreportcard.com/report/github.com/wtfutil/wtf"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/wtfutil/wtf"></a>
    <a href="https://twitter.com/wtfutil"><img alt="Twitter" src="https://img.shields.io/badge/follow-on%20twitter-blue.svg"></a>
    <a href="https://deepsource.io/gh/wtfutil/wtf/?ref=repository-badge}" target="_blank"><img alt="DeepSource" title="DeepSource" src="https://deepsource.io/gh/wtfutil/wtf.svg/?label=active+issues&show_trend=true&token=kSJAbELF2TA7rEHjK6RPUrj5"/></a>
</p>

WTF (aka 'wtfutil') is the personal information dashboard for your terminal, providing at-a-glance access to your very important but infrequently-needed stats and data.

Used by thousands of developers and tech people around the world, WTF is free and open-source. To support the continued use and development of WTF, please consider sponsoring WTF via [GitHub Sponsors](https://github.com/sponsors/senorprogrammer).

### Are you a contributor or sponsor?

Awesome! [See here](https://wtfutil.com/sponsors/exit_message/) for how you can change the exit message, the message WTF shows when quitting, to something special just for you.

## Sponsored by

<p>
	<a href="https://airbrake.io/?utm_medium=sponsor&utm_source=WTFutill&utm_content=airbrake-home-page&utm_campaign=2021-sponsorships" target=_blank>
		<img src="./images/sponsors/airbrake.png?raw=true" height="60" title="Airbrake" alt="Airbrake" />
	</a>
</p>

<hr />

<p></p>

* [Installation](#installation)
    * [Installing via Homebrew](#installing-via-homebrew)
    * [Installing via MacPorts](#installing-via-macports)
    * [Installing a Binary](#installing-a-binary)
    * [Installing from Source](#installing-from-source)
    * [Running via Docker](#running-via-docker)
* [Communication](#communication)
    * [GitHub Discussions](#github-discussions)
    * [Twitter](#twitter)
* [Documentation](#documentation)
* [Modules](#modules)
* [Getting Bugs Fixed or Features Added](#getting-bugs-fixed-or-features-added)
* [Contributing to the Source Code](#contributing-to-the-source-code)
    * [Adding Dependencies](#adding-dependencies)
* [Contributing to the Documentation](#contributing-to-the-documentation)
* [Contributors](#contributors)
* [Acknowledgements](#acknowledgments)

<p align="center">
<img src="./images/screenshot.jpg" title="screenshot" width="720" height="420" />
</p>

## Installation

### Installing via Homebrew

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

### Installing via MacPorts

You can also install via [MacPorts](https://www.macports.org/):

```console
sudo port selfupdate
sudo port install wtfutil

wtfutil
```

### Installing a Binary

[Download the latest binary](https://github.com/wtfutil/wtf/releases) from GitHub.

WTF is a stand-alone binary. Once downloaded, copy it to a location you can run executables from (ie: `/usr/local/bin/`), and set the permissions accordingly:

```bash
chmod a+x /usr/local/bin/wtfutil
```

and you should be good to go.

### Installing from Source

If you want to run the build command from within your `$GOPATH`:

```bash
# Set the Go proxy
export GOPROXY="https://proxy.golang.org,direct"

# Disable the Go checksum database
export GOSUMDB=off

# Enable Go modules
export GO111MODULE=on

go get -u github.com/wtfutil/wtf
cd $GOPATH/src/github.com/wtfutil/wtf
make install
make run
```

If you want to run the build command from a folder that is not in your `$GOPATH`:

```bash
# Set the Go proxy
export GOPROXY="https://proxy.golang.org,direct"

go get -u github.com/wtfutil/wtf
cd $GOPATH/src/github.com/wtfutil/wtf
make install
make run
```

### Installing from Source using Docker

All building is done inside a docker container. You can then copy the binary to
your local machine.

```bash
curl -o Dockerfile.build https://raw.githubusercontent.com/wtfutil/wtf/master/Dockerfile.build
docker build -f Dockerfile.build -t wtfutil --build-arg=version=master .
docker create --name wtf_build wtfutil
docker cp wtf_build:/usr/local/bin/wtfutil ~/.local/bin
docker rm wtf_build
```

**Note:** WTF is _only_ compatible with Go versions **1.16.0** or later (due to the use of Go modules and newer standard library functions). If you would like to use `gccgo` to compile, you _must_ use `gccgo-9` or later which introduces support for Go modules.

### Installing via Arch User Repository

Arch Linux users can utilise the [wtfutil](https://aur.archlinux.org/packages/wtfutil) package to build it from source, or [wtfutil-bin](https://aur.archlinux.org/packages/wtfutil-bin/) to install pre-built binaries.

## Running via Docker

You can run `wtf` inside a docker container:

```bash
# download or create the Dockerfile
curl -o Dockerfile https://raw.githubusercontent.com/wtfutil/wtf/master/Dockerfile

# build the docker container
docker build -t wtfutil .

# or for a particular tag or branch
docker build --build-arg=version=v0.25.0 -t wtfutil .

# run the container
docker run -it wtfutil

# run container with a local config file
docker run -it -v path/to/config.yml:/config/config.yml wtfutil --config=/config/config.yml
```

## Communication

### GitHub Discussions

Conversations, ideas, discussions are done on [GitHub Discussions](https://github.com/wtfutil/wtf/discussions).

Formerly they were on Slack; that channel has been deprecated.

### Twitter

Also, follow [on Twitter](https://twitter.com/wtfutil) for news and latest updates.

## Documentation

See [https://wtfutil.com](https://wtfutil.com) for the definitive
documentation. Here's some short-cuts:

* [Installation](https://wtfutil.com/quick_start/)
* [Configuration](https://wtfutil.com/configuration/files/)
* [Module Documentation](https://wtfutil.com/modules/)

## Modules

Modules are the chunks of functionality that make WTF useful. Modules are added and configured by including their configuration values in your `config.yml` file. The documentation for each module describes how to configure them.

Some interesting modules you might consider adding to get you started:

* [DigitalOcean](https://wtfutil.com/modules/digitalocean/)
* [GitHub](https://wtfutil.com/modules/github/)
* [Google Calendar](https://wtfutil.com/modules/google/gcal/)
* [HackerNews](https://wtfutil.com/modules/hackernews/)
* [Have I Been Pwned](https://wtfutil.com/modules/hibp/)
* [NewRelic](https://wtfutil.com/modules/newrelic/)
* [OpsGenie](https://wtfutil.com/modules/opsgenie/)
* [Security](https://wtfutil.com/modules/security/)
* [Transmission](https://wtfutil.com/modules/transmission/)
* [Trello](https://wtfutil.com/modules/trello/)

## Getting Bugs Fixed or Features Added

WTF is open-source software, informally maintained by a small collection of volunteers who come and go at their leisure. There are absolutely no guarantees that, even if an issue is opened for them, bugs will be fixed or features added.

If there is a bug that you really need to have fixed or a feature you really want to have implemented, you can greatly increase your chances of that happening by creating a bounty on [BountySource](https://www.bountysource.com) to provide an incentive for someone to tackle it.

## Contributing to the Source Code

First, kindly read [Talk, then code](https://dave.cheney.net/2019/02/18/talk-then-code) by Dave Cheney. It's great advice and will often save a lot of time and effort.

Next, kindly read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests.

Then create your branch, write your code, submit your PR, and join the rest of the awesome people who've contributed their time and effort towards WTF. Without their contributors, WTF wouldn't be possible.

Don't worry if you've never written Go before, or never contributed to an open source project before, or that your code won't be good enough. For a surprising number of people WTF has been their first Go project, or first open source contribution. If you're here, and you've read this far, you're the right stuff.

## Contributing to the Documentation

Documentation now lives in its own repository here: [https://github.com/wtfutil/wtfdocs](https://github.com/wtfutil/wtfdocs).

Please make all additions and updates to documentation in that repository.

### Adding Dependencies

Dependency management in WTF is handled by [Go modules](https://github.com/golang/go/wiki/Modules). Please check out that page for more details on how Go modules work.

## Contributors

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://twitter.com/senorprogrammer"><img src="https://avatars0.githubusercontent.com/u/6413?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Chris Cummer</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/anandsudhir"><img src="https://avatars2.githubusercontent.com/u/3252403?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Anand Sudhir Prayaga</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/jeangovil"><img src="https://avatars1.githubusercontent.com/u/34973359?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Hossein Mehrabi</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/Fengyalv"><img src="https://avatars0.githubusercontent.com/u/11779018?v=4?s=48" width="48px;" alt=""/><br /><sub><b>FengYa</b></sub></a><br /></td>
    <td align="center"><a href="https://fluxionnetwork.github.io/fluxion/"><img src="https://avatars2.githubusercontent.com/u/17337753?v=4?s=48" width="48px;" alt=""/><br /><sub><b>deltax</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/BillKeenan"><img src="https://avatars0.githubusercontent.com/u/1319630?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Bill Keenan</b></sub></a><br /></td>
    <td align="center"><a href="http://blog.sapara.com"><img src="https://avatars2.githubusercontent.com/u/118081?v=4?s=48" width="48px;" alt=""/><br /><sub><b>June S</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://github.com/XanthusL"><img src="https://avatars3.githubusercontent.com/u/16461061?v=4?s=48" width="48px;" alt=""/><br /><sub><b>liyiheng</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/baustinanki"><img src="https://avatars2.githubusercontent.com/u/9014288?v=4?s=48" width="48px;" alt=""/><br /><sub><b>baustinanki</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/lixin9311"><img src="https://avatars0.githubusercontent.com/u/371475?v=4?s=48" width="48px;" alt=""/><br /><sub><b>lucus lee</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/mxplusb"><img src="https://avatars1.githubusercontent.com/u/7537841?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Mike Lloyd</b></sub></a><br /></td>
    <td align="center"><a href="http://rubiojr.rbel.co"><img src="https://avatars3.githubusercontent.com/u/10998?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Sergio Rubio</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/FarhadF"><img src="https://avatars3.githubusercontent.com/u/17374492?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Farhad Farahi</b></sub></a><br /></td>
    <td align="center"><a href="http://lasantha.blogspot.com/"><img src="https://avatars1.githubusercontent.com/u/634604?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Lasantha Kularatne</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://github.com/dlom"><img src="https://avatars1.githubusercontent.com/u/823331?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Mark Old</b></sub></a><br /></td>
    <td align="center"><a href="http://flw.tools/"><img src="https://avatars0.githubusercontent.com/u/5546718?v=4?s=48" width="48px;" alt=""/><br /><sub><b>flw</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/davebarda"><img src="https://avatars0.githubusercontent.com/u/6024927?v=4?s=48" width="48px;" alt=""/><br /><sub><b>David Barda</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/matrinox"><img src="https://avatars2.githubusercontent.com/u/4261980?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Geoff Lee</b></sub></a><br /></td>
    <td align="center"><a href="http://international.github.io"><img src="https://avatars3.githubusercontent.com/u/1022918?v=4?s=48" width="48px;" alt=""/><br /><sub><b>George Opritescu</b></sub></a><br /></td>
    <td align="center"><a href="https://twitter.com/Grazfather"><img src="https://avatars3.githubusercontent.com/u/497310?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Grazfather</b></sub></a><br /></td>
    <td align="center"><a href="http://www.mikecordell.com/"><img src="https://avatars2.githubusercontent.com/u/1691120?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Michael Cordell</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="http://patrick.ibexcps.com"><img src="https://avatars2.githubusercontent.com/u/1215497?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Patrick José Pereira</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/sherodtaylor"><img src="https://avatars2.githubusercontent.com/u/1483092?v=4?s=48" width="48px;" alt=""/><br /><sub><b>sherod taylor</b></sub></a><br /></td>
    <td align="center"><a href="http://cogentia.io"><img src="https://avatars2.githubusercontent.com/u/3062663?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Andrew Scott</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/lsipii"><img src="https://avatars1.githubusercontent.com/u/12018440?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Lassi Piironen</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/BlackWebWolf"><img src="https://avatars0.githubusercontent.com/u/14799210?v=4?s=48" width="48px;" alt=""/><br /><sub><b>BlackWebWolf</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/andrewzolotukhin"><img src="https://avatars0.githubusercontent.com/u/1894885?v=4?s=48" width="48px;" alt=""/><br /><sub><b>andrewzolotukhin</b></sub></a><br /></td>
    <td align="center"><a href="https://retgits.github.io"><img src="https://avatars1.githubusercontent.com/u/8568280?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Leon Stigter</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://tamimi.se"><img src="https://avatars3.githubusercontent.com/u/21756?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Amr Tamimi</b></sub></a><br /></td>
    <td align="center"><a href="https://jagdeep.me"><img src="https://avatars3.githubusercontent.com/u/3717137?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Jagdeep Singh</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/darkSasori"><img src="https://avatars0.githubusercontent.com/u/889171?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Lineu Felipe</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/kvj"><img src="https://avatars2.githubusercontent.com/u/159124?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Konstantin</b></sub></a><br /></td>
    <td align="center"><a href="http://www.brendanoleary.com"><img src="https://avatars2.githubusercontent.com/u/6044920?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Brendan O'Leary</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/bertl4398"><img src="https://avatars2.githubusercontent.com/u/1226441?v=4?s=48" width="48px;" alt=""/><br /><sub><b>bertl4398</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/Ferenc-"><img src="https://avatars2.githubusercontent.com/u/6553695?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Ferenc-</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="http://rohanverma.net"><img src="https://avatars1.githubusercontent.com/u/952036?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Rohan Verma</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/fimtitzgerald"><img src="https://avatars1.githubusercontent.com/u/19293566?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Tim Fitzgerald</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/ruggi"><img src="https://avatars2.githubusercontent.com/u/1081051?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Federico Ruggi</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/ctwoodward"><img src="https://avatars2.githubusercontent.com/u/7293328?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Craig Woodward</b></sub></a><br /></td>
    <td align="center"><a href="https://twitter.com/ReadmeCritic"><img src="https://avatars3.githubusercontent.com/u/15367484?v=4?s=48" width="48px;" alt=""/><br /><sub><b>ReadmeCritic</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/jdevelop"><img src="https://avatars0.githubusercontent.com/u/141402?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Eugene</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/Trinergy"><img src="https://avatars1.githubusercontent.com/u/12983705?s=460&v=4?s=48" width="48px;" alt=""/><br /><sub><b>Kenny Wu</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="http://www.romeroruiz.com"><img src="https://avatars0.githubusercontent.com/u/538234?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Renán Romero</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/sticreations"><img src="https://avatars1.githubusercontent.com/u/5031240?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Bastian Groß</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/nicholas-eden"><img src="https://avatars1.githubusercontent.com/u/2496835?v=4?s=48" width="48px;" alt=""/><br /><sub><b>nicholas-eden</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/danrabinowitz"><img src="https://avatars1.githubusercontent.com/u/279390?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Dan Rabinowitz</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/dvdmssmnn"><img src="https://avatars1.githubusercontent.com/u/6897575?v=4?s=48" width="48px;" alt=""/><br /><sub><b>David Missmann</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/mweb"><img src="https://avatars2.githubusercontent.com/u/882006?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Mathias Weber</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/TheRedSpy15"><img src="https://avatars1.githubusercontent.com/u/32081703?v=4?s=48" width="48px;" alt=""/><br /><sub><b>TheRedSpy15</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://www.linkedin.com/in/harald-nordgren-44778192"><img src="https://avatars0.githubusercontent.com/u/9569897?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Harald Nordgren</b></sub></a><br /></td>
    <td align="center"><a href="http://stormfirefox1.github.io"><img src="https://avatars0.githubusercontent.com/u/11583824?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Matei Alexandru Gardus</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/Seanstoppable"><img src="https://avatars2.githubusercontent.com/u/1523955?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Sean Smith</b></sub></a><br /></td>
    <td align="center"><a href="http://kaskavalci.com"><img src="https://avatars1.githubusercontent.com/u/1646238?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Halil Kaskavalci</b></sub></a><br /></td>
    <td align="center"><a href="http://www.johandenoyer.fr"><img src="https://avatars2.githubusercontent.com/u/246715?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Johan Denoyer</b></sub></a><br /></td>
    <td align="center"><a href="https://skymeyer.be"><img src="https://avatars1.githubusercontent.com/u/593516?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Jelle Vink</b></sub></a><br /></td>
    <td align="center"><a href="http://imdevinc.com"><img src="https://avatars1.githubusercontent.com/u/3997333?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Devin Collins</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="http://danne.stayskal.com/"><img src="https://avatars3.githubusercontent.com/u/18333?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Danne Stayskal</b></sub></a><br /></td>
    <td align="center"><a href="https://www.maxbeizer.com"><img src="https://avatars1.githubusercontent.com/u/2006658?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Max Beizer</b></sub></a><br /></td>
    <td align="center"><a href="http://tinyurl.com/nwmj4as"><img src="https://avatars1.githubusercontent.com/u/194392?v=4?s=48" width="48px;" alt=""/><br /><sub><b>E:V:A</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/GaboFDC"><img src="https://avatars0.githubusercontent.com/u/1425500?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Gabriel</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/AndrewScibek"><img src="https://avatars2.githubusercontent.com/u/10111411?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Andrew Scibek</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/FriedCosey"><img src="https://avatars0.githubusercontent.com/u/29709822?v=4?s=48" width="48px;" alt=""/><br /><sub><b>FriedCosey</b></sub></a><br /></td>
    <td align="center"><a href="https://michelegera.dev/"><img src="https://avatars1.githubusercontent.com/u/3891?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Michele Gerarduzzi</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://github.com/rudolphjacksonm"><img src="https://avatars3.githubusercontent.com/u/13438569?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Jack Morris</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/foorb"><img src="https://avatars0.githubusercontent.com/u/14993807?v=4?s=48" width="48px;" alt=""/><br /><sub><b>foorb</b></sub></a><br /></td>
    <td align="center"><a href="http://researchit.las.iastate.edu"><img src="https://avatars0.githubusercontent.com/u/5819098?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Levi Baber</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/gnanderson"><img src="https://avatars0.githubusercontent.com/u/38514?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Graham Anderson</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/bosr"><img src="https://avatars2.githubusercontent.com/u/1936828?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Romain Bossart</b></sub></a><br /></td>
    <td align="center"><a href="http://eonix.ru"><img src="https://avatars0.githubusercontent.com/u/969838?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Kirill Motkov</b></sub></a><br /></td>
    <td align="center"><a href="http://www.BrianChoromanski.com"><img src="https://avatars1.githubusercontent.com/u/3665694?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Brian Choromanski</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="http://siobud.com"><img src="https://avatars0.githubusercontent.com/u/1302304?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Sean DuBois</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/gary-kim"><img src="https://avatars1.githubusercontent.com/u/47195730?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Gary Kim</b></sub></a><br /></td>
    <td align="center"><a href="https://dylanbartels.com"><img src="https://avatars1.githubusercontent.com/u/6660171?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Dylan</b></sub></a><br /></td>
    <td align="center"><a href="http://liet.me"><img src="https://avatars0.githubusercontent.com/u/1990354?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Dmytro Prokhorenkov</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/elliotrushton"><img src="https://avatars1.githubusercontent.com/u/590442?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Elliot</b></sub></a><br /></td>
    <td align="center"><a href="http://chenrui.dev"><img src="https://avatars3.githubusercontent.com/u/1580956?v=4?s=48" width="48px;" alt=""/><br /><sub><b>chenrui</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/sudermanjr"><img src="https://avatars0.githubusercontent.com/u/7624765?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Andrew Suderman</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://github.com/rwhogg"><img src="https://avatars3.githubusercontent.com/u/2373856?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Bob 'Wombat' Hogg</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/hxw"><img src="https://avatars0.githubusercontent.com/u/143462?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Christopher Hall</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/hneiva"><img src="https://avatars1.githubusercontent.com/u/3451557?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Heitor Neiva</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/herbygillot"><img src="https://avatars3.githubusercontent.com/u/618376?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Herby Gillot</b></sub></a><br /></td>
    <td align="center"><a href="http://brudil.com"><img src="https://avatars3.githubusercontent.com/u/382352?v=4?s=48" width="48px;" alt=""/><br /><sub><b>James Canning</b></sub></a><br /></td>
    <td align="center"><a href="https://twitter.com/jeffz4000"><img src="https://avatars1.githubusercontent.com/u/45892?v=4?s=48" width="48px;" alt=""/><br /><sub><b>jeffz</b></sub></a><br /></td>
    <td align="center"><a href="https://mikkeljuhl.com"><img src="https://avatars0.githubusercontent.com/u/1764035?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Mikkel Jeppesen Juhl</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://github.com/lesteenman"><img src="https://avatars1.githubusercontent.com/u/963290?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Erik</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/nyourchuck"><img src="https://avatars1.githubusercontent.com/u/155574?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Nate Yourchuck</b></sub></a><br /></td>
    <td align="center"><a href="https://cprimozic.net/"><img src="https://avatars3.githubusercontent.com/u/4335849?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Casey Primozic</b></sub></a><br /></td>
    <td align="center"><a href="http://pierdelacabeza.com/maruja"><img src="https://avatars3.githubusercontent.com/u/2430915?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Alvaro [Andor]</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/Midnight-Conqueror"><img src="https://avatars1.githubusercontent.com/u/17101621?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Joel Valentine</b></sub></a><br /></td>
    <td align="center"><a href="https://www.viktor-braun.de"><img src="https://avatars0.githubusercontent.com/u/4738210?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Viktor Braun</b></sub></a><br /></td>
    <td align="center"><a href="https://www.chrisdbrown.co.uk/"><img src="https://avatars3.githubusercontent.com/u/3877652?v=4?s=48" width="48px;" alt=""/><br /><sub><b>ChrisDBrown</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://narengowda.github.io/"><img src="https://avatars2.githubusercontent.com/u/582821?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Narendra L</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/ibaum"><img src="https://avatars1.githubusercontent.com/u/24609103?v=4?s=48" width="48px;" alt=""/><br /><sub><b>ibaum</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/noxer"><img src="https://avatars3.githubusercontent.com/u/566185?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Tim Scheuermann</b></sub></a><br /></td>
    <td align="center"><a href="https://indradhanush.github.io/"><img src="https://avatars0.githubusercontent.com/u/2682729?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Indradhanush Gupta</b></sub></a><br /></td>
    <td align="center"><a href="https://victoravelar.com"><img src="https://avatars3.githubusercontent.com/u/7926849?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Victor Hugo Avelar Ossorio</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/scw007"><img src="https://avatars3.githubusercontent.com/u/4001640?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Steven Whitehead</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/lawrencecraft"><img src="https://avatars0.githubusercontent.com/u/660580?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Lawrence Craft</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="http://avi.press"><img src="https://avatars1.githubusercontent.com/u/1388071?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Avi Press</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/Tardog"><img src="https://avatars0.githubusercontent.com/u/22562624?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Sarah Kraßnigg</b></sub></a><br /></td>
    <td align="center"><a href="http://jmks.ca"><img src="https://avatars1.githubusercontent.com/u/4923990?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Jason Schweier</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/massa1240"><img src="https://avatars2.githubusercontent.com/u/8268483?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Massa</b></sub></a><br /></td>
    <td align="center"><a href="http://boot-error.github.io"><img src="https://avatars3.githubusercontent.com/u/8546140?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Vighnesh SK</b></sub></a><br /></td>
    <td align="center"><a href="http://alexfornuto.com"><img src="https://avatars3.githubusercontent.com/u/2349184?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Alex Fornuto</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/stevenwhitehead"><img src="https://avatars0.githubusercontent.com/u/30630257?v=4?s=48" width="48px;" alt=""/><br /><sub><b>stevenwhitehead</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://github.com/jdenoy-saagie"><img src="https://avatars2.githubusercontent.com/u/55875303?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Johan Denoyer</b></sub></a><br /></td>
    <td align="center"><a href="https://albertsalim.dev"><img src="https://avatars1.githubusercontent.com/u/4749355?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Albert Salim</b></sub></a><br /></td>
    <td align="center"><a href="https://Feliciano.Tech"><img src="https://avatars1.githubusercontent.com/u/6017470?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Ricardo N Feliciano</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/3mard"><img src="https://avatars3.githubusercontent.com/u/42009880?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Omer Davutoglu</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/hemu"><img src="https://avatars0.githubusercontent.com/u/1871299?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Hemu</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/Daanikus"><img src="https://avatars0.githubusercontent.com/u/18027087?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Dan Bent</b></sub></a><br /></td>
    <td align="center"><a href="https://cizer.dev"><img src="https://avatars3.githubusercontent.com/u/20225764?v=4?s=48" width="48px;" alt=""/><br /><sub><b>C123R</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://github.com/madepolli"><img src="https://avatars1.githubusercontent.com/u/7237000?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Matjaž Depolli</b></sub></a><br /></td>
    <td align="center"><a href="https://blog.schoentoon.blue"><img src="https://avatars1.githubusercontent.com/u/417618?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Toon Schoenmakers</b></sub></a><br /></td>
    <td align="center"><a href="http://tdhttt.com"><img src="https://avatars2.githubusercontent.com/u/24703459?v=4?s=48" width="48px;" alt=""/><br /><sub><b>TDHTTTT</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/jottr"><img src="https://avatars0.githubusercontent.com/u/2744198?v=4?s=48" width="48px;" alt=""/><br /><sub><b>jottr</b></sub></a><br /></td>
    <td align="center"><a href="https://www.linkedin.com/in/nikolay-mateev-79187b167/"><img src="https://avatars3.githubusercontent.com/u/15074116?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Nikolay Mateev</b></sub></a><br /></td>
    <td align="center"><a href="https://charliewang.io"><img src="https://avatars1.githubusercontent.com/u/1320418?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Charlie Wang</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/liyiheng"><img src="https://avatars3.githubusercontent.com/u/16461061?v=4?s=48" width="48px;" alt=""/><br /><sub><b>liyiheng</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://bjoern.svbtle.com"><img src="https://avatars1.githubusercontent.com/u/1467156?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Bjoern Weidlich</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/firecat53"><img src="https://avatars1.githubusercontent.com/u/568113?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Scott Hansen</b></sub></a><br /></td>
    <td align="center"><a href="https://davidsbond.github.io"><img src="https://avatars3.githubusercontent.com/u/6227720?v=4?s=48" width="48px;" alt=""/><br /><sub><b>David Bond</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/yesnault"><img src="https://avatars3.githubusercontent.com/u/395454?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Yvonnick Esnault</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/leterio"><img src="https://avatars0.githubusercontent.com/u/15060358?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Vinícius Letério</b></sub></a><br /></td>
    <td align="center"><a href="https://adriano.fyi"><img src="https://avatars3.githubusercontent.com/u/3331648?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Adriano</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/jonhadfield"><img src="https://avatars1.githubusercontent.com/u/843944?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Jon Hadfield</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://github.com/Tdnshah"><img src="https://avatars2.githubusercontent.com/u/13272752?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Tejas Shah</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/mogensen"><img src="https://avatars2.githubusercontent.com/u/592710?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Frederik Mogensen</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/rsaarelm"><img src="https://avatars1.githubusercontent.com/u/41840?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Risto Saarelma</b></sub></a><br /></td>
    <td align="center"><a href="https://sam-github.github.io/"><img src="https://avatars2.githubusercontent.com/u/17607?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Sam Roberts</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/gerchardon"><img src="https://avatars0.githubusercontent.com/u/5973160?v=4?s=48" width="48px;" alt=""/><br /><sub><b>gerchardon</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/mryanmurphy"><img src="https://avatars2.githubusercontent.com/u/641427?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Matt</b></sub></a><br /></td>
    <td align="center"><a href="http://devco.net/"><img src="https://avatars0.githubusercontent.com/u/82342?v=4?s=48" width="48px;" alt=""/><br /><sub><b>R.I.Pienaar</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://github.com/fmotrifork"><img src="https://avatars3.githubusercontent.com/u/18327738?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Frederik Mogensen</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/aeter"><img src="https://avatars0.githubusercontent.com/u/238607?v=4?s=48" width="48px;" alt=""/><br /><sub><b>aeter</b></sub></a><br /></td>
    <td align="center"><a href="http://timhwang21.gitbook.io"><img src="https://avatars3.githubusercontent.com/u/5831434?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Tim Hwang</b></sub></a><br /></td>
    <td align="center"><a href="http://about.me/yingfan"><img src="https://avatars1.githubusercontent.com/u/10404961?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Ying Fan Chong</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/MartinJohns"><img src="https://avatars1.githubusercontent.com/u/5269069?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Martin Johns</b></sub></a><br /></td>
    <td align="center"><a href="https://www.jvt.me"><img src="https://avatars0.githubusercontent.com/u/3315059?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Jamie Tanna</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/trimble"><img src="https://avatars3.githubusercontent.com/u/371317?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Todd Trimble</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://www.mitchellhanberg.com"><img src="https://avatars2.githubusercontent.com/u/5523984?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Mitchell Hanberg</b></sub></a><br /></td>
    <td align="center"><a href="https://franga2000.com"><img src="https://avatars3.githubusercontent.com/u/3891092?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Miha Frangež</b></sub></a><br /></td>
    <td align="center"><a href="https://blog.sahilister.in/"><img src="https://avatars0.githubusercontent.com/u/52946452?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Sahil Dhiman</b></sub></a><br /></td>
    <td align="center"><a href="https://pzoo.netlify.app/"><img src="https://avatars2.githubusercontent.com/u/17727004?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Pingzhou &#124; 平舟</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/YuviGold"><img src="https://avatars0.githubusercontent.com/u/29873449?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Yuval Goldberg</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/dabcoder"><img src="https://avatars3.githubusercontent.com/u/5034531?v=4?s=48" width="48px;" alt=""/><br /><sub><b>David Bouchare</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/stone"><img src="https://avatars3.githubusercontent.com/u/29077?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Fredrik Steen</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://github.com/zye1996"><img src="https://avatars2.githubusercontent.com/u/28901953?v=4?s=48" width="48px;" alt=""/><br /><sub><b>zye1996</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/pgaxatte"><img src="https://avatars.githubusercontent.com/u/30696904?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Pierre Gaxatte</b></sub></a><br /></td>
    <td align="center"><a href="https://xntrik.wtf"><img src="https://avatars.githubusercontent.com/u/678260?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Christian Frichot</b></sub></a><br /></td>
    <td align="center"><a href="https://lukas-kaemmerling.de"><img src="https://avatars.githubusercontent.com/u/4281581?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Lukas Kämmerling</b></sub></a><br /></td>
    <td align="center"><a href="https://inetant.net/"><img src="https://avatars.githubusercontent.com/u/1765366?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Antoine Meillet</b></sub></a><br /></td>
    <td align="center"><a href="https://www.patreon.com/cclauss"><img src="https://avatars.githubusercontent.com/u/3709715?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Christian Clauss</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/GibranHL0"><img src="https://avatars.githubusercontent.com/u/61842675?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Gibran Herrera</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://hjr265.me/"><img src="https://avatars.githubusercontent.com/u/348107?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Mahmud Ridwan</b></sub></a><br /></td>
    <td align="center"><a href="https://tadeas.dev/"><img src="https://avatars.githubusercontent.com/u/33228844?v=4?s=48" width="48px;" alt=""/><br /><sub><b>tadeas</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/tnwei"><img src="https://avatars.githubusercontent.com/u/12769364?v=4?s=48" width="48px;" alt=""/><br /><sub><b>tnwei</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/Ginner"><img src="https://avatars.githubusercontent.com/u/26798615?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Ginner</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/Battleman"><img src="https://avatars.githubusercontent.com/u/6746316?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Olivier Cloux</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/dogukanturan"><img src="https://avatars.githubusercontent.com/u/32000865?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Dogukan Turan</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/devenda-avoma"><img src="https://avatars.githubusercontent.com/u/72001066?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Devendra Laulkar</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://github.com/nontw"><img src="https://avatars.githubusercontent.com/u/9658731?v=4?s=48" width="48px;" alt=""/><br /><sub><b>nont</b></sub></a><br /></td>
    <td align="center"><a href="http://kyrylo.org/"><img src="https://avatars.githubusercontent.com/u/1079123?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Kyrylo Silin</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/yosmoc"><img src="https://avatars.githubusercontent.com/u/9694?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Yoshihisa Mochihara</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/thuan1412"><img src="https://avatars.githubusercontent.com/u/36019052?v=4?s=48" width="48px;" alt=""/><br /><sub><b>thuan1412</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/siddhant94"><img src="https://avatars.githubusercontent.com/u/8606880?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Siddhant Sinha</b></sub></a><br /></td>
    <td align="center"><a href="https://resamvi.io/"><img src="https://avatars.githubusercontent.com/u/6261556?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Julien Midedji</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/hypnoglow"><img src="https://avatars.githubusercontent.com/u/4853075?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Igor Zibarev</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://juneezee.github.io/"><img src="https://avatars.githubusercontent.com/u/20135478?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Eng Zer Jun</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/Quentinchampenois"><img src="https://avatars.githubusercontent.com/u/26109239?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Quentin Champ</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/igbanam"><img src="https://avatars.githubusercontent.com/u/390059?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Igbanam Ogbuluijah</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/gcg"><img src="https://avatars.githubusercontent.com/u/811685?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Guney Can Gokoglu</b></sub></a><br /></td>
    <td align="center"><a href="https://des.wtf/"><img src="https://avatars.githubusercontent.com/u/6239776?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Des Preston</b></sub></a><br /></td>
    <td align="center"><a href="https://labesse.fr/"><img src="https://avatars.githubusercontent.com/u/5497623?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Labesse Kévin</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/asadali"><img src="https://avatars.githubusercontent.com/u/3761605?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Asad</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://mrkc.me/"><img src="https://avatars.githubusercontent.com/u/261361?v=4?s=48" width="48px;" alt=""/><br /><sub><b>markcaudill</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/fabge"><img src="https://avatars.githubusercontent.com/u/21140791?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Fabian Geiger</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/dhutty-numo"><img src="https://avatars.githubusercontent.com/u/62157262?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Duncan Hutty</b></sub></a><br /></td>
    <td align="center"><a href="https://gliptak.github.io/"><img src="https://avatars.githubusercontent.com/u/50109?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Gábor Lipták</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/Bizzaro"><img src="https://avatars.githubusercontent.com/u/10475262?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Albert Fung</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/pliski"><img src="https://avatars.githubusercontent.com/u/6731247?v=4?s=48" width="48px;" alt=""/><br /><sub><b>pliski</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/zyrre"><img src="https://avatars.githubusercontent.com/u/2995732?v=4?s=48" width="48px;" alt=""/><br /><sub><b>Peter Krantz</b></sub></a><br /></td>
  </tr>
  <tr>
    <td align="center"><a href="https://www.youtube.com/c/bashbunni"><img src="https://avatars.githubusercontent.com/u/15822994?v=4?s=48" width="48px;" alt=""/><br /><sub><b>bashbunni</b></sub></a><br /></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

## Acknowledgments

The inspiration for `WTF` came from Monica Dinculescu's
[tiny-care-terminal](https://github.com/notwaldorf/tiny-care-terminal).

WTF is built atop [tcell](https://github.com/gdamore/tcell) and [tview](https://github.com/rivo/tview), fantastic projects both. WTF is built, packaged, and deployed via [GoReleaser](https://goreleaser.com).

<p align="center">
<img src="./images/dude_wtf.png?raw=true" title="Dude WTF" width="251" height="201" />
</p>
