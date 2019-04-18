package bamboohr

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "bamboohr"

type Settings struct {
	common *cfg.Common

	apiKey    string
	subdomain string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		apiKey:    localConfig.UString("apiKey", os.Getenv("WTF_BAMBOO_HR_TOKEN")),
		subdomain: localConfig.UString("subdomain", os.Getenv("WTF_BAMBOO_HR_SUBDOMAIN")),
	}

	return &settings
}
