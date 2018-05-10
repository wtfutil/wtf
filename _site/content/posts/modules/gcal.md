---
title: "Google Calendar"
date: 2018-05-10T08:25:33-07:00
draft: false
---

## Description

Displays your upcoming Google calendar events.

<img src="/imgs/modules/gcal.png" width="320" height="389" alt="gcal screenshot" />

## Source Code

```bash
wtf/gcal/
```

## Required ENV Variables

<span class="caption">Key:</span> `WTF_GOOGLE_CAL_CLIENT_ID` <br />
<span class="caption">Value:</span> Your <a href="https://developers.google.com/calendar/auth">Google API</a> client ID.

<span class="caption">Key:</span> `WTF_GOOGLE_CAL_CLIENT_SECRET` <br />
<span class="caption">Value:</span> Your <a href="https://developers.google.com/calendar/auth">Google API</a> client secret.

## Keyboard Commands

None.

## Configuration

```yaml
gcal:
  colors:
    title: "red"
    description: "lightblue"
    highlights:
    - ['1on1|1\/11', 'green']
    - ['apple|google|aws', 'blue']
    - ['interview|meet', 'magenta']
    - ['lunch', 'yellow']
    past: "gray"
  conflictIcon: "🚨"
  currentIcon: "💥"
  enabled: true
  eventCount: 12
  position:
    top: 0
    left: 0
    height: 4
    width: 1
  refreshInterval: 300
  secretFile: "~/.wtf/gcal/client_secret.json"
```

### Attributes

`colors.title` <br />
Specifies the default colour for calendar event titles. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color</a> name.

`colors.description` <br />
Specifies the default color for calendar event descriptions. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color</a> name.

`colors.highlights` <br />
A list of arrays that define a regular expression pattern and a color.
If a calendar event title matches a regular expression, the title will
be drawn in that colour. Over-rides the default title colour. <br />
Values: [a valid regular expression, any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11 color</a> name.]

`colors.past` <br />
Specifies the color for calendar events that have passed. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color</a> name.

`conflictIcon` <br />
The icon displayed beside calendar events that have conflicting times
(they intersect or overlap in some way). <br />
Values: Any displayable unicode character.

`currentIcon` <br />
The icon displayed beside the current calendar event. <br />
Values: Any displayable unicode character.

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`eventCount` <br />
The number of calendar events to display. <br />
Values: A positive integer, `0..n`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.

`secretFile` <br />
Your <a href="https://developers.google.com/calendar/quickstart/go">Google client secret</a> JSON file. <br />
Values: A string representing a file path to the JSON secret file.
