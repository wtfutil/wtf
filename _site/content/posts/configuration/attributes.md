---
title: "Configuration: Attributes"
date: 2018-05-16T21:51:23-07:00
draft: false
---

The following top-level attributes are configurable in `config.yml`.
See this <a href="https://github.com/senorprogrammer/wtf/blob/master/_sample_configs/simple_config.yml">example config file</a> for more details.

```yaml
wtf:
  colors:
    background: "red"
    border:
      Focusable: "darkslateblue"
      focused: "orange"
      normal: "gray"
  grid:
    # How _wide_ the columns are, in terminal characters. In this case we have
    # six columns, each of which are 35 characters wide
    columns: [35, 35, 35, 35, 35, 35]

    # How _high_ the rows are, in terminal lines. In this case we have five rows
    # that support ten line of text, one of three lines, and one of four
    rows: [10, 10, 10, 10, 10, 3, 4]
  # The app redraws itself once a second
  openFileUtil: open
  refreshInterval: 1
  term: "xterm-256color"
```

### Attributes

`colors.background` <br />
The color to draw the background of the app in. Use this to match your
terminal colors. May be over-written by individual module
configurations. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`colors.border.focusable` <br />
The color in which to draw the border of widgets that can accept
keyboard focus. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`colors.border.focused` <br />
The color in which to draw the border of the widget that currently has
keyboard focus. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`colors.border.normal` <br />
The color in which to draw the borders of the widgets that cannot accept
focus. <br/>
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`grid.columns` <br />
An array that defines the widths of all the columns. <br />
Values: See <a href="https://github.com/rivo/tview/wiki/Grid">tview's
Grid</a> for details.

`grid.rows` <br />
An array that defines the heights of all the rows. <br />
Values: See <a href="https://github.com/rivo/tview/wiki/Grid">tview's
Grid</a> for details.

`openFileUtil` <br />
Command to use to open a file or URL

`refreshInterval` <br />
How often, in seconds, the UI refreshes itself. <br />
**Note:** This implementation is probably wrong and buggy and likely to
change. <br />
Values: A positive integer, `0..n`.

`term` <br />
_Optional_. <br />
Sets a custom value for the terminal type this app runs in. Leave this entry out of the config if you simply want to use your terminal's
default setting. <br />
**Note:** If an invalid value is provided for this setting, the app will
 crash with a `"terminal entry not found"` error. <br />
Values: Any valid terminal type (ie: vt100, xterm, xterm-256color, ansi,
etc.).
