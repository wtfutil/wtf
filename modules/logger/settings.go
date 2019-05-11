package logger

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "Logger"

type Settings struct {
	common *cfg.Common
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),
	}

	return &settings
}
