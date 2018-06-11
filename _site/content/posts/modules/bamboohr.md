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

## Required ENV Variables

<span class="caption">Key:</span> `WTF_BAMBOO_HR_TOKEN` <br />
<span class="caption">Value:</span> Your <a href="https://www.bamboohr.com/api/documentation/">BambooHR API</a> token.

<span class="caption">Key:</span> `WTF_BAMBOO_HR_SUBDOMAIN` <br />
<span class="caption">Value:</span> Your <a href="https://www.bamboohr.com/api/documentation/">BambooHR API</a> subdomain name.

## Keyboard Commands

None.

## Configuration

```yaml
bamboohr:
  enabled: true
  position:
    top: 0
    left: 1
    height: 2
    width: 1
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
Values: Any positive integer, `0..n`.
