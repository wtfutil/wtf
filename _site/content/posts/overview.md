---
title: "Overview"
date: 2018-05-21T16:11:58-07:00
draft: false
---

<span style="font-family: monospace; font-size: 1.6em;">WTF</span> is a personal information dashboard for your terminal, developed for those who spend most of their day in the command line.

It allows you to monitor services and systems that you otherwise might
keep browser tabs open for, the kinds of things you don't always need
visible, but might check in on every now and then.

<a href="https://github.com/senorprogrammer/wtf/releases">Download it</a> as a stand-alone, compiled binary, or install it <a href="https://github.com/senorprogrammer/wtf">from source</a>.

Once installed, edit your <a
href="/posts/configuration/">configuration file</a> and define the
modules you want to run.

Configuration instructions for each module can be found in the module
pages, listed at left.

## Command-line Options

`--config` <br />
Allows you to define a custom config file to use. See <a href="/posts/configuration/">Configuration</a> for more details.

`-h` <br />
Shows help information for the command-line arguments that WTF
takes.

`--help [module name]` <br />
Shows help information for the specific named module, if that module
supports help text. Ex: `wtf --help git`.

`--version` <br />
Shows version info.

## Keyboard Commands

<span class="caption">Key:</span> `Ctrl-R` <br />
<span class="caption">Action:</span> Force-refresh the data for all modules.

<span class="caption">Key:</span> `Esc` <br />
<span class="caption">Action:</span> Unfocus the currently-focused
widget.

<span class="caption">Key:</span> `Tab` <br />
<span class="caption">Action:</span> Move between focusable modules (`Shift-Tab` to move backwards).
