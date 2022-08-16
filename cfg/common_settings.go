package cfg

import (
	"fmt"
	"strings"
	"time"

	"github.com/olebedev/config"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const (
	defaultLanguageTag = "en-CA"
)

type Module struct {
	Name string
	Type string
}

type Sigils struct {
	Checkbox struct {
		Checked   string
		Unchecked string
	}
	Paging struct {
		Normal   string
		Selected string
	}
}

// Common defines a set of common configuration settings applicable to all modules
type Common struct {
	Module
	PositionSettings `help:"Defines where in the grid this moduleÃ¢ÂÂs widget will be displayed."`
	Sigils

	Colors ColorTheme
	Config *config.Config

	DocPath string

	Bordered        bool          `help:"Whether or not the module should be displayed with a border." values:"true, false" optional:"true" default:"true"`
	Enabled         bool          `help:"Whether or not this module is executed and if its data displayed onscreen." values:"true, false" optional:"true" default:"false"`
	Focusable       bool          `help:"Whether or  not this module is focusable." values:"true, false" optional:"true" default:"false"`
	LanguageTag     string        `help:"The BCP 47 langauge tag to localize text to." values:"Any supported BCP 47 language tag." optional:"true" default:"en-CA"`
	RefreshInterval time.Duration `help:"How often this module will update its data." values:"A positive integer followed by a time unit (ns, us or ÃÂµs, ms, s, m, h, or nothing which defaults to s)" optional:"true"`
	Title           string        `help:"The title string to show when displaying this module" optional:"true"`

	focusChar int `help:"Define one of the number keys as a short cut key to access the widget." optional:"true"`
}

// NewCommonSettingsFromModule returns a common settings configuration tailed to the given module
func NewCommonSettingsFromModule(name, defaultTitle string, defaultFocusable bool, moduleConfig *config.Config, globalConfig *config.Config) *Common {
	baseColors := NewDefaultColorTheme()

	colorsConfig, err := globalConfig.Get("wtf.colors")
	if err != nil && strings.Contains(err.Error(), "Nonexistent map") {
		// Create a default colors config to fill in for the missing one
		// This comes into play when the configuration file does not contain a `colors:` key, i.e:
		//
		//     wtf:
		//       # colors:                <- missing
		//       refreshInterval: 1
		//       openFileUtil: "open"
		//
		colorsConfig, _ = NewDefaultColorConfig()
	}

	// And finally create a third instance to be the final default fallback in case there are empty or nil values in
	// the colors extracted from the config file (aka colorsConfig)
	defaultColorTheme := NewDefaultColorTheme()

	baseColors.BorderTheme.Focusable = moduleConfig.UString("colors.border.focusable", colorsConfig.UString("border.focusable", defaultColorTheme.BorderTheme.Focusable))
	baseColors.BorderTheme.Focused = moduleConfig.UString("colors.border.focused", colorsConfig.UString("border.focused", defaultColorTheme.BorderTheme.Focused))
	baseColors.BorderTheme.Unfocusable = moduleConfig.UString("colors.border.normal", colorsConfig.UString("border.normal", defaultColorTheme.BorderTheme.Unfocusable))

	baseColors.CheckboxTheme.Checked = moduleConfig.UString("colors.checked", colorsConfig.UString("checked", defaultColorTheme.CheckboxTheme.Checked))

	baseColors.RowTheme.EvenForeground = moduleConfig.UString("colors.rows.even", colorsConfig.UString("rows.even", defaultColorTheme.RowTheme.EvenForeground))
	baseColors.RowTheme.OddForeground = moduleConfig.UString("colors.rows.odd", colorsConfig.UString("rows.odd", defaultColorTheme.RowTheme.OddForeground))

	baseColors.TextTheme.Label = moduleConfig.UString("colors.label", colorsConfig.UString("label", defaultColorTheme.TextTheme.Label))
	baseColors.TextTheme.Subheading = moduleConfig.UString("colors.subheading", colorsConfig.UString("subheading", defaultColorTheme.TextTheme.Subheading))
	baseColors.TextTheme.Text = moduleConfig.UString("colors.text", colorsConfig.UString("text", defaultColorTheme.TextTheme.Text))
	baseColors.TextTheme.Title = moduleConfig.UString("colors.title", colorsConfig.UString("title", defaultColorTheme.TextTheme.Title))

	baseColors.WidgetTheme.Background = moduleConfig.UString("colors.background", colorsConfig.UString("background", defaultColorTheme.WidgetTheme.Background))

	common := Common{
		Colors: baseColors,

		Module: Module{
			Name: name,
			Type: moduleConfig.UString("type", name),
		},

		PositionSettings: NewPositionSettingsFromYAML(moduleConfig),

		Bordered:        moduleConfig.UBool("border", true),
		Config:          moduleConfig,
		Enabled:         moduleConfig.UBool("enabled", false),
		Focusable:       moduleConfig.UBool("focusable", defaultFocusable),
		LanguageTag:     globalConfig.UString("wtf.language", defaultLanguageTag),
		RefreshInterval: ParseTimeString(moduleConfig, "refreshInterval", "300s"),
		Title:           moduleConfig.UString("title", defaultTitle),

		focusChar: moduleConfig.UInt("focusChar", -1),
	}

	sigilsPath := "wtf.sigils"
	common.Sigils.Checkbox.Checked = globalConfig.UString(sigilsPath+".checkbox.checked", "x")
	common.Sigils.Checkbox.Unchecked = globalConfig.UString(sigilsPath+".checkbox.unchecked", " ")
	common.Sigils.Paging.Normal = globalConfig.UString(sigilsPath+".paging.normal", globalConfig.UString("wtf.paging.pageSigil", "*"))
	common.Sigils.Paging.Selected = globalConfig.UString(sigilsPath+".paging.select", globalConfig.UString("wtf.paging.selectedSigil", "_"))

	return &common
}

/* -------------------- Exported Functions -------------------- */

func (common *Common) DefaultFocusedRowColor() string {
	return fmt.Sprintf(
		"%s:%s",
		common.Colors.RowTheme.HighlightedForeground,
		common.Colors.RowTheme.HighlightedBackground,
	)
}

func (common *Common) DefaultRowColor() string {
	return fmt.Sprintf(
		"%s:%s",
		common.Colors.RowTheme.EvenForeground,
		common.Colors.RowTheme.EvenBackground,
	)
}

// FocusChar returns the keyboard number assigned to the widget used to give onscreen
// focus to this widget, as a string. Focus characters can be a range between 1 and 9
func (common *Common) FocusChar() string {
	if common.focusChar <= 0 {
		return ""
	}

	if common.focusChar > 9 {
		return ""
	}

	return fmt.Sprint(common.focusChar)
}

// LocalizedPrinter returns a message.Printer instance localized to the BCP 47 language
// configuration value defined in 'wtf.language' config. If none exists, it defaults to
// 'en-CA'. Use this to format numbers, etc.
func (common *Common) LocalizedPrinter() (*message.Printer, error) {
	langTag, err := language.Parse(common.LanguageTag)
	if err != nil {
		return nil, err
	}

	prntr := message.NewPrinter(langTag)

	return prntr, nil
}

func (common *Common) RowColor(idx int) string {
	if idx%2 == 0 {
		return fmt.Sprintf(
			"%s:%s",
			common.Colors.RowTheme.EvenForeground,
			common.Colors.RowTheme.EvenBackground,
		)
	}
	return fmt.Sprintf(
		"%s:%s",
		common.Colors.RowTheme.OddForeground,
		common.Colors.RowTheme.OddBackground,
	)
}

func (common *Common) RightAlignFormat(width int) string {
	borderOffset := 2
	return fmt.Sprintf("%%%ds", width-borderOffset)
}

// PaginationMarker generates the pagination indicators that appear in the top-right corner
// of multisource widgets
func (common *Common) PaginationMarker(length, pos, width int) string {
	sigils := ""

	if length > 1 {
		sigils = strings.Repeat(common.Sigils.Paging.Normal, pos)
		sigils += common.Sigils.Paging.Selected
		sigils += strings.Repeat(common.Sigils.Paging.Normal, length-1-pos)

		sigils = "[lightblue]" + fmt.Sprintf(common.RightAlignFormat(width), sigils) + "[white]"
	}

	return sigils
}

// SetDocumentationPath is used to explicitly set the documentation path that should be opened
// when the key to open the documentation is pressed.
// Setting this is probably not necessary unless the module documentation is nested inside a
// documentation subdirectory in the /wtfutildocs repo, or the module here has a different
// name than the module's display name in the documentation (which ideally wouldn't be a thing).
func (common *Common) SetDocumentationPath(path string) {
	common.DocPath = path
}

// Validations aggregates all the validations from all the sub-sections in Common into a
// single array of validations
func (common *Common) Validations() []Validatable {
	validatables := []Validatable{}

	for _, validation := range common.PositionSettings.Validations.validations {
		validatables = append(validatables, validation)
	}

	return validatables
}
