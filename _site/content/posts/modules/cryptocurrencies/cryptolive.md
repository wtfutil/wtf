---
title: "CryptoLive"
date: 2018-06-03T20:06:40-07:00
draft: false
---

Added in `v0.0.5`.

Compare crypto currencies using [CryptoCompare](https://cryptocompare.com).

<img src="/imgs/modules/cryptolive.png" width="320" height="203" alt="cryptolive screenshot" />

## Source Code

```bash
wtf/cryptoexchanges/cryptolive/
```

## Required ENV Vars

None.

## Keyboard Commands

None.

## Configuration

```yaml
cryptolive:
  enabled: true
  position:
    top: 5
    left: 2
    height: 1
    width: 2  
  updateInterval: 15
  currencies: 
    BTC:
      displayName: Bitcoin
      to: 
        - USD
        - EUR
        - ETH
        - LTC
        - DOGE
    LTC:
      displayName: Ethereum
      to: 
        - USD
        - EUR
        - BTC
  top:
    BTC:
      displayName: Bitcoin
      limit: 5
      to:
        - USD
  colors: 
    from:
      name: coral
      displayName: grey
    to:
      name: white
      price: green
    top:
      from:
        name: grey
        displayName: coral
      to:
        name: red
        field: white
        value: green
```

### Attributes

`colors.from.name` <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`colors.from.dispayName` <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`colors.to.name` <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`colors.to.price` <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`currencies` <br />

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.
