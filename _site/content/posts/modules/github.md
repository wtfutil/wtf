---
title: "GitHub"
date: 2018-05-09T19:20:20-07:00
draft: false
---

Displays information about your git repositories hosted on Github:


#### Open Review Requests

All open code review requests assigned to you.

#### Open Pull Requests

All open pull requests created by you.

<img src="/imgs/modules/github.png" width="640" height="384" alt="github screenshot" />

## Source Code

```bash
wtf/github/
```

## Github Required ENV Variables

<span class="caption">Key:</span> `WTF_GITHUB_TOKEN` <br />
<span class="caption">Action:</span> Your <a href="https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization">Github API</a> token.

## GitHub Enterprise Required ENV Variables

<span class="caption">Key:</span> `WTF_GITHUB_TOKEN` <br />
<span class="caption">Action:</span> Your <a href="https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization">Github API</a> token.

<span class="caption">Key:</span> `WTF_GITHUB_BASE_URL` <br />
<span class="caption">Action:</span> Your <a href="https://developer.github.com/enterprise/2.13/v3/enterprise-admin/">Github Enterprise</a> API URL.

<span class="caption">Key:</span> `WTF_GITHUB_UPLOAD_URL` <br />
<span class="caption">Action:</span> Your <a href="https://developer.github.com/enterprise/2.13/v3/enterprise-admin/">Github Enterprise</a> upload URL (often the same as API URL).

## Keyboard Commands

<span class="caption">Key:</span> `/` <br />
<span class="caption">Action:</span> Open/close the widget's help window.

<span class="caption">Key:</span> `h` <br />
<span class="caption">Action:</span> Show the previous git repository.

<span class="caption">Key:</span> `l` <br />
<span class="caption">Action:</span> Show the next git repository.

<span class="caption">Key:</span> `←` <br />
<span class="caption">Action:</span> Show the previous git repository.

<span class="caption">Key:</span> `→` <br />
<span class="caption">Action:</span> Show the next git repository.

## Configuration

```yaml
github:
  enabled: true
  position:
    top: 2
    left: 3
    height: 2
    width: 2
  refreshInterval: 300
  repositories:
    wesker-api: "UmbrellaCorp"
    wtf: "senorprogrammer"
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

`repositories` <br />
A list of key/value pairs each describing a Github repository to fetch data
for. <br />
<span class="caption">Key:</span> The name of the repository. <br />
<span class="caption">Value:</span> The name of the account or organization that owns the repository.

`username` <br />
Your Github username. Used to figure out which review requests you've
been added to.
