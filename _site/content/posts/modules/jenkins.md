---
title: "Modules: Jenkins"
date: 2018-06-09T20:53:35-07:00
draft: false
---

Added in `v0.0.8`.

Displays jenkins status of given builds in a project or view

<img src="/imgs/modules/jenkins.png" alt="jenkins screenshot",
width="320" height="68" alt="jenkins screenshot" />

## Source Code

```bash
wtf/jenkins/
```

## Required ENV Variables

<span class="caption">Key:</span> `WTF_JENKINS_API_KEY` <br />
<span class="caption">Value:</span> Your <a href="https://wiki.jenkins.io/display/JENKINS/Remote+access+API">Jenkins API</a> key.

## Keyboard Commands

None.

## Configuration

```yaml
jenkins:
  enabled: true
  position:
    top: 2
    left: 3
    height: 2
    width: 3
  refreshInterval: 300
  url: "https://jenkins.domain.com/jenkins/view_url"
  user: "username"
```

### Attributes

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed.

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.

`user` <br />
Your Jenkins username. <br />

`url` <br />
The url to your Jenkins project or view. <br />
Values: A valid URI.

