package ping

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "Pings"
)

type Target struct {
	Name string `help: "Name: The label to use for the host you want to ping"`
	Host string `help: "Host: IP address or hostname to ping"`
	Up bool // not meant to be set by user
}

type Settings struct {
	common *cfg.Common
	targets []Target
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
        common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		targets:  buildTargets(ymlConfig),
	}

	return &settings
}

func buildTargets(ymlConfig *config.Config) []Target {
	targets := []Target{}
	yaml := ymlConfig.UList("targets")
	for _, target := range yaml {
		if target,ok := target.(map[string]interface{}); ok {
			for k,v := range target {
				name := k
				host := v.(string)
				t := Target{Name: name, Host: host, Up: false}
				targets = append(targets, t)
			}
		}
	}
	return targets
}

