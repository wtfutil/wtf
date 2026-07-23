package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WtfConfigDir_EnvOverride(t *testing.T) {
	t.Run("WTF_CONFIG_DIR takes priority", func(t *testing.T) {
		customDir := t.TempDir()
		t.Setenv("WTF_CONFIG_DIR", customDir)
		t.Setenv("XDG_CONFIG_HOME", "/should/not/use/this")

		dir, err := WtfConfigDir()
		assert.NoError(t, err)
		assert.Equal(t, customDir, dir)
	})

	t.Run("falls back to XDG_CONFIG_HOME when WTF_CONFIG_DIR unset", func(t *testing.T) {
		t.Setenv("WTF_CONFIG_DIR", "")
		t.Setenv("XDG_CONFIG_HOME", t.TempDir())

		dir, err := WtfConfigDir()
		assert.NoError(t, err)
		assert.Contains(t, dir, "wtf")
	})
}
