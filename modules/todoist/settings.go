package todoist

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = true
	defaultTitle     = "Todoist"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	apiKey   string `help:"Your Todoist API key"`
	projects []uint
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:   ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_TODOIST_TOKEN"))),
		projects: utils.IntsToUints(utils.ToInts(ymlConfig.UList("projects"))),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).Load()

	return &settings
}
