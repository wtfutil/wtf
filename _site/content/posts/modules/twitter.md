---
title: "Twitter"
date: 2018-07-31T20:21:37-07:00
draft: false
---

Connects to the Twitter API and displays a single user's tweets
NOTE: This only works for single-application developer accounts for now

## Source Code

```bash
wtf/twitter/
```
## Required ENV Variables

<span class="caption">Key:</span> `WTF_TWITTER_BEARER_TOKEN` <br />
<span class="caption">Action:</span> Your <a href="https://developer.twitter.com/en/docs/basics/authentication/overview/application-only.html">Twitter single-application Bearer Token</a><br />

## Keyboard Commands

None.

## Configuration

```yaml
twitter:
  enabled: true
  position:
    top: 0
    left: 1
    height: 1
    width: 1
  refreshInterval: 20000
```

### Attributes

`enabled` <br />
Whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: Any positive integer, `0..n`.
