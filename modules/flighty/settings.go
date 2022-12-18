package flighty

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = true
	defaultTitle     = "Flights"
)

type Settings struct {
	*cfg.Common

	username string
	password string
	aircraft []string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := &Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		username: ymlConfig.UString("authentication.username"),
		password: ymlConfig.UString("authentication.password"),
		aircraft: utils.ToStrs(ymlConfig.UList("aircraft")),
	}

	settings.SetDocumentationPath("flighty")

	return settings
}
