package power

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "power"

type Settings struct {
	common *cfg.Common

	filePath string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),
	}

	return &settings
}
