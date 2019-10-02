package todo

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Todo"
)

type Settings struct {
	common *cfg.Common

	filePath  string
	checked   string
	unchecked string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	common := cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig)

	settings := Settings{
		common: common,

		filePath:  ymlConfig.UString("filename"),
		checked:   ymlConfig.UString("checkedIcon", common.Checkbox.Checked),
		unchecked: ymlConfig.UString("uncheckedIcon", common.Checkbox.Unchecked),
	}

	return &settings
}
