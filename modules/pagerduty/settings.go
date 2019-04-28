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

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		apiKey:           ymlConfig.UString("apiKey", os.Getenv("WTF_PAGERDUTY_API_KEY")),
		escalationFilter: ymlConfig.UList("escalationFilter"),
		showIncidents:    ymlConfig.UBool("showIncidents", true),
		showSchedules:    ymlConfig.UBool("showSchedules", true),
	}

	return &settings
}
