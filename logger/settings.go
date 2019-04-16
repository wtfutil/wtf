package logger

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	common *cfg.Common
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),
	}

	return &settings
}
