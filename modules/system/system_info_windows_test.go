//go:build windows

package system

import (
	"testing"

	"gotest.tools/assert"
)

func Test_windowsProductVersion(t *testing.T) {
	tests := []struct {
		name            string
		rawVersion      string
		expectedProduct string
		expectedBuild   string
	}{
		{
			name:            "windows 11",
			rawVersion:      "10.0.26200",
			expectedProduct: "Windows 11",
			expectedBuild:   "26200",
		},
		{
			name:            "windows 11, first build",
			rawVersion:      "10.0.22000",
			expectedProduct: "Windows 11",
			expectedBuild:   "22000",
		},
		{
			name:            "windows 10",
			rawVersion:      "10.0.19045",
			expectedProduct: "Windows 10",
			expectedBuild:   "19045",
		},
		{
			name:            "windows 10, build just below the windows 11 threshold",
			rawVersion:      "10.0.21999",
			expectedProduct: "Windows 10",
			expectedBuild:   "21999",
		},
		{
			name:            "with trailing whitespace",
			rawVersion:      "10.0.26200\r\n",
			expectedProduct: "Windows 11",
			expectedBuild:   "26200",
		},
		{
			name:            "non-windows-10 major version",
			rawVersion:      "6.1.7601",
			expectedProduct: "Windows 6.1",
			expectedBuild:   "7601",
		},
		{
			name:            "malformed version",
			rawVersion:      "not-a-version",
			expectedProduct: "Windows not-a-version",
			expectedBuild:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, build := windowsProductVersion(tt.rawVersion)

			assert.Equal(t, tt.expectedProduct, product)
			assert.Equal(t, tt.expectedBuild, build)
		})
	}
}
