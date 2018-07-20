---
title: "Weather"
date: 2018-05-09T11:44:13-07:00
draft: false
---

Displays a configurable list of current weather report, including
current temperature, sunrise time, and sunset time.

<img src="/imgs/modules/weather.png" width="320" height="187" alt="weather screenshot" />

## Source Code

```bash
wtf/weather/
```

## Required ENV Variables



<span class="caption">Key:</span> `WTF_OWM_API_KEY` <br />
<span class="caption">Action:</span> Your <a href="https://openweathermap.org/appid">OpenWeatherMap API</a> key. <br />
<span class="caption">Note:</span> DEPRECATED. See the `apiKey` config value, below.

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
  apiKey: "2dfb3e3650a1950adddb6badf5ba1aaa"
  # From http://openweathermap.org/help/city_list.txt
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

`apiKey` <br />
Your <a href="https://openweathermap.org/appid">OpenWeatherMap API</a> key.

`cityids` <br />
A list of the <a
href="http://openweathermap.org/help/city_list.txt">OpenWeatherMap city
IDs</a> for the cities you want to view. <br />
Values: A list of positive integers, `0..n`

`colors.current` <br />
The color to highlight the current temperature in. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`language` <br />
The human language in which to present the weather data. <br />
Values: Any <a href="https://openweathermap.org/current">language identifier</a> specified by OpenWeatherMap.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.

`tempUnit` <br />
The temperature scale in which to display temperature values. <br />
Values: `F` for Fahrenheit, `C` for Celcius.
