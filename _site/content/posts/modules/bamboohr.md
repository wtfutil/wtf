---
title: "BambooHR"
date: 2018-05-07T20:17:37-07:00
draft: false
---

Connects to the BambooHR API and displays who will be Away today.

## Source Code

```bash
wtf/bamboohr/
```

## Keyboard Commands

None.

## Configuration

```yaml
bamboohr:
  apiKey: "3276d7155dd9ee27b8b14f8743a408a9"
  enabled: true
  position:
    top: 0
    left: 1
    height: 2
    width: 1
  refreshInterval: 900
  subdomain: "testco"
```

### Attributes

`apiKey` <br />
Value: Your <a href="https://www.bamboohr.com/api/documentation/">BambooHR API</a> token.

`enabled` <br />
Whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: Any positive integer, `0..n`.

`subdomain` <br />
Value: Your <a href="https://www.bamboohr.com/api/documentation/">BambooHR API</a> subdomain name.
