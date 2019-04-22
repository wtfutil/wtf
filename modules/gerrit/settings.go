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

const configKey = "gerrit"

type Settings struct {
	colors
	common *cfg.Common

	domain                  string
	password                string
	projects                []interface{}
	username                string
	verifyServerCertificate bool
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		domain:                  localConfig.UString("domain", ""),
		password:                localConfig.UString("password", os.Getenv("WTF_GERRIT_PASSWORD")),
		projects:                localConfig.UList("projects"),
		username:                localConfig.UString("username", ""),
		verifyServerCertificate: localConfig.UBool("verifyServerCertificate", true),
	}

	settings.colors.rows.even = localConfig.UString("colors.rows.even", "white")
	settings.colors.rows.odd = localConfig.UString("colors.rows.odd", "blue")

	return &settings
}
