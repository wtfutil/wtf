package jenkins

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Jenkins"
)

type Settings struct {
	common *cfg.Common

	apiKey                  string `help:"Your Jenkins API key."`
	jobNameRegex            string `help:"A regex that filters the jobs shown in the widget." optional:"true"`
	successBallColor        string `help:"Changes the default color of successful Jenkins jobs to the color of your choosing." values:"blue, green, purple, yellow, etc." optional:"true"`
	url                     string `help:"The url to your Jenkins project or view."`
	user                    string `help:"Your Jenkins username."`
	verifyServerCertificate bool   `help:"Determines whether or not the server’s certificate chain and host name are verified." values:"true or false" optional:"true"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:                  ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_JENKINS_API_KEY"))),
		jobNameRegex:            ymlConfig.UString("jobNameRegex", ".*"),
		successBallColor:        ymlConfig.UString("successBallColor", "blue"),
		url:                     ymlConfig.UString("url"),
		user:                    ymlConfig.UString("user"),
		verifyServerCertificate: ymlConfig.UBool("verifyServerCertificate", true),
	}

	return &settings
}
