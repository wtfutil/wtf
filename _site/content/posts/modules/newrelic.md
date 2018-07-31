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

## Keyboard Commands

None.

## Configuration

```yaml
newrelic:
  apiKey: "3276d7155dd9ee27b8b14f8743a408a9"
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

`apiKey` <br />
Value: Your <a href="https://docs.newrelic.com/docs/apis/getting-started/intro-apis/access-rest-api-keys">New Relic API</a> token.

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
