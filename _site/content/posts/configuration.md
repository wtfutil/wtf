---
title: "Configuration"
date: 2018-04-15T21:17:16-07:00
draft: false
---

By default WTF looks in a `~/.wtf/` directory for a YAML file called
`config.yml`. If the `~/.wtf/` directory doesn't exist, WTF will create that directory
on start-up, and then display instructions for creating a new
configuration file.

In other words, WTF expects to have a YAML config file at: `~/.wtf/config.yml`.

## Example Configuration Files

A couple of example config files are provided in the `_sample_configs/`
directory of the Git repository.

To try out WTF quickly, copy
`simple_config.yml` into `~/.wtf/` as `config.yml` and relaunch WTF. You
should see the app launch and display the <a href="/posts/modules/security/">Security</a>,
<a href="/posts/modules/clocks/">Clocks</a> and <a href="/posts/modules/status/">Status</a> widgets onscreen.

## Custom Configuration Files

To try out different configurations (or run multiple instances of WTF),
you can pass the path to a config file via command line arguments on
start-up.

To load a custom configuration file (ie: one that's not
`~/.wtf/config.yml`), pass in the path to configuration file as a
parameter on launch:

```bash
    $> wtf --config=path/to/custom/config.yml
```
