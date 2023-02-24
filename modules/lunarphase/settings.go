package lunarphase

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Phase of the Moon"
	dateFormat       = "2006-01-02"
	phaseFormat      = "01-02-2006"
)

type Settings struct {
	*cfg.Common

	language       string
	requestTimeout int
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		language:       ymlConfig.UString("language", "en"),
		requestTimeout: ymlConfig.UInt("timeout", 30),
	}

	settings.SetDocumentationPath("lunarphase")

	return &settings
}
