package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewDefaultColorTheme(t *testing.T) {
	theme := NewDefaultColorTheme()

	assert.Equal(t, "orange", theme.BorderTheme.Focused)
	assert.Equal(t, "red", theme.TextTheme.Subheading)
	assert.Equal(t, "transparent", theme.WidgetTheme.Background)
}

func Test_NewDefaultColorConfig(t *testing.T) {
	cfg, err := NewDefaultColorConfig()

	assert.Nil(t, err)

	assert.Equal(t, "orange", cfg.UString("bordertheme.focused"))
	assert.Equal(t, "red", cfg.UString("texttheme.subheading"))
	assert.Equal(t, "transparent", cfg.UString("widgettheme.background"))
	assert.Equal(t, "", cfg.UString("widgettheme.missing"))
}
