package cfg

import (
	"testing"
	"time"

	"github.com/olebedev/config"
	"github.com/stretchr/testify/assert"
)

var (
	testYaml = `
wtf:
  colors:
`

	moduleConfig, _   = config.ParseYaml(testYaml)
	globalSettings, _ = config.ParseYaml(testYaml)

	testCfg = NewCommonSettingsFromModule(
		"test",
		"Test Config",
		true,
		moduleConfig,
		globalSettings,
	)
)

func Test_NewCommonSettingsFromModule(t *testing.T) {
	assert.Equal(t, true, testCfg.Bordered)
	assert.Equal(t, false, testCfg.Enabled)
	assert.Equal(t, true, testCfg.Focusable)
	assert.Equal(t, "test", testCfg.Module.Name)
	assert.Equal(t, "test", testCfg.Module.Type)
	assert.Equal(t, "", testCfg.FocusChar())
	assert.Equal(t, 300*time.Second, testCfg.RefreshInterval)
	assert.Equal(t, "Test Config", testCfg.Title)
}

func Test_DefaultFocusedRowColor(t *testing.T) {
	assert.Equal(t, "black:green", testCfg.DefaultFocusedRowColor())
}

func Test_DefaultRowColor(t *testing.T) {
	assert.Equal(t, "white:transparent", testCfg.DefaultRowColor())
}

func Test_FocusChar(t *testing.T) {
	tests := []struct {
		name         string
		before       func(testCfg *Common)
		expectedChar string
	}{
		{
			name: "with negative focus char",
			before: func(testCfg *Common) {
				testCfg.focusChar = -1
			},
			expectedChar: "",
		},
		{
			name: "with positive focus char",
			before: func(testCfg *Common) {
				testCfg.focusChar = 3
			},
			expectedChar: "3",
		},
		{
			name: "with zero focus char",
			before: func(testCfg *Common) {
				testCfg.focusChar = 0
			},
			expectedChar: "",
		},
		{
			name: "with large focus char",
			before: func(testCfg *Common) {
				testCfg.focusChar = 10
			},
			expectedChar: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before(testCfg)

			assert.Equal(t, tt.expectedChar, testCfg.FocusChar())
		})
	}
}

func Test_RowColor(t *testing.T) {
	tests := []struct {
		name          string
		idx           int
		expectedColor string
	}{
		{
			name:          "odd rows, default",
			idx:           3,
			expectedColor: "lightblue:transparent",
		},
		{
			name:          "even rows, default",
			idx:           8,
			expectedColor: "white:transparent",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedColor, testCfg.RowColor(tt.idx))
		})
	}
}

func Test_RightAlignFormat(t *testing.T) {
	tests := []struct {
		name     string
		width    int
		expected string
	}{
		{
			name:     "with zero",
			width:    0,
			expected: "%-2s",
		},
		{
			name:     "with positive integer",
			width:    3,
			expected: "%1s",
		},
		{
			name:     "with negative integer",
			width:    -3,
			expected: "%-5s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, testCfg.RightAlignFormat(tt.width))
		})
	}
}

func Test_PaginationMarker(t *testing.T) {
	tests := []struct {
		name     string
		len      int
		pos      int
		width    int
		expected string
	}{
		{
			name:     "with zero pages",
			len:      0,
			pos:      1,
			width:    5,
			expected: "",
		},
		{
			name:     "with one page",
			len:      1,
			pos:      1,
			width:    5,
			expected: "",
		},
		{
			name:     "with multiple pages",
			len:      3,
			pos:      1,
			width:    5,
			expected: "[lightblue]*_*[white]",
		},
		{
			name:     "with negative pages",
			len:      -3,
			pos:      1,
			width:    5,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, testCfg.PaginationMarker(tt.len, tt.pos, tt.width))
		})
	}
}

func Test_Validations(t *testing.T) {
	assert.Equal(t, 4, len(testCfg.Validations()))
}
