---
title: "Pretty Weather"
date: 2018-06-02T05:32:04-07:00
draft: false
---

Displays weather information as ASCII art from
[Wttr.in](http://wttr.in).

<img src="/imgs/modules/prettyweather.png" width="320" height="191" alt="prettyweather screenshot" />

## Source Code

```bash
wtf/prettyweather/
```

## Required ENV Variables

None.

## Keyboard Commands

None.

## Configuration

```yaml
    prettyweather:
      enabled: true
      city: "tehran"
      position:
        top: 3
        left: 5
        height: 1
        width: 1
      refreshInterval: 300
      unit: "c"
      view: 0
      language: "en"
```

### Attributes

`city` <br />
_Optional_. It will grab the current location from your IP address if
omitted.<br />
Values: The name of any city supported by [Wttr.in](http://wttr.in).

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.

`unit` <br />
_Optional_. It will use metric if you are out of US and imperial for US.<br />
The temperature scale in which to display temperature values. <br />
Values: `F` for Fahrenheit, `C` for Celcius.

`view` <br />
_Optional_ Wttr.in view configuration. <br />
Values: See `curl wttr.in/:help` for more details.

`language` <br />
_Optional_ Wttr.in language configuration. <br />
Values: See `curl wttr.in/:translation` for more details.
