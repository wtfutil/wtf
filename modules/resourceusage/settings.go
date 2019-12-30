package resourceusage

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "ResourceUsage"
)

type Settings struct {
	common      *cfg.Common
	cpuCombined bool
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common:      cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		cpuCombined: ymlConfig.UBool("cpuCombined", false),
	}

	return &settings
}
