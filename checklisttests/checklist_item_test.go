package wtftests

import (
	"testing"

	"github.com/olebedev/config"
	. "github.com/senorprogrammer/wtf/checklist"
	"github.com/senorprogrammer/wtf/wtf"
	. "github.com/stretchr/testify/assert"
)

/* -------------------- CheckMark -------------------- */

func TestCheckMark(t *testing.T) {
	loadConfig()

	item := ChecklistItem{}
	Equal(t, " ", item.CheckMark())

	item = ChecklistItem{Checked: true}
	Equal(t, "x", item.CheckMark())
}

/* -------------------- Toggle -------------------- */

func TestToggle(t *testing.T) {
	loadConfig()

	item := ChecklistItem{}
	Equal(t, false, item.Checked)

	item.Toggle()
	Equal(t, true, item.Checked)

	item.Toggle()
	Equal(t, false, item.Checked)
}

/* -------------------- helpers -------------------- */

func loadConfig() {
	wtf.Config, _ = config.ParseYamlFile("../_sample_configs/simple_config.yml")
}
