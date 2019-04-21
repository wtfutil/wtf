package jenkins

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "jenkins"

type Settings struct {
	common *cfg.Common

	apiKey                  string
	jobNameRegex            string
	successBallColor        string
	url                     string
	user                    string
	verifyServerCertificate bool
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		apiKey:                  localConfig.UString("apiKey", os.Getenv("WTF_JENKINS_API_KEY")),
		jobNameRegex:            localConfig.UString("jobNameRegex", ".*"),
		successBallColor:        localConfig.UString("successBallColor", "blue"),
		url:                     localConfig.UString("url"),
		user:                    localConfig.UString("user"),
		verifyServerCertificate: localConfig.UBool("verifyServerCertificate", true),
	}

	return &settings
}
