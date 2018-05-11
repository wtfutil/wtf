---
title: "Todo"
date: 2018-05-10T12:41:50-07:00
draft:  false
---

## Description

<img src="/imgs/modules/todo.png" width="320" height="388" alt="todo screenshot" />

## Source Code

```bash
wtf/todo/
```

## Required ENV Variables

None.

## Keyboard Commands

A basic, interactive todo list.

## Configuration

```yaml
todo:
  checkedIcon: "X"
  colors:
    checked: gray
    highlight:
      fore: "black"
      back: "orange"
  enabled: true
  filename: "todo.yml"
  position:
    top: 2
    left: 2
    height: 2
    width: 1
  refreshInterval: 3600
```

### Attributes

`checkedIcon` <br />
The icon used to denote a "checked" todo item. <br />
Values: Any displayable unicode character.

`colors.checked` <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11 color</a> name.

`colors.highlight.fore` <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11 color</a> name.

`colors.highlight.back` <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11 color</a> name.

`enabled` <br />
Whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`filename` <br />
The name for the todo file. <br />
*Note:* Currently this file *must* reside in the `~/.wtf/` directory.
This is a <a href="https://github.com/senorprogrammer/wtf/issues/35">known bug</a>. <br />
Values: Any valid filename, ideally ending in `yml`.

`position` <br />
Where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.
