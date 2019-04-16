package textfile

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	common *cfg.Common

	filePaths   []interface{}
	format      bool
	formatStyle string
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.textfile")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		filePaths:   localConfig.UList("filePaths"),
		format:      localConfig.UBool("format", false),
		formatStyle: localConfig.UString("formatStyle", "vim"),
	}

	return &settings
}
