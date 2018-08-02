---
title: "Gerrit"
date: 2018-06-27T15:55:42-07:00
draft: false
---

Displays information about your projects hosted on Gerrit:

#### Open Incoming Reviews

All open reviews that are requesting your approval.

#### My Outgoing Reviews

All open reviews created by you.

<img src="/imgs/modules/gerrit.png" width="640" height="167" alt="gerrit screenshot" />

## Source Code

```bash
wtf/gerrit/
```

## Keyboard Commands

<span class="caption">Key:</span> `/` <br />
<span class="caption">Action:</span> Open/close the widget's help window.

<span class="caption">Key:</span> `h` <br />
<span class="caption">Action:</span> Show the previous project.

<span class="caption">Key:</span> `l` <br />
<span class="caption">Action:</span> Show the next project.

<span class="caption">Key:</span> `j` <br />
<span class="caption">Action:</span> Select the next review in the list.

<span class="caption">Key:</span> `k` <br />
<span class="caption">Action:</span> Select the previous review in the list.

<span class="caption">Key:</span> `r` <br />
<span class="caption">Action:</span> Refresh the data.

<span class="caption">Key:</span> `←` <br />
<span class="caption">Action:</span> Show the previous project.

<span class="caption">Key:</span> `→` <br />
<span class="caption">Action:</span> Show the next project.

<span class="caption">Key:</span> `↓` <br />
<span class="caption">Action:</span> Select the next review in the list.

<span class="caption">Key:</span> `↑` <br />
<span class="caption">Action:</span> Select the previous review in the list.

<span class="caption">Key:</span> `[return]` <br />
<span class="caption">Action:</span> Open the selected review in the browser.

## Configuration

```yaml
gerrit:
  colors:
    rows:
      even: "lightblue"
      odd: "white"
  domain: https://gerrit-review.googlesource.com
  enabled: true
  password: "mypassword"
  position:
    top: 2
    left: 3
    height: 2
    width: 2
  projects:
  - org/test-project"
  - dotfiles
  refreshInterval: 300
  username: "myname"
  verifyServerCertificate: false
```

### Attributes

`colors.rows.even` <br />
Define the foreground color for even-numbered rows. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`colors.rows.odd` <br />
Define the foreground color for odd-numbered rows. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`domain` <br />
Your Gerrit corporate domain. <br />
Values: A valid URI.

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`password` <br />
Value: Your <a href="https://gerrit-review.googlesource.com/Documentation/user-upload.html#http">Gerrit HTTP Password</a>.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`projects` <br />
A list of Gerrit project names to fetch data for. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.

`username` <br />
Your Gerrit username.

`verifyServerCertificate` <br />
_Optional_ <br />
Determines whether or not the server's certificate chain and host name are verified. <br />
Values: `true`, `false`.
