---
title: "Zendesk"
date: 2018-07-23T18:55:37-08:00
draft: false
---

Displays tickets in the "New" status - i.e. have not yet been assigned.

## Source Code

```bash
wtf/zendesk/
```

## Required ENV Variables

<span class="caption">Key:</span> `ZENDESK_API` <br />
<span class="caption">Value:</span> Your Zendesk API Token

<span class="caption">Key:</span> `ZENDESK_DOMAIN` <br />
<span class="caption">Value:</span> Your Zendesk subdomain

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
      enabled: true
      username: "your_email@acme.com"
      status: "new"
      position:
        top: 0
        left: 2
        height: 1
        width: 1
```

### Attributes

`username` <br />
Your Zendesk username
Values: A valid Zendesk username (usually an email address).

`status` <br />
The status of tickets you want to retrieve.
Values: `new`, `open`, `pending`, `hold`.
