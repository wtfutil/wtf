---
title: "Gitter"
date: 2018-08-02T12:36:08-04:00
draft: false
---

Added in `v0.2.1`.

Displays chat messages from Gitter.

<img src="/imgs/modules/gitter.png" width="847" height="160" alt="gitter screenshot" />

## Source Code

```bash
wtf/gitter/
```

## Keyboard Commands

<span class="caption">Key:</span> `j` <br />
<span class="caption">Action:</span> Select the next message in the list.

<span class="caption">Key:</span> `k` <br />
<span class="caption">Action:</span> Select the previous message in the list.

<span class="caption">Key:</span> `r` <br />
<span class="caption">Action:</span> Refresh the data.

<span class="caption">Key:</span> `↓` <br />
<span class="caption">Action:</span> Select the next message in the list.

<span class="caption">Key:</span> `↑` <br />
<span class="caption">Action:</span> Select the previous message in the list.

## Configuration

```yaml
gitter:
  apiToken: "ab345546asdfasb465234fgjgh068f39a35c3e4139ee383f7"
  enabled: true
  numberOfMessages: 10
  position:
    top: 4
    left: 1
    height: 1
    width: 4
  roomUri: wtfutil/Lobby
  refreshInterval: 300
```

### Attributes

`apiToken` <br />
Value: Your <a href="https://developer.gitter.im/apps">Gitter</a>Personal Access Token.

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`numberOfMessages` <br />
_Optional_ <br />
Maximum number of _(newest)_ messages to be displayed. Default is `10`<br />

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`roomUri` <br />
_Optional_ <br /
URI of the room you would like to see the chat messages from. Default is `wtfutil/Lobby`<br />
Values: `new`, `top`, `job`, `ask`

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.
