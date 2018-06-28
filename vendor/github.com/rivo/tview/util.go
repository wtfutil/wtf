package tview

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/gdamore/tcell"
	runewidth "github.com/mattn/go-runewidth"
)

// Text alignment within a box.
const (
	AlignLeft = iota
	AlignCenter
	AlignRight
)

// Common regular expressions.
var (
	colorPattern     = regexp.MustCompile(`\[([a-zA-Z]+|#[0-9a-zA-Z]{6}|\-)?(:([a-zA-Z]+|#[0-9a-zA-Z]{6}|\-)?(:([lbdru]+|\-)?)?)?\]`)
	regionPattern    = regexp.MustCompile(`\["([a-zA-Z0-9_,;: \-\.]*)"\]`)
	escapePattern    = regexp.MustCompile(`\[([a-zA-Z0-9_,;: \-\."#]+)\[(\[*)\]`)
	nonEscapePattern = regexp.MustCompile(`(\[[a-zA-Z0-9_,;: \-\."#]+\[*)\]`)
	boundaryPattern  = regexp.MustCompile("([[:punct:]]\\s*|\\s+)")
	spacePattern     = regexp.MustCompile(`\s+`)
)

// Positions of substrings in regular expressions.
const (
	colorForegroundPos = 1
	colorBackgroundPos = 3
	colorFlagPos       = 5
)

// Predefined InputField acceptance functions.
var (
	// InputFieldInteger accepts integers.
	InputFieldInteger func(text string, ch rune) bool

	// InputFieldFloat accepts floating-point numbers.
	InputFieldFloat func(text string, ch rune) bool

	// InputFieldMaxLength returns an input field accept handler which accepts
	// input strings up to a given length. Use it like this:
	//
	//   inputField.SetAcceptanceFunc(InputFieldMaxLength(10)) // Accept up to 10 characters.
	InputFieldMaxLength func(maxLength int) func(text string, ch rune) bool
)

// Package initialization.
func init() {
	// Initialize the predefined input field handlers.
	InputFieldInteger = func(text string, ch rune) bool {
		if text == "-" {
			return true
		}
		_, err := strconv.Atoi(text)
		return err == nil
	}
	InputFieldFloat = func(text string, ch rune) bool {
		if text == "-" || text == "." || text == "-." {
			return true
		}
		_, err := strconv.ParseFloat(text, 64)
		return err == nil
	}
	InputFieldMaxLength = func(maxLength int) func(text string, ch rune) bool {
		return func(text string, ch rune) bool {
			return len([]rune(text)) <= maxLength
		}
	}
}

// styleFromTag takes the given style, defined by a foreground color (fgColor),
// a background color (bgColor), and style attributes, and modifies it based on
// the substrings (tagSubstrings) extracted by the regular expression for color
// tags. The new colors and attributes are returned where empty strings mean
// "don't modify" and a dash ("-") means "reset to default".
func styleFromTag(fgColor, bgColor, attributes string, tagSubstrings []string) (newFgColor, newBgColor, newAttributes string) {
	if tagSubstrings[colorForegroundPos] != "" {
		color := tagSubstrings[colorForegroundPos]
		if color == "-" {
			fgColor = "-"
		} else if color != "" {
			fgColor = color
		}
	}

	if tagSubstrings[colorBackgroundPos-1] != "" {
		color := tagSubstrings[colorBackgroundPos]
		if color == "-" {
			bgColor = "-"
		} else if color != "" {
			bgColor = color
		}
	}

	if tagSubstrings[colorFlagPos-1] != "" {
		flags := tagSubstrings[colorFlagPos]
		if flags == "-" {
			attributes = "-"
		} else if flags != "" {
			attributes = flags
		}
	}

	return fgColor, bgColor, attributes
}

// overlayStyle mixes a background color with a foreground color (fgColor),
// a (possibly new) background color (bgColor), and style attributes, and
// returns the resulting style. For a definition of the colors and attributes,
// see styleFromTag(). Reset instructions cause the corresponding part of the
// default style to be used.
func overlayStyle(background tcell.Color, defaultStyle tcell.Style, fgColor, bgColor, attributes string) tcell.Style {
	defFg, defBg, defAttr := defaultStyle.Decompose()
	style := defaultStyle.Background(background)

	style = style.Foreground(defFg)
	if fgColor != "" {
		style = style.Foreground(tcell.GetColor(fgColor))
	}

	if bgColor == "-" || bgColor == "" && defBg != tcell.ColorDefault {
		style = style.Background(defBg)
	} else if bgColor != "" {
		style = style.Background(tcell.GetColor(bgColor))
	}

	if attributes == "-" {
		style = style.Bold(defAttr&tcell.AttrBold > 0)
		style = style.Blink(defAttr&tcell.AttrBlink > 0)
		style = style.Reverse(defAttr&tcell.AttrReverse > 0)
		style = style.Underline(defAttr&tcell.AttrUnderline > 0)
		style = style.Dim(defAttr&tcell.AttrDim > 0)
	} else if attributes != "" {
		style = style.Normal()
		for _, flag := range attributes {
			switch flag {
			case 'l':
				style = style.Blink(true)
			case 'b':
				style = style.Bold(true)
			case 'd':
				style = style.Dim(true)
			case 'r':
				style = style.Reverse(true)
			case 'u':
				style = style.Underline(true)
			}
		}
	}

	return style
}

