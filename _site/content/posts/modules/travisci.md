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

## Keyboard Commands

<span class="caption">Key:</span> `[return]` <br />
<span class="caption">Action:</span> Open the selected build in the browser.

<span class="caption">Key:</span> `j` <br />
<span class="caption">Action:</span> Select the next build in the list.

<span class="caption">Key:</span> `k` <br />
<span class="caption">Action:</span> Select the previous build in the list.

<span class="caption">Key:</span> `r` <br />
<span class="caption">Action:</span> Refresh the data.

<span class="caption">Key:</span> `↓` <br />
<span class="caption">Action:</span> Select the next build in the list.

<span class="caption">Key:</span> `↑` <br />
<span class="caption">Action:</span> Select the previous build in the list.

## Configuration

```yaml
travisci:
  apiKey: "3276d7155dd9ee27b8b14f8743a408a9"
  enabled: true
  position:
    top: 4
    left: 1
    height: 1
    width: 2
  pro: false
  refreshInterval: 900
```

### Attributes

`apiKey` <br />
Value: Your <a href="https://developer.travis-ci.org/authentication">Travis CI API</a> access token.

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.

`pro` <br />
Determines whether or not this module will use the Pro version of Travis CI.<br />
Values: `true`, `false`.
