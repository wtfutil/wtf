package utils

import (
	"os/exec"
	"reflect"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func Test_DoesNotInclude(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		val      string
		expected bool
	}{
		{
			name:     "when included",
			strs:     []string{"a", "b", "c"},
			val:      "b",
			expected: false,
		},
		{
			name:     "when not included",
			strs:     []string{"a", "b", "c"},
			val:      "f",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := DoesNotInclude(tt.strs, tt.val)

			if tt.expected != actual {
				t.Errorf("\nexpected: %t\n     got: %t", tt.expected, actual)
			}
		})
	}
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

func Test_Includes(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		val      string
		expected bool
	}{
		{
			name:     "when included",
			strs:     []string{"a", "b", "c"},
			val:      "b",
			expected: true,
		},
		{
			name:     "when not included",
			strs:     []string{"a", "b", "c"},
			val:      "f",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Includes(tt.strs, tt.val)

			if tt.expected != actual {
				t.Errorf("\nexpected: %t\n     got: %t", tt.expected, actual)
			}
		})
	}
}

func Test_ReadFileBytes(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		expected []byte
	}{
		{
			name:     "with non-existant file",
			file:     "/tmp/junk-daa6bf613f4c.md",
			expected: []byte{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, _ := ReadFileBytes(tt.file)

			if reflect.DeepEqual(tt.expected, actual) == false {
				t.Errorf("\nexpected: %q\n     got: %q", tt.expected, actual)
			}
		})
	}
}
