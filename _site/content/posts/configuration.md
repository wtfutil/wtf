---
title: "Configuration"
date: 2018-04-15T21:17:16-07:00
draft: false
---

## Index

* [Configuration Files](#configuration-files)
* [Environment (ENV) Variables](#environment-env-variables)
* [Grid Layout](#grid-layout)

## Configuration Files

By default WTF looks in a `~/.config/wtf/` directory for a YAML file called
`config.yml`. If the `~/.config/wtf/` directory doesn't exist, WTF will create that directory
on start-up, and then display instructions for creating a new
configuration file.

In other words, WTF expects to have a YAML config file at: `~/.config/wtf/config.yml`.

#### Example Configuration Files

A couple of example config files are provided in the `_sample_configs/`
directory of the Git repository.

To try out WTF quickly, copy
`simple_config.yml` into `~/.config/wtf/` as `config.yml` and relaunch WTF. You
should see the app launch and display the <a href="/posts/modules/security/">Security</a>,
<a href="/posts/modules/clocks/">Clocks</a> and <a href="/posts/modules/status/">Status</a> widgets onscreen.

#### Custom Configuration Files

To try out different configurations (or run multiple instances of WTF),
you can pass the path to a config file via command line arguments on
start-up.

To load a custom configuration file (ie: one that's not
`~/.config/wtf/config.yml`), pass in the path to configuration file as a
parameter on launch:

```bash
    $> wtf --config=path/to/custom/config.yml
```

#### Configuration Attributes

A number of top-level attributes can be set to customize your WTF
install. See <a href="/posts/configuration/attributes/">Attributes</a> for detials.

## Environment (ENV) Variables

Some modules require the presence of environment variables to function
properly. Usually these are API keys or other sensitive data that one
wouldn't want to have laying about in the config files.

For modules that require them, the name of the required environment
variable(s) can be found in that module's "Required ENV Variables"
section of the documentation. See <a href="/posts/modules/opsgenie/">OpsGenie</a> for an example.

## Grid Layout

WTF uses the `Grid` layout system from [tview](https://github.com/rivo/tview/blob/master/grid.go) to position widgets
onscreen. It's not immediately obvious how this works, so here's an
explanation:

Think of your terminal screen as a matrix of letter positions, say `100` chrs wide and `58` chrs tall.

Columns breaks up the width of the screen into chunks, each chunk a specified number of characters wide. use

`[10, 10, 10, 10, 10, 10, 10, 10, 10, 10]`

Ten columns that are ten characters wide

Rows break up the height of the screen into chunks, each chunk a specified number of characters tall. If we wanted to have five rows:

`[10, 10, 10, 10, 18]`

The co-ordinate system starts at top-left and defines how wide and tall a widget is. If we wanted to put a 2-col, 2-row widget in the bottom of the screen, we'd position it at:

```
  top: 4     // top starts in the 4th row
  left: 9    // left starts in the 9th column
  height: 2  // span down rows 4 & 5 (18 characters in size, total)
  width: 2   // span across cols 9 & 10 (20 characters in size, total)
```
