---
title: "Blockfolio"
date: 2018-06-13T09:29:59-07:00
draft: false
---

Added in `v0.0.8`.

Display your Blockfolio crypto holdings.

<img src="/imgs/modules/blockfolio.png" width="320" height="185" alt="blockfolio screenshot" />

## Source

```bash
wtf/blockfolio/
```

## Required ENV Variables

None.

## Keyboard Commands

None.

## Configuration

```yaml
blockfolio:
  colors:
    name: blue
    grows: green
    drop: red
  device_token: "device token"
  displayHoldings: true
  enabled: true
  position:
    top: 3
    left: 1
    width: 1
    height: 1
  refreshInterval: 400
```

### Attributes

`colors.name` <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`colors.grows` <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`colors.drop` <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`device_token` <br />
Value: See [this gist](https://github.com/bob6664569/blockfolio-api-client) for
details on how to get your Blockfolio API token.

`displayHoldings` <br />

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.
