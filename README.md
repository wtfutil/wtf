
<img src="https://travis-ci.com/senorprogrammer/wtf.svg?branch=master" /> <a href="https://gitter.im/wtfutil/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge"><img src="https://badges.gitter.im/wtfutil/Lobby.svg" /></a> <a href="https://twitter.com/wtfutil"><img src="https://img.shields.io/badge/follow-on%20twitter-blue.svg" /></a> <a href="#contributors"><img src="https://img.shields.io/badge/all_contributors-27-orange.svg?style=flat-square" /></a> <img src="https://goreportcard.com/badge/github.com/senorprogrammer/wtf" />

# WTF

A personal terminal-based dashboard utility, designed for
displaying infrequently-needed, but very important, daily data.

<p align="center">
<img src="./docs/img/screenshot.jpg" title="screenshot" width="720" height="420" />
</p>

## Quick Start

[Download and run the latest binary](https://github.com/senorprogrammer/wtf/releases) or install from source:

```bash
go get -u github.com/senorprogrammer/wtf
cd $GOPATH/src/github.com/senorprogrammer/wtf
make install
make run
```

**Note:** WTF is _only_ compatible with Go versions **1.9.2** or later. It currently _does not_ compile with `gccgo`.

## Documentation

See [https://wtfutil.com](https://wtfutil.com) for the definitive
documentation. Here's some short-cuts:

* [Installation](http://wtfutil.com/posts/installation/)
* [Configuration](http://wtfutil.com/posts/configuration/)
* [Module Documentation](http://wtfutil.com/posts/modules/)

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests.

### Adding Dependencies

Dependency management in WTF is handled by [dep](https://github.com/golang/dep). See that page for installation and usage details.

If the work you're doing requires the addition of a new dependency,
please be sure to use `dep` to [vendor your dependencies](https://golang.github.io/dep/docs/daily-dep.html#adding-a-new-dependency).

## Contributors


Thanks goes to these wonderful people:

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore -->
| [<img src="https://avatars0.githubusercontent.com/u/6413?v=4" width="48px;"/><br /><sub><b>Chris Cummer</b></sub>](https://twitter.com/senorprogrammer)<br /> | [<img src="https://avatars1.githubusercontent.com/u/34973359?v=4" width="48px;"/><br /><sub><b>Hossein Mehrabi</b></sub>](https://github.com/jeangovil)<br /> | [<img src="https://avatars0.githubusercontent.com/u/11779018?v=4" width="48px;"/><br /><sub><b>FengYa</b></sub>](https://github.com/Fengyalv)<br /> | [<img src="https://avatars2.githubusercontent.com/u/17337753?v=4" width="48px;"/><br /><sub><b>deltax</b></sub>](https://fluxionnetwork.github.io/fluxion/)<br /> | [<img src="https://avatars0.githubusercontent.com/u/1319630?v=4" width="48px;"/><br /><sub><b>Bill Keenan</b></sub>](https://github.com/BillKeenan)<br /> | [<img src="https://avatars2.githubusercontent.com/u/118081?v=4" width="48px;"/><br /><sub><b>June S</b></sub>](http://blog.sapara.com)<br /> | [<img src="https://avatars3.githubusercontent.com/u/16461061?v=4" width="48px;"/><br /><sub><b>liyiheng</b></sub>](https://github.com/XanthusL)<br /> |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| [<img src="https://avatars2.githubusercontent.com/u/9014288?v=4" width="48px;"/><br /><sub><b>baustinanki</b></sub>](https://github.com/baustinanki)<br /> | [<img src="https://avatars0.githubusercontent.com/u/371475?v=4" width="48px;"/><br /><sub><b>lucus lee</b></sub>](https://github.com/lixin9311)<br /> | [<img src="https://avatars1.githubusercontent.com/u/7537841?v=4" width="48px;"/><br /><sub><b>Mike Lloyd</b></sub>](https://github.com/mxplusb)<br /> | [<img src="https://avatars3.githubusercontent.com/u/10998?v=4" width="48px;"/><br /><sub><b>Sergio Rubio</b></sub>](http://rubiojr.rbel.co)<br /> | [<img src="https://avatars3.githubusercontent.com/u/17374492?v=4" width="48px;"/><br /><sub><b>Farhad Farahi</b></sub>](https://github.com/FarhadF)<br /> | [<img src="https://avatars1.githubusercontent.com/u/634604?v=4" width="48px;"/><br /><sub><b>Lasantha Kularatne</b></sub>](http://lasantha.blogspot.com/)<br /> | [<img src="https://avatars1.githubusercontent.com/u/823331?v=4" width="48px;"/><br /><sub><b>Mark Old</b></sub>](https://github.com/dlom)<br /> |
| [<img src="https://avatars0.githubusercontent.com/u/5546718?v=4" width="48px;"/><br /><sub><b>flw</b></sub>](http://flw.tools/)<br /> | [<img src="https://avatars0.githubusercontent.com/u/6024927?v=4" width="48px;"/><br /><sub><b>David Barda</b></sub>](https://github.com/davebarda)<br /> | [<img src="https://avatars2.githubusercontent.com/u/4261980?v=4" width="48px;"/><br /><sub><b>Geoff Lee</b></sub>](https://github.com/matrinox)<br /> | [<img src="https://avatars3.githubusercontent.com/u/1022918?v=4" width="48px;"/><br /><sub><b>George Opritescu</b></sub>](http://international.github.io)<br /> | [<img src="https://avatars3.githubusercontent.com/u/497310?v=4" width="48px;"/><br /><sub><b>Grazfather</b></sub>](https://twitter.com/Grazfather)<br /> | [<img src="https://avatars2.githubusercontent.com/u/1691120?v=4" width="48px;"/><br /><sub><b>Michael Cordell</b></sub>](http://www.mikecordell.com/)<br /> | [<img src="https://avatars2.githubusercontent.com/u/1215497?v=4" width="48px;"/><br /><sub><b>Patrick José Pereira</b></sub>](http://patrick.ibexcps.com)<br /> |
| [<img src="https://avatars2.githubusercontent.com/u/1483092?v=4" width="48px;"/><br /><sub><b>sherod taylor</b></sub>](https://github.com/sherodtaylor)<br /> | [<img src="https://avatars2.githubusercontent.com/u/3062663?v=4" width="48px;"/><br /><sub><b>Andrew Scott</b></sub>](http://cogentia.io)<br /> | [<img src="https://avatars2.githubusercontent.com/u/3252403?v=4" width="48px;"/><br /><sub><b>Anand Sudhir Prayaga</b></sub>](https://github.com/anandsudhir)<br /> | [<img src="https://avatars1.githubusercontent.com/u/12018440?v=4" width="48px;"/><br /><sub><b>Lassi Piironen</b></sub>](https://github.com/lsipii)<br /> | [<img src="https://avatars0.githubusercontent.com/u/14799210?v=4" width="48px;"/><br /><sub><b>BlackWebWolf</b></sub>](https://github.com/BlackWebWolf)<br /> | [<img src="https://avatars0.githubusercontent.com/u/1894885?v=4" width="48px;"/><br /><sub><b>andrewzolotukhin</b></sub>](https://github.com/andrewzolotukhin)<br /> | [<img src="https://avatars1.githubusercontent.com/u/8568280?v=4" width="48px;"/><br /><sub><b>Leon Stigter</b></sub>](https://retgits.github.io)<br /> |
| [<img src="https://avatars3.githubusercontent.com/u/21756?v=4" width="48px;"/><br /><sub><b>Amr Tamimi</b></sub>](https://tamimi.se)<br /> | [<img src="https://avatars3.githubusercontent.com/u/3717137?v=4" width="48px;"/><br /><sub><b>Jagdeep Singh</b></sub>](https://jagdeep.me)<br /> |
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/kentcdodds/all-contributors) specification. Contributions of any kind welcome!

## Acknowledgments

The inspiration for `WTF` came from Monica Dinculescu's
[tiny-care-terminal](https://github.com/notwaldorf/tiny-care-terminal).

Many thanks to <a href="https://lendesk.com">Lendesk</a> for supporting this project by
providing time to develop it.

The following open-source libraries were used in the creation of `WTF`.
Many thanks to all these developers.

* [calendar](https://google.golang.org/api/calendar/v3)
* [config](https://github.com/olebedev/config)
* [go-gerrit](https://github.com/andygrunwald/go-gerrit)
* [go-github](https://github.com/google/go-github)
* [goreleaser](https://github.com/goreleaser/goreleaser)
* [newrelic](https://github.com/yfronto/newrelic)
* [openweathermap](https://github.com/briandowns/openweathermap)
* [tcell](https://github.com/gdamore/tcell)
* [tview](https://github.com/rivo/tview)

<p align="center">
<img src="./docs/img/dude_wtf.png?raw=true" title="Dude WTF" width="251" height="201" />
</p>
