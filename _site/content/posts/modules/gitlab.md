---
title: "Gitlab"
date: 2018-06-08T13:14:11-07:00
draft: false
---

Added in `v0.0.8`.

<img src="/imgs/modules/gitlab.png" width="640" height="384" alt="gitlab screenshot" />

Displays information about your projects hosted on Gitlab:

#### Open Approval Requests

All open merge requests that are requesting your approval.

#### Open Merge Requests

All open merge requests created by you.

## Source Code

```bash
wtf/gitlab/
```

## Required ENV Variables

<span class="caption">Key:</span> `WTF_GITLAB_TOKEN` <br />
<span class="caption">Action:</span> A <a href="https://docs.gitlab.com/ce/user/profile/personal_access_tokens.html">Gitlab personal access token</a>. Requires at least `api` access.

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
gitlab:
  enabled: true
  position:
    top: 2
    left: 3
    height: 2
    width: 2
  refreshInterval: 300
  projects:
    tasks: "gitlab-org/release"
    gitlab-ce: "gitlab-org"
  username: "senorprogrammer"
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
_Optional_. Your Gitlab corporate domain. <br />
Values: A valid URI.

`projects` <br />
A list of key/value pairs each describing a Gitlab project to fetch data
for. <br />
<span class="caption">Key:</span> The name of the project. <br />
<span class="caption">Value:</span> The namespace of the project.

`username` <br />
Your Gitlab username. Used to figure out which requests require your approval
