package pagerduty

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "PagerDuty"

type Settings struct {
	common *cfg.Common

	apiKey           string
	escalationFilter []interface{}
	scheduleIDs      []interface{}
	showIncidents    bool
	showSchedules    bool
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		apiKey:           ymlConfig.UString("apiKey", os.Getenv("WTF_PAGERDUTY_API_KEY")),
		escalationFilter: ymlConfig.UList("escalationFilter"),
		scheduleIDs:      ymlConfig.UList("scheduleIDs", []interface{}{}),
		showIncidents:    ymlConfig.UBool("showIncidents", true),
		showSchedules:    ymlConfig.UBool("showSchedules", true),
	}

	return &settings
}
