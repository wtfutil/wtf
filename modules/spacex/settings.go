package spacex

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
)

type Settings struct {
	common *cfg.Common
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	spacex := ymlConfig.UString("spacex")
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, spacex, defaultFocusable, ymlConfig, globalConfig),
	}
	return &settings
}
