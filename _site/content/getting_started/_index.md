---
title: "Getting Started"
date: 2018-05-21T16:11:58-07:00
draft: false
weight: 1
---

## Quick Start

1. <a href="https://github.com/senorprogrammer/wtf/releases">Download</a> the stand-alone, compiled binary.
2. Unzip the downloaded file.
3. From the command line, `cd` into the newly-created `/wtf` directory.
4. From the command line, run the app: `./wtf`

This should launch the app in your terminal using the default simple
configuration. See <a href="/posts/configuration/">Configuration</a> for
more details.

## Command-line Options

`--config, -c` <br />
Allows you to define a custom config file to use. See <a href="/posts/configuration/">Configuration</a> for more details.

`--help, -h` <br />
Shows help information for the command-line arguments that WTF
takes.

`--module, -m` <br />
Shows help information for the specific named module, if that module
supports help text. <br />
Example: `wtf --module=todo`.

`--version, -v` <br />
Shows version info.

## Keyboard Commands

<span class="caption">Key:</span> `Ctrl-R` <br />
<span class="caption">Action:</span> Force-refresh the data for all modules.

<span class="caption">Key:</span> `Esc` <br />
<span class="caption">Action:</span> Unfocus the currently-focused
widget.

<span class="caption">Key:</span> `Tab` <br />
<span class="caption">Action:</span> Move between focusable modules (`Shift-Tab` to move backwards).
