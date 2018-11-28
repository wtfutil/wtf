package wtftests

import (
	"testing"

	"github.com/olebedev/config"
	. "github.com/stretchr/testify/assert"
	. "github.com/wtfutil/wtf/checklist"
	"github.com/wtfutil/wtf/wtf"
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
