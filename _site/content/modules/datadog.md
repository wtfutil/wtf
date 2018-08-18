---
title: "Datadog"
date: 2018-08-18T00:00:00Z
draft: false
weight: 160
---

<img class="screenshot" src="/imgs/modules/newrelic.png" width="640" height="189" alt="newrelic screenshot" />

Connects to the Datadog API and displays alerting modules

## Source Code

```bash
wtf/datadog/
```

## Configuration

```yaml
datadog:
  apiKey: "<yourapikey>"
  applicationKey: "<yourapplicationkey>"
  enabled: true
  position:
    top: 4
    left: 3
    height: 1
    width: 2
  monitors:
    tags:
      - "team:ops"
```

### Attributes

`apiKey` <br />
Value: Your <a href="https://docs.datadoghq.com/api/?lang=python#overview">Datadog API</a> key.

`applicationKey` <br />
The integer ID of the New Relic application you wish to report on. <br/>
Value: Your <a href="https://docs.datadoghq.com/api/?lang=python#overview">Datadog Application</a> key.

`monitors` <br />
Configuration for the monitors functionality.

`tags` <br />
Array of tags you want to query monitors by.

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.
