---
title: "Security"
date: 2018-05-08T20:33:28-07:00
draft: false
---

## Description

Displays some general information about the state of the machine's wifi
connection, firewall, and DNS settings.

#### Wifi Network

<ul class="list-ornate">
  <li>The name of the current network</li>
  <li>Whether or not the network uses <a href="https://www.howtogeek.com/167783/htg-explains-the-difference-between-wep-wpa-and-wpa2-wireless-encryption-and-why-it-matters/">encryption</a> and if so, what flavour</li>
</ul>

#### Firewall

<ul class="list-ornate">
<li>Whether or not the <a href="https://support.apple.com/en-ca/HT201642">firewall</a> is enabled</li>
<li>Whether or not <a href="https://support.apple.com/en-ca/HT201642">Stealth Mode</a> is enabled</li>
</ul>

#### DNS

<ul class="list-ornate">
<li>Which <a hre="https://developers.cloudflare.com/1.1.1.1/what-is-1.1.1.1/">DNS resolvers</a> (servers) the machine is configured to use</li>
</ul>

<img src="/imgs/modules/security.png" width="320" height="192" alt="clocks screenshot" />

## Location

```bash
wtf/security
```

## Required ENV Variables

None.

## Keyboard Commands

None.

## Configuration

```yaml
security:
  enabled: true
  position:
    top: 1
    left: 2
    height: 1
    width: 1
  refreshInterval: 3600
```

### Attributes

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: Any positive integer, `0...n`.
