---
title: "Spotify"
date: 2018-10-02T17:13:11-07:00
draft: false
weight: 205
---

Control the Spotify client.

## Source Code

```bash
wtf/spotify/
```

## Keyboard Commands

<span class="caption">Key:</span> `/` <br />
<span class="caption">Action:</span> Open/close the widget's help window.

<span class="caption">Key:</span> `space` <br />
<span class="caption">Action:</span> Play/pause the currently-selected
track

<span class="caption">Key:</span> `h` <br />
<span class="caption">Action:</span> Play previous song

<span class="caption">Key:</span> `l` <br />
<span class="caption">Action:</span> Play next song

## Configuration

```yaml
spotify:
  enabled: true
  position:
    top: 1
    left: 2
    height: 1
    width: 1
  refreshInterval: 0
```

## Attributes

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.
