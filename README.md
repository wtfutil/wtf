<p align="right">
  
  <img src="https://travis-ci.com/senorprogrammer/wtf.svg?branch=master" />
</p>

<p align="center">
<img src="./docs/img/wtf.jpg?raw=true" title="WTF" width="852" height="240" />
</p>

A personal terminal-based dashboard utility, designed for
displaying infrequently-needed, but very important, daily data.

<p align="center">
<img src="./docs/img/screenshot.jpg" title="screenshot" width="720" height="420" />
</p>

## Quick Start

### Installation from Source

**Note:** WTF is _only_ compatible with Go versions **1.9.2** or later. It currently _does not_ compile with `gccgo`.

```bash
go get -u github.com/senorprogrammer/wtf
cd $GOPATH/src/github.com/senorprogrammer/wtf
make install
make run
```

Or [download the latest binary](https://github.com/senorprogrammer/wtf/releases).

## Support

Chat on Gitter
[![Join the chat at https://gitter.im/wtfutil/Lobby](https://badges.gitter.im/wtfutil/Lobby.svg)](https://gitter.im/wtfutil/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## Documentation

See [https://wtfutil.com](https://wtfutil.com) for the definitive
documentation. Here's some short-cuts:

* [Installation](http://wtfutil.com/posts/installation/)
* [Configuration](http://wtfutil.com/posts/configuration/)
* [Module Documentation](http://wtfutil.com/posts/modules/)

And a "probably up-to-date" list of currently-implemented modules:

* [BambooHR](http://wtfutil.com/posts/modules/bamboohr/)
* [World Clocks](http://wtfutil.com/posts/modules/clocks/)
* [Command Runner](http://wtfutil.com/posts/modules/cmdrunner/)
* [Google Calendar](http://wtfutil.com/posts/modules/gcal/)
* [Git](http://wtfutil.com/posts/modules/git/)
* [GitHub](http://wtfutil.com/posts/modules/github/)
* [IPInfo](http://wtfutil.com/posts/modules/ipinfo/)
* [Jira](http://wtfutil.com/posts/modules/jira/)
* [New Relic](http://wtfutil.com/posts/modules/newrelic/)
* [OpsGenie](http://wtfutil.com/posts/modules/opsgenie)
* [Power](http://wtfutil.com/posts/modules/power/)
* [PrettyWeather](http://wtfutil.com/posts/modules/prettyweather/)*
* [Security](http://wtfutil.com/posts/modules/security/)
* [Textfile](http://wtfutil.com/posts/modules/textfile/)
* [Todo List](http://wtfutil.com/posts/modules/todo/)
* [Weather](http://wtfutil.com/posts/modules/weather/)

*experimental

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests.

## Authors

* Chris Cummer, [senorprogrammer](https://github.com/senorprogrammer)

## License

See [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments

The inspiration for `WTF` came from Monica Dinculescu's
[tiny-care-terminal](https://github.com/notwaldorf/tiny-care-terminal).

Many thanks to <a href="https://lendesk.com">Lendesk</a> for supporting this project by
providing time to develop it.

The following open-source libraries were used in the creation of `WTF`.
Many thanks to all these developers.

* [calendar](https://google.golang.org/api/calendar/v3)
* [config](https://github.com/olebedev/config)
* [go-github](https://github.com/google/go-github)
* [goreleaser](https://github.com/goreleaser/goreleaser)
* [newrelic](https://github.com/yfronto/newrelic)
* [openweathermap](https://github.com/briandowns/openweathermap)
* [tcell](https://github.com/gdamore/tcell)
* [tview](https://github.com/rivo/tview)

<p align="center">
<img src="./docs/img/dude_wtf.png?raw=true" title="Dude WTF" width="251" height="201" />
</p>

