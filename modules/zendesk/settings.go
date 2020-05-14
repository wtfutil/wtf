package zendesk

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Zendesk"
)

type Settings struct {
	common *cfg.Common

	apiKey    string
	status    string
	subdomain string
	username  string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:    ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("ZENDESK_API"))),
		status:    ymlConfig.UString("status"),
		subdomain: ymlConfig.UString("subdomain", os.Getenv("ZENDESK_SUBDOMAIN")),
		username:  ymlConfig.UString("username"),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).Load()

	return &settings
}
