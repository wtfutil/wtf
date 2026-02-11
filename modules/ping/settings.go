package ping

import (
	"fmt"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = false
	defaultTitle     = "Pings"
)

type Host struct {
	Label    string `help:"Label: The name to use for the host you want to ping. Uses hostname if blank."`
	Hostname string `help:"Hostname: IP address or hostname to ping"`
	Up       bool   // not meant to be set by user
}

type Settings struct {
	common *cfg.Common
	hosts  []Host
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		hosts:  buildhosts(ymlConfig),
	}

	return &settings
}

func (widget *Widget) ConfigText() string {
	return utils.HelpFromInterface(Settings{})
}

func buildhosts(ymlConfig *config.Config) []Host {

	hosts := []Host{}
	yaml := ymlConfig.UList("hosts")

	// Iterate through each host in the config
	for _, rawHost := range yaml {

		host, ok := rawHost.(map[string]interface{})
		if !ok {
			continue // bad host, skip
		}

		hostname, ok := host["hostname"].(string)
		if !ok {
			continue // hostname is required, skip
		}

		if hostname == "" {
			continue // hostname is required, skip
		}

		label := hostname // a default if missing from config
		if value, ok := host["label"]; ok {
			// Using Sprintf here instead of a string assert. This is to cover the
			// case where someone puts a number as the label instead of a YAML string.
			// Weird case, yes, but wanted to prevent runtime errors.
			label = fmt.Sprintf("%v", value)
		}

		hosts = append(hosts, Host{Label: label, Hostname: hostname, Up: false})
	}
	return hosts
}
