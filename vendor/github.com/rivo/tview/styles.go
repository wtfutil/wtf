package tview

import "github.com/gdamore/tcell"

// Theme defines the colors used when primitives are initialized.
type Theme struct {
	PrimitiveBackgroundColor    tcell.Color // Main background color for primitives.
	ContrastBackgroundColor     tcell.Color // Background color for contrasting elements.
	MoreContrastBackgroundColor tcell.Color // Background color for even more contrasting elements.
	BorderColor                 tcell.Color // Box borders.
	TitleColor                  tcell.Color // Box titles.
	GraphicsColor               tcell.Color // Graphics.
	PrimaryTextColor            tcell.Color // Primary text.
	SecondaryTextColor          tcell.Color // Secondary text (e.g. labels).
	TertiaryTextColor           tcell.Color // Tertiary text (e.g. subtitles, notes).
	InverseTextColor            tcell.Color // Text on primary-colored backgrounds.
	ContrastSecondaryTextColor  tcell.Color // Secondary text on ContrastBackgroundColor-colored backgrounds.
}

// Styles defines the theme for applications. The default is for a black
// background and some basic colors: black, white, yellow, green, cyan, and
// blue.
var Styles = Theme{
	PrimitiveBackgroundColor:    tcell.ColorBlack,
	ContrastBackgroundColor:     tcell.ColorBlue,
	MoreContrastBackgroundColor: tcell.ColorGreen,
	BorderColor:                 tcell.ColorWhite,
	TitleColor:                  tcell.ColorWhite,
	GraphicsColor:               tcell.ColorWhite,
	PrimaryTextColor:            tcell.ColorWhite,
	SecondaryTextColor:          tcell.ColorYellow,
	TertiaryTextColor:           tcell.ColorGreen,
	InverseTextColor:            tcell.ColorBlue,
	ContrastSecondaryTextColor:  tcell.ColorDarkCyan,
}
