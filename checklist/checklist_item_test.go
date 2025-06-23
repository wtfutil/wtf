package checklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testChecklistItem() *ChecklistItem {
	item := NewChecklistItem(
		false,
		nil,
		make([]string, 0),
		"test",
		"",
		"",
	)
	return item
}

func Test_CheckMark(t *testing.T) {
	item := testChecklistItem()
	assert.Equal(t, " ", item.CheckMark())

	item.Toggle()
	assert.Equal(t, "x", item.CheckMark())
}

func Test_Toggle(t *testing.T) {
	item := testChecklistItem()
	assert.Equal(t, false, item.Checked)

	item.Toggle()
	assert.Equal(t, true, item.Checked)

	item.Toggle()
	assert.Equal(t, false, item.Checked)
}
