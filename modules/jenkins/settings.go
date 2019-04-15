package jenkins

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	common *cfg.Common

	apiKey                  string
	successBallColor        string
	url                     string
	user                    string
	verifyServerCertificate bool
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.jenkins")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		apiKey:                  localConfig.UString("apiKey", os.Getenv("WTF_JENKINS_API_KEY")),
		successBallColor:        localConfig.UString("successBallColor", "blue"),
		url:                     localConfig.UString("url"),
		user:                    localConfig.UString("user"),
		verifyServerCertificate: localConfig.UBool("verifyServerCertificate", true),
	}

	return &settings
}
