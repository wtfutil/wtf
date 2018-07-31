---
title: "Trello"
date: 2018-05-10T10:44:35-07:00
draft: false
---

Displays all Trello cards on specified lists.

<img src="/imgs/modules/trello.png" width="640" height="188" alt="trello screenshot" />

## Source Code

```bash
wtf/trello/
```

## Keyboard Commands

None.

## Configuration

### Single Trello List

```yaml
trello:
  accessToken: "7b8b14f8743a408a93276d7155dd9ee2"
  apiKey: "3276d7155dd9ee27b8b14f8743a408a9"
  board: Main
  enabled: true
  list: "Todo"
  position:
    height: 1
    left: 2
    top: 0
    width: 1
  refreshInterval: 3600
  username: myname
```

### Multiple Trello Lists

If you want to monitor multiple Trello lists, use the following
configuration (note the difference in `list`):

```yaml
trello:
  accessToken: "7b8b14f8743a408a93276d7155dd9ee2"
  apiKey: "3276d7155dd9ee27b8b14f8743a408a9"
  board: Main
  enabled: true
  list: ["Todo", "Done"]
  position:
    height: 1
    left: 2
    top: 0
    width: 1
  refreshInterval: 3600
  username: myname
```

### Attributes

`accessToken` <br />
Value: Your Trello access token.

`apiKey` <br />
Value: Your Trello API key.

`board` <br />
The name of the Trello board. <br />

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`list` <br />
The Trello lists to fetch cards from. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.

`username` <br />
Your Trello username. <br />

`position` <br />
Where in the grid this module's widget will be displayed. <br />
