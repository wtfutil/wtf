package flighty

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Flights"
)

type Settings struct {
	*cfg.Common

	username string
	password string
	aircraft map[string]string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := &Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		username: ymlConfig.UString("authentication.username"),
		password: ymlConfig.UString("authentication.password"),

		aircraft: buildAircraft(ymlConfig),
	}

	settings.SetDocumentationPath("flighty")

	return settings
}

/* -------------------- Unexported Functions -------------------- */

func buildAircraft(ymlConfig *config.Config) map[string]string {
	flights := make(map[string]string)

	aircrafts, err := ymlConfig.Map("aircraft")
	if err == nil {
		for name, icao24 := range aircrafts {
			flights[name] = icao24.(string)
		}
	}

	return flights
}
