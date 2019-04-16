package spotifyweb

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	common *cfg.Common

	callbackPort string
	clientID     string
	secretKey    string
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.spotifyweb")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		callbackPort: localConfig.UString("callbackPort", "8080"),
		clientID:     localConfig.UString("clientID", os.Getenv("SPOTIFY_ID")),
		secretKey:    localConfig.UString("secretKey", os.Getenv("SPOTIFY_SECRET")),
	}

	return &settings
}
