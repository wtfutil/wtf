package utils

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExpandHomeDir(t *testing.T) {
	tests := []struct {
		name             string
		path             string
		expectedContains string
		expectedError    error
	}{
		{
			name:             "with empty path",
			path:             "",
			expectedContains: "",
			expectedError:    nil,
		},
		{
			name:             "with relative path",
			path:             "~/test",
			expectedContains: "test",
			expectedError:    nil,
		},
		{
			name:             "with absolute path",
			path:             "/Users/test",
			expectedContains: "test",
			expectedError:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ExpandHomeDir(tt.path)

			if len(tt.path) > 0 && tt.path[0] == '~' {
				// Tilde-expanded paths should be absolute
				assert.True(t, filepath.IsAbs(actual), "expected absolute path, got: %s", actual)
			}

			assert.Contains(t, actual, tt.expectedContains)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
