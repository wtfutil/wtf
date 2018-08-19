---
title: "Textfile"
date: 2018-05-09T11:13:11-07:00
draft: false
weight: 210
---

<img class="screenshot" src="/imgs/modules/textfile.png" width="320" height="133" alt="textfile screenshot" />

Displays the contents of the specified text file in the widget.

## Source Code

```bash
wtf/textfile/
```

## Keyboard Commands

<span class="caption">Key:</span> `/` <br />
<span class="caption">Action:</span> Open/close the widget's help window.

<span class="caption">Key:</span> `o` <br />
<span class="caption">Action:</span> Opens the text file in whichever text editor is associated  with that file type.

## Configuration

```yaml
textfile:
  enabled: true
  filePath: "~/Desktop/notes.md"
  format: true
  formatStyle: "dracula"
  position:
    top: 5
    left: 4
    height: 2
    width: 1
  refreshInterval: 15
```

### Attributes

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`filePath` <br />
The path to the file to be displayed in the widget. <br />

`format` <br />
Whether or not to try and format and syntax highlight the displayedtext. <br />
Values: `true`, `false`. <br />
Default: `false`.

`formatStyle` <br />
The style of syntax highlighting to format the text with. <br />
Values: See [Chroma styles](https://github.com/alecthomas/chroma/tree/master/styles) for all
valid options. <br />
Default: `vim`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.
