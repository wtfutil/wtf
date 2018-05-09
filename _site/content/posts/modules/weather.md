---
title: "Weather"
date: 2018-05-09T11:44:13-07:00
draft: false
---

## Description

## Source Code

```bash
wtf/weather/
```

## Required ENV Variables

<span class="caption">Key:</span> `WTF_OWM_API_KEY` <br />
<span class="caption">Action:</span> Your <a href="https://openweathermap.org/appid">OpenWeatherMap API</a> key.

## Keyboard Commands

<span class="caption">Key:</span> `/` <br />
<span class="caption">Action:</span> Open/close the widget's help window.

<span class="caption">Key:</span> `h` <br />
<span class="caption">Action:</span> Show the previous weather location.

<span class="caption">Key:</span> `l` <br />
<span class="caption">Action:</span> Show the next weather location.

<span class="caption">Key:</span> `←` <br />
<span class="caption">Action:</span> Show the previous weather location.

<span class="caption">Key:</span> `→` <br />
<span class="caption">Action:</span> Show the next weather location.

## Configuration

```yaml
weather:
  cityids:
  - 6173331
  - 3128760
  - 6167865
  - 6176823
  colors:
    current: "lightblue"
  enabled: true
  language: "EN"
  position:
    top: 0
    left: 2
    height: 1
    width: 1
  refreshInterval: 900
  tempUnit: "C"
```

### Attributes
