package cfg

import (
	"github.com/olebedev/config"
	"gopkg.in/yaml.v2"
)

// BorderTheme defines the default color scheme for drawing widget borders
type BorderTheme struct {
	Focusable   string
	Focused     string
	Unfocusable string
}

// CheckboxTheme defines the default color scheme for drawing checkable rows in widgets
type CheckboxTheme struct {
	Checked string
}

// RowTheme defines the default color scheme for row text
type RowTheme struct {
	EvenBackground string
	EvenForeground string

	OddBackground string
	OddForeground string

	HighlightedBackground string
	HighlightedForeground string
}

// TextTheme defines the default color scheme for text rendering
type TextTheme struct {
	Label      string
	Subheading string
	Text       string
	Title      string
}

// WidgetTheme defines the default color scheme for the widget rect itself
type WidgetTheme struct {
	Background string
}

// ColorTheme is an alamgam of all the default color settings
type ColorTheme struct {
	BorderTheme
	CheckboxTheme
	RowTheme
	TextTheme
	WidgetTheme
}

// NewDefaultColorTheme creates and returns an instance of DefaultColorTheme
func NewDefaultColorTheme() ColorTheme {
	defaultTheme := ColorTheme{
		BorderTheme: BorderTheme{
			Focusable:   "blue",
			Focused:     "orange",
			Unfocusable: "gray",
		},

		CheckboxTheme: CheckboxTheme{
			Checked: "gray",
		},

		RowTheme: RowTheme{
			EvenBackground: "transparent",
			EvenForeground: "white",

			OddBackground: "transparent",
			OddForeground: "lightblue",

			HighlightedForeground: "black",
			HighlightedBackground: "green",
		},

		TextTheme: TextTheme{
			Label:      "lightblue",
			Subheading: "red",
			Text:       "white",
			Title:      "green",
		},

		WidgetTheme: WidgetTheme{
			Background: "transparent",
		},
	}

	return defaultTheme
}

// NewDefaultColorConfig creates and returns a config.Config-compatible configuration struct
// using a DefaultColorTheme to pre-populate all the relevant values
func NewDefaultColorConfig() (*config.Config, error) {
	colorTheme := NewDefaultColorTheme()

	yamlBytes, err := yaml.Marshal(colorTheme)
	if err != nil {
		return nil, err
	}

	cfg, err := config.ParseYamlBytes(yamlBytes)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
