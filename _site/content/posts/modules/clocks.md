---
title: "Clocks"
date: 2018-05-07T19:47:31-07:00
draft: false
---

Displays a configurable list of world clocks, the local time, and date.

<img src="/imgs/modules/clocks.png" width="320" height="191" alt="clocks screenshot" />

## Source Code

```bash
wtf/clocks/
```

## Required ENV Variables

None.

## Keyboard Commands

None.

## Configuration

```yaml
clocks:
  colors:
    rows:
      even: "lightblue"
      odd: "white"
  enabled: true
  locations:
    # From https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
    Avignon: "Europe/Paris"
    Barcelona: "Europe/Madrid"
    Dubai: "Asia/Dubai"
    New York: "America/New York"
    Toronto: "America/Toronto"
    UTC: "Etc/UTC"
    Vancouver: "America/Vancouver"
  position:
    top: 4
    left: 0
    height: 1
    width: 1
  refreshInterval: 15
  # Valid options are: alphabetical, chronological
  sort: "alphabetical"
```
### Attributes

`colors.rows.even` <br />
The foreground color for even-numbered rows. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`colors.rows.odd` <br />
The foreground color for the odd-numbered rows. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`locations` <br />
Defines the timezones for the world clocks that you want to display.
`key` is a unique label that will be displayed in the UI. `value` is a
timezone name. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/List_of_tz_database_time_zones">TZ database timezone</a>.

`position` <br />
Defines where in the grid this module's widget will be displayed.

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.

`sort` <br />
Defines the display order of the clocks in the widget. <br />
Values: `alphabetical` or `chronological`. `alphabetical` will sort in
acending order by `key`, `chronological` will sort in ascending order by
date/time.
