package krisinformation

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "Krisinformation"
	defaultRadius    = -1
	defaultCountry   = true
	defaultCounty    = ""
	defaultMaxItems  = -1
	defaultMaxAge    = 720
)

// Settings defines the configuration properties for this module
type Settings struct {
	common    *cfg.Common
	latitude  float64 `help:"The latitude of the position from which the widget should look for messages." optional:"true"`
	longitude float64 `help:"The longitude of the position from which the widget should look for messages." optional:"true"`
	radius    int     `help:"The radius in km from your position that the widget should look for messages. need latitude/longitude setting,Default 10" optional:"true"`
	county    string  `help:"The county from where to display messages" optional:"true"`
	country   bool    `help:"Only display country wide messages" optional:"true"`
	maxitems  int     `help:"Only display X number of latest messages" optional:"true"`
	maxage    int     `help:"Only show messages younger than maxage" optional:"true"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common:    cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		latitude:  ymlConfig.UFloat64("latitude", -1),
		longitude: ymlConfig.UFloat64("longitude", -1),
		radius:    ymlConfig.UInt("radius", defaultRadius),
		country:   ymlConfig.UBool("country", defaultCountry),
		county:    ymlConfig.UString("county", defaultCounty),
		maxitems:  ymlConfig.UInt("items", defaultMaxItems),
		maxage:    ymlConfig.UInt("maxages", defaultMaxAge),
	}

	return &settings
}
