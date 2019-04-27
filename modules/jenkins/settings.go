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

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		apiKey:                  ymlConfig.UString("apiKey", os.Getenv("WTF_JENKINS_API_KEY")),
		jobNameRegex:            ymlConfig.UString("jobNameRegex", ".*"),
		successBallColor:        ymlConfig.UString("successBallColor", "blue"),
		url:                     ymlConfig.UString("url"),
		user:                    ymlConfig.UString("user"),
		verifyServerCertificate: ymlConfig.UBool("verifyServerCertificate", true),
	}

	return &settings
}
