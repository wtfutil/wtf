package toggl

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "Toggl"
)

type Settings struct {
	common *cfg.Common

	apiKey              string `help:"Your Toggl API token."`
	runningEntryColor   string `help:"Terminal color for running timer entries"`
	completedEntryColor string `help:"Terminal color for completed timer entries"`
	durationColor       string `help:"Terminal color for timer entry durations"`
	numberOfEntries     int    `help:"Number of Toggle entries to display"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:              ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_TOGGL_API_KEY"))),
		runningEntryColor:   ymlConfig.UString("runningEntryColor", ymlConfig.UString("runningEntryColor", "red")),
		completedEntryColor: ymlConfig.UString("completedEntryColor", ymlConfig.UString("completedEntryColor", "green")),
		durationColor:       ymlConfig.UString("durationColor", ymlConfig.UString("durationColor", "darkwhite")),
		numberOfEntries:     ymlConfig.UInt("numberOfEntries", ymlConfig.UInt("numberOfEntries", 10)),
	}

	return &settings
}
