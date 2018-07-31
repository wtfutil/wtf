---
title: "Todoist"
date: 2018-07-05T22:55:55-03:00
draft: false
---

Added in `v0.0.11`.

Displays all items on specified project.

<img src="/imgs/modules/todoist.png" alt="todoist screenshot" />

## Source Code

```bash
wtf/todoist/
```

## Keyboard Commands

<span class="caption">Key:</span> `h` <br />
<span class="caption">Action:</span> Show the previous project.

<span class="caption">Key:</span> `←` <br />
<span class="caption">Action:</span> Show the previous project.

<span class="caption">Key:</span> `l` <br />
<span class="caption">Action:</span> Show the next project.

<span class="caption">Key:</span> `→` <br />
<span class="caption">Action:</span> Show the next project.

<span class="caption">Key:</span> `j` <br />
<span class="caption">Action:</span> Select the next item in the list.

<span class="caption">Key:</span> `↓` <br />
<span class="caption">Action:</span> Select the next item in the list.

<span class="caption">Key:</span> `k` <br />
<span class="caption">Action:</span> Select the previous item in the list.

<span class="caption">Key:</span> `↑` <br />
<span class="caption">Action:</span> Select the previous item in the list.

<span class="caption">Key:</span> `c` <br />
<span class="caption">Action:</span> Close current item.

<span class="caption">Key:</span> `d` <br />
<span class="caption">Action:</span> Delete current item.

<span class="caption">Key:</span> `r` <br />
<span class="caption">Action:</span> Reload all projects.

## Configuration

```yaml
todoist:
  apiKey: "3276d7155dd9ee27b8b14f8743a408a9"
  enabled: true
  position:
    top: 0
    left: 2
    height: 1
    width: 1
  projects:
    - 122247497
  refreshInterval: 3600
```

### Attributes

`apiKey` <br />
Value: Your <a href="https://developer.todoist.com/sync/v7/">Todoist API</a> token.

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Where in the grid this module's widget will be displayed. <br />

`projects` <br />
The todoist projects to fetch items from. <br />
Values: The integer ID of the project.

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.
