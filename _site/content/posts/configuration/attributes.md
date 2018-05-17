---
title: "Attributes"
date: 2018-05-16T21:51:23-07:00
draft: false
---

This page describes the top-level attributes that can be added to
`config.yml`.

```yaml
wtf:
  colors:
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
  refreshInterval: 1
```

### Attributes

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

`refreshInterval` <br />
How often, in seconds, the UI refreshes itself. <br />
**Note:** This implementation is probably wrong and buggy and likely to
change. <br />
Values: A positive integer, `0..n`.
