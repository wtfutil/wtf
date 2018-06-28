---
title: "Gerrit"
date: 2018-06-27T15:55:42-07:00
draft: false
---

<img src="/imgs/modules/gerrit.png" width="640" height="384" alt="gerrit screenshot" />

Displays information about your projects hosted on Gerrit:

#### Open Incoming Reviews

All open reviews that are requesting your approval.

#### My Outgoing Reviews

All open reviews created by you.

## Source Code

```bash
wtf/gerrit/
```

## Required ENV Variables

<span class="caption">Key:</span> `WTF_GERRIT_PASSWORD` <br />
<span class="caption">Action:</span> Your <a href="https://gerrit-review.googlesource.com/Documentation/user-upload.html#http">Gerrit HTTP Password</a>.

## Keyboard Commands

<span class="caption">Key:</span> `/` <br />
<span class="caption">Action:</span> Open/close the widget's help window.

<span class="caption">Key:</span> `h` <br />
<span class="caption">Action:</span> Show the previous project.

<span class="caption">Key:</span> `l` <br />
<span class="caption">Action:</span> Show the next project.

<span class="caption">Key:</span> `←` <br />
<span class="caption">Action:</span> Show the previous project.

<span class="caption">Key:</span> `→` <br />
<span class="caption">Action:</span> Show the next project.

## Configuration

```yaml
gerrit:
  enabled: true
  position:
    top: 2
    left: 3
    height: 2
    width: 2
  refreshInterval: 300
  projects:
  - org/test-project"
  - dotfiles
  username: "myname"
```

### Attributes

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.

`domain` <br />
_Optional_. Your Gerrit corporate domain. <br />
Values: A valid URI.

`projects` <br />
A list of Gerrit project names to fetch data for. <br />

`username` <br />
Your Gerrit username. <br />

`verifyServerCertificate` <br />
_Optional_ <br />
Determines whether or not the server's certificate chain and host name are verified. <br />
Values: `true`, `false`.
