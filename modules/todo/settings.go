package todo

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Todo"
)

// Settings defines the configuration properties for this module
type Settings struct {
	*cfg.Common

	filePath          string
	checked           string
	unchecked         string
	newPos            string
	checkedPos        string
	parseDates        bool
	dateColor         string
	switchToInDaysIn  int
	undatedAsDays     int
	hideYearIfCurrent bool
	dateFormat        string
	parseTags         bool
	tagColor          string
	tagsAtEnd         bool
	hideTags          []interface{}
	hiddenNumInTitle  bool
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	common := cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig)

	settings := Settings{
		Common: common,

		filePath:          ymlConfig.UString("filename"),
		checked:           ymlConfig.UString("checkedIcon", common.Checkbox.Checked),
		unchecked:         ymlConfig.UString("uncheckedIcon", common.Checkbox.Unchecked),
		newPos:            ymlConfig.UString("newPos", "first"),
		checkedPos:        ymlConfig.UString("checkedPos", "last"),
		parseDates:        ymlConfig.UBool("dates.enabled", true),
		dateColor:         ymlConfig.UString("colors.date", "chartreuse"),
		switchToInDaysIn:  ymlConfig.UInt("dates.switchToInDaysIn", 7),
		undatedAsDays:     ymlConfig.UInt("dates.undatedAsDays", 7),
		hideYearIfCurrent: ymlConfig.UBool("dates.hideYearIfCurrent", true),
		dateFormat:        ymlConfig.UString("dates.format", "yyyy-mm-dd"),
		parseTags:         ymlConfig.UBool("tags.enabled", true),
		tagColor:          ymlConfig.UString("colors.tags", "khaki"),
		tagsAtEnd:         ymlConfig.UString("tags.pos", "end") == "end",
		hideTags:          ymlConfig.UList("tags.hide"),
		hiddenNumInTitle:  ymlConfig.UBool("tags.hiddenInTitle", true),
	}

	switch settings.newPos {
	case "first", "last":
	default:
		settings.newPos = "last"
	}
	switch settings.checkedPos {
	case "first", "last", "none":
	default:
		settings.checkedPos = "last"
	}
	switch settings.dateFormat {
	case "yyyy-mm-dd", "yy-mm-dd", "dd-mm-yyyy", "dd-mm-yy", "dd M yy", "dd M yyyy":
	default:
		settings.dateFormat = "yyyy-mm-dd"
	}

	return &settings
}
