package cfg

import (
	"fmt"
	"strings"

	"github.com/olebedev/config"
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

type Common struct {
	Module
	PositionSettings `help:"Defines where in the grid this module’s widget will be displayed."`
	Sigils

	Colors          ColorTheme
	Bordered        bool   `help:"Whether or not the module should be displayed with a border." values:"true, false" optional:"true" default:"true"`
	Enabled         bool   `help:"Whether or not this module is executed and if its data displayed onscreen." values:"true, false" optional:"true" default:"false"`
	Focusable       bool   `help:"Whether or  not this module is focusable." values:"true, false" optional:"true" default:"false"`
	RefreshInterval int    `help:"How often, in seconds, this module will update its data." values:"A positive integer, 0..n." optional:"true"`
	Title           string `help:"The title string to show when displaying this module" optional:"true"`
	Config          *config.Config

	focusChar int `help:"Define one of the number keys as a short cut key to access the widget." optional:"true"`
}

// NewCommonSettingsFromModule returns a common settings configuration tailed to the given module
func NewCommonSettingsFromModule(name, defaultTitle string, defaultFocusable bool, moduleConfig *config.Config, globalSettings *config.Config) *Common {
	baseColors := NewDefaultColorTheme()

	colorsConfig, err := globalSettings.Get("wtf.colors")
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

		PositionSettings: NewPositionSettingsFromYAML(name, moduleConfig),

		Bordered:        moduleConfig.UBool("border", true),
		Config:          moduleConfig,
		Enabled:         moduleConfig.UBool("enabled", false),
		Focusable:       moduleConfig.UBool("focusable", defaultFocusable),
		RefreshInterval: moduleConfig.UInt("refreshInterval", 300),
		Title:           moduleConfig.UString("title", defaultTitle),

		focusChar: moduleConfig.UInt("focusChar", -1),
	}

	sigilsPath := "wtf.sigils"

	common.Sigils.Checkbox.Checked = globalSettings.UString(sigilsPath+".checkbox.checked", "x")
	common.Sigils.Checkbox.Unchecked = globalSettings.UString(sigilsPath+".checkbox.unchecked", " ")
	common.Sigils.Paging.Normal = globalSettings.UString(sigilsPath+".paging.normal", globalSettings.UString("wtf.paging.pageSigil", "*"))
	common.Sigils.Paging.Selected = globalSettings.UString(sigilsPath+".paging.select", globalSettings.UString("wtf.paging.selectedSigil", "_"))

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

func (common *Common) FocusChar() string {
	if common.focusChar <= -1 {
		return ""
	}

	return string('0' + common.focusChar)
}

func (common *Common) RowColor(idx int) string {
	if idx%2 == 0 {
		return common.Colors.RowTheme.EvenForeground
	}

	return common.Colors.RowTheme.OddForeground
}

func (common *Common) RightAlignFormat(width int) string {
	borderOffset := 2
	return fmt.Sprintf("%%%ds", width-borderOffset)
}

func (common *Common) SigilStr(len, pos, width int) string {
	sigils := ""

	if len > 1 {
		sigils = strings.Repeat(common.Sigils.Paging.Normal, pos)
		sigils += common.Sigils.Paging.Selected
		sigils += strings.Repeat(common.Sigils.Paging.Normal, len-1-pos)

		sigils = "[lightblue]" + fmt.Sprintf(common.RightAlignFormat(width), sigils) + "[white]"
	}

	return sigils
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
