package weather

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Weather"
)

type colors struct {
	current string
}

type Settings struct {
	colors
	common *cfg.Common

	apiKey   string
	cityIDs  []interface{}
	language string
	tempUnit string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:   ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_OWM_API_KEY"))),
		cityIDs:  ymlConfig.UList("cityids"),
		language: ymlConfig.UString("language", "EN"),
		tempUnit: ymlConfig.UString("tempUnit", "C"),
	}

	settings.colors.current = ymlConfig.UString("colors.current", "green")

	return &settings
}
