package bamboohr

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "BambooHR"
)

type Settings struct {
	common *cfg.Common

	apiKey    string `help:"Your BambooHR API token."`
	subdomain string `help:"Your BambooHR API subdomain name."`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:    ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_BAMBOO_HR_TOKEN"))),
		subdomain: ymlConfig.UString("subdomain", os.Getenv("WTF_BAMBOO_HR_SUBDOMAIN")),
	}

	cfg.ConfigureSecret(
		globalConfig,
		"",
		name,
		&settings.subdomain,
		&settings.apiKey,
	)

	return &settings
}