// decomposeString returns information about a string which may contain color
// tags. It returns the indices of the color tags (as returned by
// re.FindAllStringIndex()), the color tags themselves (as returned by
// re.FindAllStringSubmatch()), the indices of an escaped tags, the string
// stripped by any color tags and escaped, and the screen width of the stripped
// string.
func decomposeString(text string) (colorIndices [][]int, colors [][]string, escapeIndices [][]int, stripped string, width int) {
	// Get positions of color and escape tags.
	colorIndices = colorPattern.FindAllStringIndex(text, -1)
	colors = colorPattern.FindAllStringSubmatch(text, -1)
	escapeIndices = escapePattern.FindAllStringIndex(text, -1)

	// Because the color pattern detects empty tags, we need to filter them out.
	for i := len(colorIndices) - 1; i >= 0; i-- {
		if colorIndices[i][1]-colorIndices[i][0] == 2 {
			colorIndices = append(colorIndices[:i], colorIndices[i+1:]...)
			colors = append(colors[:i], colors[i+1:]...)
		}
	}

	// Remove the color tags from the original string.
	var from int
	buf := make([]byte, 0, len(text))
	for _, indices := range colorIndices {
		buf = append(buf, []byte(text[from:indices[0]])...)
		from = indices[1]
	}
	buf = append(buf, text[from:]...)

	// Escape string.
	stripped = string(escapePattern.ReplaceAll(buf, []byte("[$1$2]")))

	// Get the width of the stripped string.
	width = runewidth.StringWidth(stripped)

	return
}

// Print prints text onto the screen into the given box at (x,y,maxWidth,1),
// not exceeding that box. "align" is one of AlignLeft, AlignCenter, or
// AlignRight. The screen's background color will not be changed.
//
// You can change the colors and text styles mid-text by inserting a color tag.
// See the package description for details.
//
// Returns the number of actual runes printed (not including color tags) and the
// actual width used for the printed runes.
func Print(screen tcell.Screen, text string, x, y, maxWidth, align int, color tcell.Color) (int, int) {
	return printWithStyle(screen, text, x, y, maxWidth, align, tcell.StyleDefault.Foreground(color))
}

