package datadog

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "DataDog"

type Settings struct {
	common *cfg.Common

	apiKey         string
	applicationKey string
	tags           []interface{}
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		apiKey:         ymlConfig.UString("apiKey", os.Getenv("WTF_DATADOG_API_KEY")),
		applicationKey: ymlConfig.UString("applicationKey", os.Getenv("WTF_DATADOG_APPLICATION_KEY")),
		tags:           ymlConfig.UList("monitors.tags"),
	}

	return &settings
}
