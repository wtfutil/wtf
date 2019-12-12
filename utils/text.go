package utils

import (
	"fmt"
	"strings"

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
func Truncate(src string, len int, withEllipse bool) string {
	var numRunes = 0

	for index := range src {
		numRunes++
		if numRunes > len {
			if withEllipse == true {
				return src[:index-1] + "…"
			}

			return src[:index]
		}
	}
	return src
}
