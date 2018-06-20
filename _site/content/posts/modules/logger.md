---
title: "Logger"
date: 2018-06-16T14:22:18-07:00
draft: false
---

Displays the contents of the WTF log file.

To log to this file in your own modules:

```golang
require "github.com/senorprogrammer/wtf/logger"
 logger.Log("This is a log entry")
```

## Source Code

```bash
wtf/logger/
```

## Required ENV Variables

None.

## Keyboard Commands

Arrow keys scroll through the log file.

## Configuration

```yaml
logger:
  enabled: true
  position:
    top: 5
    left: 4
    height: 2
    width: 1
  refreshInterval: 1
```

### Attributes

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
**Note:** If you're using logging and logging is _disabled_, your logs
will still be written to file, the widget just won't be shown onscreen.
If you have `logger.Log` calls in your code, regardless of this setting,
they will be written out. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.
