---
title: "Bittrex"
date: 2018-06-04T20:06:40-07:00
draft: false
---

Added in `v0.0.5`.

Get the last 24 hour summary of cryptocurrencies market using [Bittrex](https://bittrex.com).

<img src="/imgs/modules/bittrex.png" width="320" height="412" alt="bittrex screenshot" />

## Source Code

```bash
wtf/cryptoexchanges/bittrex/
```

## Required ENV Variables

None.

## Keyboard Commands

None.

## Configuration

```yaml
bittrex:
    enabled: true
    position:
        top: 1
        left: 2
        height: 3
        width: 1
    refreshInterval: 5
    summary:
        BTC:
            displayName: Bitcoin
            market:
            - LTC
            - ETH
    colors:
        base:
            name: orange
            displayName: red
        market:
            name: red
            field: white
            value: green
```

### Attributes

`colors.base.name` <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`colors.base.dispayName` <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`colors.market.name` <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`colors.market.field` <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`colors.market.value` <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`summary` <br />

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.
