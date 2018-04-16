---
title: "Configuration"
date: 2018-04-15T21:17:16-07:00
draft: false
---

By default WTF looks in a `~/.wtf/` directory for a YAML file called
`config.yml`. If `~/.wtf/` doesn't exist, WTF will create that directory
on start-up, and then display instructions for creating the
configuration file.

## Config.yml

## Example Config Files

A few example config files are provided in the `_sample_configs/`
directory of the Git repository. To try out WTF quickly, copy
`simple_config.yml` into `~/.wtf/` as `config.yml` and relaunch WTF. You
should see the app launch and display the _Security_ and _Status_
modules.

## Custom Configuration Files

To load a custom configuration file (ie: one that's not
`~/.wtf/config.yml`), pass in the path to configuration file as a
parameter on launch:
```bash
    $> wtf --config=path/to/custom/config.yml
```
Example:
```bash
    %> wtf --config=~/Documents/monitoring.yml
```

This is also the easiest way to run multiple instances of WTF, should
you want to run multiple independent dashboards.
