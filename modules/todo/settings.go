package todo

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	common *cfg.Common

	filePath string
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.todo")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		filePath: localConfig.UString("filename"),
	}

	return &settings
}
