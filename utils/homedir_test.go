package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExpandHomeDir(t *testing.T) {
	tests := []struct {
		name             string
		path             string
		expectedStart    string
		expectedContains string
		expectedError    error
	}{
		{
			name:             "with empty path",
			path:             "",
			expectedStart:    "",
			expectedContains: "",
			expectedError:    nil,
		},
		{
			name:             "with relative path",
			path:             "~/test",
			expectedStart:    "/",
			expectedContains: "/test",
			expectedError:    nil,
		},
		{
			name:             "with absolute path",
			path:             "/Users/test",
			expectedStart:    "/",
			expectedContains: "/test",
			expectedError:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ExpandHomeDir(tt.path)

			if len(tt.path) > 0 {
				assert.Equal(t, tt.expectedStart, string(actual[0]))
			}

			assert.Contains(t, actual, tt.expectedContains)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
