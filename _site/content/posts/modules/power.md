---
title: "Power"
date: 2018-05-26T19:26:23-07:00
draft: false
---

Displays information about the current power source.

For battery, also displays the current charge, estimated time remaining,
and whether it is charging or discharging.

<img src="/imgs/modules/power.png" width="320" height="129" alt="power screenshot" />

## Source Code
```bash
wtf/power/
```

## Required ENV Variables

None.

## Keyboard Commands

None.

## Configuration

```yaml
power:
  enabled: true
  position:
    top: 5
    left: 0
    height: 2
    width: 1
  refreshInterval: 15
```

### Attributes

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.
