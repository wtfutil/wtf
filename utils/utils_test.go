package utils

import (
	"os/exec"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func Test_Init(t *testing.T) {
	Init("cats")
	Equal(t, OpenFileUtil, "cats")
}

func Test_CenterText(t *testing.T) {
	Equal(t, "cat", CenterText("cat", -9))
	Equal(t, "cat", CenterText("cat", 0))
	Equal(t, "   cat   ", CenterText("cat", 9))
}

func Test_ExecuteCommand(t *testing.T) {
	tests := []struct {
		name     string
		cmd      *exec.Cmd
		expected string
	}{
		{
			name:     "with nil command",
			cmd:      nil,
			expected: "",
		},
		{
			name:     "with defined command",
			cmd:      exec.Command("echo", "cats"),
			expected: "cats\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ExecuteCommand(tt.cmd)

			if tt.expected != actual {
				t.Errorf("\nexpected: %s\n     got: %s", tt.expected, actual)
			}
		})
	}
}

func Test_FindMatch(t *testing.T) {
	var result [][]string

	expected := [][]string([][]string{[]string{"SSID: 7E5B5C", "7E5B5C"}})
	result = FindMatch(`s*SSID: (.+)s*`, "SSID: 7E5B5C")
	Equal(t, expected, result)
}

func Test_ExcludeWhenTrue(t *testing.T) {
	Equal(t, true, Exclude([]string{"cat", "dog", "rat"}, "bat"))
	Equal(t, false, Exclude([]string{"cat", "dog", "rat"}, "dog"))
}

func Test_NameFromEmail(t *testing.T) {
	Equal(t, "", NameFromEmail(""))
	Equal(t, "Chris Cummer", NameFromEmail("chris.cummer@me.com"))
}

func Test_NamesFromEmails(t *testing.T) {
	var result []string

	result = NamesFromEmails([]string{})
	Equal(t, []string{}, result)

	result = NamesFromEmails([]string{"chris.cummer@me.com", "chriscummer@me.com"})
	Equal(t, []string{"Chris Cummer", "Chriscummer"}, result)
}

func Test_PadRow(t *testing.T) {
	Equal(t, "", PadRow(0, 0))
	Equal(t, "", PadRow(5, 2))
	Equal(t, " ", PadRow(1, 2))
}

func Test_MapToStrs(t *testing.T) {
	expected := map[string]string{
		"a": "a",
		"b": "b",
		"c": "c",
	}

	source := make(map[string]interface{})
	for _, val := range expected {
		source[val] = val
	}

	Equal(t, expected, MapToStrs(source))
}

func Test_ToInts(t *testing.T) {
	expected := []int{1, 2, 3}

	source := make([]interface{}, len(expected))
	for idx, val := range expected {
		source[idx] = val
	}

	Equal(t, expected, ToInts(source))
}

func Test_ToStrs(t *testing.T) {
	expected := []string{"cat", "dog", "rat"}

	source := make([]interface{}, len(expected))
	for idx, val := range expected {
		source[idx] = val
	}

	Equal(t, expected, ToStrs(source))
}
