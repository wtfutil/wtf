package bamboohr

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "BambooHR"

type Settings struct {
	common *cfg.Common

	apiKey    string `help:"Your BambooHR API token."`
	subdomain string `help:"Your BambooHR API subdomain name."`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		apiKey:    ymlConfig.UString("apiKey", os.Getenv("WTF_BAMBOO_HR_TOKEN")),
		subdomain: ymlConfig.UString("subdomain", os.Getenv("WTF_BAMBOO_HR_SUBDOMAIN")),
	}

	return &settings
}
