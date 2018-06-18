---
title: "Git"
date: 2018-05-09T14:20:48-07:00
draft: false
---

Displays information about local git repositories: branch, changed
files, and recent commits.

<img src="/imgs/modules/git.png" width="720" height="292" alt="git screenshot" />

#### Branch

The name of the currently-active git branch.

#### Changed Files

A list of all the files that have changed since the last
commit, and their status.

#### Recent Commits

A list of `n` recent commits, who committed it, and when.

## Source Code

```bash
wtf/git/
```

## Required ENV Variables

None.

## Keyboard Commands

<span class="caption">Key:</span> `/` <br />
<span class="caption">Action:</span> Open/close the widget's help window.

<span class="caption">Key:</span> `h` <br />
<span class="caption">Action:</span> Show the previous git repository.

<span class="caption">Key:</span> `l` <br />
<span class="caption">Action:</span> Show the next git repository.

<span class="caption">Key:</span> `←` <br />
<span class="caption">Action:</span> Show the previous git repository.

<span class="caption">Key:</span> `→` <br />
<span class="caption">Action:</span> Show the next git repository.

## Configuration

```yaml
git:
  commitCount: 5
  commitFormat: "[forestgreen]%h [grey]%cd [white]%s [grey]%an[white]"
  dateFormat: "%H:%M %d %b %y"
  enabled: true
  position:
    top: 0
    left: 3
    height: 2
    width: 2
  refreshInterval: 8
  repositories:
  - "/Users/chris/go/src/github.com/senorprogrammer/wtf"
  - "/Users/user/fakeapp"
```

### Attributes

`commitCount` <br />
The number of past commits to display. <br />
Values: A positive integer, `0..n`.

`commitFormat` <br />
_Optional_ The string format for the commit message. <br />

`dateFormat` <br />
_Optional_ The string format for the date/time in the commit message.
<br />

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.

`repositories` <br />
Defines which git repositories to watch. <br />
Values: A list of zero or more local file paths pointing to valid git repositories.
