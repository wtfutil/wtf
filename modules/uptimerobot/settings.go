package uptimerobot

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Uptime Robot"
)

type Settings struct {
	common *cfg.Common

	apiKey        string `help:"An UptimeRobot API key."`
	uptimePeriods string `help:"The periods over which to display uptime (in days, dash-separated)." optional:"true"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:        ymlConfig.UString("apiKey", os.Getenv("WTF_UPTIMEROBOT_APIKEY")),
		uptimePeriods: ymlConfig.UString("uptimePeriods", "30"),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).
		Service("https://api.uptimerobot.com").Load()

	return &settings
}
