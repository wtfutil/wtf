package gitter

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "gitter"

type Settings struct {
	common *cfg.Common

	apiToken         string
	numberOfMessages int
	roomURI          string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		apiToken:         localConfig.UString("apiToken", os.Getenv("WTF_GITTER_API_TOKEN")),
		numberOfMessages: localConfig.UInt("numberOfMessages", 10),
		roomURI:          localConfig.UString("roomUri", "wtfutil/Lobby"),
	}

	return &settings
}
