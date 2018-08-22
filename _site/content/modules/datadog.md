---
title: "Datadog"
date: 2018-08-18T00:00:00Z
draft: false
weight: 60
---

Connects to the Datadog API and displays alerting modules.

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
  monitors:
    tags:
      - "team:ops"
  position:
    top: 4
    left: 3
    height: 1
    width: 2
```

### Attributes

`apiKey` <br />
Value: Your <a href="https://docs.datadoghq.com/api/?lang=python#overview">Datadog API</a> key.

`applicationKey` <br />
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
