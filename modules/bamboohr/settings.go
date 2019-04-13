package bamboohr

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	Common *cfg.Common

	APIKey    string
	Subdomain string
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.bamboohr")

	settings := Settings{
		Common:    cfg.NewCommonSettingsFromYAML(ymlConfig),
		APIKey:    localConfig.UString("apiKey", os.Getenv("WTF_BAMBOO_HR_TOKEN")),
		Subdomain: localConfig.UString("subdomain", os.Getenv("WTF_BAMBOO_HR_SUBDOMAIN")),
	}

	return &settings
}
