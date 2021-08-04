package healthchecks

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = true
	defaultTitle     = "Healthchecks.io"
)

type Settings struct {
	*cfg.Common

	apiKey string   `help:"An healthchecks API key." optional:"false"`
	apiURL string   `help:"Base URL for API" optional:"true"`
	tags   []string `help:"Filters the checks and returns only the checks that are tagged with the specified value"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey: ymlConfig.UString("apiKey", os.Getenv("WTF_HEALTHCHECKS_APIKEY")),
		apiURL: ymlConfig.UString("apiURL", "https://hc-ping.com/"),
		tags:   utils.ToStrs(ymlConfig.UList("tags")),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).
		Service(settings.apiURL).Load()

	return &settings
}
