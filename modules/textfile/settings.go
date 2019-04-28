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

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		filePaths:   ymlConfig.UList("filePaths"),
		format:      ymlConfig.UBool("format", false),
		formatStyle: ymlConfig.UString("formatStyle", "vim"),
	}

	return &settings
}
