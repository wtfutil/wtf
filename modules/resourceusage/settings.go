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
	showCPU     bool
	showMem     bool
	showSwp     bool
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common:      cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		cpuCombined: ymlConfig.UBool("cpuCombined", false),
		showCPU:     ymlConfig.UBool("showCPU", true),
		showMem:     ymlConfig.UBool("showMem", true),
		showSwp:     ymlConfig.UBool("showSwp", true),
	}

	return &settings
}
