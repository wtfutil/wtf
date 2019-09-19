package docker

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "docker"
)

// Settings defines the configuration options for this module
type Settings struct {
	common     *cfg.Common
	labelColor string
}

// NewSettingsFromYAML creates and returns an instance of Settings with configuration options populated
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common:     cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		labelColor: ymlConfig.UString("labelColor", "white"),
	}

	return &settings
}
