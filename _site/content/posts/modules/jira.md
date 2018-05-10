---
title: "Jira"
date: 2018-05-10T10:44:35-07:00
draft: false
---

## Description

Displays all Jira issues assigned to you for the specified project.

<img src="/imgs/modules/jira.png" width="640" height="188" alt="jira screenshot" />

## Source Code

```bash
wtf/jira/
```

## Required ENV Variables

<span class="caption">Key:</span> `WTF_JIRA_API_KEY` <br />
<span class="caption">Value:</span> Your <a href="https://confluence.atlassian.com/cloud/api-tokens-938839638.html">Jira API</a> key.

## Keyboard Commands

None.

## Configuration

```yaml
jira:
  domain: "https://umbrellacorp.atlassian.net"
  email: "chriscummer@me.com"
  enabled: true
  position:
    top: 4
    left: 1
    height: 1
    width: 2
  project: "CORE"
  refreshInterval: 900
  username: "chris.cummer"
```

### Attributes

`domain` <br />

`email` <br />

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`project` <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.

`username` <br />
