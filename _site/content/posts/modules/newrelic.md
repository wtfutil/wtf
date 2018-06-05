---
title: "New Relic"
date: 2018-05-09T09:01:14-07:00
draft: false
---

Connects to the New Relic API and displays the last n deploys of the
monitored application: deploy ID, deploy time, and who deployed it.

<img src="/imgs/modules/newrelic.png" width="640" height="189" alt="newrelic screenshot" />

## Source Code

```bash
wtf/newrelic/
```

## Required ENV Variables

<span class="caption">Key:</span> `WTF_NEW_RELIC_API_KEY` <br />
<span class="caption">Value:</span> Your <a href="">New Relic API</a>
token.

## Keyboard Commands

None.

## Configuration

```yaml
newrelic:
  applicationId: 10549735
  deployCount: 6
  enabled: true
  position:
    top: 4
    left: 3
    height: 1
    width: 2
  refreshInterval: 900
```

### Attributes

`applicationId` <br />
The integer ID of the New Relic application you wish to report on. <br
/>
Values: A positive integer, `0..n`.

`deployCount` <br />
The number of past deploys to display on screen. <br />
Values: A positive integer, `0..n`.

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.
