---
title: "Twitter"
date: 2018-07-31T20:21:37-07:00
draft: false
weight: 260
---

Added in `v0.2.0`.

Connects to the Twitter API and displays a single user's tweets.

NOTE: This only works for single-application developer accounts for now.

To make this work, you'll need a couple of things:

1. A [Twitter developer account](https://developer.twitter.com/content/developer-twitter/en.html)
2. A [Twitter bearer token](https://developer.twitter.com/en/docs/basics/authentication/overview/application-only).

Once you have your developer account, a relatively painless way to get a
bearer token is to use [TBT](https://github.com/Trinergy/twitter_bearer_token).

## Source Code

```bash
wtf/twitter/
```

## Configuration

```yaml
twitter:
  bearerToken: "3276d7155dd9ee27b8b14f8743a408a9"
  enabled: true
  position:
    top: 0
    left: 1
    height: 1
    width: 1
  refreshInterval: 20000
  screenName: "wtfutil"
```

### Attributes

`bearerToken` <br />
Value: Your <a href="https://developer.twitter.com/en/docs/basics/authentication/overview/application-only.html">Twitter single-application Bearer Token</a>

`enabled` <br />
Whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: Any positive integer, `0..n`.

`screenName` <br />
The screen name of the Twitter user who's tweets you want to follow. <br />
Values: Any valid Twitter user's screen name.
