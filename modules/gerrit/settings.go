package gerrit

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Gerrit"
)

type colors struct {
	rows struct {
		even string `help:"Define the foreground color for even-numbered rows." values:"Any X11 color name." optional:"true"`
		odd  string `help:"Define the foreground color for odd-numbered rows." values:"Any X11 color name." optional:"true"`
	}
}

type Settings struct {
	colors
	common *cfg.Common

	domain                  string        `help:"Your Gerrit corporate domain."`
	password                string        `help:"Your Gerrit HTTP Password."`
	projects                []interface{} `help:"A list of Gerrit project names to fetch data for."`
	username                string        `help:"Your Gerrit username."`
	verifyServerCertificate bool          `help:"Determines whether or not the server’s certificate chain and host name are verified." values:"true or false" optional:"true"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		domain:                  ymlConfig.UString("domain", ""),
		password:                ymlConfig.UString("password", os.Getenv("WTF_GERRIT_PASSWORD")),
		projects:                ymlConfig.UList("projects"),
		username:                ymlConfig.UString("username", ""),
		verifyServerCertificate: ymlConfig.UBool("verifyServerCertificate", true),
	}

	cfg.ConfigureSecret(
		globalConfig,
		settings.domain,
		name,
		nil, // Seems like it should be mandatory, but its optional above.
		&settings.password,
	)

	settings.colors.rows.even = ymlConfig.UString("colors.rows.even", "white")
	settings.colors.rows.odd = ymlConfig.UString("colors.rows.odd", "blue")

	return &settings
}
