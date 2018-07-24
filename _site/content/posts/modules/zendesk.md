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

## Configuration

```yaml
zendesk:
      enabled: true
      username: "your_email@acme.com"
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
