package zendesk

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	common *cfg.Common

	apiKey    string
	status    string
	subdomain string
	username  string
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.zendesk")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		apiKey:    localConfig.UString("apiKey", os.Getenv("ZENDESK_API")),
		status:    localConfig.UString("status"),
		subdomain: localConfig.UString("subdomain", os.Getenv("ZENDESK_SUBDOMAIN")),
		username:  localConfig.UString("username"),
	}

	return &settings
}
