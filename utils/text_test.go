package utils

import (
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
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
