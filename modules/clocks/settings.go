package clocks

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = false
	defaultTitle     = "Clocks"
)

// Settings defines the configuration properties for this module
type Settings struct {
	*cfg.Common

	dateFormat string  `help:"The format of the date string for all clocks." values:"Any valid Go date layout which is handled by Time.Format. Defaults to Jan 2."`
	timeFormat string  `help:"The format of the time string for all clocks." values:"Any valid Go time layout which is handled by Time.Format. Defaults to 15:04 MST."`
	locations  []Clock `help:"Defines the timezones for the world clocks that you want to display. key is a unique label that will be displayed in the UI. value is a timezone name." values:"Any TZ database timezone."`
	sort       string  `help:"Defines the display order of the clocks in the widget." values:"'alphabetical', 'chronological', or 'natural. 'alphabetical' will sort in ascending order by key, 'chronological' will sort in ascending order by date/time, 'natural' will keep ordering as per the config."`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		dateFormat: ymlConfig.UString("dateFormat", utils.SimpleDateFormat),
		timeFormat: ymlConfig.UString("timeFormat", utils.SimpleTimeFormat),
		locations:  buildLocations(ymlConfig),
		sort:       ymlConfig.UString("sort"),
	}

	return &settings
}

func buildLocations(ymlConfig *config.Config) []Clock {
	clocks := []Clock{}
	locations, err := ymlConfig.Map("locations")
	if err == nil {
		for k, v := range locations {
			name := k
			zone := v.(string)
			clock, err := BuildClock(name, zone)
			if err == nil {
				clocks = append(clocks, clock)
			}
		}
		return clocks
	}

	listLocations := ymlConfig.UList("locations")
	for _, location := range listLocations {
		if location, ok := location.(map[string]interface{}); ok {
			for k, v := range location {
				name := k
				zone := v.(string)
				clock, err := BuildClock(name, zone)
				if err == nil {
					clocks = append(clocks, clock)
				}
			}
		}
	}
	return clocks
}
