package pagerduty

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "PagerDuty"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	apiKey           string        `help:"Your PagerDuty API key."`
	escalationFilter []interface{} `help:"An array of schedule names you want to filter on."`
	myName           string        `help:"The name to highlight when on-call in PagerDuty."`
	scheduleIDs      []interface{} `help:"An array of schedule IDs you want to restrict the query to."`
	showIncidents    bool          `help:"Whether or not to list incidents." optional:"true"`
	showSchedules    bool          `help:"Whether or not to show schedules." optional:"true"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:           ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_PAGERDUTY_API_KEY"))),
		escalationFilter: ymlConfig.UList("escalationFilter"),
		myName:           ymlConfig.UString("myName"),
		scheduleIDs:      ymlConfig.UList("scheduleIDs", []interface{}{}),
		showIncidents:    ymlConfig.UBool("showIncidents", true),
		showSchedules:    ymlConfig.UBool("showSchedules", true),
	}

	return &settings
}
