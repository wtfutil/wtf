---
title: "CmdRunner"
date: 2018-05-17T17:17:10-07:00
draft: false
---

Runs a terminal command on a schedule.

## Source Code

```bash
wtf/cmdrunner/
```

## Required ENV Variables

None.

## Keyboard Commands

None.

## Configuration

```yaml
cmdrunner:
  args: ["-g", "batt"]
  cmd: "pmset"
  enabled: true
  position:
    top: 6
    left: 1
    height: 1
    width: 3
  refreshInterval: 30
```

### Attributes

`args` <br />
The arguments to the command, with each item as an element in an array.
Example: for `curl -I cisco.com`, the arguments array would be `["-I", "cisco.com"]`.

`cmd` <br />
The terminal command to be run, withouth the arguments. Ie: `ping`,
`whoami`, `curl`. <br />


`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed.

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.


