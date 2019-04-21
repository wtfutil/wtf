package textfile

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "textfile"

type Settings struct {
	common *cfg.Common

	filePaths   []interface{}
	format      bool
	formatStyle string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		filePaths:   localConfig.UList("filePaths"),
		format:      localConfig.UBool("format", false),
		formatStyle: localConfig.UString("formatStyle", "vim"),
	}

	return &settings
}
