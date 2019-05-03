package opsgenie

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "OpsGenie"

type Settings struct {
	common *cfg.Common

	apiKey                 string
	displayEmpty           bool
	schedule               []string
	scheduleIdentifierType string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		apiKey:                 ymlConfig.UString("apiKey", os.Getenv("WTF_OPS_GENIE_API_KEY")),
		displayEmpty:           ymlConfig.UBool("displayEmpty", true),
		scheduleIdentifierType: ymlConfig.UString("scheduleIdentifierType", "id"),
	}

	settings.schedule = settings.arrayifySchedules(ymlConfig, globalConfig)

	return &settings
}

// arrayifySchedules figures out if we're dealing with a single project or an array of projects
func (settings *Settings) arrayifySchedules(ymlConfig *config.Config, globalConfig *config.Config) []string {
	schedules := []string{}

	// Single schedule
	schedule, err := ymlConfig.String("schedule")
	if err == nil {
		schedules = append(schedules, schedule)
		return schedules
	}

	// Array of schedules
	scheduleList := ymlConfig.UList("schedule")
	for _, scheduleName := range scheduleList {
		if schedule, ok := scheduleName.(string); ok {
			schedules = append(schedules, schedule)
		}
	}

	return schedules
}
