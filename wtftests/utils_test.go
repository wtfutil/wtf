package wtftests

import (
	"testing"

	. "github.com/senorprogrammer/wtf/wtf"
	. "github.com/stretchr/testify/assert"
)

/* -------------------- CenterText() -------------------- */

func TestCenterText(t *testing.T) {
	Equal(t, "cat", CenterText("cat", -9))
	Equal(t, "cat", CenterText("cat", 0))
	Equal(t, "   cat   ", CenterText("cat", 9))
}

/* -------------------- FindMatch() -------------------- */

func TestFindMatch(t *testing.T) {
	var result [][]string

	expected := [][]string([][]string{[]string{"SSID: 7E5B5C", "7E5B5C"}})
	result = FindMatch(`s*SSID: (.+)s*`, "SSID: 7E5B5C")
	Equal(t, expected, result)
}

/* -------------------- Exclude() -------------------- */

func TestExcludeWhenTrue(t *testing.T) {
	Equal(t, true, Exclude([]string{"cat", "dog", "rat"}, "bat"))
	Equal(t, false, Exclude([]string{"cat", "dog", "rat"}, "dog"))
}

/* -------------------- NameFromEmail() -------------------- */

func TestNameFromEmail(t *testing.T) {
	Equal(t, "", NameFromEmail(""))
	Equal(t, "Chris Cummer", NameFromEmail("chris.cummer@me.com"))
}

/* -------------------- NamesFromEmails() -------------------- */

func TestNamesFromEmails(t *testing.T) {
	var result []string

	result = NamesFromEmails([]string{})
	Equal(t, []string{}, result)

	result = NamesFromEmails([]string{"chris.cummer@me.com", "chriscummer@me.com"})
	Equal(t, []string{"Chris Cummer", "Chriscummer"}, result)
}

/* -------------------- PadRow() -------------------- */

func TestPadRow(t *testing.T) {
	Equal(t, "", PadRow(0, 0))
	Equal(t, "", PadRow(5, 2))
	Equal(t, " ", PadRow(1, 2))
}

/* -------------------- ToInts() -------------------- */

func TestToInts(t *testing.T) {
	expected := []int{1, 2, 3}

	source := make([]interface{}, len(expected))
	for idx, val := range expected {
		source[idx] = val
	}

	Equal(t, expected, ToInts(source))
}

/* -------------------- ToStrs() -------------------- */

func TestToStrs(t *testing.T) {
	expected := []string{"cat", "dog", "rat"}

	source := make([]interface{}, len(expected))
	for idx, val := range expected {
		source[idx] = val
	}

	Equal(t, expected, ToStrs(source))
}
