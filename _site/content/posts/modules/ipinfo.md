---
title: "Module: IPInfo"
date: 2018-06-01T23:18:48-07:00
draft: false
---

Displays your current IP address information, from ipinfo.io.

<img src="/imgs/modules/ipinfo.png" width="320" height="199" alt="ipinfo screenshot" />

## Source Code

```bash
wtf/ipinfo/
```

## Required ENV Variables

None.

## Keyboard Commands

None.

## Configuration

```yaml
ipinfo:
  enabled: true
  position:
    top: 1
    left: 2
    height: 1
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