// printWithStyle works like Print() but it takes a style instead of just a
// foreground color.
func printWithStyle(screen tcell.Screen, text string, x, y, maxWidth, align int, style tcell.Style) (int, int) {
	if maxWidth <= 0 || len(text) == 0 {
		return 0, 0
	}

	// Decompose the text.
	colorIndices, colors, escapeIndices, strippedText, _ := decomposeString(text)

	// We deal with runes, not with bytes.
	runes := []rune(strippedText)

	// This helper function takes positions for a substring of "runes" and returns
	// a new string corresponding to this substring, making sure printing that
	// substring will observe color tags.
	substring := func(from, to int) string {
		var (
			colorPos, escapePos, runePos, startPos       int
			foregroundColor, backgroundColor, attributes string
		)
		if from >= len(runes) {
			return ""
		}
		for pos := range text {
			// Handle color tags.
			if colorPos < len(colorIndices) && pos >= colorIndices[colorPos][0] && pos < colorIndices[colorPos][1] {
				if pos == colorIndices[colorPos][1]-1 {
					if runePos <= from {
						foregroundColor, backgroundColor, attributes = styleFromTag(foregroundColor, backgroundColor, attributes, colors[colorPos])
					}
					colorPos++
				}
				continue
			}

			// Handle escape tags.
			if escapePos < len(escapeIndices) && pos >= escapeIndices[escapePos][0] && pos < escapeIndices[escapePos][1] {
				if pos == escapeIndices[escapePos][1]-1 {
					escapePos++
				} else if pos == escapeIndices[escapePos][1]-2 {
					continue
				}
			}

			// Check boundaries.
			if runePos == from {
				startPos = pos
			} else if runePos >= to {
				return fmt.Sprintf(`[%s:%s:%s]%s`, foregroundColor, backgroundColor, attributes, text[startPos:pos])
			}

			runePos++
		}

		return fmt.Sprintf(`[%s:%s:%s]%s`, foregroundColor, backgroundColor, attributes, text[startPos:])
	}

	// We want to reduce everything to AlignLeft.
	if align == AlignRight {
		width := 0
		start := len(runes)
		for index := start - 1; index >= 0; index-- {
			w := runewidth.RuneWidth(runes[index])
			if width+w > maxWidth {
				break
			}
			width += w
			start = index
		}
		for start < len(runes) && runewidth.RuneWidth(runes[start]) == 0 {
			start++
		}
		return printWithStyle(screen, substring(start, len(runes)), x+maxWidth-width, y, width, AlignLeft, style)
	} else if align == AlignCenter {
		width := runewidth.StringWidth(strippedText)
		if width == maxWidth {
			// Use the exact space.
			return printWithStyle(screen, text, x, y, maxWidth, AlignLeft, style)
		} else if width < maxWidth {
			// We have more space than we need.
			half := (maxWidth - width) / 2
			return printWithStyle(screen, text, x+half, y, maxWidth-half, AlignLeft, style)
		} else {
			// Chop off runes until we have a perfect fit.
			var choppedLeft, choppedRight, leftIndex, rightIndex int
			rightIndex = len(runes) - 1
			for rightIndex > leftIndex && width-choppedLeft-choppedRight > maxWidth {
				if choppedLeft < choppedRight {
					leftWidth := runewidth.RuneWidth(runes[leftIndex])
					choppedLeft += leftWidth
					leftIndex++
					for leftIndex < len(runes) && leftIndex < rightIndex && runewidth.RuneWidth(runes[leftIndex]) == 0 {
						leftIndex++
					}
				} else {
					rightWidth := runewidth.RuneWidth(runes[rightIndex])
					choppedRight += rightWidth
					rightIndex--
				}
			}
			return printWithStyle(screen, substring(leftIndex, rightIndex), x, y, maxWidth, AlignLeft, style)
		}
	}

	// Draw text.
	drawn := 0
	drawnWidth := 0
	var (
		colorPos, escapePos                          int
		foregroundColor, backgroundColor, attributes string
	)
	runeSequence := make([]rune, 0, 10)
	runeSeqWidth := 0
	flush := func() {
		if len(runeSequence) == 0 {
			return // Nothing to flush.
		}

		// Print the rune sequence.
		finalX := x + drawnWidth
		_, _, finalStyle, _ := screen.GetContent(finalX, y)
		_, background, _ := finalStyle.Decompose()
		finalStyle = overlayStyle(background, style, foregroundColor, backgroundColor, attributes)
		var comb []rune
		if len(runeSequence) > 1 && !unicode.IsControl(runeSequence[1]) {
			// Allocate space for the combining characters only when necessary.
			comb = make([]rune, len(runeSequence)-1)
			copy(comb, runeSequence[1:])
		}
		for offset := 0; offset < runeSeqWidth; offset++ {
			// To avoid undesired effects, we place the same character in all cells.
			screen.SetContent(finalX+offset, y, runeSequence[0], comb, finalStyle)
		}

		// Advance and reset.
		drawn += len(runeSequence)
		drawnWidth += runeSeqWidth
		runeSequence = runeSequence[:0]
		runeSeqWidth = 0
	}
	for pos, ch := range text {
		// Handle color tags.
		if colorPos < len(colorIndices) && pos >= colorIndices[colorPos][0] && pos < colorIndices[colorPos][1] {
			flush()
			if pos == colorIndices[colorPos][1]-1 {
				foregroundColor, backgroundColor, attributes = styleFromTag(foregroundColor, backgroundColor, attributes, colors[colorPos])
				colorPos++
			}
			continue
		}

		// Handle escape tags.
		if escapePos < len(escapeIndices) && pos >= escapeIndices[escapePos][0] && pos < escapeIndices[escapePos][1] {
			flush()
			if pos == escapeIndices[escapePos][1]-1 {
				escapePos++
			} else if pos == escapeIndices[escapePos][1]-2 {
				continue
			}
		}

		// Check if we have enough space for this rune.
		chWidth := runewidth.RuneWidth(ch)
		if drawnWidth+chWidth > maxWidth {
			break // No. We're done then.
		}

		// Put this rune in the queue.
		if chWidth == 0 {
			// If this is not a modifier, we treat it as a space character.
			if len(runeSequence) == 0 {
				ch = ' '
				chWidth = 1
			}
		} else {
			// We have a character. Flush all previous runes.
			flush()
		}
		runeSequence = append(runeSequence, ch)
		runeSeqWidth += chWidth
	}
	if drawnWidth+runeSeqWidth <= maxWidth {
		flush()
	}

	return drawn, drawnWidth
}

