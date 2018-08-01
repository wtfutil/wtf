---
title: "OpsGenie"
date: 2018-05-08T20:53:40-07:00
draft: false
---

Connects to the OpsGenie API and displays all your scheduled rotations
and who's currently on call.

<img src="/imgs/modules/opsgenie.png" width="320" height="389" alt="opsgenie screenshot" />

## Source Code

```bash
wtf/opsgenie/
```

## Keyboard Commands

None.

## Configuration

```yaml
opsgenie:
  apiKey: "3276d7155dd9ee27b8b14f8743a408a9"
  displayEmpty: false
  enabled: true
  position:
    top: 2
    left: 1
    height: 2
    width: 1
  refreshInterval: 21600
```

### Attributes

`apiKey` < br />
Value: Your <a href="https://docs.opsgenie.com/docs/api-integration">OpsGenie API</a> token.

`displayEmpty` <br />
Whether schedules with no assigned person on-call should be displayed. <br />
Values:  `true`, `false`.

`enabled` <br />
Whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.
