---
title: "CircleCI"
date: 2018-06-10T19:26:08-04:00
draft: false
---

Added in `v0.0.7`.

Displays build information for your CircleCI account.

<img src="/imgs/modules/circleci.png" width="609" height="150" alt="circleci screenshot" />

## Source Code

```bash
wtf/circleci/
```

## Required ENV Variables

<span class="caption">Key:</span> `WTF_CIRCLE_API_KEY` <br />
<span class="caption">Value:</span> Your <a href="https://circleci.com/account/api">CircleCI API</a>
token.

## Keyboard Commands

None.

## Configuration

```yaml
circleci:
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
