package ipinfo

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "IPInfo"
)

type Settings struct {
	*cfg.Common

	apiToken string `help:"An api token" optional:"true"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiToken: ymlConfig.UString("apiToken", ""),
	}

	settings.SetDocumentationPath("ipaddress/ipinfo")

	return &settings
}
