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

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		apiKey:    ymlConfig.UString("apiKey", os.Getenv("WTF_BAMBOO_HR_TOKEN")),
		subdomain: ymlConfig.UString("subdomain", os.Getenv("WTF_BAMBOO_HR_SUBDOMAIN")),
	}

	return &settings
}
