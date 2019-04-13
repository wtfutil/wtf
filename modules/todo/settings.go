package todo

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	Common *cfg.Common

	FilePath string
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.todo")

	settings := Settings{
		Common:   cfg.NewCommonSettingsFromYAML(ymlConfig),
		FilePath: localConfig.UString("filename"),
	}

	return &settings
}
