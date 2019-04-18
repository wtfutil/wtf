package pagerduty

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "pagerduty"

type Settings struct {
	common *cfg.Common

	apiKey           string
	escalationFilter []interface{}
	showIncidents    bool
	showSchedules    bool
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		apiKey:           localConfig.UString("apiKey", os.Getenv("WTF_PAGERDUTY_API_KEY")),
		escalationFilter: localConfig.UList("escalationFilter"),
		showIncidents:    localConfig.UBool("showIncidents", true),
		showSchedules:    localConfig.UBool("showSchedules", true),
	}

	return &settings
}
