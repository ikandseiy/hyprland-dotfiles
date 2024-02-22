# Powerbar

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/adamveld12/powerbar)
[![Build Status](https://veldhousen.semaphoreci.com/badges/powerbar.svg?style=shields)](https://veldhousen.semaphoreci.com/projects/powerbar)
[![Go Report Card](https://goreportcard.com/badge/github.com/adamveld12/powerbar)](https://goreportcard.com/report/github.com/adamveld12/powerbar)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/adamveld12/powerbar)
[![AUR version](https://img.shields.io/aur/version/powerbar)](https://aur.archlinux.org/packages/powerbar/)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Configurable battery status module for [Waybar](). Uses UPower and [DBus](https://upower.freedesktop.org/docs/Device.html).

## Usage

Command line:

```text
𝝺 powerbar --help
Usage of powerbar:
  -charging string
     format string (default "{state} {capacity}% - {usage}W - {H}h {M}m")
  -discharging string
     format string (default "{state} {capacity}% - {usage}W - {H}h {M}m")
  -full string
     format string (default "{state} {capacity}% - {usage}W - {H}h {M}m")
  -waybar
     enable waybar mode
```

Variables:

- `{capacity}`: Percentage of the battery capacity left
- `{H}`, `{M}`: Hour and Minutes left when `state` is discharging, if charging this indicates how long until 100%.
- `{usage}`: How much power is used in watts when `state` is discharging, If charging this indicates how much the battery is receiving.
- `{state}`: The state of the battery. Can be `Charging`, `Discharging`, `Unknown` or `Fully Charged`

### Using with Waybar

An example waybar config:

```json
{
    "layer": "top",
    "position": "top",
    "modules-left": [],
    "modules-center": [],
    "modules-right": [ "custom/powerbar"],

    "custom/powerbar": {
        "return-type": "json",
        "interval": 1,
        "exec": "$HOME/.config/waybar/powerbar -full 'FULL - {usage}W' -charging '{state} {capacity}% - {usage}W - {H}h {M}m' -waybar 2> /dev/null"
    },
}
```

Classes are lower cased, hypened variants of the possible values for `state`

```css
#custom-powerbar {
    border-bottom: 0px solid rgba(0, 255, 0, 1);
}

#custom-powerbar.charging,
#custom-powerbar.fully-charged {
    border-bottom: 2px solid rgb(0, 255, 0);
}

#custom-powerbar.discharging {
    border-bottom: 2px solid rgb(255, 217, 0);
}
```

## LICENSE

GPL-V3
