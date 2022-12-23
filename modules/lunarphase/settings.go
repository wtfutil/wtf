package lunarphase

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "Phase of the Moon"
)

type Settings struct {
	*cfg.Common

	language string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		language: ymlConfig.UString("language", "en"),
	}

	settings.SetDocumentationPath("lunarphase")

	return &settings
}
