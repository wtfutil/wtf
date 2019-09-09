package digitalclock

import (
	"fmt"
	"strings"
)

// ClockFontInterface to makes sure all fonts implement join and get methods
type ClockFontInterface interface {
	join() string
	get() string
}

// ClockFont struct to hold the font info
type ClockFont struct {
	fontRows int
	fonts    map[string][]string
}

// function to join fonts, since the fonts have multi rows
func fontsJoin(fontCharArray [][]string, rows int, color string) string {
	outString := ""

	for i := 0; i < rows; i++ {
		outString += fmt.Sprintf("[%s]", color)
		for _, charFont := range fontCharArray {
			outString += " " + fmt.Sprintf("[%s]%s", color, charFont[i])
		}
		outString += "\n"
	}
	return strings.TrimSuffix(outString, "\n")
}

func (font *ClockFont) get(char string) []string {
	return font.fonts[char]
}

func getDigitalFont() ClockFont {
	fontsMap := map[string][]string{
		"1": {"▄█ ", " █ ", "▄█▄"},
		"2": {"█▀█", " ▄▀", "█▄▄"},
		"3": {"█▀▀█", "  ▀▄", "█▄▄█"},
		"4": {" █▀█ ", "█▄▄█▄", "   █ "},
		"5": {"█▀▀", "▀▀▄", "▄▄▀"},
		"6": {"▄▀▀▄", "█▄▄ ", "▀▄▄▀"},
		"7": {"▀▀▀█", "  █ ", " ▐▌ "},
		"8": {"▄▀▀▄", "▄▀▀▄", "▀▄▄▀"},
		"9": {"▄▀▀▄", "▀▄▄█", " ▄▄▀"},
		"0": {"█▀▀█", "█  █", "█▄▄█"},
		":": {"█", " ", "█"},
		" ": {" ", " ", " "},
		"A": {"", "", "AM"},
		"P": {"", "", "PM"},
	}

	digitalFont := ClockFont{fontRows: 3, fonts: fontsMap}
	return digitalFont
}

func getBigFont() ClockFont {
	fontsMap := map[string][]string{
		"1": {" ┏┓ ", "┏┛┃ ", "┗┓┃ ", " ┃┃ ", "┏┛┗┓", "┗━━┛"},
		"2": {"┏━━━┓", "┃┏━┓┃", "┗┛┏┛┃", "┏━┛┏┛", "┃ ┗━┓", "┗━━━┛"},
		"3": {"┏━━━┓", "┃┏━┓┃", "┗┛┏┛┃", "┏┓┗┓┃", "┃┗━┛┃", "┗━━━┛"},
		"4": {"┏┓ ┏┓", "┃┃ ┃┃", "┃┗━┛┃", "┗━━┓┃", "   ┃┃", "   ┗┛"},
		"5": {"┏━━━┓", "┃┏━━┛", "┃┗━━┓", "┗━━┓┃", "┏━━┛┃", "┗━━━┛"},
		"6": {"┏━━━┓", "┃┏━━┛", "┃┗━━┓", "┃┏━┓┃", "┃┗━┛┃", "┗━━━┛"},
		"7": {"┏━━━┓", "┃┏━┓┃", "┗┛┏┛┃", "  ┃┏┛", "  ┃┃ ", "  ┗┛ "},
		"8": {"┏━━━┓", "┃┏━┓┃", "┃┗━┛┃", "┃┏━┓┃", "┃┗━┛┃", "┗━━━┛"},
		"9": {"┏━━━┓", "┃┏━┓┃", "┃┗━┛┃", "┗━━┓┃", "┏━━┛┃", "┗━━━┛"},
		"0": {"┏━━━┓", "┃┏━┓┃", "┃┃ ┃┃", "┃┃ ┃┃", "┃┗━┛┃", "┗━━━┛"},
		":": {"   ", "┏━┓", "┗━┛", "┏━┓", "┗━┛", "   "},
		" ": {"   ", "   ", "   ", "   ", "   ", "   "},
		"A": {"", "", "", "", "", "AM"},
		"P": {"", "", "", "", "", "PM"},
	}

	bigFont := ClockFont{fontRows: 6, fonts: fontsMap}
	return bigFont
}

// getFont returns appropriate font map based on the font settings
func getFont(widgetSettings Settings) ClockFont {
	if strings.ToLower(widgetSettings.font) == "digitalfont" {
		return getDigitalFont()
	}
	return getBigFont()
}
