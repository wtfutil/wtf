package utils

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/rivo/tview"
)

// CenterText takes a string and a width and pads the left and right of the string with
// empty spaces to ensure that the string is in the middle of the returned value
//
// Example:
//
//    x := CenterText("cat", 11)
//    > "    cat    "
//
func CenterText(str string, width int) string {
	if width < 0 {
		width = 0
	}

	return fmt.Sprintf("%[1]*s", -width, fmt.Sprintf("%[1]*s", (width+len(str))/2, str))
}

// HighlightableHelper pads the given text with blank spaces to the width of the view
// containing it. This is helpful for extending row highlighting across the entire width
// of the view
func HighlightableHelper(view *tview.TextView, input string, idx, offset int) string {
	_, _, w, _ := view.GetInnerRect()

	fmtStr := fmt.Sprintf(`["%d"][""]`, idx)
	fmtStr += input
	fmtStr += RowPadding(offset, w)
	fmtStr += `[""]` + "\n"

	return fmtStr
}

// RowPadding returns a padding for a row to make it the full width of the containing widget.
// Useful for ensuring row highlighting spans the full width (I suspect tcell has a better
// way to do this, but I haven't yet found it)
func RowPadding(offset int, max int) string {
	padSize := max - offset
	if padSize < 0 {
		padSize = 0
	}

	return strings.Repeat(" ", padSize)
}

// Truncate chops a given string at len length. Appends an ellipse character if warranted
func Truncate(src string, maxLen int, withEllipse bool) string {
	if len(src) < 1 || maxLen < 1 {
		return ""
	}

	if maxLen == 1 {
		return src[:1]
	}

	var runeCount = 0
	for idx := range src {
		runeCount++
		if runeCount > maxLen {
			if withEllipse {
				return src[:idx-1] + "…"
			}

			return src[:idx]
		}
	}
	return src
}

// Formats number as string with 1000 delimiters and, if necessary, rounds it to 2 decimals
func PrettyNumber(number float64) string {
	p := message.NewPrinter(language.English)
	if number == math.Trunc(number) {
		return p.Sprintf("%.0f", number)
	} else {
		return p.Sprintf("%.2f", number)
	}
}
