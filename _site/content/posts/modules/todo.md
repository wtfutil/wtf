---
title: "Todo"
date: 2018-05-10T12:41:50-07:00
draft:  false
---

An interactive todo list.

<img src="/imgs/modules/todo.png" width="320" height="388" alt="todo screenshot" />

## Source Code

```bash
wtf/todo/
```

## Required ENV Variables

None.

## Keyboard Commands

<span class="caption">Key:</span> `[return]` <br />
<span class="caption">Action:</span> Edit the selected item. <br />
<span class="caption">Action:</span> Close the modal item dialog and save changes. <br />

<span class="caption">Key:</span> `[esc]` <br />
<span class="caption">Action:</span> Remove focus from the selected item. <br />
<span class="caption">Action:</span> Close the modal item dialog without saving changes.

<span class="caption">Key:</span> `[space]` <br />
<span class="caption">Action:</span> Check/uncheck the selected item.

<span class="caption">Key:</span> `/` <br />
<span class="caption">Action:</span> Open/close the widget's help window.

<span class="caption">Key:</span> `j` <br />
<span class="caption">Action:</span> Select the next item in the list.

<span class="caption">Key:</span> `k` <br />
<span class="caption">Action:</span> Select the previous item in the list.

<span class="caption">Key:</span> `n` <br />
<span class="caption">Action:</span> Create a new list item.

<span class="caption">Key:</span> `o` <br />
<span class="caption">Action:</span> Opens the todo list file in
whichever text editor is associated with that file type.

<span class="caption">Key:</span> `↓` <br />
<span class="caption">Action:</span> Select the next item in the list.

<span class="caption">Key:</span> `↑` <br />
<span class="caption">Action:</span> Select the previous item in the list.

<span class="caption">Key:</span> `Ctrl-d` <br />
<span class="caption">Action:</span> Delete the selected item.

<span class="caption">Key:</span> `Ctrl-J` <br />
<span class="caption">Action:</span> Move the selected item down the list.

<span class="caption">Key:</span> `Ctrl-K` <br />
<span class="caption">Action:</span> Move the selected item up the list.

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
The foreground color for the currently-selected row. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11 color</a> name.

`colors.highlight.back` <br />
The background color for the currently-selected row. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11 color</a> name.

`enabled` <br />
Whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`filename` <br />
The name for the todo file. <br />
Values: Any valid filename, ideally ending in `yml`.

`position` <br />
Where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.
