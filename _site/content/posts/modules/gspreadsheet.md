---
title: "Google Spreadsheets"
date: 2018-06-10T18:26:26-04:00
draft: false
---

Added in `v0.0.7`.

Display information from cells in a Google Spreadsheet.

```bash
wtf/gspreadsheets/
```

## Required ENV Variables

None.

## Keyboard Commands

None.

## Configuration

```yaml
gspreadsheets:
  colors:
    values: "green"
  cells:
    names:
    - "Cell 1 name"
    - "Cell 2 name"
    addresses:
    - "A1"
    - "A2"
  enabled: true
  position:
    top: 0
    left: 0
    width: 1
    height: 1
  refreshInterval: "300"
  secretFile: "~/.config/wtf/gspreadsheets/client_secret.json"
  sheetId: "id_of_google_spreadsheet"
```

### Attributes

`colors.values` <br />
The color to display the cell values in. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11 color</a> name.

`cells.names` <br />

`cells.addresses` <br />

`enabled` <br />
Whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.

`secretFile` <br />
Your <a href="https://developers.google.com/sheets/api/quickstart/go">Google client secret</a> JSON file. <br />
Values: A string representing a file path to the JSON secret file.
