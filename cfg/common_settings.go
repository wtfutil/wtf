package cfg

import (
	"fmt"
	"strings"

	"github.com/olebedev/config"
)

type Colors struct {
	Background      string
	BorderFocusable string
	BorderFocused   string
	BorderNormal    string
	Checked         string
	Foreground      string
	HighlightFore   string
	HighlightBack   string
	Text            string
	Title           string

	Rows struct {
		Even string
		Odd  string
	}
}

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
	Colors
	Module
	PositionSettings `help:"Defines where in the grid this module’s widget will be displayed."`
	Sigils

	Bordered        bool   `help:"Whether or not the module should be displayed with a border." values:"true, false" optional:"true" default:"true"`
	Enabled         bool   `help:"Whether or not this module is executed and if its data displayed onscreen." values:"true, false" optional:"true" default:"false"`
	Focusable       bool   `help:"Whether or  not this module is focusable." values:"true, false" optional:"true" default:"false"`
	RefreshInterval int    `help:"How often, in seconds, this module will update its data." values:"A positive integer, 0..n." optional:"true"`
	Title           string `help:"The title string to show when displaying this module" optional:"true"`
	Config          *config.Config

	focusChar int `help:"Define one of the number keys as a short cut key to access the widget." optional:"true"`
}

func NewCommonSettingsFromModule(name, defaultTitle string, defaultFocusable bool, moduleConfig *config.Config, globalSettings *config.Config) *Common {
	colorsConfig, _ := globalSettings.Get("wtf.colors")
	sigilsPath := "wtf.sigils"

	common := Common{
		Colors: Colors{
			Background:      moduleConfig.UString("colors.background", colorsConfig.UString("background", "transparent")),
			BorderFocusable: moduleConfig.UString("colors.border.focusable", colorsConfig.UString("border.focusable", "red")),
			BorderFocused:   moduleConfig.UString("colors.border.focused", colorsConfig.UString("border.focused", "orange")),
			BorderNormal:    moduleConfig.UString("colors.border.normal", colorsConfig.UString("border.normal", "gray")),
			Checked:         moduleConfig.UString("colors.checked", colorsConfig.UString("checked", "white")),
			Foreground:      moduleConfig.UString("colors.foreground", colorsConfig.UString("foreground", "white")),
			HighlightFore:   moduleConfig.UString("colors.highlight.fore", colorsConfig.UString("highlight.fore", "black")),
			HighlightBack:   moduleConfig.UString("colors.highlight.back", colorsConfig.UString("highlight.back", "green")),
			Text:            moduleConfig.UString("colors.text", colorsConfig.UString("text", "white")),
			Title:           moduleConfig.UString("colors.title", colorsConfig.UString("title", "white")),
		},

		Module: Module{
			Name: name,
			Type: moduleConfig.UString("type", name),
		},

		PositionSettings: NewPositionSettingsFromYAML(name, moduleConfig),

		Bordered:        moduleConfig.UBool("border", true),
		Enabled:         moduleConfig.UBool("enabled", false),
		Focusable:       moduleConfig.UBool("focusable", defaultFocusable),
		RefreshInterval: moduleConfig.UInt("refreshInterval", 300),
		Title:           moduleConfig.UString("title", defaultTitle),
		Config:          moduleConfig,

		focusChar: moduleConfig.UInt("focusChar", -1),
	}

	common.Colors.Rows.Even = moduleConfig.UString("colors.rows.even", colorsConfig.UString("rows.even", "white"))
	common.Colors.Rows.Odd = moduleConfig.UString("colors.rows.odd", colorsConfig.UString("rows.odd", "lightblue"))

	common.Sigils.Checkbox.Checked = globalSettings.UString(sigilsPath+".checkbox.checked", "x")
	common.Sigils.Checkbox.Unchecked = globalSettings.UString(sigilsPath+".checkbox.unchecked", " ")

	common.Sigils.Paging.Normal = globalSettings.UString(sigilsPath+".paging.normal", globalSettings.UString("wtf.paging.pageSigil", "*"))
	common.Sigils.Paging.Selected = globalSettings.UString(sigilsPath+".paging.select", globalSettings.UString("wtf.paging.selectedSigil", "_"))

	return &common
}

/* -------------------- Exported Functions -------------------- */

func (common *Common) DefaultFocusedRowColor() string {
	return fmt.Sprintf("%s:%s", common.Colors.HighlightFore, common.Colors.HighlightBack)
}

func (common *Common) DefaultRowColor() string {
	return fmt.Sprintf("%s:%s", common.Colors.Foreground, common.Colors.Background)
}

func (common *Common) FocusChar() string {
	if common.focusChar <= -1 {
		return ""
	}

	return string('0' + common.focusChar)
}

func (common *Common) RowColor(idx int) string {
	if idx%2 == 0 {
		return common.Colors.Rows.Even
	}

	return common.Colors.Rows.Odd
}

func (common *Common) RightAlignFormat(width int) string {
	borderOffset := 2
	return fmt.Sprintf("%%%ds", width-borderOffset)
}

func (common *Common) SigilStr(len, pos int, width int) string {
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
