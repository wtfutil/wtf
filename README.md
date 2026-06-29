<p align="center">
    <img src="./images/logo_transparent.png?raw=true" title="WTF" alt="WTF" width="560" height="560" />
</p>

[![GitHub Release](https://img.shields.io/github/v/release/wtfutil/wtf?logo=github&style=for-the-badge)](https://github.com/wtfutil/wtf/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/wtfutil/wtf?style=for-the-badge)](https://goreportcard.com/report/github.com/wtfutil/wtf)
[![Discourse Status](https://img.shields.io/discourse/status?server=https%3A%2F%2Fdiscuss.linodians.com&logo=discourse&style=for-the-badge)](https://discuss.linodians.com/c/projects/wtf/7)
[![Bluesky followers](https://img.shields.io/bluesky/followers/wtfutil.bsky.social?logo=bluesky&style=for-the-badge)](https://bsky.app/profile/wtfutil.bsky.social)
[![Mastodon followers](https://img.shields.io/mastodon/follow/115007297910718188?domain=social.linodians.com&logo=mastodon&style=for-the-badge)](https://social.linodians.com/@WTFutil)
![Static Badge](https://img.shields.io/badge/LICENSE-MPL--2.0-orange?style=for-the-badge)

---

:rotating_light: :warning: This project is going through some changes as we prepare for v1.0. This includes
a rename from WTF to **Tessera**! Stay tuned! :warning: :rotating_light:

WTF (aka 'wtfutil') is the personal information dashboard for your terminal, providing at-a-glance access to your very important but infrequently-needed stats and data.

Used by thousands of developers and tech people around the world, WTF is free and open-source. To support the continued use and development of WTF, please consider sponsoring WTF via [GitHub Sponsors](https://github.com/sponsors/FelicianoTech).

### Are you a contributor or sponsor?

Awesome! [See here](https://wtfutil.com/sponsors/exit_message/) for how you can change the exit message, the message WTF shows when quitting, to something special just for you.

---

* [Installation](#installation)
    * [Installing via Homebrew](#installing-via-homebrew)
    * [Installing via `go install`](#installing-via-go-install)
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
brew install linodians/tap/wtfutil

wtfutil
```

### Installing via `go install`

Just run

```sh
go install github.com/wtfutil/wtf@latest
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

### Installing via Arch User Repository

Arch Linux users can utilise the [wtfutil](https://aur.archlinux.org/packages/wtfutil) package to build it from source, or [wtfutil-bin](https://aur.archlinux.org/packages/wtfutil-bin/) to install pre-built binaries.


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


## Acknowledgments

The inspiration for `WTF` came from Monica Dinculescu's
[tiny-care-terminal](https://github.com/notwaldorf/tiny-care-terminal).

WTF is built atop [tcell](https://github.com/gdamore/tcell) and [tview](https://github.com/rivo/tview), fantastic projects both. WTF is built, packaged, and deployed via [GoReleaser](https://goreleaser.com).

<p align="center">
<img src="./images/dude_wtf.png?raw=true" title="Dude WTF" width="251" height="201" />
</p>
