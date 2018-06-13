---
title: "IP-API"
date: 2018-06-10T19:41:52-04:00
draft: false
---

Added in `v0.0.7`.

Displays your current IP address information, from [IP-APIcom](http://ip-api.com).

**Note:** IP-API.com has a free-plan rate limit of 120 requests per
minute.

## Source Code

```bash
wtf/ipapi/
```

## Required ENV Variables

None.

## Keyboard Commands

None.

## Configuration

```yaml
ipinfo:
  colors:
    name: red
    value: white
  enabled: true
  position:
    top: 1
    left: 2
    height: 1
    width: 1
  refreshInterval: 150
```
### Attributes

`colors.name` <br />
The default colour for the row names. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11 color</a> name.

`colors.value` <br />
The default colour for the row values. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11 color</a> name.

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.
