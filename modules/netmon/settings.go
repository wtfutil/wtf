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

	showOnly string

	ignoreLoopback bool
	ignoreEthernet bool
	ignoreWireless bool
	ignoreBridges  bool
	ignoreDocker   bool
	ignoreVeth     bool
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		showOnly: ymlConfig.UString("showOnly"),

		ignoreLoopback: ymlConfig.UBool("ignoreLoopback"),
		ignoreEthernet: ymlConfig.UBool("ignoreEthernet"),
		ignoreWireless: ymlConfig.UBool("ignoreWireless"),
		ignoreBridges:  ymlConfig.UBool("ignoreBridges"),
		ignoreDocker:   ymlConfig.UBool("ignoreDocker"),
		ignoreVeth:     ymlConfig.UBool("ignoreVeth"),
	}

	return &settings
}
