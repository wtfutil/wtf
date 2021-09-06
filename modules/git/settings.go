package git

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = true
	defaultTitle     = "Git"
)

type Settings struct {
	*cfg.Common

	commitCount      int           `help:"The number of past commits to display." values:"A positive integer, 0..n." optional:"true"`
	sections         []interface{} `help:"Sections to show" optional:"true"`
	showModuleName   bool          `help:"Whether to show 'Git - ' before information in title" optional:"true" default:"true"`
	branchInTitle    bool          `help:"Whether to show branch name in title instead of the widget body itself" optional:"true" default:"false"`
	showFilesIfEmpty bool          `help:"Whether to show Changed Files section if no changed files" optional:"true" default:"true"`
	lastFolderTitle  bool          `help:"Whether to show only last part of directory path instead of full path" optional:"true" default:"false"`
	commitFormat     string        `help:"The string format for the commit message." optional:"true"`
	dateFormat       string        `help:"The string format for the date/time in the commit message." optional:"true"`
	repositories     []interface{} `help:"Defines which git repositories to watch." values:"A list of zero or more local file paths pointing to valid git repositories."`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		commitCount:      ymlConfig.UInt("commitCount", 10),
		sections:         ymlConfig.UList("sections"),
		showModuleName:   ymlConfig.UBool("showModuleName", true),
		branchInTitle:    ymlConfig.UBool("branchInTitle", false),
		showFilesIfEmpty: ymlConfig.UBool("showFilesIfEmpty", true),
		lastFolderTitle:  ymlConfig.UBool("lastFolderTitle", false),
		commitFormat:     ymlConfig.UString("commitFormat", "[forestgreen]%h [white]%s [grey]%an on %cd[white]"),
		dateFormat:       ymlConfig.UString("dateFormat", "%b %d, %Y"),
		repositories:     ymlConfig.UList("repositories"),
	}
	if len(settings.sections) == 0 {
		for _, v := range []string{"branch", "files", "commits"} {
			settings.sections = append(settings.sections, v)
		}
	}

	return &settings
}

func (widget *Widget) ConfigText() string {
	return utils.HelpFromInterface(Settings{})
}