// PrintSimple prints white text to the screen at the given position.
func PrintSimple(screen tcell.Screen, text string, x, y int) {
	Print(screen, text, x, y, math.MaxInt32, AlignLeft, Styles.PrimaryTextColor)
}

// StringWidth returns the width of the given string needed to print it on
// screen. The text may contain color tags which are not counted.
func StringWidth(text string) int {
	_, _, _, _, width := decomposeString(text)
	return width
}

// WordWrap splits a text such that each resulting line does not exceed the
// given screen width. Possible split points are after any punctuation or
// whitespace. Whitespace after split points will be dropped.
//
// This function considers color tags to have no width.
//
// Text is always split at newline characters ('\n').
func WordWrap(text string, width int) (lines []string) {
	colorTagIndices, _, escapeIndices, strippedText, _ := decomposeString(text)

	// Find candidate breakpoints.
	breakPoints := boundaryPattern.FindAllStringIndex(strippedText, -1)

	// This helper function adds a new line to the result slice. The provided
	// positions are in stripped index space.
	addLine := func(from, to int) {
		// Shift indices back to original index space.
		var colorTagIndex, escapeIndex int
		for colorTagIndex < len(colorTagIndices) && to >= colorTagIndices[colorTagIndex][0] ||
			escapeIndex < len(escapeIndices) && to >= escapeIndices[escapeIndex][0] {
			past := 0
			if colorTagIndex < len(colorTagIndices) {
				tagWidth := colorTagIndices[colorTagIndex][1] - colorTagIndices[colorTagIndex][0]
				if colorTagIndices[colorTagIndex][0] < from {
					from += tagWidth
					to += tagWidth
					colorTagIndex++
				} else if colorTagIndices[colorTagIndex][0] < to {
					to += tagWidth
					colorTagIndex++
				} else {
					past++
				}
			} else {
				past++
			}
			if escapeIndex < len(escapeIndices) {
				tagWidth := escapeIndices[escapeIndex][1] - escapeIndices[escapeIndex][0]
				if escapeIndices[escapeIndex][0] < from {
					from += tagWidth
					to += tagWidth
					escapeIndex++
				} else if escapeIndices[escapeIndex][0] < to {
					to += tagWidth
					escapeIndex++
				} else {
					past++
				}
			} else {
				past++
			}
			if past == 2 {
				break // All other indices are beyond the requested string.
			}
		}
		lines = append(lines, text[from:to])
	}

	// Determine final breakpoints.
	var start, lastEnd, newStart, breakPoint int
	for {
		// What's our candidate string?
		var candidate string
		if breakPoint < len(breakPoints) {
			candidate = text[start:breakPoints[breakPoint][1]]
		} else {
			candidate = text[start:]
		}
		candidate = strings.TrimRightFunc(candidate, unicode.IsSpace)

		if runewidth.StringWidth(candidate) >= width {
			// We're past the available width.
			if lastEnd > start {
				// Use the previous candidate.
				addLine(start, lastEnd)
				start = newStart
			} else {
				// We have no previous candidate. Make a hard break.
				var lineWidth int
				for index, ch := range text {
					if index < start {
						continue
					}
					chWidth := runewidth.RuneWidth(ch)
					if lineWidth > 0 && lineWidth+chWidth >= width {
						addLine(start, index)
						start = index
						break
					}
					lineWidth += chWidth
				}
			}
		} else {
			// We haven't hit the right border yet.
			if breakPoint >= len(breakPoints) {
				// It's the last line. We're done.
				if len(candidate) > 0 {
					addLine(start, len(strippedText))
				}
				break
			} else {
				// We have a new candidate.
				lastEnd = start + len(candidate)
				newStart = breakPoints[breakPoint][1]
				breakPoint++
			}
		}
	}

	return
}

// Escape escapes the given text such that color and/or region tags are not
// recognized and substituted by the print functions of this package. For
// example, to include a tag-like string in a box title or in a TextView:
//
//   box.SetTitle(tview.Escape("[squarebrackets]"))
//   fmt.Fprint(textView, tview.Escape(`["quoted"]`))
func Escape(text string) string {
	return nonEscapePattern.ReplaceAllString(text, "$1[]")
}
