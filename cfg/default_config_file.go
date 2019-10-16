package cfg

const defaultConfigFile = `wtf:
  colors:
    border:
      focusable: darkslateblue
      focused: orange
      normal: gray
  grid:
    columns: [32, 32, 32, 32, 90]
    rows: [10, 10, 10, 4, 4, 90]
  refreshInterval: 1
  mods:
    clocks_a:
      colors:
        rows:
          even: "lightblue"
          odd: "white"
      enabled: true
      locations:
        Vancouver: "America/Vancouver"
        Toronto: "America/Toronto"
      position:
        top: 0
        left: 1
        height: 1
        width: 1
      refreshInterval: 15
      sort: "alphabetical"
      title: "Clocks A"
      type: "clocks"
    clocks_b:
      colors:
        rows:
          even: "lightblue"
          odd: "white"
      enabled: true
      locations:
        Avignon: "Europe/Paris"
        Barcelona: "Europe/Madrid"
        Dubai: "Asia/Dubai"
      position:
        top: 0
        left: 2
        height: 1
        width: 1
      refreshInterval: 15
      sort: "alphabetical"
      title: "Clocks B"
      type: "clocks"
    feedreader:
      enabled: true
      feeds:
      - http://wtfutil.com/blog/index.xml
      feedLimit: 10
      position:
        top: 1
        left: 1
        width: 2
        height: 1
      updateInterval: 14400
    ipinfo:
      colors:
        name: "lightblue"
        value: "white"
      enabled: true
      position:
        top: 2
        left: 1
        height: 1
        width: 1
      refreshInterval: 150
    power:
      enabled: true
      position:
        top: 2
        left: 2
        height: 1
        width: 1
      refreshInterval: 15
      title: "⚡️"
    textfile:
      enabled: true
      filePath: "~/.config/wtf/config.yml"
      format: true
      position:
        top: 0
        left: 0
        height: 4
        width: 1
      refreshInterval: 30
      wrapText: false
    uptime:
      args: [""]
      cmd: "uptime"
      enabled: true
      position:
        top: 3
        left: 1
        height: 1
        width: 2
      refreshInterval: 30
      type: cmdrunner
`
