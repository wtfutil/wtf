---
title: "Bamboohr"
date: 2018-05-07T20:17:37-07:00
draft: false
---

## Description

Connects to the BambooHR API and displays who will be Away today.

## Location

```bash
wtf/bamboohr
```

## Required ENV Variables

`WTF_BAMBOO_HR_TOKEN` <br />
Your <a href="https://www.bamboohr.com/api/documentation/">BambooHR API</a> token.

`WTF_BAMBOO_HR_SUBDOMAIN` <br />
Your <a href="https://www.bamboohr.com/api/documentation/">BambooHR API</a> subdomain name.

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
Values: Any positive integer, `0...n`.
