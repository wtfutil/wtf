package wtftests

import (
	"testing"

	. "github.com/stretchr/testify/assert"
	. "github.com/wtfutil/wtf/checklist"
)

/* -------------------- CheckMark -------------------- */

func TestCheckMark(t *testing.T) {
	item := ChecklistItem{}
	Equal(t, " ", item.CheckMark())

	item = ChecklistItem{Checked: true}
	Equal(t, "x", item.CheckMark())
}

/* -------------------- Toggle -------------------- */

func TestToggle(t *testing.T) {
	item := ChecklistItem{}
	Equal(t, false, item.Checked)

	item.Toggle()
	Equal(t, true, item.Checked)

	item.Toggle()
	Equal(t, false, item.Checked)
}
