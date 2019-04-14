package datadog

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	common *cfg.Common

	apiKey         string
	applicationKey string
	tags           []interface{}
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.todo")

	settings := Settings{
		common:         cfg.NewCommonSettingsFromYAML(ymlConfig),
		apiKey:         localConfig.UString("apiKey", os.Getenv("WTF_DATADOG_API_KEY")),
		applicationKey: localConfig.UString("applicationKey", os.Getenv("WTF_DATADOG_APPLICATION_KEY")),
		tags:           localConfig.UList("monitors.tags"),
	}

	return &settings
}
