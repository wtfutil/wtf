---
title: "TravisCI"
date: 2018-07-18T14:36:08-04:00
draft: false
---

Added in `v0.0.12`.

Displays build information for your Travis CI account.

<img src="/imgs/modules/travisci.png" width="609" height="150" alt="travisci screenshot" />

## Source Code

```bash
wtf/travisci/
```

## Required ENV Variables

<span class="caption">Key:</span> `WTF_TRAVIS_API_TOKEN` <br />
<span class="caption">Value:</span> Your <a href="https://developer.travis-ci.org/authentication">Travis CI API</a> access token.

## Keyboard Commands

None.

## Configuration

```yaml
travisci:
  enabled: true
  position:
    top: 4
    left: 1
    height: 1
    width: 2
  refreshInterval: 900
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
