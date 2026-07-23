package digitalclock

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func TestGetFont(t *testing.T) {
	tests := []struct {
		name         string
		font         string
		expectedRows int
	}{
		{"digital font", "digitalfont", 3},
		{"bold font", "boldfont", 5},
		{"big font (default)", "bigfont", 6},
		{"unknown defaults to big", "unknown", 6},
		{"empty defaults to big", "", 6},
		{"case insensitive", "DigitalFont", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := Settings{font: tt.font}
			font := getFont(settings)
			if font.fontRows != tt.expectedRows {
				t.Errorf("getFont(%q).fontRows = %d, want %d", tt.font, font.fontRows, tt.expectedRows)
			}
		})
	}
}

func TestClockFontGet(t *testing.T) {
	font := getDigitalFont()

	// Verify all expected characters exist
	expectedChars := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ":", " ", "A", "P"}
	for _, char := range expectedChars {
		result := font.get(char)
		if result == nil {
			t.Errorf("font.get(%q) returned nil", char)
		}
		if len(result) != font.fontRows {
			t.Errorf("font.get(%q) has %d rows, want %d", char, len(result), font.fontRows)
		}
	}
}

func TestFontsJoin(t *testing.T) {
	font := getDigitalFont()
	chars := [][]string{
		font.get("1"),
		font.get("2"),
	}

	result := fontsJoin(chars, font.fontRows, "white")

	// Should have fontRows lines
	lines := strings.Split(result, "\n")
	if len(lines) != font.fontRows {
		t.Errorf("fontsJoin produced %d lines, want %d", len(lines), font.fontRows)
	}

	// Each line should contain the color directive
	for _, line := range lines {
		if !strings.Contains(line, "[white]") {
			t.Errorf("line missing color: %q", line)
		}
	}
}

func TestGetDigitalFontCompleteness(t *testing.T) {
	font := getDigitalFont()
	if font.fontRows != 3 {
		t.Errorf("digital font rows = %d, want 3", font.fontRows)
	}
}

func TestGetBigFontCompleteness(t *testing.T) {
	font := getBigFont()
	if font.fontRows != 6 {
		t.Errorf("big font rows = %d, want 6", font.fontRows)
	}

	// All digits should have 6 rows
	for i := 0; i <= 9; i++ {
		char := string(rune('0' + i))
		rows := font.get(char)
		if len(rows) != 6 {
			t.Errorf("big font char %q has %d rows, want 6", char, len(rows))
		}
	}
}

func TestGetBoldFontCompleteness(t *testing.T) {
	font := getBoldFont()
	if font.fontRows != 5 {
		t.Errorf("bold font rows = %d, want 5", font.fontRows)
	}
}

// digitChars are the glyphs that must render at a fixed column width per font.
var digitChars = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

// TestFontDigitWidthsUniform guards against the misalignment bug where a
// digit glyph narrower or wider than its peers throws off column alignment.
func TestFontDigitWidthsUniform(t *testing.T) {
	fonts := map[string]ClockFont{
		"digitalfont": getDigitalFont(),
		"bigfont":     getBigFont(),
		"boldfont":    getBoldFont(),
	}

	for fontName, font := range fonts {
		t.Run(fontName, func(t *testing.T) {
			wantWidth := -1
			for _, char := range digitChars {
				rows := font.get(char)
				if len(rows) != font.fontRows {
					t.Fatalf("char %q has %d rows, want %d", char, len(rows), font.fontRows)
				}
				for rowIdx, row := range rows {
					width := utf8.RuneCountInString(row)
					if wantWidth == -1 {
						wantWidth = width
					}
					if width != wantWidth {
						t.Errorf("char %q row %d width = %d, want %d (digit glyphs must share one width per font)", char, rowIdx, width, wantWidth)
					}
				}
			}
		})
	}
}

// TestFontGlyphRowsSelfConsistent ensures that within a single glyph, every
// row is the same width.
func TestFontGlyphRowsSelfConsistent(t *testing.T) {
	fonts := map[string]ClockFont{
		"digitalfont": getDigitalFont(),
		"bigfont":     getBigFont(),
		"boldfont":    getBoldFont(),
	}

	for fontName, font := range fonts {
		t.Run(fontName, func(t *testing.T) {
			for char, rows := range font.fonts {
				if char == "A" || char == "P" || len(rows) == 0 {
					continue
				}
				width := utf8.RuneCountInString(rows[0])
				for rowIdx, row := range rows {
					if w := utf8.RuneCountInString(row); w != width {
						t.Errorf("char %q row %d width = %d, want %d (all rows of one glyph must match)", char, rowIdx, w, width)
					}
				}
			}
		})
	}
}
