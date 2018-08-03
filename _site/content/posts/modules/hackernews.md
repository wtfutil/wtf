---
title: "Hacker News"
date: 2018-08-02T16:36:08-04:00
draft: false
---

Added in `v0.1.2`.

Displays stories from Hacker News.

<img src="/imgs/modules/hackernews.png" width="843" height="201" alt="hackernews screenshot" />

## Source Code

```bash
wtf/hackernews/
```

## Keyboard Commands

<span class="caption">Key:</span> `[return]` <br />
<span class="caption">Action:</span> Open the selected story in the browser.

<span class="caption">Key:</span> `j` <br />
<span class="caption">Action:</span> Select the next story in the list.

<span class="caption">Key:</span> `k` <br />
<span class="caption">Action:</span> Select the previous story in the list.

<span class="caption">Key:</span> `r` <br />
<span class="caption">Action:</span> Refresh the data.

<span class="caption">Key:</span> `↓` <br />
<span class="caption">Action:</span> Select the next story in the list.

<span class="caption">Key:</span> `↑` <br />
<span class="caption">Action:</span> Select the previous story in the list.

## Configuration

```yaml
hackernews:
  enabled: true
  numberOfStories: 10
  position:
    top: 4
    left: 1
    height: 1
    width: 2
  storyType: top
  refreshInterval: 900
```

### Attributes

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`numberOfStories` <br />
_Optional_ <br />
Defines number of stories to be displayed. Default is `10`<br />

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`storyType` <br />
_Optional_ <br /
Defines type of stories to be displayed. Default is `top` stories<br />
Values: `new`, `top`, `job`, `ask`

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.
