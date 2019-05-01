package gerrit

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type colors struct {
	rows struct {
		even string
		odd  string
	}
}

const defaultTitle = "Gerrit"

type Settings struct {
	colors
	common *cfg.Common

	domain                  string
	password                string
	projects                []interface{}
	username                string
	verifyServerCertificate bool
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		domain:                  ymlConfig.UString("domain", ""),
		password:                ymlConfig.UString("password", os.Getenv("WTF_GERRIT_PASSWORD")),
		projects:                ymlConfig.UList("projects"),
		username:                ymlConfig.UString("username", ""),
		verifyServerCertificate: ymlConfig.UBool("verifyServerCertificate", true),
	}

	settings.colors.rows.even = ymlConfig.UString("colors.rows.even", "white")
	settings.colors.rows.odd = ymlConfig.UString("colors.rows.odd", "blue")

	return &settings
}
