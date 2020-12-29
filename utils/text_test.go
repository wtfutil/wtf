package utils

import (
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Test_CenterText(t *testing.T) {
	assert.Equal(t, "cat", CenterText("cat", -9))
	assert.Equal(t, "cat", CenterText("cat", 0))
	assert.Equal(t, "   cat   ", CenterText("cat", 9))
}

func Test_HighlightableHelper(t *testing.T) {
	view := tview.NewTextView()
	actual := HighlightableHelper(view, "cats", 0, 5)

	assert.Equal(t, "[\"0\"][\"\"]cats          [\"\"]\n", actual)
}

func Test_RowPadding(t *testing.T) {
	assert.Equal(t, "", RowPadding(0, 0))
	assert.Equal(t, "", RowPadding(5, 2))
	assert.Equal(t, " ", RowPadding(1, 2))
	assert.Equal(t, "     ", RowPadding(0, 5))
}

func Test_Truncate(t *testing.T) {
	assert.Equal(t, "", Truncate("cat", 0, false))
	assert.Equal(t, "c", Truncate("cat", 1, false))
	assert.Equal(t, "ca", Truncate("cat", 2, false))
	assert.Equal(t, "cat", Truncate("cat", 3, false))
	assert.Equal(t, "cat", Truncate("cat", 4, false))

	assert.Equal(t, "", Truncate("cat", 0, true))
	assert.Equal(t, "c", Truncate("cat", 1, true))
	assert.Equal(t, "c…", Truncate("cat", 2, true))
	assert.Equal(t, "cat", Truncate("cat", 3, true))
	assert.Equal(t, "cat", Truncate("cat", 4, true))

	// Only supports non-ellipsed emoji
	assert.Equal(t, "🌮🚙", Truncate("🌮🚙💥👾", 2, false))
}

func Test_PrettyNumber(t *testing.T) {
	locPrinter := message.NewPrinter(language.English)

	assert.Equal(t, "1,000,000", PrettyNumber(locPrinter, 1000000))
	assert.Equal(t, "1,000,000.99", PrettyNumber(locPrinter, 1000000.99))
	assert.Equal(t, "1,000,000", PrettyNumber(locPrinter, 1000000.00))
	assert.Equal(t, "100,000", PrettyNumber(locPrinter, 100000))
	assert.Equal(t, "100,000.01", PrettyNumber(locPrinter, 100000.009))
	assert.Equal(t, "10,000", PrettyNumber(locPrinter, 10000))
	assert.Equal(t, "1,000", PrettyNumber(locPrinter, 1000))
	assert.Equal(t, "1,000", PrettyNumber(locPrinter, 1000))
	assert.Equal(t, "100", PrettyNumber(locPrinter, 100))
	assert.Equal(t, "0", PrettyNumber(locPrinter, 0))
	assert.Equal(t, "0.10", PrettyNumber(locPrinter, 0.1))
}
