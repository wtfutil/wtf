package opsgenie

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "opsgenie"

type Settings struct {
	common *cfg.Common

	apiKey                 string
	displayEmpty           bool
	schedule               []string
	scheduleIdentifierType string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		apiKey:                 localConfig.UString("apiKey", os.Getenv("WTF_OPS_GENIE_API_KEY")),
		displayEmpty:           localConfig.UBool("displayEmpty", true),
		scheduleIdentifierType: localConfig.UString("scheduleIdentifierType", "id"),
	}

	settings.schedule = settings.arrayifySchedules(localConfig)

	return &settings
}

// arrayifySchedules figures out if we're dealing with a single project or an array of projects
func (settings *Settings) arrayifySchedules(localConfig *config.Config) []string {
	schedules := []string{}

	// Single schedule
	schedule, err := localConfig.String("schedule")
	if err == nil {
		schedules = append(schedules, schedule)
		return schedules
	}

	// Array of schedules
	scheduleList := localConfig.UList("schedule")
	for _, scheduleName := range scheduleList {
		if schedule, ok := scheduleName.(string); ok {
			schedules = append(schedules, schedule)
		}
	}

	return schedules
}
