package netmon

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "Network Monitor"
)

type Settings struct {
	*cfg.Common

	ignoreLoopback bool
	ignoreBridges  bool
	ignoreDocker   bool
	ignoreVETH     bool
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		ignoreLoopback: ymlConfig.UBool("ignoreLoopback"),
		ignoreBridges:  ymlConfig.UBool("ignoreBridges"),
		ignoreDocker:   ymlConfig.UBool("ignoreDocker"),
		ignoreVETH:     ymlConfig.UBool("ignoreVETH"),
	}

	return &settings
}
