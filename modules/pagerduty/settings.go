package pagerduty

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	common *cfg.Common

	apiKey           string
	escalationFilter []interface{}
	showIncidents    bool
	showSchedules    bool
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.pagerduty")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		apiKey:           localConfig.UString("apiKey", os.Getenv("WTF_PAGERDUTY_API_KEY")),
		escalationFilter: localConfig.UList("escalationFilter"),
		showIncidents:    localConfig.UBool("showIncidents", true),
		showSchedules:    localConfig.UBool("showSchedules", true),
	}

	return &settings
}
