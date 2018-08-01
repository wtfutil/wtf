---
title: "Zendesk"
date: 2018-07-23T18:55:37-08:00
draft: false
---

Added in `v0.1.0`.

Displays tickets in the "New" status - i.e. have not yet been assigned.

## Source Code

```bash
wtf/zendesk/
```

## Keyboard Commands

<span class="caption">Key:</span> `[return]` <br />
<span class="caption">Action:</span> Open the selected ticket in the browser.

<span class="caption">Key:</span> `j` <br />
<span class="caption">Action:</span> Select the next item in the list.

<span class="caption">Key:</span> `k` <br />
<span class="caption">Action:</span> Select the previous item in the list.

<span class="caption">Key:</span> `↓` <br />
<span class="caption">Action:</span> Scroll down the list.

<span class="caption">Key:</span> `↑` <br />
<span class="caption">Action:</span> Scroll up the list.

## Configuration

```yaml
zendesk:
  apiKey: "3276d7155dd9ee27b8b14f8743a408a9"
  enabled: true
  position:
    top: 0
    left: 2
    height: 1
    width: 1
  status: "new"
  subdomain: "your_domain"
  username: "your_email@acme.com"
```

### Attributes

`apiKey` <br />
Value: Your Zendesk API token.

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`status` <br />
The status of tickets you want to retrieve.
Values: `new`, `open`, `pending`, `hold`.

`subdomain` <br />
Value: Your Zendesk subdomain.

`username` <br />
Your Zendesk username
Values: A valid Zendesk username (usually an email address).

