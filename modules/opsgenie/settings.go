package opsgenie

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "OpsGenie"
)

type Settings struct {
	*cfg.Common

	apiKey                 string   `help:"Your OpsGenie API token."`
	region                 string   `help:"Defines region to use. Possible options: us (by default), eu." optional:"true"`
	displayEmpty           bool     `help:"Whether schedules with no assigned person on-call should be displayed." optional:"true"`
	schedule               []string `help:"A list of names of the schedule(s) to retrieve."`
	scheduleIdentifierType string   `help:"Type of the schedule identifier." values:"id or name" optional:"true"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:                 ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_OPS_GENIE_API_KEY"))),
		region:                 ymlConfig.UString("region", "us"),
		displayEmpty:           ymlConfig.UBool("displayEmpty", true),
		scheduleIdentifierType: ymlConfig.UString("scheduleIdentifierType", "id"),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).Load()

	settings.schedule = settings.arrayifySchedules(ymlConfig)

	return &settings
}

// arrayifySchedules figures out if we're dealing with a single project or an array of projects
func (settings *Settings) arrayifySchedules(ymlConfig *config.Config) []string {
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
